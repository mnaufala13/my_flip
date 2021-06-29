package usecase

import (
	"context"
	"database/sql"
	"flip/biglip/repository/postgres"
	"flip/biglip/repository/svc"
	"flip/domain"
	"flip/models"
)

type BigFlipUsecase interface {
	Disburse(ctx context.Context, withdrawalId string, payload domain.DisbursePayload) (*models.BigflipLog, error)
	Sync(ctx context.Context) error
}

type bigflipUC struct {
	db       *sql.DB
	svcFlip  svc.BigFlipSvcRepository
	psqlFlip postgres.BigFlipPsqlRepository
}

// Disburse to call api bigflip
func (b *bigflipUC) Disburse(ctx context.Context, withdrawalId string, payload domain.DisbursePayload) (*models.BigflipLog, error) {
	lg, err := b.Disburse(ctx, withdrawalId, payload)
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
	logsSuccess := models.BigflipLogSlice{}
	for _, l := range logs {
		trx, _ := b.svcFlip.Status(ctx, int(l.TransactionID))
		if trx.Status == domain.WithdrawSuccess {
			logsSuccess = append(logsSuccess, l)
		}
	}
	err = b.psqlFlip.SetSuccessBulk(ctx, logsSuccess)
	if err != nil {
		return err
	}
	return nil
}
