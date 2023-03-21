package dao

import (
	"context"
	"database/sql"
	"fmt"

	account_model "github.com/fabl3ss/banking_system/modules/account/internal/model"
	"github.com/fabl3ss/banking_system/pkg/customerrors"
	"github.com/google/uuid"
)

type DepositAccountProjectionDAO struct {
	db             *sql.DB
	tableName      string
	prepStatements *creditDepositAccountProjectionDAOPrepStatements
	accountDao     *AccountProjectionDAO
}

type creditDepositAccountProjectionDAOPrepStatements struct {
	GetById *sql.Stmt
}

type depositAccountQueryObject struct {
	accountId     uuid.UUID
	depositAmount int64
	annualRate    int
}

func NewDepositAccountProjectionDAO(db *sql.DB) *DepositAccountProjectionDAO {
	dao := &DepositAccountProjectionDAO{
		db:        db,
		tableName: "deposit_accounts",
	}
	dao.prepareStatements()

	return dao
}

func (dao *DepositAccountProjectionDAO) GetByID(ctx context.Context, id uuid.UUID) (*account_model.DepositAccount, error) {
	queryObject := new(depositAccountQueryObject)

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
			&queryObject.depositAmount,
			&queryObject.annualRate,
		); err != nil {
			return nil, err
		}
	}

	account, err := dao.accountDao.GetByID(ctx, queryObject.accountId)
	if err != nil {
		return nil, err
	}

	return account_model.NewDepositAccount(
		*account,
		queryObject.depositAmount,
		queryObject.annualRate,
	), nil
}

func (dao *DepositAccountProjectionDAO) prepareStatements() {
	getByIdQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE account_id = $1`,
		dao.tableName,
	)

	getById, err := dao.db.Prepare(getByIdQuery)
	if err != nil {
		panic(err)
	}

	dao.prepStatements = &creditDepositAccountProjectionDAOPrepStatements{
		GetById: getById,
	}
}
