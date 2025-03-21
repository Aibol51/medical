// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/appointment"
)

// AppointmentCreate is the builder for creating a Appointment entity.
type AppointmentCreate struct {
	config
	mutation *AppointmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AppointmentCreate) SetCreatedAt(t time.Time) *AppointmentCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableCreatedAt(t *time.Time) *AppointmentCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AppointmentCreate) SetUpdatedAt(t time.Time) *AppointmentCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableUpdatedAt(t *time.Time) *AppointmentCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetPatientName sets the "patient_name" field.
func (ac *AppointmentCreate) SetPatientName(s string) *AppointmentCreate {
	ac.mutation.SetPatientName(s)
	return ac
}

// SetNillablePatientName sets the "patient_name" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillablePatientName(s *string) *AppointmentCreate {
	if s != nil {
		ac.SetPatientName(*s)
	}
	return ac
}

// SetPhoneNumber sets the "phone_number" field.
func (ac *AppointmentCreate) SetPhoneNumber(s string) *AppointmentCreate {
	ac.mutation.SetPhoneNumber(s)
	return ac
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillablePhoneNumber(s *string) *AppointmentCreate {
	if s != nil {
		ac.SetPhoneNumber(*s)
	}
	return ac
}

// SetIDCard sets the "id_card" field.
func (ac *AppointmentCreate) SetIDCard(s string) *AppointmentCreate {
	ac.mutation.SetIDCard(s)
	return ac
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableIDCard(s *string) *AppointmentCreate {
	if s != nil {
		ac.SetIDCard(*s)
	}
	return ac
}

// SetGender sets the "gender" field.
func (ac *AppointmentCreate) SetGender(i int32) *AppointmentCreate {
	ac.mutation.SetGender(i)
	return ac
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableGender(i *int32) *AppointmentCreate {
	if i != nil {
		ac.SetGender(*i)
	}
	return ac
}

// SetAge sets the "age" field.
func (ac *AppointmentCreate) SetAge(i int32) *AppointmentCreate {
	ac.mutation.SetAge(i)
	return ac
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableAge(i *int32) *AppointmentCreate {
	if i != nil {
		ac.SetAge(*i)
	}
	return ac
}

// SetAppointmentTime sets the "appointment_time" field.
func (ac *AppointmentCreate) SetAppointmentTime(i int64) *AppointmentCreate {
	ac.mutation.SetAppointmentTime(i)
	return ac
}

// SetNillableAppointmentTime sets the "appointment_time" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableAppointmentTime(i *int64) *AppointmentCreate {
	if i != nil {
		ac.SetAppointmentTime(*i)
	}
	return ac
}

// SetSymptoms sets the "symptoms" field.
func (ac *AppointmentCreate) SetSymptoms(s string) *AppointmentCreate {
	ac.mutation.SetSymptoms(s)
	return ac
}

// SetNillableSymptoms sets the "symptoms" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableSymptoms(s *string) *AppointmentCreate {
	if s != nil {
		ac.SetSymptoms(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *AppointmentCreate) SetStatus(i int32) *AppointmentCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableStatus(i *int32) *AppointmentCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// SetRemarks sets the "remarks" field.
func (ac *AppointmentCreate) SetRemarks(s string) *AppointmentCreate {
	ac.mutation.SetRemarks(s)
	return ac
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableRemarks(s *string) *AppointmentCreate {
	if s != nil {
		ac.SetRemarks(*s)
	}
	return ac
}

// SetUserID sets the "user_id" field.
func (ac *AppointmentCreate) SetUserID(s string) *AppointmentCreate {
	ac.mutation.SetUserID(s)
	return ac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableUserID(s *string) *AppointmentCreate {
	if s != nil {
		ac.SetUserID(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AppointmentCreate) SetID(u uuid.UUID) *AppointmentCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AppointmentCreate) SetNillableID(u *uuid.UUID) *AppointmentCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// Mutation returns the AppointmentMutation object of the builder.
func (ac *AppointmentCreate) Mutation() *AppointmentMutation {
	return ac.mutation
}

// Save creates the Appointment in the database.
func (ac *AppointmentCreate) Save(ctx context.Context) (*Appointment, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppointmentCreate) SaveX(ctx context.Context) *Appointment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AppointmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AppointmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AppointmentCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := appointment.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := appointment.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Status(); !ok {
		v := appointment.DefaultStatus
		ac.mutation.SetStatus(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := appointment.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AppointmentCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Appointment.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Appointment.updated_at"`)}
	}
	return nil
}

func (ac *AppointmentCreate) sqlSave(ctx context.Context) (*Appointment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AppointmentCreate) createSpec() (*Appointment, *sqlgraph.CreateSpec) {
	var (
		_node = &Appointment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(appointment.Table, sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(appointment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(appointment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.PatientName(); ok {
		_spec.SetField(appointment.FieldPatientName, field.TypeString, value)
		_node.PatientName = value
	}
	if value, ok := ac.mutation.PhoneNumber(); ok {
		_spec.SetField(appointment.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := ac.mutation.IDCard(); ok {
		_spec.SetField(appointment.FieldIDCard, field.TypeString, value)
		_node.IDCard = value
	}
	if value, ok := ac.mutation.Gender(); ok {
		_spec.SetField(appointment.FieldGender, field.TypeInt32, value)
		_node.Gender = value
	}
	if value, ok := ac.mutation.Age(); ok {
		_spec.SetField(appointment.FieldAge, field.TypeInt32, value)
		_node.Age = value
	}
	if value, ok := ac.mutation.AppointmentTime(); ok {
		_spec.SetField(appointment.FieldAppointmentTime, field.TypeInt64, value)
		_node.AppointmentTime = value
	}
	if value, ok := ac.mutation.Symptoms(); ok {
		_spec.SetField(appointment.FieldSymptoms, field.TypeString, value)
		_node.Symptoms = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(appointment.FieldStatus, field.TypeInt32, value)
		_node.Status = value
	}
	if value, ok := ac.mutation.Remarks(); ok {
		_spec.SetField(appointment.FieldRemarks, field.TypeString, value)
		_node.Remarks = value
	}
	if value, ok := ac.mutation.UserID(); ok {
		_spec.SetField(appointment.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	return _node, _spec
}

// AppointmentCreateBulk is the builder for creating many Appointment entities in bulk.
type AppointmentCreateBulk struct {
	config
	err      error
	builders []*AppointmentCreate
}

// Save creates the Appointment entities in the database.
func (acb *AppointmentCreateBulk) Save(ctx context.Context) ([]*Appointment, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Appointment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppointmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AppointmentCreateBulk) SaveX(ctx context.Context) []*Appointment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AppointmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AppointmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
