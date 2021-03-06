// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BigflipLog is an object representing the database table.
type BigflipLog struct {
	ID              string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	TransactionID   int64     `boil:"transaction_id" json:"transaction_id" toml:"transaction_id" yaml:"transaction_id"`
	Amount          int       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Status          string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	TRXTimestamp    null.Time `boil:"trx_timestamp" json:"trx_timestamp,omitempty" toml:"trx_timestamp" yaml:"trx_timestamp,omitempty"`
	BankCode        string    `boil:"bank_code" json:"bank_code" toml:"bank_code" yaml:"bank_code"`
	AccountNumber   string    `boil:"account_number" json:"account_number" toml:"account_number" yaml:"account_number"`
	BeneficiaryName string    `boil:"beneficiary_name" json:"beneficiary_name" toml:"beneficiary_name" yaml:"beneficiary_name"`
	Remark          string    `boil:"remark" json:"remark" toml:"remark" yaml:"remark"`
	Receipt         string    `boil:"receipt" json:"receipt" toml:"receipt" yaml:"receipt"`
	TimeServed      null.Time `boil:"time_served" json:"time_served,omitempty" toml:"time_served" yaml:"time_served,omitempty"`
	Fee             int       `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	WithdrawalID    string    `boil:"withdrawal_id" json:"withdrawal_id" toml:"withdrawal_id" yaml:"withdrawal_id"`

	R *bigflipLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bigflipLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BigflipLogColumns = struct {
	ID              string
	TransactionID   string
	Amount          string
	Status          string
	TRXTimestamp    string
	BankCode        string
	AccountNumber   string
	BeneficiaryName string
	Remark          string
	Receipt         string
	TimeServed      string
	Fee             string
	WithdrawalID    string
}{
	ID:              "id",
	TransactionID:   "transaction_id",
	Amount:          "amount",
	Status:          "status",
	TRXTimestamp:    "trx_timestamp",
	BankCode:        "bank_code",
	AccountNumber:   "account_number",
	BeneficiaryName: "beneficiary_name",
	Remark:          "remark",
	Receipt:         "receipt",
	TimeServed:      "time_served",
	Fee:             "fee",
	WithdrawalID:    "withdrawal_id",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BigflipLogWhere = struct {
	ID              whereHelperstring
	TransactionID   whereHelperint64
	Amount          whereHelperint
	Status          whereHelperstring
	TRXTimestamp    whereHelpernull_Time
	BankCode        whereHelperstring
	AccountNumber   whereHelperstring
	BeneficiaryName whereHelperstring
	Remark          whereHelperstring
	Receipt         whereHelperstring
	TimeServed      whereHelpernull_Time
	Fee             whereHelperint
	WithdrawalID    whereHelperstring
}{
	ID:              whereHelperstring{field: "\"bigflip_logs\".\"id\""},
	TransactionID:   whereHelperint64{field: "\"bigflip_logs\".\"transaction_id\""},
	Amount:          whereHelperint{field: "\"bigflip_logs\".\"amount\""},
	Status:          whereHelperstring{field: "\"bigflip_logs\".\"status\""},
	TRXTimestamp:    whereHelpernull_Time{field: "\"bigflip_logs\".\"trx_timestamp\""},
	BankCode:        whereHelperstring{field: "\"bigflip_logs\".\"bank_code\""},
	AccountNumber:   whereHelperstring{field: "\"bigflip_logs\".\"account_number\""},
	BeneficiaryName: whereHelperstring{field: "\"bigflip_logs\".\"beneficiary_name\""},
	Remark:          whereHelperstring{field: "\"bigflip_logs\".\"remark\""},
	Receipt:         whereHelperstring{field: "\"bigflip_logs\".\"receipt\""},
	TimeServed:      whereHelpernull_Time{field: "\"bigflip_logs\".\"time_served\""},
	Fee:             whereHelperint{field: "\"bigflip_logs\".\"fee\""},
	WithdrawalID:    whereHelperstring{field: "\"bigflip_logs\".\"withdrawal_id\""},
}

// BigflipLogRels is where relationship names are stored.
var BigflipLogRels = struct {
	Withdrawal string
}{
	Withdrawal: "Withdrawal",
}

// bigflipLogR is where relationships are stored.
type bigflipLogR struct {
	Withdrawal *Withdrawal `boil:"Withdrawal" json:"Withdrawal" toml:"Withdrawal" yaml:"Withdrawal"`
}

// NewStruct creates a new relationship struct
func (*bigflipLogR) NewStruct() *bigflipLogR {
	return &bigflipLogR{}
}

// bigflipLogL is where Load methods for each relationship are stored.
type bigflipLogL struct{}

var (
	bigflipLogAllColumns            = []string{"id", "transaction_id", "amount", "status", "trx_timestamp", "bank_code", "account_number", "beneficiary_name", "remark", "receipt", "time_served", "fee", "withdrawal_id"}
	bigflipLogColumnsWithoutDefault = []string{"transaction_id", "status", "trx_timestamp", "bank_code", "account_number", "beneficiary_name", "remark", "receipt", "time_served", "withdrawal_id"}
	bigflipLogColumnsWithDefault    = []string{"id", "amount", "fee"}
	bigflipLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// BigflipLogSlice is an alias for a slice of pointers to BigflipLog.
	// This should generally be used opposed to []BigflipLog.
	BigflipLogSlice []*BigflipLog
	// BigflipLogHook is the signature for custom BigflipLog hook methods
	BigflipLogHook func(context.Context, boil.ContextExecutor, *BigflipLog) error

	bigflipLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bigflipLogType                 = reflect.TypeOf(&BigflipLog{})
	bigflipLogMapping              = queries.MakeStructMapping(bigflipLogType)
	bigflipLogPrimaryKeyMapping, _ = queries.BindMapping(bigflipLogType, bigflipLogMapping, bigflipLogPrimaryKeyColumns)
	bigflipLogInsertCacheMut       sync.RWMutex
	bigflipLogInsertCache          = make(map[string]insertCache)
	bigflipLogUpdateCacheMut       sync.RWMutex
	bigflipLogUpdateCache          = make(map[string]updateCache)
	bigflipLogUpsertCacheMut       sync.RWMutex
	bigflipLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bigflipLogBeforeInsertHooks []BigflipLogHook
var bigflipLogBeforeUpdateHooks []BigflipLogHook
var bigflipLogBeforeDeleteHooks []BigflipLogHook
var bigflipLogBeforeUpsertHooks []BigflipLogHook

var bigflipLogAfterInsertHooks []BigflipLogHook
var bigflipLogAfterSelectHooks []BigflipLogHook
var bigflipLogAfterUpdateHooks []BigflipLogHook
var bigflipLogAfterDeleteHooks []BigflipLogHook
var bigflipLogAfterUpsertHooks []BigflipLogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BigflipLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BigflipLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BigflipLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BigflipLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BigflipLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BigflipLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BigflipLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BigflipLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BigflipLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bigflipLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBigflipLogHook registers your hook function for all future operations.
func AddBigflipLogHook(hookPoint boil.HookPoint, bigflipLogHook BigflipLogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bigflipLogBeforeInsertHooks = append(bigflipLogBeforeInsertHooks, bigflipLogHook)
	case boil.BeforeUpdateHook:
		bigflipLogBeforeUpdateHooks = append(bigflipLogBeforeUpdateHooks, bigflipLogHook)
	case boil.BeforeDeleteHook:
		bigflipLogBeforeDeleteHooks = append(bigflipLogBeforeDeleteHooks, bigflipLogHook)
	case boil.BeforeUpsertHook:
		bigflipLogBeforeUpsertHooks = append(bigflipLogBeforeUpsertHooks, bigflipLogHook)
	case boil.AfterInsertHook:
		bigflipLogAfterInsertHooks = append(bigflipLogAfterInsertHooks, bigflipLogHook)
	case boil.AfterSelectHook:
		bigflipLogAfterSelectHooks = append(bigflipLogAfterSelectHooks, bigflipLogHook)
	case boil.AfterUpdateHook:
		bigflipLogAfterUpdateHooks = append(bigflipLogAfterUpdateHooks, bigflipLogHook)
	case boil.AfterDeleteHook:
		bigflipLogAfterDeleteHooks = append(bigflipLogAfterDeleteHooks, bigflipLogHook)
	case boil.AfterUpsertHook:
		bigflipLogAfterUpsertHooks = append(bigflipLogAfterUpsertHooks, bigflipLogHook)
	}
}

// One returns a single bigflipLog record from the query.
func (q bigflipLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BigflipLog, error) {
	o := &BigflipLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bigflip_logs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BigflipLog records from the query.
func (q bigflipLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (BigflipLogSlice, error) {
	var o []*BigflipLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BigflipLog slice")
	}

	if len(bigflipLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BigflipLog records in the query.
func (q bigflipLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bigflip_logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bigflipLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bigflip_logs exists")
	}

	return count > 0, nil
}

// Withdrawal pointed to by the foreign key.
func (o *BigflipLog) Withdrawal(mods ...qm.QueryMod) withdrawalQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.WithdrawalID),
	}

	queryMods = append(queryMods, mods...)

	query := Withdrawals(queryMods...)
	queries.SetFrom(query.Query, "\"withdrawals\"")

	return query
}

// LoadWithdrawal allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (bigflipLogL) LoadWithdrawal(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBigflipLog interface{}, mods queries.Applicator) error {
	var slice []*BigflipLog
	var object *BigflipLog

	if singular {
		object = maybeBigflipLog.(*BigflipLog)
	} else {
		slice = *maybeBigflipLog.(*[]*BigflipLog)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bigflipLogR{}
		}
		args = append(args, object.WithdrawalID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bigflipLogR{}
			}

			for _, a := range args {
				if a == obj.WithdrawalID {
					continue Outer
				}
			}

			args = append(args, obj.WithdrawalID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`withdrawals`),
		qm.WhereIn(`withdrawals.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Withdrawal")
	}

	var resultSlice []*Withdrawal
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Withdrawal")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for withdrawals")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for withdrawals")
	}

	if len(bigflipLogAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Withdrawal = foreign
		if foreign.R == nil {
			foreign.R = &withdrawalR{}
		}
		foreign.R.BigflipLog = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.WithdrawalID == foreign.ID {
				local.R.Withdrawal = foreign
				if foreign.R == nil {
					foreign.R = &withdrawalR{}
				}
				foreign.R.BigflipLog = local
				break
			}
		}
	}

	return nil
}

// SetWithdrawal of the bigflipLog to the related item.
// Sets o.R.Withdrawal to related.
// Adds o to related.R.BigflipLog.
func (o *BigflipLog) SetWithdrawal(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Withdrawal) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"bigflip_logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"withdrawal_id"}),
		strmangle.WhereClause("\"", "\"", 2, bigflipLogPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.WithdrawalID = related.ID
	if o.R == nil {
		o.R = &bigflipLogR{
			Withdrawal: related,
		}
	} else {
		o.R.Withdrawal = related
	}

	if related.R == nil {
		related.R = &withdrawalR{
			BigflipLog: o,
		}
	} else {
		related.R.BigflipLog = o
	}

	return nil
}

// BigflipLogs retrieves all the records using an executor.
func BigflipLogs(mods ...qm.QueryMod) bigflipLogQuery {
	mods = append(mods, qm.From("\"bigflip_logs\""))
	return bigflipLogQuery{NewQuery(mods...)}
}

// FindBigflipLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBigflipLog(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*BigflipLog, error) {
	bigflipLogObj := &BigflipLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bigflip_logs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bigflipLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bigflip_logs")
	}

	return bigflipLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BigflipLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bigflip_logs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bigflipLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bigflipLogInsertCacheMut.RLock()
	cache, cached := bigflipLogInsertCache[key]
	bigflipLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bigflipLogAllColumns,
			bigflipLogColumnsWithDefault,
			bigflipLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bigflipLogType, bigflipLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bigflipLogType, bigflipLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bigflip_logs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bigflip_logs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into bigflip_logs")
	}

	if !cached {
		bigflipLogInsertCacheMut.Lock()
		bigflipLogInsertCache[key] = cache
		bigflipLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BigflipLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BigflipLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bigflipLogUpdateCacheMut.RLock()
	cache, cached := bigflipLogUpdateCache[key]
	bigflipLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bigflipLogAllColumns,
			bigflipLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bigflip_logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bigflip_logs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bigflipLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bigflipLogType, bigflipLogMapping, append(wl, bigflipLogPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update bigflip_logs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bigflip_logs")
	}

	if !cached {
		bigflipLogUpdateCacheMut.Lock()
		bigflipLogUpdateCache[key] = cache
		bigflipLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bigflipLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bigflip_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bigflip_logs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BigflipLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bigflipLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bigflip_logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bigflipLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bigflipLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bigflipLog")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BigflipLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bigflip_logs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bigflipLogColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bigflipLogUpsertCacheMut.RLock()
	cache, cached := bigflipLogUpsertCache[key]
	bigflipLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bigflipLogAllColumns,
			bigflipLogColumnsWithDefault,
			bigflipLogColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bigflipLogAllColumns,
			bigflipLogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bigflip_logs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bigflipLogPrimaryKeyColumns))
			copy(conflict, bigflipLogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"bigflip_logs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bigflipLogType, bigflipLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bigflipLogType, bigflipLogMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert bigflip_logs")
	}

	if !cached {
		bigflipLogUpsertCacheMut.Lock()
		bigflipLogUpsertCache[key] = cache
		bigflipLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BigflipLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BigflipLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BigflipLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bigflipLogPrimaryKeyMapping)
	sql := "DELETE FROM \"bigflip_logs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bigflip_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bigflip_logs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bigflipLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bigflipLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bigflip_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bigflip_logs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BigflipLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bigflipLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bigflipLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bigflip_logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bigflipLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bigflipLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bigflip_logs")
	}

	if len(bigflipLogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BigflipLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBigflipLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BigflipLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BigflipLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bigflipLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bigflip_logs\".* FROM \"bigflip_logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bigflipLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BigflipLogSlice")
	}

	*o = slice

	return nil
}

// BigflipLogExists checks if the BigflipLog row exists.
func BigflipLogExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bigflip_logs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bigflip_logs exists")
	}

	return exists, nil
}
