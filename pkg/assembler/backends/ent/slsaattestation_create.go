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
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/builder"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/slsaattestation"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// SLSAAttestationCreate is the builder for creating a SLSAAttestation entity.
type SLSAAttestationCreate struct {
	config
	mutation *SLSAAttestationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBuildType sets the "build_type" field.
func (sac *SLSAAttestationCreate) SetBuildType(s string) *SLSAAttestationCreate {
	sac.mutation.SetBuildType(s)
	return sac
}

// SetBuiltByID sets the "built_by_id" field.
func (sac *SLSAAttestationCreate) SetBuiltByID(i int) *SLSAAttestationCreate {
	sac.mutation.SetBuiltByID(i)
	return sac
}

// SetSubjectID sets the "subject_id" field.
func (sac *SLSAAttestationCreate) SetSubjectID(i int) *SLSAAttestationCreate {
	sac.mutation.SetSubjectID(i)
	return sac
}

// SetSlsaPredicate sets the "slsa_predicate" field.
func (sac *SLSAAttestationCreate) SetSlsaPredicate(mp []*model.SLSAPredicate) *SLSAAttestationCreate {
	sac.mutation.SetSlsaPredicate(mp)
	return sac
}

// SetSlsaVersion sets the "slsa_version" field.
func (sac *SLSAAttestationCreate) SetSlsaVersion(s string) *SLSAAttestationCreate {
	sac.mutation.SetSlsaVersion(s)
	return sac
}

// SetStartedOn sets the "started_on" field.
func (sac *SLSAAttestationCreate) SetStartedOn(t time.Time) *SLSAAttestationCreate {
	sac.mutation.SetStartedOn(t)
	return sac
}

// SetNillableStartedOn sets the "started_on" field if the given value is not nil.
func (sac *SLSAAttestationCreate) SetNillableStartedOn(t *time.Time) *SLSAAttestationCreate {
	if t != nil {
		sac.SetStartedOn(*t)
	}
	return sac
}

// SetFinishedOn sets the "finished_on" field.
func (sac *SLSAAttestationCreate) SetFinishedOn(t time.Time) *SLSAAttestationCreate {
	sac.mutation.SetFinishedOn(t)
	return sac
}

// SetNillableFinishedOn sets the "finished_on" field if the given value is not nil.
func (sac *SLSAAttestationCreate) SetNillableFinishedOn(t *time.Time) *SLSAAttestationCreate {
	if t != nil {
		sac.SetFinishedOn(*t)
	}
	return sac
}

// SetOrigin sets the "origin" field.
func (sac *SLSAAttestationCreate) SetOrigin(s string) *SLSAAttestationCreate {
	sac.mutation.SetOrigin(s)
	return sac
}

// SetCollector sets the "collector" field.
func (sac *SLSAAttestationCreate) SetCollector(s string) *SLSAAttestationCreate {
	sac.mutation.SetCollector(s)
	return sac
}

// SetBuiltFromHash sets the "built_from_hash" field.
func (sac *SLSAAttestationCreate) SetBuiltFromHash(s string) *SLSAAttestationCreate {
	sac.mutation.SetBuiltFromHash(s)
	return sac
}

// AddBuiltFromIDs adds the "built_from" edge to the Artifact entity by IDs.
func (sac *SLSAAttestationCreate) AddBuiltFromIDs(ids ...int) *SLSAAttestationCreate {
	sac.mutation.AddBuiltFromIDs(ids...)
	return sac
}

// AddBuiltFrom adds the "built_from" edges to the Artifact entity.
func (sac *SLSAAttestationCreate) AddBuiltFrom(a ...*Artifact) *SLSAAttestationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sac.AddBuiltFromIDs(ids...)
}

// SetBuiltBy sets the "built_by" edge to the Builder entity.
func (sac *SLSAAttestationCreate) SetBuiltBy(b *Builder) *SLSAAttestationCreate {
	return sac.SetBuiltByID(b.ID)
}

// SetSubject sets the "subject" edge to the Artifact entity.
func (sac *SLSAAttestationCreate) SetSubject(a *Artifact) *SLSAAttestationCreate {
	return sac.SetSubjectID(a.ID)
}

// Mutation returns the SLSAAttestationMutation object of the builder.
func (sac *SLSAAttestationCreate) Mutation() *SLSAAttestationMutation {
	return sac.mutation
}

// Save creates the SLSAAttestation in the database.
func (sac *SLSAAttestationCreate) Save(ctx context.Context) (*SLSAAttestation, error) {
	return withHooks(ctx, sac.sqlSave, sac.mutation, sac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sac *SLSAAttestationCreate) SaveX(ctx context.Context) *SLSAAttestation {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sac *SLSAAttestationCreate) Exec(ctx context.Context) error {
	_, err := sac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sac *SLSAAttestationCreate) ExecX(ctx context.Context) {
	if err := sac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *SLSAAttestationCreate) check() error {
	if _, ok := sac.mutation.BuildType(); !ok {
		return &ValidationError{Name: "build_type", err: errors.New(`ent: missing required field "SLSAAttestation.build_type"`)}
	}
	if _, ok := sac.mutation.BuiltByID(); !ok {
		return &ValidationError{Name: "built_by_id", err: errors.New(`ent: missing required field "SLSAAttestation.built_by_id"`)}
	}
	if _, ok := sac.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject_id", err: errors.New(`ent: missing required field "SLSAAttestation.subject_id"`)}
	}
	if _, ok := sac.mutation.SlsaVersion(); !ok {
		return &ValidationError{Name: "slsa_version", err: errors.New(`ent: missing required field "SLSAAttestation.slsa_version"`)}
	}
	if _, ok := sac.mutation.Origin(); !ok {
		return &ValidationError{Name: "origin", err: errors.New(`ent: missing required field "SLSAAttestation.origin"`)}
	}
	if _, ok := sac.mutation.Collector(); !ok {
		return &ValidationError{Name: "collector", err: errors.New(`ent: missing required field "SLSAAttestation.collector"`)}
	}
	if _, ok := sac.mutation.BuiltFromHash(); !ok {
		return &ValidationError{Name: "built_from_hash", err: errors.New(`ent: missing required field "SLSAAttestation.built_from_hash"`)}
	}
	if _, ok := sac.mutation.BuiltByID(); !ok {
		return &ValidationError{Name: "built_by", err: errors.New(`ent: missing required edge "SLSAAttestation.built_by"`)}
	}
	if _, ok := sac.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required edge "SLSAAttestation.subject"`)}
	}
	return nil
}

func (sac *SLSAAttestationCreate) sqlSave(ctx context.Context) (*SLSAAttestation, error) {
	if err := sac.check(); err != nil {
		return nil, err
	}
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sac.mutation.id = &_node.ID
	sac.mutation.done = true
	return _node, nil
}

func (sac *SLSAAttestationCreate) createSpec() (*SLSAAttestation, *sqlgraph.CreateSpec) {
	var (
		_node = &SLSAAttestation{config: sac.config}
		_spec = sqlgraph.NewCreateSpec(slsaattestation.Table, sqlgraph.NewFieldSpec(slsaattestation.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sac.conflict
	if value, ok := sac.mutation.BuildType(); ok {
		_spec.SetField(slsaattestation.FieldBuildType, field.TypeString, value)
		_node.BuildType = value
	}
	if value, ok := sac.mutation.SlsaPredicate(); ok {
		_spec.SetField(slsaattestation.FieldSlsaPredicate, field.TypeJSON, value)
		_node.SlsaPredicate = value
	}
	if value, ok := sac.mutation.SlsaVersion(); ok {
		_spec.SetField(slsaattestation.FieldSlsaVersion, field.TypeString, value)
		_node.SlsaVersion = value
	}
	if value, ok := sac.mutation.StartedOn(); ok {
		_spec.SetField(slsaattestation.FieldStartedOn, field.TypeTime, value)
		_node.StartedOn = &value
	}
	if value, ok := sac.mutation.FinishedOn(); ok {
		_spec.SetField(slsaattestation.FieldFinishedOn, field.TypeTime, value)
		_node.FinishedOn = &value
	}
	if value, ok := sac.mutation.Origin(); ok {
		_spec.SetField(slsaattestation.FieldOrigin, field.TypeString, value)
		_node.Origin = value
	}
	if value, ok := sac.mutation.Collector(); ok {
		_spec.SetField(slsaattestation.FieldCollector, field.TypeString, value)
		_node.Collector = value
	}
	if value, ok := sac.mutation.BuiltFromHash(); ok {
		_spec.SetField(slsaattestation.FieldBuiltFromHash, field.TypeString, value)
		_node.BuiltFromHash = value
	}
	if nodes := sac.mutation.BuiltFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   slsaattestation.BuiltFromTable,
			Columns: slsaattestation.BuiltFromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.BuiltByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   slsaattestation.BuiltByTable,
			Columns: []string{slsaattestation.BuiltByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(builder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BuiltByID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   slsaattestation.SubjectTable,
			Columns: []string{slsaattestation.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubjectID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SLSAAttestation.Create().
//		SetBuildType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SLSAAttestationUpsert) {
//			SetBuildType(v+v).
//		}).
//		Exec(ctx)
func (sac *SLSAAttestationCreate) OnConflict(opts ...sql.ConflictOption) *SLSAAttestationUpsertOne {
	sac.conflict = opts
	return &SLSAAttestationUpsertOne{
		create: sac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SLSAAttestation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sac *SLSAAttestationCreate) OnConflictColumns(columns ...string) *SLSAAttestationUpsertOne {
	sac.conflict = append(sac.conflict, sql.ConflictColumns(columns...))
	return &SLSAAttestationUpsertOne{
		create: sac,
	}
}

type (
	// SLSAAttestationUpsertOne is the builder for "upsert"-ing
	//  one SLSAAttestation node.
	SLSAAttestationUpsertOne struct {
		create *SLSAAttestationCreate
	}

	// SLSAAttestationUpsert is the "OnConflict" setter.
	SLSAAttestationUpsert struct {
		*sql.UpdateSet
	}
)

// SetBuildType sets the "build_type" field.
func (u *SLSAAttestationUpsert) SetBuildType(v string) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldBuildType, v)
	return u
}

// UpdateBuildType sets the "build_type" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateBuildType() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldBuildType)
	return u
}

// SetBuiltByID sets the "built_by_id" field.
func (u *SLSAAttestationUpsert) SetBuiltByID(v int) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldBuiltByID, v)
	return u
}

// UpdateBuiltByID sets the "built_by_id" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateBuiltByID() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldBuiltByID)
	return u
}

// SetSubjectID sets the "subject_id" field.
func (u *SLSAAttestationUpsert) SetSubjectID(v int) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldSubjectID, v)
	return u
}

// UpdateSubjectID sets the "subject_id" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateSubjectID() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldSubjectID)
	return u
}

// SetSlsaPredicate sets the "slsa_predicate" field.
func (u *SLSAAttestationUpsert) SetSlsaPredicate(v []*model.SLSAPredicate) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldSlsaPredicate, v)
	return u
}

// UpdateSlsaPredicate sets the "slsa_predicate" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateSlsaPredicate() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldSlsaPredicate)
	return u
}

// ClearSlsaPredicate clears the value of the "slsa_predicate" field.
func (u *SLSAAttestationUpsert) ClearSlsaPredicate() *SLSAAttestationUpsert {
	u.SetNull(slsaattestation.FieldSlsaPredicate)
	return u
}

// SetSlsaVersion sets the "slsa_version" field.
func (u *SLSAAttestationUpsert) SetSlsaVersion(v string) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldSlsaVersion, v)
	return u
}

// UpdateSlsaVersion sets the "slsa_version" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateSlsaVersion() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldSlsaVersion)
	return u
}

// SetStartedOn sets the "started_on" field.
func (u *SLSAAttestationUpsert) SetStartedOn(v time.Time) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldStartedOn, v)
	return u
}

// UpdateStartedOn sets the "started_on" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateStartedOn() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldStartedOn)
	return u
}

// ClearStartedOn clears the value of the "started_on" field.
func (u *SLSAAttestationUpsert) ClearStartedOn() *SLSAAttestationUpsert {
	u.SetNull(slsaattestation.FieldStartedOn)
	return u
}

// SetFinishedOn sets the "finished_on" field.
func (u *SLSAAttestationUpsert) SetFinishedOn(v time.Time) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldFinishedOn, v)
	return u
}

// UpdateFinishedOn sets the "finished_on" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateFinishedOn() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldFinishedOn)
	return u
}

// ClearFinishedOn clears the value of the "finished_on" field.
func (u *SLSAAttestationUpsert) ClearFinishedOn() *SLSAAttestationUpsert {
	u.SetNull(slsaattestation.FieldFinishedOn)
	return u
}

// SetOrigin sets the "origin" field.
func (u *SLSAAttestationUpsert) SetOrigin(v string) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldOrigin, v)
	return u
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateOrigin() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldOrigin)
	return u
}

// SetCollector sets the "collector" field.
func (u *SLSAAttestationUpsert) SetCollector(v string) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldCollector, v)
	return u
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateCollector() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldCollector)
	return u
}

// SetBuiltFromHash sets the "built_from_hash" field.
func (u *SLSAAttestationUpsert) SetBuiltFromHash(v string) *SLSAAttestationUpsert {
	u.Set(slsaattestation.FieldBuiltFromHash, v)
	return u
}

// UpdateBuiltFromHash sets the "built_from_hash" field to the value that was provided on create.
func (u *SLSAAttestationUpsert) UpdateBuiltFromHash() *SLSAAttestationUpsert {
	u.SetExcluded(slsaattestation.FieldBuiltFromHash)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.SLSAAttestation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SLSAAttestationUpsertOne) UpdateNewValues() *SLSAAttestationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SLSAAttestation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SLSAAttestationUpsertOne) Ignore() *SLSAAttestationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SLSAAttestationUpsertOne) DoNothing() *SLSAAttestationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SLSAAttestationCreate.OnConflict
// documentation for more info.
func (u *SLSAAttestationUpsertOne) Update(set func(*SLSAAttestationUpsert)) *SLSAAttestationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SLSAAttestationUpsert{UpdateSet: update})
	}))
	return u
}

// SetBuildType sets the "build_type" field.
func (u *SLSAAttestationUpsertOne) SetBuildType(v string) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetBuildType(v)
	})
}

// UpdateBuildType sets the "build_type" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateBuildType() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateBuildType()
	})
}

// SetBuiltByID sets the "built_by_id" field.
func (u *SLSAAttestationUpsertOne) SetBuiltByID(v int) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetBuiltByID(v)
	})
}

// UpdateBuiltByID sets the "built_by_id" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateBuiltByID() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateBuiltByID()
	})
}

// SetSubjectID sets the "subject_id" field.
func (u *SLSAAttestationUpsertOne) SetSubjectID(v int) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetSubjectID(v)
	})
}

// UpdateSubjectID sets the "subject_id" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateSubjectID() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateSubjectID()
	})
}

// SetSlsaPredicate sets the "slsa_predicate" field.
func (u *SLSAAttestationUpsertOne) SetSlsaPredicate(v []*model.SLSAPredicate) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetSlsaPredicate(v)
	})
}

// UpdateSlsaPredicate sets the "slsa_predicate" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateSlsaPredicate() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateSlsaPredicate()
	})
}

// ClearSlsaPredicate clears the value of the "slsa_predicate" field.
func (u *SLSAAttestationUpsertOne) ClearSlsaPredicate() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.ClearSlsaPredicate()
	})
}

// SetSlsaVersion sets the "slsa_version" field.
func (u *SLSAAttestationUpsertOne) SetSlsaVersion(v string) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetSlsaVersion(v)
	})
}

// UpdateSlsaVersion sets the "slsa_version" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateSlsaVersion() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateSlsaVersion()
	})
}

// SetStartedOn sets the "started_on" field.
func (u *SLSAAttestationUpsertOne) SetStartedOn(v time.Time) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetStartedOn(v)
	})
}

// UpdateStartedOn sets the "started_on" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateStartedOn() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateStartedOn()
	})
}

// ClearStartedOn clears the value of the "started_on" field.
func (u *SLSAAttestationUpsertOne) ClearStartedOn() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.ClearStartedOn()
	})
}

// SetFinishedOn sets the "finished_on" field.
func (u *SLSAAttestationUpsertOne) SetFinishedOn(v time.Time) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetFinishedOn(v)
	})
}

// UpdateFinishedOn sets the "finished_on" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateFinishedOn() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateFinishedOn()
	})
}

// ClearFinishedOn clears the value of the "finished_on" field.
func (u *SLSAAttestationUpsertOne) ClearFinishedOn() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.ClearFinishedOn()
	})
}

// SetOrigin sets the "origin" field.
func (u *SLSAAttestationUpsertOne) SetOrigin(v string) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateOrigin() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *SLSAAttestationUpsertOne) SetCollector(v string) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateCollector() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateCollector()
	})
}

// SetBuiltFromHash sets the "built_from_hash" field.
func (u *SLSAAttestationUpsertOne) SetBuiltFromHash(v string) *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetBuiltFromHash(v)
	})
}

// UpdateBuiltFromHash sets the "built_from_hash" field to the value that was provided on create.
func (u *SLSAAttestationUpsertOne) UpdateBuiltFromHash() *SLSAAttestationUpsertOne {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateBuiltFromHash()
	})
}

// Exec executes the query.
func (u *SLSAAttestationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SLSAAttestationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SLSAAttestationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SLSAAttestationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SLSAAttestationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SLSAAttestationCreateBulk is the builder for creating many SLSAAttestation entities in bulk.
type SLSAAttestationCreateBulk struct {
	config
	builders []*SLSAAttestationCreate
	conflict []sql.ConflictOption
}

// Save creates the SLSAAttestation entities in the database.
func (sacb *SLSAAttestationCreateBulk) Save(ctx context.Context) ([]*SLSAAttestation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*SLSAAttestation, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SLSAAttestationMutation)
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
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *SLSAAttestationCreateBulk) SaveX(ctx context.Context) []*SLSAAttestation {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sacb *SLSAAttestationCreateBulk) Exec(ctx context.Context) error {
	_, err := sacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sacb *SLSAAttestationCreateBulk) ExecX(ctx context.Context) {
	if err := sacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SLSAAttestation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SLSAAttestationUpsert) {
//			SetBuildType(v+v).
//		}).
//		Exec(ctx)
func (sacb *SLSAAttestationCreateBulk) OnConflict(opts ...sql.ConflictOption) *SLSAAttestationUpsertBulk {
	sacb.conflict = opts
	return &SLSAAttestationUpsertBulk{
		create: sacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SLSAAttestation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sacb *SLSAAttestationCreateBulk) OnConflictColumns(columns ...string) *SLSAAttestationUpsertBulk {
	sacb.conflict = append(sacb.conflict, sql.ConflictColumns(columns...))
	return &SLSAAttestationUpsertBulk{
		create: sacb,
	}
}

// SLSAAttestationUpsertBulk is the builder for "upsert"-ing
// a bulk of SLSAAttestation nodes.
type SLSAAttestationUpsertBulk struct {
	create *SLSAAttestationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SLSAAttestation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SLSAAttestationUpsertBulk) UpdateNewValues() *SLSAAttestationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SLSAAttestation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SLSAAttestationUpsertBulk) Ignore() *SLSAAttestationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SLSAAttestationUpsertBulk) DoNothing() *SLSAAttestationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SLSAAttestationCreateBulk.OnConflict
// documentation for more info.
func (u *SLSAAttestationUpsertBulk) Update(set func(*SLSAAttestationUpsert)) *SLSAAttestationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SLSAAttestationUpsert{UpdateSet: update})
	}))
	return u
}

// SetBuildType sets the "build_type" field.
func (u *SLSAAttestationUpsertBulk) SetBuildType(v string) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetBuildType(v)
	})
}

// UpdateBuildType sets the "build_type" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateBuildType() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateBuildType()
	})
}

// SetBuiltByID sets the "built_by_id" field.
func (u *SLSAAttestationUpsertBulk) SetBuiltByID(v int) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetBuiltByID(v)
	})
}

// UpdateBuiltByID sets the "built_by_id" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateBuiltByID() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateBuiltByID()
	})
}

// SetSubjectID sets the "subject_id" field.
func (u *SLSAAttestationUpsertBulk) SetSubjectID(v int) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetSubjectID(v)
	})
}

// UpdateSubjectID sets the "subject_id" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateSubjectID() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateSubjectID()
	})
}

// SetSlsaPredicate sets the "slsa_predicate" field.
func (u *SLSAAttestationUpsertBulk) SetSlsaPredicate(v []*model.SLSAPredicate) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetSlsaPredicate(v)
	})
}

// UpdateSlsaPredicate sets the "slsa_predicate" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateSlsaPredicate() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateSlsaPredicate()
	})
}

// ClearSlsaPredicate clears the value of the "slsa_predicate" field.
func (u *SLSAAttestationUpsertBulk) ClearSlsaPredicate() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.ClearSlsaPredicate()
	})
}

// SetSlsaVersion sets the "slsa_version" field.
func (u *SLSAAttestationUpsertBulk) SetSlsaVersion(v string) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetSlsaVersion(v)
	})
}

// UpdateSlsaVersion sets the "slsa_version" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateSlsaVersion() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateSlsaVersion()
	})
}

// SetStartedOn sets the "started_on" field.
func (u *SLSAAttestationUpsertBulk) SetStartedOn(v time.Time) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetStartedOn(v)
	})
}

// UpdateStartedOn sets the "started_on" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateStartedOn() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateStartedOn()
	})
}

// ClearStartedOn clears the value of the "started_on" field.
func (u *SLSAAttestationUpsertBulk) ClearStartedOn() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.ClearStartedOn()
	})
}

// SetFinishedOn sets the "finished_on" field.
func (u *SLSAAttestationUpsertBulk) SetFinishedOn(v time.Time) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetFinishedOn(v)
	})
}

// UpdateFinishedOn sets the "finished_on" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateFinishedOn() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateFinishedOn()
	})
}

// ClearFinishedOn clears the value of the "finished_on" field.
func (u *SLSAAttestationUpsertBulk) ClearFinishedOn() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.ClearFinishedOn()
	})
}

// SetOrigin sets the "origin" field.
func (u *SLSAAttestationUpsertBulk) SetOrigin(v string) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetOrigin(v)
	})
}

// UpdateOrigin sets the "origin" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateOrigin() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateOrigin()
	})
}

// SetCollector sets the "collector" field.
func (u *SLSAAttestationUpsertBulk) SetCollector(v string) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetCollector(v)
	})
}

// UpdateCollector sets the "collector" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateCollector() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateCollector()
	})
}

// SetBuiltFromHash sets the "built_from_hash" field.
func (u *SLSAAttestationUpsertBulk) SetBuiltFromHash(v string) *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.SetBuiltFromHash(v)
	})
}

// UpdateBuiltFromHash sets the "built_from_hash" field to the value that was provided on create.
func (u *SLSAAttestationUpsertBulk) UpdateBuiltFromHash() *SLSAAttestationUpsertBulk {
	return u.Update(func(s *SLSAAttestationUpsert) {
		s.UpdateBuiltFromHash()
	})
}

// Exec executes the query.
func (u *SLSAAttestationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SLSAAttestationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SLSAAttestationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SLSAAttestationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
