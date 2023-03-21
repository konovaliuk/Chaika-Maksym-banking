package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	account_model "github.com/fabl3ss/banking_system/modules/account/internal/model"
	"github.com/fabl3ss/banking_system/pkg/customerrors"
	"github.com/google/uuid"
)

type AccountProjectionDAO struct {
	db             *sql.DB
	tableName      string
	prepStatements *accountProjectionDAOPrepStatements
}

type accountProjectionDAOPrepStatements struct {
	GetById *sql.Stmt
}

type accountQueryObject struct {
	id           uuid.UUID
	holderId     uuid.UUID
	currencyCode string
	amount       int64
	openedAt     time.Time
	expiryDate   time.Time
}

func NewAccountProjectionDAO(db *sql.DB) *AccountProjectionDAO {
	dao := &AccountProjectionDAO{
		db:        db,
		tableName: "account",
	}
	dao.prepareStatements()

	return dao
}

func (dao *AccountProjectionDAO) GetByID(ctx context.Context, id uuid.UUID) (*account_model.Account, error) {
	queryObject := new(accountQueryObject)

	rows, err := dao.prepStatements.GetById.QueryContext(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.NewNotFoundError("account not found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryObject.id,
			&queryObject.holderId,
			&queryObject.currencyCode,
			&queryObject.amount,
			&queryObject.openedAt,
			&queryObject.expiryDate,
		); err != nil {
			return nil, err
		}
	}

	return account_model.NewAccount(
		queryObject.id,
		queryObject.holderId,
		queryObject.currencyCode,
		queryObject.amount,
		queryObject.openedAt,
		queryObject.expiryDate,
	), nil
}

func (dao *AccountProjectionDAO) prepareStatements() {
	getByIdQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE id = $1`,
		dao.tableName,
	)

	getById, err := dao.db.Prepare(getByIdQuery)
	if err != nil {
		panic(err)
	}

	dao.prepStatements = &accountProjectionDAOPrepStatements{
		GetById: getById,
	}
}
