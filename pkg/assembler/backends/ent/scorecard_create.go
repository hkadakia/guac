// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certifyscorecard"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/scorecard"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// ScorecardCreate is the builder for creating a Scorecard entity.
type ScorecardCreate struct {
	config
	mutation *ScorecardMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetChecks sets the "checks" field.
func (sc *ScorecardCreate) SetChecks(mc []*model.ScorecardCheck) *ScorecardCreate {
	sc.mutation.SetChecks(mc)
	return sc
}

// SetAggregateScore sets the "aggregate_score" field.
func (sc *ScorecardCreate) SetAggregateScore(f float64) *ScorecardCreate {
	sc.mutation.SetAggregateScore(f)
	return sc
}

// SetNillableAggregateScore sets the "aggregate_score" field if the given value is not nil.
func (sc *ScorecardCreate) SetNillableAggregateScore(f *float64) *ScorecardCreate {
	if f != nil {
		sc.SetAggregateScore(*f)
	}
	return sc
}

// SetTimeScanned sets the "time_scanned" field.
func (sc *ScorecardCreate) SetTimeScanned(t time.Time) *ScorecardCreate {
	sc.mutation.SetTimeScanned(t)
	return sc
}

// SetNillableTimeScanned sets the "time_scanned" field if the given value is not nil.
func (sc *ScorecardCreate) SetNillableTimeScanned(t *time.Time) *ScorecardCreate {
	if t != nil {
		sc.SetTimeScanned(*t)
	}
	return sc
}

// SetScorecardVersion sets the "scorecard_version" field.
func (sc *ScorecardCreate) SetScorecardVersion(s string) *ScorecardCreate {
	sc.mutation.SetScorecardVersion(s)
	return sc
}

// SetScorecardCommit sets the "scorecard_commit" field.
func (sc *ScorecardCreate) SetScorecardCommit(s string) *ScorecardCreate {
	sc.mutation.SetScorecardCommit(s)
	return sc
}

// SetOrigin sets the "origin" field.
func (sc *ScorecardCreate) SetOrigin(s string) *ScorecardCreate {
	sc.mutation.SetOrigin(s)
	return sc
}

// SetCollector sets the "collector" field.
func (sc *ScorecardCreate) SetCollector(s string) *ScorecardCreate {
	sc.mutation.SetCollector(s)
	return sc
}

// AddCertificationIDs adds the "certifications" edge to the CertifyScorecard entity by IDs.
func (sc *ScorecardCreate) AddCertificationIDs(ids ...int) *ScorecardCreate {
	sc.mutation.AddCertificationIDs(ids...)
	return sc
}

// AddCertifications adds the "certifications" edges to the CertifyScorecard entity.
func (sc *ScorecardCreate) AddCertifications(c ...*CertifyScorecard) *ScorecardCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCertificationIDs(ids...)
}

// Mutation returns the ScorecardMutation object of the builder.
func (sc *ScorecardCreate) Mutation() *ScorecardMutation {
	return sc.mutation
}

// Save creates the Scorecard in the database.
func (sc *ScorecardCreate) Save(ctx context.Context) (*Scorecard, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScorecardCreate) SaveX(ctx context.Context) *Scorecard {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScorecardCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScorecardCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScorecardCreate) defaults() {
	if _, ok := sc.mutation.AggregateScore(); !ok {
		v := scorecard.DefaultAggregateScore
		sc.mutation.SetAggregateScore(v)
	}
	if _, ok := sc.mutation.TimeScanned(); !ok {
		v := scorecard.DefaultTimeScanned()
		sc.mutation.SetTimeScanned(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScorecardCreate) check() error {
	if _, ok := sc.mutation.Checks(); !ok {
		return &ValidationError{Name: "checks", err: errors.New(`ent: missing required field "Scorecard.checks"`)}
	}
	if _, ok := sc.mutation.AggregateScore(); !ok {
		return &ValidationError{Name: "aggregate_score", err: errors.New(`ent: missing required field "Scorecard.aggregate_score"`)}
	}
	if _, ok := sc.mutation.TimeScanned(); !ok {
		return &ValidationError{Name: "time_scanned", err: errors.New(`ent: missing required field "Scorecard.time_scanned"`)}
	}
	if _, ok := sc.mutation.ScorecardVersion(); !ok {
		return &ValidationError{Name: "scorecard_version", err: errors.New(`ent: missing required field "Scorecard.scorecard_version"`)}
	}
	if _, ok := sc.mutation.ScorecardCommit(); !ok {
		return &ValidationError{Name: "scorecard_commit", err: errors.New(`ent: missing required field "Scorecard.scorecard_commit"`)}
	}
	if _, ok := sc.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "Scorecard.origin"`)}
	}
	if _, ok := sc.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "Scorecard.collector"`)}
	}
	return nil
}

func (sc *ScorecardCreate) sqlSave(ctx context.Context) (*Scorecard, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScorecardCreate) createSpec() (*Scorecard, *sqlgraph.CreateSpec) {
	var (
		_node = &Scorecard{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(scorecard.Table, sqlgraph.NewFieldSpec(scorecard.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Checks(); ok {
		_spec.SetField(scorecard.FieldChecks, field.TypeJSON, value)
		_node.Checks = value
	}
	if value, ok := sc.mutation.AggregateScore(); ok {
		_spec.SetField(scorecard.FieldAggregateScore, field.TypeFloat64, value)
		_node.AggregateScore = value
	}
	if value, ok := sc.mutation.TimeScanned(); ok {
		_spec.SetField(scorecard.FieldTimeScanned, field.TypeTime, value)
		_node.TimeScanned = value
	}
	if value, ok := sc.mutation.ScorecardVersion(); ok {
		_spec.SetField(scorecard.FieldScorecardVersion, field.TypeString, value)
		_node.ScorecardVersion = value
	}
	if value, ok := sc.mutation.ScorecardCommit(); ok {
		_spec.SetField(scorecard.FieldScorecardCommit, field.TypeString, value)
		_node.ScorecardCommit = value
	}
	if value, ok := sc.mutation.Origin(); ok {
		_spec.SetField(scorecard.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := sc.mutation.Collector(); ok {
		_spec.SetField(scorecard.FieldCollector, field.TypeString, value)
		_node.Collector = value
	}
	if nodes := sc.mutation.CertificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scorecard.CertificationsTable,
			Columns: []string{scorecard.CertificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(certifyscorecard.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Scorecard.Create().
//		SetChecks(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ScorecardUpsert) {
//			SetChecks(v+v).
//		}).
//		Exec(ctx)
func (sc *ScorecardCreate) OnConflict(opts ...sql.ConflictOption) *ScorecardUpsertOne {
	sc.conflict = opts
	return &ScorecardUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Scorecard.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *ScorecardCreate) OnConflictColumns(columns ...string) *ScorecardUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ScorecardUpsertOne{
		create: sc,
	}
}

type (
	// ScorecardUpsertOne is the builder for "upsert"-ing
	//  one Scorecard node.
	ScorecardUpsertOne struct {
		create *ScorecardCreate
	}

	// ScorecardUpsert is the "OnConflict" setter.
	ScorecardUpsert struct {
		*sql.UpdateSet
	}
)

// SetChecks sets the "checks" field.
func (u *ScorecardUpsert) SetChecks(v []*model.ScorecardCheck) *ScorecardUpsert {
	u.Set(scorecard.FieldChecks, v)
	return u
}

// UpdateChecks sets the "checks" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateChecks() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldChecks)
	return u
}

// SetAggregateScore sets the "aggregate_score" field.
func (u *ScorecardUpsert) SetAggregateScore(v float64) *ScorecardUpsert {
	u.Set(scorecard.FieldAggregateScore, v)
	return u
}

// UpdateAggregateScore sets the "aggregate_score" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateAggregateScore() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldAggregateScore)
	return u
}

// AddAggregateScore adds v to the "aggregate_score" field.
func (u *ScorecardUpsert) AddAggregateScore(v float64) *ScorecardUpsert {
	u.Add(scorecard.FieldAggregateScore, v)
	return u
}

// SetTimeScanned sets the "time_scanned" field.
func (u *ScorecardUpsert) SetTimeScanned(v time.Time) *ScorecardUpsert {
	u.Set(scorecard.FieldTimeScanned, v)
	return u
}

// UpdateTimeScanned sets the "time_scanned" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateTimeScanned() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldTimeScanned)
	return u
}

// SetScorecardVersion sets the "scorecard_version" field.
func (u *ScorecardUpsert) SetScorecardVersion(v string) *ScorecardUpsert {
	u.Set(scorecard.FieldScorecardVersion, v)
	return u
}

// UpdateScorecardVersion sets the "scorecard_version" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateScorecardVersion() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldScorecardVersion)
	return u
}

// SetScorecardCommit sets the "scorecard_commit" field.
func (u *ScorecardUpsert) SetScorecardCommit(v string) *ScorecardUpsert {
	u.Set(scorecard.FieldScorecardCommit, v)
	return u
}

// UpdateScorecardCommit sets the "scorecard_commit" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateScorecardCommit() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldScorecardCommit)
	return u
}

// SetOrigin sets the "origin" field.
func (u *ScorecardUpsert) SetOrigin(v string) *ScorecardUpsert {
	u.Set(scorecard.FieldOrigin, v)
	return u
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateOrigin() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldOrigin)
	return u
}

// SetCollector sets the "collector" field.
func (u *ScorecardUpsert) SetCollector(v string) *ScorecardUpsert {
	u.Set(scorecard.FieldCollector, v)
	return u
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *ScorecardUpsert) UpdateCollector() *ScorecardUpsert {
	u.SetExcluded(scorecard.FieldCollector)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Scorecard.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ScorecardUpsertOne) UpdateNewValues() *ScorecardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Scorecard.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ScorecardUpsertOne) Ignore() *ScorecardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ScorecardUpsertOne) DoNothing() *ScorecardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ScorecardCreate.OnConflict
// documentation for more info.
func (u *ScorecardUpsertOne) Update(set func(*ScorecardUpsert)) *ScorecardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ScorecardUpsert{UpdateSet: update})
	}))
	return u
}

// SetChecks sets the "checks" field.
func (u *ScorecardUpsertOne) SetChecks(v []*model.ScorecardCheck) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetChecks(v)
	})
}

// UpdateChecks sets the "checks" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateChecks() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateChecks()
	})
}

// SetAggregateScore sets the "aggregate_score" field.
func (u *ScorecardUpsertOne) SetAggregateScore(v float64) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetAggregateScore(v)
	})
}

// AddAggregateScore adds v to the "aggregate_score" field.
func (u *ScorecardUpsertOne) AddAggregateScore(v float64) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.AddAggregateScore(v)
	})
}

// UpdateAggregateScore sets the "aggregate_score" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateAggregateScore() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateAggregateScore()
	})
}

// SetTimeScanned sets the "time_scanned" field.
func (u *ScorecardUpsertOne) SetTimeScanned(v time.Time) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetTimeScanned(v)
	})
}

// UpdateTimeScanned sets the "time_scanned" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateTimeScanned() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateTimeScanned()
	})
}

// SetScorecardVersion sets the "scorecard_version" field.
func (u *ScorecardUpsertOne) SetScorecardVersion(v string) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetScorecardVersion(v)
	})
}

// UpdateScorecardVersion sets the "scorecard_version" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateScorecardVersion() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateScorecardVersion()
	})
}

// SetScorecardCommit sets the "scorecard_commit" field.
func (u *ScorecardUpsertOne) SetScorecardCommit(v string) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetScorecardCommit(v)
	})
}

// UpdateScorecardCommit sets the "scorecard_commit" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateScorecardCommit() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateScorecardCommit()
	})
}

// SetOrigin sets the "origin" field.
func (u *ScorecardUpsertOne) SetOrigin(v string) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateOrigin() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *ScorecardUpsertOne) SetCollector(v string) *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *ScorecardUpsertOne) UpdateCollector() *ScorecardUpsertOne {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateCollector()
	})
}

// Exec executes the query.
func (u *ScorecardUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ScorecardCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ScorecardUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ScorecardUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ScorecardUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ScorecardCreateBulk is the builder for creating many Scorecard entities in bulk.
type ScorecardCreateBulk struct {
	config
	builders []*ScorecardCreate
	conflict []sql.ConflictOption
}

// Save creates the Scorecard entities in the database.
func (scb *ScorecardCreateBulk) Save(ctx context.Context) ([]*Scorecard, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Scorecard, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScorecardMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScorecardCreateBulk) SaveX(ctx context.Context) []*Scorecard {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScorecardCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScorecardCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Scorecard.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ScorecardUpsert) {
//			SetChecks(v+v).
//		}).
//		Exec(ctx)
func (scb *ScorecardCreateBulk) OnConflict(opts ...sql.ConflictOption) *ScorecardUpsertBulk {
	scb.conflict = opts
	return &ScorecardUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Scorecard.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *ScorecardCreateBulk) OnConflictColumns(columns ...string) *ScorecardUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ScorecardUpsertBulk{
		create: scb,
	}
}

// ScorecardUpsertBulk is the builder for "upsert"-ing
// a bulk of Scorecard nodes.
type ScorecardUpsertBulk struct {
	create *ScorecardCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Scorecard.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ScorecardUpsertBulk) UpdateNewValues() *ScorecardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Scorecard.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ScorecardUpsertBulk) Ignore() *ScorecardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ScorecardUpsertBulk) DoNothing() *ScorecardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ScorecardCreateBulk.OnConflict
// documentation for more info.
func (u *ScorecardUpsertBulk) Update(set func(*ScorecardUpsert)) *ScorecardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ScorecardUpsert{UpdateSet: update})
	}))
	return u
}

// SetChecks sets the "checks" field.
func (u *ScorecardUpsertBulk) SetChecks(v []*model.ScorecardCheck) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetChecks(v)
	})
}

// UpdateChecks sets the "checks" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateChecks() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateChecks()
	})
}

// SetAggregateScore sets the "aggregate_score" field.
func (u *ScorecardUpsertBulk) SetAggregateScore(v float64) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetAggregateScore(v)
	})
}

// AddAggregateScore adds v to the "aggregate_score" field.
func (u *ScorecardUpsertBulk) AddAggregateScore(v float64) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.AddAggregateScore(v)
	})
}

// UpdateAggregateScore sets the "aggregate_score" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateAggregateScore() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateAggregateScore()
	})
}

// SetTimeScanned sets the "time_scanned" field.
func (u *ScorecardUpsertBulk) SetTimeScanned(v time.Time) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetTimeScanned(v)
	})
}

// UpdateTimeScanned sets the "time_scanned" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateTimeScanned() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateTimeScanned()
	})
}

// SetScorecardVersion sets the "scorecard_version" field.
func (u *ScorecardUpsertBulk) SetScorecardVersion(v string) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetScorecardVersion(v)
	})
}

// UpdateScorecardVersion sets the "scorecard_version" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateScorecardVersion() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateScorecardVersion()
	})
}

// SetScorecardCommit sets the "scorecard_commit" field.
func (u *ScorecardUpsertBulk) SetScorecardCommit(v string) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetScorecardCommit(v)
	})
}

// UpdateScorecardCommit sets the "scorecard_commit" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateScorecardCommit() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateScorecardCommit()
	})
}

// SetOrigin sets the "origin" field.
func (u *ScorecardUpsertBulk) SetOrigin(v string) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateOrigin() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *ScorecardUpsertBulk) SetCollector(v string) *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *ScorecardUpsertBulk) UpdateCollector() *ScorecardUpsertBulk {
	return u.Update(func(s *ScorecardUpsert) {
		s.UpdateCollector()
	})
}

// Exec executes the query.
func (u *ScorecardUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ScorecardCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ScorecardCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ScorecardUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
