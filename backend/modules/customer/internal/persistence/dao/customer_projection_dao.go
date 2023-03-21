package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	customer_model "github.com/fabl3ss/banking_system/modules/customer/internal/model"
	"github.com/fabl3ss/banking_system/pkg/customerrors"
)

type CustomerProjectionDAO struct {
	db             *sql.DB
	tableName      string
	prepStatements *customerProjectionDAOPrepStatements
}

type customerProjectionDAOPrepStatements struct {
	GetByEmail *sql.Stmt
}

type customerQueryObject struct {
	id           uuid.UUID
	email        string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewCustomerProjectionDAO(db *sql.DB) *CustomerProjectionDAO {
	dao := &CustomerProjectionDAO{
		db:        db,
		tableName: "customers",
	}
	dao.prepareStatements()

	return dao
}

func (dao *CustomerProjectionDAO) GetByEmail(ctx context.Context, email string) (*customer_model.Customer, error) {
	queryObject := new(customerQueryObject)

	rows, err := dao.prepStatements.GetByEmail.QueryContext(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.NewNotFoundError("customer not found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryObject.id,
			&queryObject.email,
			&queryObject.passwordHash,
			&queryObject.createdAt,
			&queryObject.updatedAt,
		); err != nil {
			return nil, err
		}
	}

	return customer_model.NewCustomer(
		queryObject.id,
		queryObject.email,
		queryObject.passwordHash,
		queryObject.createdAt,
		queryObject.createdAt,
	), nil
}

func (dao *CustomerProjectionDAO) prepareStatements() {
	getByEmailQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE email = $1`,
		dao.tableName,
	)

	getByEmail, err := dao.db.Prepare(getByEmailQuery)
	if err != nil {
		panic(err)
	}

	dao.prepStatements = &customerProjectionDAOPrepStatements{
		GetByEmail: getByEmail,
	}
}
