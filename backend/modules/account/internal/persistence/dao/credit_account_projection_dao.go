package dao

import (
	"context"
	"database/sql"
	"fmt"

	account_model "github.com/fabl3ss/banking_system/modules/account/internal/model"
	"github.com/fabl3ss/banking_system/pkg/customerrors"
	"github.com/google/uuid"
)

type CreditAccountProjectionDAO struct {
	db             *sql.DB
	tableName      string
	prepStatements *creditCreditAccountProjectionDAOPrepStatements
	accountDao     *AccountProjectionDAO
}

type creditCreditAccountProjectionDAOPrepStatements struct {
	GetById *sql.Stmt
}

type creditAccountQueryObject struct {
	accountId       uuid.UUID
	creditLimit     int64
	debt            int64
	accruedInterest int
	creditRate      int
}

func NewCreditAccountProjectionDAO(db *sql.DB) *CreditAccountProjectionDAO {
	dao := &CreditAccountProjectionDAO{
		db:        db,
		tableName: "credit_accounts",
	}
	dao.prepareStatements()

	return dao
}

func (dao *CreditAccountProjectionDAO) GetByID(ctx context.Context, id uuid.UUID) (*account_model.CreditAccount, error) {
	queryObject := new(creditAccountQueryObject)

	rows, err := dao.prepStatements.GetById.QueryContext(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.NewNotFoundError("credit account not found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryObject.accountId,
			&queryObject.creditLimit,
			&queryObject.debt,
			&queryObject.accruedInterest,
			&queryObject.creditRate,
		); err != nil {
			return nil, err
		}
	}

	account, err := dao.accountDao.GetByID(ctx, queryObject.accountId)
	if err != nil {
		return nil, err
	}

	return account_model.NewCreditAccount(
		*account,
		queryObject.creditLimit,
		queryObject.debt,
		queryObject.accruedInterest,
		queryObject.creditRate,
	), nil
}

func (dao *CreditAccountProjectionDAO) prepareStatements() {
	getByIdQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE account_id = $1`,
		dao.tableName,
	)

	getById, err := dao.db.Prepare(getByIdQuery)
	if err != nil {
		panic(err)
	}

	dao.prepStatements = &creditCreditAccountProjectionDAOPrepStatements{
		GetById: getById,
	}
}
