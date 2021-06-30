package postgres

import (
	"context"
	"database/sql"
	"flip/domain"
	"flip/models"
	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BigFlipPsqlRepository interface {
	Get(ctx context.Context, id string) (*models.BigflipLog, error)
	Insert(ctx context.Context, exec boil.ContextExecutor, log models.BigflipLog) (*models.BigflipLog, error)
	Update(ctx context.Context, log models.BigflipLog) (*models.BigflipLog, error)
	FetchPending(ctx context.Context) (models.BigflipLogSlice, error)
	SetSuccessBulk(ctx context.Context, logs models.BigflipLogSlice) error
}

type flipper struct {
	db *sql.DB
}

func NewFlipper(db *sql.DB) BigFlipPsqlRepository {
	return &flipper{db: db}
}

func (f *flipper) SetSuccessBulk(ctx context.Context, logs models.BigflipLogSlice) error {
	_, err := logs.UpdateAll(ctx, f.db, models.M{"status": domain.WithdrawSuccess})
	if err != nil {
		return errors.Wrap(err, "error update status success bulk mode")
	}
	return nil
}

func (f *flipper) Get(ctx context.Context, id string) (*models.BigflipLog, error) {
	l, err := models.FindBigflipLog(ctx, f.db, id)
	if err != nil {
		return nil, errors.Wrap(err, "error get bigflip log")
	}
	return l, nil
}

func (f *flipper) Insert(ctx context.Context, exec boil.ContextExecutor, log models.BigflipLog) (*models.BigflipLog, error) {
	log.ID = uuid.NewString()
	err := log.Insert(ctx, exec, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "error insert bigflip log")
	}
	return &log, nil
}

func (f *flipper) Update(ctx context.Context, log models.BigflipLog) (*models.BigflipLog, error) {
	_, err := log.Update(ctx, f.db, boil.Infer())
	if err != nil {
		return nil, errors.Wrap(err, "error update bigflip log")
	}
	return &log, nil
}

func (f *flipper) FetchPending(ctx context.Context) (models.BigflipLogSlice, error) {
	logs, err := models.BigflipLogs(models.BigflipLogWhere.Status.EQ(domain.WithdrawPending)).All(ctx, f.db)
	if err != nil {
		return nil, errors.Wrap(err, "error fetch pending bigflip")
	}
	return logs, nil
}
