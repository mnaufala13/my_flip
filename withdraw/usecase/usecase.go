package usecase

import (
	"context"
	"database/sql"
	"flip/biglip/usecase"
	"flip/domain"
	"flip/models"
	"flip/withdraw/repository/postgres"
	"log"
)

type WithdrawUsecase interface {
	Create(ctx context.Context, request domain.WithdrawRequest) (*models.Withdrawal, error)
	SetSuccessDisburse(ctx context.Context, id string) (*models.Withdrawal, error)
}

type withdrawUC struct {
	db       *sql.DB
	withdraw postgres.WithdrawRepository
	bigflip  usecase.BigFlipUsecase
}

func NewWithdrawUC(db *sql.DB, withdraw postgres.WithdrawRepository, bigflip usecase.BigFlipUsecase) WithdrawUsecase {
	return &withdrawUC{db: db, withdraw: withdraw, bigflip: bigflip}
}

func (w withdrawUC) Create(ctx context.Context, request domain.WithdrawRequest) (*models.Withdrawal, error) {
	wm, err := w.withdraw.Insert(ctx, models.Withdrawal{
		AccountNumber: request.AccountNumber,
		BankCode:      request.BankCode,
		Remark:        request.Remark,
		Amount:        request.Amount,
	})
	if err != nil {
		return nil, err
	}
	go func() {
		_, err := w.bigflip.Disburse(ctx, wm.ID, domain.DisbursePayload{
			BankCode:      wm.BankCode,
			AccountNumber: wm.AccountNumber,
			Amount:        wm.Amount,
			Remark:        wm.Remark,
		})
		if err != nil {
			log.Println(err)
			return
		}
		_, err = w.SetSuccessDisburse(ctx, wm.ID)
		if err != nil {
			log.Println(err)
			return
		}
	}()
	return wm, nil
}

func (w withdrawUC) SetSuccessDisburse(ctx context.Context, id string) (*models.Withdrawal, error) {
	wm, err := w.withdraw.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	wm.IsSuccess = true
	_, err = w.withdraw.Update(ctx, *wm)
	if err != nil {
		return nil, err
	}
	return wm, nil
}
