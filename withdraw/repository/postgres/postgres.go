package postgres

import (
	"context"
	"database/sql"
	"flip/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type WithdrawRepository interface {
	Insert(ctx context.Context, withdrawal models.Withdrawal) (*models.Withdrawal, error)
	Update(ctx context.Context, withdrawal models.Withdrawal) (*models.Withdrawal, error)
	Get(ctx context.Context, id string) (*models.Withdrawal, error)
}

type withdrawer struct {
	db *sql.DB
}

func NewWithdrawer(db *sql.DB) WithdrawRepository {
	return &withdrawer{db: db}
}

func (w *withdrawer) Insert(ctx context.Context, withdraw models.Withdrawal) (*models.Withdrawal, error) {
	withdraw.ID = uuid.NewString()
	err := withdraw.Insert(ctx, w.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "error insert withdrawal")
	}
	return &withdraw, nil
}

func (w *withdrawer) Update(ctx context.Context, withdraw models.Withdrawal) (*models.Withdrawal, error) {
	_, err := withdraw.Update(ctx, w.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "error update withdrawal")
	}
	return &withdraw, nil
}

func (w *withdrawer) Get(ctx context.Context, id string) (*models.Withdrawal, error) {
	withdrawal, err := models.FindWithdrawal(ctx, w.db, id)
	if err != nil {
		return nil, errors.Wrap(err, "error get withdrawal")
	}
	return withdrawal, nil
}
