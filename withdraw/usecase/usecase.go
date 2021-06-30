package usecase

import (
	"context"
	"database/sql"
	"flip/biglip/usecase"
	"flip/domain"
	"flip/models"
	"flip/withdraw/repository/postgres"
	"log"
	"time"
)

type WithdrawUsecase interface {
	Create(ctx context.Context, request domain.WithdrawRequest) (*models.Withdrawal, error)
	History(ctx context.Context) ([]*domain.History, error)
	SetSuccessDisburse(ctx context.Context, id string) (*models.Withdrawal, error)
	SyncWithdrawal(ctx context.Context) error
}

type withdrawUC struct {
	db       *sql.DB
	withdraw postgres.WithdrawRepository
	bigflip  usecase.BigFlipUsecase
}

func NewWithdrawUC(db *sql.DB, withdraw postgres.WithdrawRepository, bigflip usecase.BigFlipUsecase) WithdrawUsecase {
	return &withdrawUC{db: db, withdraw: withdraw, bigflip: bigflip}
}

func (w withdrawUC) History(ctx context.Context) ([]*domain.History, error) {
	hh, err := w.withdraw.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	histories := []*domain.History{}
	for _, h := range hh {
		ht := &domain.History{
			Amount:        h.Amount,
			BankCode:      h.BankCode,
			AccountNumber: h.AccountNumber,
			Remark:        h.Remark,
		}
		if h.IsSuccess && h.R.BigflipLog != nil {
			ht.Status = h.R.BigflipLog.Status
			ht.Fee = h.R.BigflipLog.Fee
			ht.Receipt = h.R.BigflipLog.Receipt
			if !h.R.BigflipLog.TimeServed.Time.IsZero() {
				ht.TimeServed = h.R.BigflipLog.TimeServed.Time.Format(time.RFC3339)
			}
		}
		histories = append(histories, ht)
	}
	return histories, nil
}

func (w withdrawUC) SyncWithdrawal(ctx context.Context) error {
	withdrawals, err := w.withdraw.FetchNotSuccess(ctx)
	if err != nil {
		return err
	}
	if len(withdrawals) == 0 {
		return nil
	}
	successWithdrawals := models.WithdrawalSlice{}
	for _, wm := range withdrawals {
		_, err := w.bigflip.Disburse(ctx, wm.ID, domain.DisbursePayload{
			BankCode:      wm.BankCode,
			AccountNumber: wm.AccountNumber,
			Amount:        wm.Amount,
			Remark:        wm.Remark,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		successWithdrawals = append(successWithdrawals, wm)
	}
	if len(successWithdrawals) == 0 {
		return nil
	}
	err = w.withdraw.SetSuccessStatusBulk(ctx, successWithdrawals)
	if err != nil {
		return err
	}
	return nil
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
