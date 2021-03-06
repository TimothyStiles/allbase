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

// CompoundStructure is an object representing the database table.
type CompoundStructure struct {
	Molregno         int64       `boil:"molregno" json:"molregno" toml:"molregno" yaml:"molregno"`
	Molfile          null.String `boil:"molfile" json:"molfile,omitempty" toml:"molfile" yaml:"molfile,omitempty"`
	StandardInchi    null.String `boil:"standard_inchi" json:"standard_inchi,omitempty" toml:"standard_inchi" yaml:"standard_inchi,omitempty"`
	StandardInchiKey string      `boil:"standard_inchi_key" json:"standard_inchi_key" toml:"standard_inchi_key" yaml:"standard_inchi_key"`
	CanonicalSmiles  null.String `boil:"canonical_smiles" json:"canonical_smiles,omitempty" toml:"canonical_smiles" yaml:"canonical_smiles,omitempty"`

	R *compoundStructureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L compoundStructureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompoundStructureColumns = struct {
	Molregno         string
	Molfile          string
	StandardInchi    string
	StandardInchiKey string
	CanonicalSmiles  string
}{
	Molregno:         "molregno",
	Molfile:          "molfile",
	StandardInchi:    "standard_inchi",
	StandardInchiKey: "standard_inchi_key",
	CanonicalSmiles:  "canonical_smiles",
}

var CompoundStructureTableColumns = struct {
	Molregno         string
	Molfile          string
	StandardInchi    string
	StandardInchiKey string
	CanonicalSmiles  string
}{
	Molregno:         "compound_structures.molregno",
	Molfile:          "compound_structures.molfile",
	StandardInchi:    "compound_structures.standard_inchi",
	StandardInchiKey: "compound_structures.standard_inchi_key",
	CanonicalSmiles:  "compound_structures.canonical_smiles",
}

// Generated where

var CompoundStructureWhere = struct {
	Molregno         whereHelperint64
	Molfile          whereHelpernull_String
	StandardInchi    whereHelpernull_String
	StandardInchiKey whereHelperstring
	CanonicalSmiles  whereHelpernull_String
}{
	Molregno:         whereHelperint64{field: "\"compound_structures\".\"molregno\""},
	Molfile:          whereHelpernull_String{field: "\"compound_structures\".\"molfile\""},
	StandardInchi:    whereHelpernull_String{field: "\"compound_structures\".\"standard_inchi\""},
	StandardInchiKey: whereHelperstring{field: "\"compound_structures\".\"standard_inchi_key\""},
	CanonicalSmiles:  whereHelpernull_String{field: "\"compound_structures\".\"canonical_smiles\""},
}

// CompoundStructureRels is where relationship names are stored.
var CompoundStructureRels = struct {
	MolregnoMoleculeDictionary string
}{
	MolregnoMoleculeDictionary: "MolregnoMoleculeDictionary",
}

// compoundStructureR is where relationships are stored.
type compoundStructureR struct {
	MolregnoMoleculeDictionary *MoleculeDictionary `boil:"MolregnoMoleculeDictionary" json:"MolregnoMoleculeDictionary" toml:"MolregnoMoleculeDictionary" yaml:"MolregnoMoleculeDictionary"`
}

// NewStruct creates a new relationship struct
func (*compoundStructureR) NewStruct() *compoundStructureR {
	return &compoundStructureR{}
}

func (r *compoundStructureR) GetMolregnoMoleculeDictionary() *MoleculeDictionary {
	if r == nil {
		return nil
	}
	return r.MolregnoMoleculeDictionary
}

// compoundStructureL is where Load methods for each relationship are stored.
type compoundStructureL struct{}

var (
	compoundStructureAllColumns            = []string{"molregno", "molfile", "standard_inchi", "standard_inchi_key", "canonical_smiles"}
	compoundStructureColumnsWithoutDefault = []string{"molregno", "standard_inchi_key"}
	compoundStructureColumnsWithDefault    = []string{"molfile", "standard_inchi", "canonical_smiles"}
	compoundStructurePrimaryKeyColumns     = []string{"molregno"}
	compoundStructureGeneratedColumns      = []string{}
)

type (
	// CompoundStructureSlice is an alias for a slice of pointers to CompoundStructure.
	// This should almost always be used instead of []CompoundStructure.
	CompoundStructureSlice []*CompoundStructure
	// CompoundStructureHook is the signature for custom CompoundStructure hook methods
	CompoundStructureHook func(context.Context, boil.ContextExecutor, *CompoundStructure) error

	compoundStructureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	compoundStructureType                 = reflect.TypeOf(&CompoundStructure{})
	compoundStructureMapping              = queries.MakeStructMapping(compoundStructureType)
	compoundStructurePrimaryKeyMapping, _ = queries.BindMapping(compoundStructureType, compoundStructureMapping, compoundStructurePrimaryKeyColumns)
	compoundStructureInsertCacheMut       sync.RWMutex
	compoundStructureInsertCache          = make(map[string]insertCache)
	compoundStructureUpdateCacheMut       sync.RWMutex
	compoundStructureUpdateCache          = make(map[string]updateCache)
	compoundStructureUpsertCacheMut       sync.RWMutex
	compoundStructureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var compoundStructureAfterSelectHooks []CompoundStructureHook

var compoundStructureBeforeInsertHooks []CompoundStructureHook
var compoundStructureAfterInsertHooks []CompoundStructureHook

var compoundStructureBeforeUpdateHooks []CompoundStructureHook
var compoundStructureAfterUpdateHooks []CompoundStructureHook

var compoundStructureBeforeDeleteHooks []CompoundStructureHook
var compoundStructureAfterDeleteHooks []CompoundStructureHook

var compoundStructureBeforeUpsertHooks []CompoundStructureHook
var compoundStructureAfterUpsertHooks []CompoundStructureHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CompoundStructure) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CompoundStructure) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CompoundStructure) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CompoundStructure) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CompoundStructure) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CompoundStructure) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CompoundStructure) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CompoundStructure) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CompoundStructure) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range compoundStructureAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompoundStructureHook registers your hook function for all future operations.
func AddCompoundStructureHook(hookPoint boil.HookPoint, compoundStructureHook CompoundStructureHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		compoundStructureAfterSelectHooks = append(compoundStructureAfterSelectHooks, compoundStructureHook)
	case boil.BeforeInsertHook:
		compoundStructureBeforeInsertHooks = append(compoundStructureBeforeInsertHooks, compoundStructureHook)
	case boil.AfterInsertHook:
		compoundStructureAfterInsertHooks = append(compoundStructureAfterInsertHooks, compoundStructureHook)
	case boil.BeforeUpdateHook:
		compoundStructureBeforeUpdateHooks = append(compoundStructureBeforeUpdateHooks, compoundStructureHook)
	case boil.AfterUpdateHook:
		compoundStructureAfterUpdateHooks = append(compoundStructureAfterUpdateHooks, compoundStructureHook)
	case boil.BeforeDeleteHook:
		compoundStructureBeforeDeleteHooks = append(compoundStructureBeforeDeleteHooks, compoundStructureHook)
	case boil.AfterDeleteHook:
		compoundStructureAfterDeleteHooks = append(compoundStructureAfterDeleteHooks, compoundStructureHook)
	case boil.BeforeUpsertHook:
		compoundStructureBeforeUpsertHooks = append(compoundStructureBeforeUpsertHooks, compoundStructureHook)
	case boil.AfterUpsertHook:
		compoundStructureAfterUpsertHooks = append(compoundStructureAfterUpsertHooks, compoundStructureHook)
	}
}

// One returns a single compoundStructure record from the query.
func (q compoundStructureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CompoundStructure, error) {
	o := &CompoundStructure{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for compound_structures")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CompoundStructure records from the query.
func (q compoundStructureQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompoundStructureSlice, error) {
	var o []*CompoundStructure

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CompoundStructure slice")
	}

	if len(compoundStructureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CompoundStructure records in the query.
func (q compoundStructureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count compound_structures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q compoundStructureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if compound_structures exists")
	}

	return count > 0, nil
}

// MolregnoMoleculeDictionary pointed to by the foreign key.
func (o *CompoundStructure) MolregnoMoleculeDictionary(mods ...qm.QueryMod) moleculeDictionaryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"molregno\" = ?", o.Molregno),
	}

	queryMods = append(queryMods, mods...)

	return MoleculeDictionaries(queryMods...)
}

// LoadMolregnoMoleculeDictionary allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (compoundStructureL) LoadMolregnoMoleculeDictionary(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCompoundStructure interface{}, mods queries.Applicator) error {
	var slice []*CompoundStructure
	var object *CompoundStructure

	if singular {
		object = maybeCompoundStructure.(*CompoundStructure)
	} else {
		slice = *maybeCompoundStructure.(*[]*CompoundStructure)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &compoundStructureR{}
		}
		args = append(args, object.Molregno)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &compoundStructureR{}
			}

			for _, a := range args {
				if a == obj.Molregno {
					continue Outer
				}
			}

			args = append(args, obj.Molregno)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`molecule_dictionary`),
		qm.WhereIn(`molecule_dictionary.molregno in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load MoleculeDictionary")
	}

	var resultSlice []*MoleculeDictionary
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice MoleculeDictionary")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for molecule_dictionary")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for molecule_dictionary")
	}

	if len(compoundStructureAfterSelectHooks) != 0 {
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
		object.R.MolregnoMoleculeDictionary = foreign
		if foreign.R == nil {
			foreign.R = &moleculeDictionaryR{}
		}
		foreign.R.MolregnoCompoundStructure = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Molregno == foreign.Molregno {
				local.R.MolregnoMoleculeDictionary = foreign
				if foreign.R == nil {
					foreign.R = &moleculeDictionaryR{}
				}
				foreign.R.MolregnoCompoundStructure = local
				break
			}
		}
	}

	return nil
}

// SetMolregnoMoleculeDictionary of the compoundStructure to the related item.
// Sets o.R.MolregnoMoleculeDictionary to related.
// Adds o to related.R.MolregnoCompoundStructure.
func (o *CompoundStructure) SetMolregnoMoleculeDictionary(ctx context.Context, exec boil.ContextExecutor, insert bool, related *MoleculeDictionary) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"compound_structures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"molregno"}),
		strmangle.WhereClause("\"", "\"", 0, compoundStructurePrimaryKeyColumns),
	)
	values := []interface{}{related.Molregno, o.Molregno}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Molregno = related.Molregno
	if o.R == nil {
		o.R = &compoundStructureR{
			MolregnoMoleculeDictionary: related,
		}
	} else {
		o.R.MolregnoMoleculeDictionary = related
	}

	if related.R == nil {
		related.R = &moleculeDictionaryR{
			MolregnoCompoundStructure: o,
		}
	} else {
		related.R.MolregnoCompoundStructure = o
	}

	return nil
}

// CompoundStructures retrieves all the records using an executor.
func CompoundStructures(mods ...qm.QueryMod) compoundStructureQuery {
	mods = append(mods, qm.From("\"compound_structures\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"compound_structures\".*"})
	}

	return compoundStructureQuery{q}
}

// FindCompoundStructure retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompoundStructure(ctx context.Context, exec boil.ContextExecutor, molregno int64, selectCols ...string) (*CompoundStructure, error) {
	compoundStructureObj := &CompoundStructure{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"compound_structures\" where \"molregno\"=?", sel,
	)

	q := queries.Raw(query, molregno)

	err := q.Bind(ctx, exec, compoundStructureObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from compound_structures")
	}

	if err = compoundStructureObj.doAfterSelectHooks(ctx, exec); err != nil {
		return compoundStructureObj, err
	}

	return compoundStructureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CompoundStructure) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no compound_structures provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(compoundStructureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	compoundStructureInsertCacheMut.RLock()
	cache, cached := compoundStructureInsertCache[key]
	compoundStructureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			compoundStructureAllColumns,
			compoundStructureColumnsWithDefault,
			compoundStructureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(compoundStructureType, compoundStructureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(compoundStructureType, compoundStructureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"compound_structures\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"compound_structures\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into compound_structures")
	}

	if !cached {
		compoundStructureInsertCacheMut.Lock()
		compoundStructureInsertCache[key] = cache
		compoundStructureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CompoundStructure.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CompoundStructure) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	compoundStructureUpdateCacheMut.RLock()
	cache, cached := compoundStructureUpdateCache[key]
	compoundStructureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			compoundStructureAllColumns,
			compoundStructurePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update compound_structures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"compound_structures\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, compoundStructurePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(compoundStructureType, compoundStructureMapping, append(wl, compoundStructurePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update compound_structures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for compound_structures")
	}

	if !cached {
		compoundStructureUpdateCacheMut.Lock()
		compoundStructureUpdateCache[key] = cache
		compoundStructureUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q compoundStructureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for compound_structures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for compound_structures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompoundStructureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), compoundStructurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"compound_structures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, compoundStructurePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in compoundStructure slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all compoundStructure")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CompoundStructure) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no compound_structures provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(compoundStructureColumnsWithDefault, o)

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

	compoundStructureUpsertCacheMut.RLock()
	cache, cached := compoundStructureUpsertCache[key]
	compoundStructureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			compoundStructureAllColumns,
			compoundStructureColumnsWithDefault,
			compoundStructureColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			compoundStructureAllColumns,
			compoundStructurePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert compound_structures, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(compoundStructurePrimaryKeyColumns))
			copy(conflict, compoundStructurePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"compound_structures\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(compoundStructureType, compoundStructureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(compoundStructureType, compoundStructureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert compound_structures")
	}

	if !cached {
		compoundStructureUpsertCacheMut.Lock()
		compoundStructureUpsertCache[key] = cache
		compoundStructureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CompoundStructure record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CompoundStructure) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CompoundStructure provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), compoundStructurePrimaryKeyMapping)
	sql := "DELETE FROM \"compound_structures\" WHERE \"molregno\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from compound_structures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for compound_structures")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q compoundStructureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no compoundStructureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from compound_structures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for compound_structures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompoundStructureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(compoundStructureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), compoundStructurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"compound_structures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, compoundStructurePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from compoundStructure slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for compound_structures")
	}

	if len(compoundStructureAfterDeleteHooks) != 0 {
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
func (o *CompoundStructure) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompoundStructure(ctx, exec, o.Molregno)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompoundStructureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompoundStructureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), compoundStructurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"compound_structures\".* FROM \"compound_structures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, compoundStructurePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompoundStructureSlice")
	}

	*o = slice

	return nil
}

// CompoundStructureExists checks if the CompoundStructure row exists.
func CompoundStructureExists(ctx context.Context, exec boil.ContextExecutor, molregno int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"compound_structures\" where \"molregno\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, molregno)
	}
	row := exec.QueryRowContext(ctx, sql, molregno)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if compound_structures exists")
	}

	return exists, nil
}
