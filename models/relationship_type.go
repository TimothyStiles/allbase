// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// RelationshipType is an object representing the database table.
type RelationshipType struct {
	RelationshipType string      `boil:"relationship_type" json:"relationship_type" toml:"relationship_type" yaml:"relationship_type"`
	RelationshipDesc null.String `boil:"relationship_desc" json:"relationship_desc,omitempty" toml:"relationship_desc" yaml:"relationship_desc,omitempty"`

	R *relationshipTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L relationshipTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RelationshipTypeColumns = struct {
	RelationshipType string
	RelationshipDesc string
}{
	RelationshipType: "relationship_type",
	RelationshipDesc: "relationship_desc",
}

var RelationshipTypeTableColumns = struct {
	RelationshipType string
	RelationshipDesc string
}{
	RelationshipType: "relationship_type.relationship_type",
	RelationshipDesc: "relationship_type.relationship_desc",
}

// Generated where

var RelationshipTypeWhere = struct {
	RelationshipType whereHelperstring
	RelationshipDesc whereHelpernull_String
}{
	RelationshipType: whereHelperstring{field: "\"relationship_type\".\"relationship_type\""},
	RelationshipDesc: whereHelpernull_String{field: "\"relationship_type\".\"relationship_desc\""},
}

// RelationshipTypeRels is where relationship names are stored.
var RelationshipTypeRels = struct {
	Assays string
}{
	Assays: "Assays",
}

// relationshipTypeR is where relationships are stored.
type relationshipTypeR struct {
	Assays AssaySlice `boil:"Assays" json:"Assays" toml:"Assays" yaml:"Assays"`
}

// NewStruct creates a new relationship struct
func (*relationshipTypeR) NewStruct() *relationshipTypeR {
	return &relationshipTypeR{}
}

func (r *relationshipTypeR) GetAssays() AssaySlice {
	if r == nil {
		return nil
	}
	return r.Assays
}

// relationshipTypeL is where Load methods for each relationship are stored.
type relationshipTypeL struct{}

var (
	relationshipTypeAllColumns            = []string{"relationship_type", "relationship_desc"}
	relationshipTypeColumnsWithoutDefault = []string{"relationship_type"}
	relationshipTypeColumnsWithDefault    = []string{"relationship_desc"}
	relationshipTypePrimaryKeyColumns     = []string{"relationship_type"}
	relationshipTypeGeneratedColumns      = []string{}
)

type (
	// RelationshipTypeSlice is an alias for a slice of pointers to RelationshipType.
	// This should almost always be used instead of []RelationshipType.
	RelationshipTypeSlice []*RelationshipType
	// RelationshipTypeHook is the signature for custom RelationshipType hook methods
	RelationshipTypeHook func(context.Context, boil.ContextExecutor, *RelationshipType) error

	relationshipTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	relationshipTypeType                 = reflect.TypeOf(&RelationshipType{})
	relationshipTypeMapping              = queries.MakeStructMapping(relationshipTypeType)
	relationshipTypePrimaryKeyMapping, _ = queries.BindMapping(relationshipTypeType, relationshipTypeMapping, relationshipTypePrimaryKeyColumns)
	relationshipTypeInsertCacheMut       sync.RWMutex
	relationshipTypeInsertCache          = make(map[string]insertCache)
	relationshipTypeUpdateCacheMut       sync.RWMutex
	relationshipTypeUpdateCache          = make(map[string]updateCache)
	relationshipTypeUpsertCacheMut       sync.RWMutex
	relationshipTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var relationshipTypeAfterSelectHooks []RelationshipTypeHook

var relationshipTypeBeforeInsertHooks []RelationshipTypeHook
var relationshipTypeAfterInsertHooks []RelationshipTypeHook

var relationshipTypeBeforeUpdateHooks []RelationshipTypeHook
var relationshipTypeAfterUpdateHooks []RelationshipTypeHook

var relationshipTypeBeforeDeleteHooks []RelationshipTypeHook
var relationshipTypeAfterDeleteHooks []RelationshipTypeHook

var relationshipTypeBeforeUpsertHooks []RelationshipTypeHook
var relationshipTypeAfterUpsertHooks []RelationshipTypeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RelationshipType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RelationshipType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RelationshipType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RelationshipType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RelationshipType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RelationshipType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RelationshipType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RelationshipType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RelationshipType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range relationshipTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRelationshipTypeHook registers your hook function for all future operations.
func AddRelationshipTypeHook(hookPoint boil.HookPoint, relationshipTypeHook RelationshipTypeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		relationshipTypeAfterSelectHooks = append(relationshipTypeAfterSelectHooks, relationshipTypeHook)
	case boil.BeforeInsertHook:
		relationshipTypeBeforeInsertHooks = append(relationshipTypeBeforeInsertHooks, relationshipTypeHook)
	case boil.AfterInsertHook:
		relationshipTypeAfterInsertHooks = append(relationshipTypeAfterInsertHooks, relationshipTypeHook)
	case boil.BeforeUpdateHook:
		relationshipTypeBeforeUpdateHooks = append(relationshipTypeBeforeUpdateHooks, relationshipTypeHook)
	case boil.AfterUpdateHook:
		relationshipTypeAfterUpdateHooks = append(relationshipTypeAfterUpdateHooks, relationshipTypeHook)
	case boil.BeforeDeleteHook:
		relationshipTypeBeforeDeleteHooks = append(relationshipTypeBeforeDeleteHooks, relationshipTypeHook)
	case boil.AfterDeleteHook:
		relationshipTypeAfterDeleteHooks = append(relationshipTypeAfterDeleteHooks, relationshipTypeHook)
	case boil.BeforeUpsertHook:
		relationshipTypeBeforeUpsertHooks = append(relationshipTypeBeforeUpsertHooks, relationshipTypeHook)
	case boil.AfterUpsertHook:
		relationshipTypeAfterUpsertHooks = append(relationshipTypeAfterUpsertHooks, relationshipTypeHook)
	}
}

// One returns a single relationshipType record from the query.
func (q relationshipTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RelationshipType, error) {
	o := &RelationshipType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for relationship_type")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RelationshipType records from the query.
func (q relationshipTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (RelationshipTypeSlice, error) {
	var o []*RelationshipType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RelationshipType slice")
	}

	if len(relationshipTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RelationshipType records in the query.
func (q relationshipTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count relationship_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q relationshipTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if relationship_type exists")
	}

	return count > 0, nil
}

// Assays retrieves all the assay's Assays with an executor.
func (o *RelationshipType) Assays(mods ...qm.QueryMod) assayQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"assays\".\"relationship_type\"=?", o.RelationshipType),
	)

	return Assays(queryMods...)
}

// LoadAssays allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (relationshipTypeL) LoadAssays(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRelationshipType interface{}, mods queries.Applicator) error {
	var slice []*RelationshipType
	var object *RelationshipType

	if singular {
		object = maybeRelationshipType.(*RelationshipType)
	} else {
		slice = *maybeRelationshipType.(*[]*RelationshipType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &relationshipTypeR{}
		}
		args = append(args, object.RelationshipType)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &relationshipTypeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.RelationshipType) {
					continue Outer
				}
			}

			args = append(args, obj.RelationshipType)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`assays`),
		qm.WhereIn(`assays.relationship_type in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load assays")
	}

	var resultSlice []*Assay
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice assays")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on assays")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for assays")
	}

	if len(assayAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Assays = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &assayR{}
			}
			foreign.R.AssayRelationshipType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.RelationshipType, foreign.RelationshipType) {
				local.R.Assays = append(local.R.Assays, foreign)
				if foreign.R == nil {
					foreign.R = &assayR{}
				}
				foreign.R.AssayRelationshipType = local
				break
			}
		}
	}

	return nil
}

// AddAssays adds the given related objects to the existing relationships
// of the relationship_type, optionally inserting them as new records.
// Appends related to o.R.Assays.
// Sets related.R.AssayRelationshipType appropriately.
func (o *RelationshipType) AddAssays(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Assay) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.RelationshipType, o.RelationshipType)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"assays\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"relationship_type"}),
				strmangle.WhereClause("\"", "\"", 0, assayPrimaryKeyColumns),
			)
			values := []interface{}{o.RelationshipType, rel.AssayID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.RelationshipType, o.RelationshipType)
		}
	}

	if o.R == nil {
		o.R = &relationshipTypeR{
			Assays: related,
		}
	} else {
		o.R.Assays = append(o.R.Assays, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &assayR{
				AssayRelationshipType: o,
			}
		} else {
			rel.R.AssayRelationshipType = o
		}
	}
	return nil
}

// SetAssays removes all previously related items of the
// relationship_type replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.AssayRelationshipType's Assays accordingly.
// Replaces o.R.Assays with related.
// Sets related.R.AssayRelationshipType's Assays accordingly.
func (o *RelationshipType) SetAssays(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Assay) error {
	query := "update \"assays\" set \"relationship_type\" = null where \"relationship_type\" = ?"
	values := []interface{}{o.RelationshipType}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Assays {
			queries.SetScanner(&rel.RelationshipType, nil)
			if rel.R == nil {
				continue
			}

			rel.R.AssayRelationshipType = nil
		}
		o.R.Assays = nil
	}

	return o.AddAssays(ctx, exec, insert, related...)
}

// RemoveAssays relationships from objects passed in.
// Removes related items from R.Assays (uses pointer comparison, removal does not keep order)
// Sets related.R.AssayRelationshipType.
func (o *RelationshipType) RemoveAssays(ctx context.Context, exec boil.ContextExecutor, related ...*Assay) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.RelationshipType, nil)
		if rel.R != nil {
			rel.R.AssayRelationshipType = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("relationship_type")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Assays {
			if rel != ri {
				continue
			}

			ln := len(o.R.Assays)
			if ln > 1 && i < ln-1 {
				o.R.Assays[i] = o.R.Assays[ln-1]
			}
			o.R.Assays = o.R.Assays[:ln-1]
			break
		}
	}

	return nil
}

// RelationshipTypes retrieves all the records using an executor.
func RelationshipTypes(mods ...qm.QueryMod) relationshipTypeQuery {
	mods = append(mods, qm.From("\"relationship_type\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"relationship_type\".*"})
	}

	return relationshipTypeQuery{q}
}

// FindRelationshipType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRelationshipType(ctx context.Context, exec boil.ContextExecutor, relationshipType string, selectCols ...string) (*RelationshipType, error) {
	relationshipTypeObj := &RelationshipType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"relationship_type\" where \"relationship_type\"=?", sel,
	)

	q := queries.Raw(query, relationshipType)

	err := q.Bind(ctx, exec, relationshipTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from relationship_type")
	}

	if err = relationshipTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return relationshipTypeObj, err
	}

	return relationshipTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RelationshipType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no relationship_type provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(relationshipTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	relationshipTypeInsertCacheMut.RLock()
	cache, cached := relationshipTypeInsertCache[key]
	relationshipTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			relationshipTypeAllColumns,
			relationshipTypeColumnsWithDefault,
			relationshipTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(relationshipTypeType, relationshipTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(relationshipTypeType, relationshipTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"relationship_type\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"relationship_type\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into relationship_type")
	}

	if !cached {
		relationshipTypeInsertCacheMut.Lock()
		relationshipTypeInsertCache[key] = cache
		relationshipTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RelationshipType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RelationshipType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	relationshipTypeUpdateCacheMut.RLock()
	cache, cached := relationshipTypeUpdateCache[key]
	relationshipTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			relationshipTypeAllColumns,
			relationshipTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update relationship_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"relationship_type\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, relationshipTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(relationshipTypeType, relationshipTypeMapping, append(wl, relationshipTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update relationship_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for relationship_type")
	}

	if !cached {
		relationshipTypeUpdateCacheMut.Lock()
		relationshipTypeUpdateCache[key] = cache
		relationshipTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q relationshipTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for relationship_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for relationship_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RelationshipTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), relationshipTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"relationship_type\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, relationshipTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in relationshipType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all relationshipType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RelationshipType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no relationship_type provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(relationshipTypeColumnsWithDefault, o)

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

	relationshipTypeUpsertCacheMut.RLock()
	cache, cached := relationshipTypeUpsertCache[key]
	relationshipTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			relationshipTypeAllColumns,
			relationshipTypeColumnsWithDefault,
			relationshipTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			relationshipTypeAllColumns,
			relationshipTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert relationship_type, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(relationshipTypePrimaryKeyColumns))
			copy(conflict, relationshipTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"relationship_type\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(relationshipTypeType, relationshipTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(relationshipTypeType, relationshipTypeMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert relationship_type")
	}

	if !cached {
		relationshipTypeUpsertCacheMut.Lock()
		relationshipTypeUpsertCache[key] = cache
		relationshipTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RelationshipType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RelationshipType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RelationshipType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), relationshipTypePrimaryKeyMapping)
	sql := "DELETE FROM \"relationship_type\" WHERE \"relationship_type\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from relationship_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for relationship_type")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q relationshipTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no relationshipTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from relationship_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for relationship_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RelationshipTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(relationshipTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), relationshipTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"relationship_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, relationshipTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from relationshipType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for relationship_type")
	}

	if len(relationshipTypeAfterDeleteHooks) != 0 {
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
func (o *RelationshipType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRelationshipType(ctx, exec, o.RelationshipType)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RelationshipTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RelationshipTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), relationshipTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"relationship_type\".* FROM \"relationship_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, relationshipTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RelationshipTypeSlice")
	}

	*o = slice

	return nil
}

// RelationshipTypeExists checks if the RelationshipType row exists.
func RelationshipTypeExists(ctx context.Context, exec boil.ContextExecutor, relationshipType string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"relationship_type\" where \"relationship_type\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, relationshipType)
	}
	row := exec.QueryRowContext(ctx, sql, relationshipType)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if relationship_type exists")
	}

	return exists, nil
}