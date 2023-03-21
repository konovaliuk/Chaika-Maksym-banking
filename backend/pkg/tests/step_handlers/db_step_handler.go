package step_handlers

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v16"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DBStepHandler struct {
	db *sql.DB
}

func NewDBStepHandler(db *sql.DB) *DBStepHandler {
	return &DBStepHandler{
		db: db,
	}
}

func (h *DBStepHandler) RegisterSteps(sc *godog.ScenarioContext) {
	sc.Step(`^the next records exist in "([^"]*)" table:$`, h.insertRecords)
	sc.Step(`^I see next records in "([^"]*)" table:$`, h.checkRecordsEquals)
}

func (h *DBStepHandler) insertRecords(tableName string, tableVars *godog.Table) error {
	columns := lo.Map(tableVars.Rows[0].Cells, func(cell *messages.PickleTableCell, _ int) string {
		return cell.Value
	})
	placeholders := lo.Map(columns, func(column string, i int) string {
		return fmt.Sprintf("$%d", i+1)
	})

	for _, tableRow := range tableVars.Rows[1:] {
		values := lo.Map(tableRow.Cells, func(cell *messages.PickleTableCell, _ int) interface{} {
			return cell.Value
		})

		_, err := h.db.Exec(
			fmt.Sprintf(
				"INSERT INTO %s (%s) VALUES (%s)",
				tableName,
				strings.Join(columns, ","),
				strings.Join(placeholders, ","),
			),
			values...,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *DBStepHandler) checkRecordsEquals(tableName string, tableVars *godog.Table) error {
	// Build the SQL query based on the table name and column names
	columns := make([]string, 0, len(tableVars.Rows[0].Cells))
	for _, cell := range tableVars.Rows[0].Cells {
		columns = append(columns, cell.Value)
	}
	sqlQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), tableName)

	// Execute the SQL query and fetch the results
	rows, err := h.db.Query(sqlQuery)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	// Build a slice of maps representing the database rows
	dbRows := make([]map[string]interface{}, 0)
	for rows.Next() {
		rowValues := make([]interface{}, 0, len(tableVars.Rows[0].Cells))
		for range tableVars.Rows[0].Cells {
			rowValues = append(rowValues, new(interface{}))
		}
		err := rows.Scan(rowValues...)
		if err != nil {
			return fmt.Errorf("error scanning database row: %w", err)
		}

		rowMap := make(map[string]interface{})
		for i, cell := range tableVars.Rows[0].Cells {
			if time, ok := (*rowValues[i].(*interface{})).(time.Time); ok {
				rowMap[cell.Value] = time.UTC()
				continue
			}

			rowMap[cell.Value] = *(rowValues[i].(*interface{}))
		}
		dbRows = append(dbRows, rowMap)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error fetching database rows: %w", err)
	}

	// Convert the table variables to a slice of maps
	tableRows := make([]map[string]interface{}, 0)
	for _, row := range tableVars.Rows[1:] {
		rowMap := make(map[string]interface{})
		for j, cell := range row.Cells {
			if _, err := uuid.Parse(cell.Value); err == nil {
				rowMap[tableVars.Rows[0].Cells[j].Value] = []byte(cell.Value)
				continue
			}
			if v, err := time.Parse(time.RFC3339, cell.Value); err == nil {
				rowMap[tableVars.Rows[0].Cells[j].Value] = v
				continue
			}

			rowMap[tableVars.Rows[0].Cells[j].Value] = cell.Value
		}
		tableRows = append(tableRows, rowMap)
	}

	if !reflect.DeepEqual(dbRows, tableRows) {
		return fmt.Errorf("mismatch between expected and actual database rows, Acutal: %v\nExpected: %v", dbRows, tableRows)
	}

	return nil
}
