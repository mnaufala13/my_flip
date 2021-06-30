package usecase

import (
	"context"
	"database/sql"
	"flip/biglip/repository/postgres"
	"flip/biglip/repository/svc"
	"flip/domain"
	"log"
)

type BigFlipUsecase interface {
	Disburse(ctx context.Context, withdrawalId string, payload domain.DisbursePayload) (*domain.FlipTransaction, error)
	Sync(ctx context.Context) error
}

type bigflipUC struct {
	db       *sql.DB
	svcFlip  svc.BigFlipSvcRepository
	psqlFlip postgres.BigFlipPsqlRepository
}

func NewBigflipUC(db *sql.DB, svcFlip svc.BigFlipSvcRepository, psqlFlip postgres.BigFlipPsqlRepository) BigFlipUsecase {
	return &bigflipUC{db: db, svcFlip: svcFlip, psqlFlip: psqlFlip}
}

// Disburse to call api bigflip
func (b *bigflipUC) Disburse(ctx context.Context, withdrawalId string, payload domain.DisbursePayload) (*domain.FlipTransaction, error) {
	lg, err := b.svcFlip.Disburse(ctx, withdrawalId, payload)
	if err != nil {
		return nil, err
	}
	return lg, nil
}

// Sync to sync pending bigflip transactions
func (b *bigflipUC) Sync(ctx context.Context) error {
	logs, err := b.psqlFlip.FetchPending(ctx)
	if err != nil {
		return err
	}
	if len(logs) == 0 {
		return nil
	}
	for _, l := range logs {
		trx, _ := b.svcFlip.Status(ctx, int(l.TransactionID))
		if trx.Status != domain.WithdrawSuccess {
			continue
		}
		trxModel, err := domain.FlipTransactionToModel(*trx)
		if err != nil {
			log.Println(err)
			continue
		}
		l.Status = trxModel.Status
		l.TimeServed = trxModel.TimeServed
		l.Receipt = trx.Receipt
		_, err = b.psqlFlip.Update(ctx, *l)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
