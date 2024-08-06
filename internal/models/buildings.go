// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Building is an object representing the database table.
type Building struct {
	ID             int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	VillageID      int       `boil:"village_id" json:"village_id" toml:"village_id" yaml:"village_id"`
	Name           string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	BuildingType   int       `boil:"building_type" json:"building_type" toml:"building_type" yaml:"building_type"`
	Level          int       `boil:"level" json:"level" toml:"level" yaml:"level"`
	ProductionRate null.Int  `boil:"production_rate" json:"production_rate,omitempty" toml:"production_rate" yaml:"production_rate,omitempty"`
	BuildTime      null.Time `boil:"build_time" json:"build_time,omitempty" toml:"build_time" yaml:"build_time,omitempty"`
	LastUpgrade    null.Time `boil:"last_upgrade" json:"last_upgrade,omitempty" toml:"last_upgrade" yaml:"last_upgrade,omitempty"`

	R *buildingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildingColumns = struct {
	ID             string
	VillageID      string
	Name           string
	BuildingType   string
	Level          string
	ProductionRate string
	BuildTime      string
	LastUpgrade    string
}{
	ID:             "id",
	VillageID:      "village_id",
	Name:           "name",
	BuildingType:   "building_type",
	Level:          "level",
	ProductionRate: "production_rate",
	BuildTime:      "build_time",
	LastUpgrade:    "last_upgrade",
}

var BuildingTableColumns = struct {
	ID             string
	VillageID      string
	Name           string
	BuildingType   string
	Level          string
	ProductionRate string
	BuildTime      string
	LastUpgrade    string
}{
	ID:             "buildings.id",
	VillageID:      "buildings.village_id",
	Name:           "buildings.name",
	BuildingType:   "buildings.building_type",
	Level:          "buildings.level",
	ProductionRate: "buildings.production_rate",
	BuildTime:      "buildings.build_time",
	LastUpgrade:    "buildings.last_upgrade",
}

// Generated where

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

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
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

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
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

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var BuildingWhere = struct {
	ID             whereHelperint
	VillageID      whereHelperint
	Name           whereHelperstring
	BuildingType   whereHelperint
	Level          whereHelperint
	ProductionRate whereHelpernull_Int
	BuildTime      whereHelpernull_Time
	LastUpgrade    whereHelpernull_Time
}{
	ID:             whereHelperint{field: "\"buildings\".\"id\""},
	VillageID:      whereHelperint{field: "\"buildings\".\"village_id\""},
	Name:           whereHelperstring{field: "\"buildings\".\"name\""},
	BuildingType:   whereHelperint{field: "\"buildings\".\"building_type\""},
	Level:          whereHelperint{field: "\"buildings\".\"level\""},
	ProductionRate: whereHelpernull_Int{field: "\"buildings\".\"production_rate\""},
	BuildTime:      whereHelpernull_Time{field: "\"buildings\".\"build_time\""},
	LastUpgrade:    whereHelpernull_Time{field: "\"buildings\".\"last_upgrade\""},
}

// BuildingRels is where relationship names are stored.
var BuildingRels = struct {
	Village string
}{
	Village: "Village",
}

// buildingR is where relationships are stored.
type buildingR struct {
	Village *Village `boil:"Village" json:"Village" toml:"Village" yaml:"Village"`
}

// NewStruct creates a new relationship struct
func (*buildingR) NewStruct() *buildingR {
	return &buildingR{}
}

func (r *buildingR) GetVillage() *Village {
	if r == nil {
		return nil
	}
	return r.Village
}

// buildingL is where Load methods for each relationship are stored.
type buildingL struct{}

var (
	buildingAllColumns            = []string{"id", "village_id", "name", "building_type", "level", "production_rate", "build_time", "last_upgrade"}
	buildingColumnsWithoutDefault = []string{"village_id", "name", "building_type", "level"}
	buildingColumnsWithDefault    = []string{"id", "production_rate", "build_time", "last_upgrade"}
	buildingPrimaryKeyColumns     = []string{"id"}
	buildingGeneratedColumns      = []string{}
)

type (
	// BuildingSlice is an alias for a slice of pointers to Building.
	// This should almost always be used instead of []Building.
	BuildingSlice []*Building
	// BuildingHook is the signature for custom Building hook methods
	BuildingHook func(context.Context, boil.ContextExecutor, *Building) error

	buildingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildingType                 = reflect.TypeOf(&Building{})
	buildingMapping              = queries.MakeStructMapping(buildingType)
	buildingPrimaryKeyMapping, _ = queries.BindMapping(buildingType, buildingMapping, buildingPrimaryKeyColumns)
	buildingInsertCacheMut       sync.RWMutex
	buildingInsertCache          = make(map[string]insertCache)
	buildingUpdateCacheMut       sync.RWMutex
	buildingUpdateCache          = make(map[string]updateCache)
	buildingUpsertCacheMut       sync.RWMutex
	buildingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var buildingAfterSelectMu sync.Mutex
var buildingAfterSelectHooks []BuildingHook

var buildingBeforeInsertMu sync.Mutex
var buildingBeforeInsertHooks []BuildingHook
var buildingAfterInsertMu sync.Mutex
var buildingAfterInsertHooks []BuildingHook

var buildingBeforeUpdateMu sync.Mutex
var buildingBeforeUpdateHooks []BuildingHook
var buildingAfterUpdateMu sync.Mutex
var buildingAfterUpdateHooks []BuildingHook

var buildingBeforeDeleteMu sync.Mutex
var buildingBeforeDeleteHooks []BuildingHook
var buildingAfterDeleteMu sync.Mutex
var buildingAfterDeleteHooks []BuildingHook

var buildingBeforeUpsertMu sync.Mutex
var buildingBeforeUpsertHooks []BuildingHook
var buildingAfterUpsertMu sync.Mutex
var buildingAfterUpsertHooks []BuildingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Building) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Building) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Building) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Building) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Building) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Building) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Building) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Building) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Building) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range buildingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBuildingHook registers your hook function for all future operations.
func AddBuildingHook(hookPoint boil.HookPoint, buildingHook BuildingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		buildingAfterSelectMu.Lock()
		buildingAfterSelectHooks = append(buildingAfterSelectHooks, buildingHook)
		buildingAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		buildingBeforeInsertMu.Lock()
		buildingBeforeInsertHooks = append(buildingBeforeInsertHooks, buildingHook)
		buildingBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		buildingAfterInsertMu.Lock()
		buildingAfterInsertHooks = append(buildingAfterInsertHooks, buildingHook)
		buildingAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		buildingBeforeUpdateMu.Lock()
		buildingBeforeUpdateHooks = append(buildingBeforeUpdateHooks, buildingHook)
		buildingBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		buildingAfterUpdateMu.Lock()
		buildingAfterUpdateHooks = append(buildingAfterUpdateHooks, buildingHook)
		buildingAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		buildingBeforeDeleteMu.Lock()
		buildingBeforeDeleteHooks = append(buildingBeforeDeleteHooks, buildingHook)
		buildingBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		buildingAfterDeleteMu.Lock()
		buildingAfterDeleteHooks = append(buildingAfterDeleteHooks, buildingHook)
		buildingAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		buildingBeforeUpsertMu.Lock()
		buildingBeforeUpsertHooks = append(buildingBeforeUpsertHooks, buildingHook)
		buildingBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		buildingAfterUpsertMu.Lock()
		buildingAfterUpsertHooks = append(buildingAfterUpsertHooks, buildingHook)
		buildingAfterUpsertMu.Unlock()
	}
}

// One returns a single building record from the query.
func (q buildingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Building, error) {
	o := &Building{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for buildings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Building records from the query.
func (q buildingQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildingSlice, error) {
	var o []*Building

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Building slice")
	}

	if len(buildingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Building records in the query.
func (q buildingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count buildings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q buildingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if buildings exists")
	}

	return count > 0, nil
}

// Village pointed to by the foreign key.
func (o *Building) Village(mods ...qm.QueryMod) villageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.VillageID),
	}

	queryMods = append(queryMods, mods...)

	return Villages(queryMods...)
}

// LoadVillage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (buildingL) LoadVillage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuilding interface{}, mods queries.Applicator) error {
	var slice []*Building
	var object *Building

	if singular {
		var ok bool
		object, ok = maybeBuilding.(*Building)
		if !ok {
			object = new(Building)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBuilding)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBuilding))
			}
		}
	} else {
		s, ok := maybeBuilding.(*[]*Building)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBuilding)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBuilding))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &buildingR{}
		}
		args[object.VillageID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildingR{}
			}

			args[obj.VillageID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`villages`),
		qm.WhereIn(`villages.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Village")
	}

	var resultSlice []*Village
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Village")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for villages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for villages")
	}

	if len(villageAfterSelectHooks) != 0 {
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
		object.R.Village = foreign
		if foreign.R == nil {
			foreign.R = &villageR{}
		}
		foreign.R.Buildings = append(foreign.R.Buildings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.VillageID == foreign.ID {
				local.R.Village = foreign
				if foreign.R == nil {
					foreign.R = &villageR{}
				}
				foreign.R.Buildings = append(foreign.R.Buildings, local)
				break
			}
		}
	}

	return nil
}

// SetVillage of the building to the related item.
// Sets o.R.Village to related.
// Adds o to related.R.Buildings.
func (o *Building) SetVillage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Village) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"buildings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"village_id"}),
		strmangle.WhereClause("\"", "\"", 2, buildingPrimaryKeyColumns),
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

	o.VillageID = related.ID
	if o.R == nil {
		o.R = &buildingR{
			Village: related,
		}
	} else {
		o.R.Village = related
	}

	if related.R == nil {
		related.R = &villageR{
			Buildings: BuildingSlice{o},
		}
	} else {
		related.R.Buildings = append(related.R.Buildings, o)
	}

	return nil
}

// Buildings retrieves all the records using an executor.
func Buildings(mods ...qm.QueryMod) buildingQuery {
	mods = append(mods, qm.From("\"buildings\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"buildings\".*"})
	}

	return buildingQuery{q}
}

// FindBuilding retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuilding(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Building, error) {
	buildingObj := &Building{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"buildings\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, buildingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from buildings")
	}

	if err = buildingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return buildingObj, err
	}

	return buildingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Building) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildings provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(buildingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildingInsertCacheMut.RLock()
	cache, cached := buildingInsertCache[key]
	buildingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildingAllColumns,
			buildingColumnsWithDefault,
			buildingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildingType, buildingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildingType, buildingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"buildings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"buildings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into buildings")
	}

	if !cached {
		buildingInsertCacheMut.Lock()
		buildingInsertCache[key] = cache
		buildingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Building.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Building) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	buildingUpdateCacheMut.RLock()
	cache, cached := buildingUpdateCache[key]
	buildingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildingAllColumns,
			buildingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update buildings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"buildings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, buildingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildingType, buildingMapping, append(wl, buildingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update buildings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for buildings")
	}

	if !cached {
		buildingUpdateCacheMut.Lock()
		buildingUpdateCache[key] = cache
		buildingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q buildingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for buildings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for buildings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BuildingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"buildings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, buildingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in building slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all building")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Building) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no buildings provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(buildingColumnsWithDefault, o)

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

	buildingUpsertCacheMut.RLock()
	cache, cached := buildingUpsertCache[key]
	buildingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			buildingAllColumns,
			buildingColumnsWithDefault,
			buildingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			buildingAllColumns,
			buildingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert buildings, could not build update column list")
		}

		ret := strmangle.SetComplement(buildingAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(buildingPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert buildings, could not build conflict column list")
			}

			conflict = make([]string, len(buildingPrimaryKeyColumns))
			copy(conflict, buildingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"buildings\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(buildingType, buildingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildingType, buildingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert buildings")
	}

	if !cached {
		buildingUpsertCacheMut.Lock()
		buildingUpsertCache[key] = cache
		buildingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Building record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Building) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Building provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildingPrimaryKeyMapping)
	sql := "DELETE FROM \"buildings\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from buildings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for buildings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q buildingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BuildingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(buildingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"buildings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from building slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildings")
	}

	if len(buildingAfterDeleteHooks) != 0 {
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
func (o *Building) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuilding(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BuildingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"buildings\".* FROM \"buildings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildingSlice")
	}

	*o = slice

	return nil
}

// BuildingExists checks if the Building row exists.
func BuildingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"buildings\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if buildings exists")
	}

	return exists, nil
}

// Exists checks if the Building row exists.
func (o *Building) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BuildingExists(ctx, exec, o.ID)
}
