package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	transaction_model "github.com/fabl3ss/banking_system/modules/transfer/internal/model"
	"github.com/google/uuid"

	"github.com/fabl3ss/banking_system/pkg/customerrors"
)

type TransferProjectionDAO struct {
	db             *sql.DB
	tableName      string
	prepStatements *customerProjectionDAOPrepStatements
}

type customerProjectionDAOPrepStatements struct {
	GetById *sql.Stmt
}

type transferQueryObject struct {
	id                 uuid.UUID
	senderAccountId    uuid.UUID
	recipientAccountId uuid.UUID
	amount             int64
	currencyCode       string
	createdAt          time.Time
	updatedAt          time.Time
	status             string
}

func NewTransferProjectionDAO(db *sql.DB) *TransferProjectionDAO {
	dao := &TransferProjectionDAO{
		db:        db,
		tableName: "transfers",
	}
	dao.prepareStatements()

	return dao
}

func (dao *TransferProjectionDAO) GetById(ctx context.Context, id uuid.UUID) (*transaction_model.Transfer, error) {
	queryObject := new(transferQueryObject)

	rows, err := dao.prepStatements.GetById.QueryContext(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerrors.NewNotFoundError("transfer not found")
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryObject.id,
			&queryObject.senderAccountId,
			&queryObject.recipientAccountId,
			&queryObject.amount,
			&queryObject.currencyCode,
			&queryObject.createdAt,
			&queryObject.updatedAt,
			&queryObject.status,
		); err != nil {
			return nil, err
		}
	}

	return transaction_model.NewTransfer(
		queryObject.id,
		queryObject.senderAccountId,
		queryObject.recipientAccountId,
		queryObject.amount,
		queryObject.currencyCode,
		queryObject.createdAt,
		queryObject.updatedAt,
		mapTransferStatus(queryObject.status),
	), nil
}

func (dao *TransferProjectionDAO) prepareStatements() {
	getByIdQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE id = $1`,
		dao.tableName,
	)

	getById, err := dao.db.Prepare(getByIdQuery)
	if err != nil {
		panic(err)
	}

	dao.prepStatements = &customerProjectionDAOPrepStatements{
		GetById: getById,
	}
}

func mapTransferStatus(status string) transaction_model.TransferStatus {
	switch status {
	case "pending":
		return transaction_model.TransferStatusPending
	case "approved":
		return transaction_model.TransferStatusApproved
	case "declined":
		return transaction_model.TransferStatusDeclined
	}

	return "undefined"
}
