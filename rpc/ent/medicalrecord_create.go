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
	"github.com/suyuan32/simple-admin-core/rpc/ent/medicalrecord"
)

// MedicalRecordCreate is the builder for creating a MedicalRecord entity.
type MedicalRecordCreate struct {
	config
	mutation *MedicalRecordMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mrc *MedicalRecordCreate) SetCreatedAt(t time.Time) *MedicalRecordCreate {
	mrc.mutation.SetCreatedAt(t)
	return mrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableCreatedAt(t *time.Time) *MedicalRecordCreate {
	if t != nil {
		mrc.SetCreatedAt(*t)
	}
	return mrc
}

// SetUpdatedAt sets the "updated_at" field.
func (mrc *MedicalRecordCreate) SetUpdatedAt(t time.Time) *MedicalRecordCreate {
	mrc.mutation.SetUpdatedAt(t)
	return mrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableUpdatedAt(t *time.Time) *MedicalRecordCreate {
	if t != nil {
		mrc.SetUpdatedAt(*t)
	}
	return mrc
}

// SetPatientName sets the "patient_name" field.
func (mrc *MedicalRecordCreate) SetPatientName(s string) *MedicalRecordCreate {
	mrc.mutation.SetPatientName(s)
	return mrc
}

// SetGender sets the "gender" field.
func (mrc *MedicalRecordCreate) SetGender(i int32) *MedicalRecordCreate {
	mrc.mutation.SetGender(i)
	return mrc
}

// SetAge sets the "age" field.
func (mrc *MedicalRecordCreate) SetAge(i int32) *MedicalRecordCreate {
	mrc.mutation.SetAge(i)
	return mrc
}

// SetIDCardNumber sets the "id_card_number" field.
func (mrc *MedicalRecordCreate) SetIDCardNumber(s string) *MedicalRecordCreate {
	mrc.mutation.SetIDCardNumber(s)
	return mrc
}

// SetNillableIDCardNumber sets the "id_card_number" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableIDCardNumber(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetIDCardNumber(*s)
	}
	return mrc
}

// SetPhoneNumber sets the "phone_number" field.
func (mrc *MedicalRecordCreate) SetPhoneNumber(s string) *MedicalRecordCreate {
	mrc.mutation.SetPhoneNumber(s)
	return mrc
}

// SetChiefComplaint sets the "chief_complaint" field.
func (mrc *MedicalRecordCreate) SetChiefComplaint(s string) *MedicalRecordCreate {
	mrc.mutation.SetChiefComplaint(s)
	return mrc
}

// SetNillableChiefComplaint sets the "chief_complaint" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableChiefComplaint(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetChiefComplaint(*s)
	}
	return mrc
}

// SetPresentIllness sets the "present_illness" field.
func (mrc *MedicalRecordCreate) SetPresentIllness(s string) *MedicalRecordCreate {
	mrc.mutation.SetPresentIllness(s)
	return mrc
}

// SetNillablePresentIllness sets the "present_illness" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillablePresentIllness(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetPresentIllness(*s)
	}
	return mrc
}

// SetPastHistory sets the "past_history" field.
func (mrc *MedicalRecordCreate) SetPastHistory(s string) *MedicalRecordCreate {
	mrc.mutation.SetPastHistory(s)
	return mrc
}

// SetNillablePastHistory sets the "past_history" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillablePastHistory(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetPastHistory(*s)
	}
	return mrc
}

// SetSmokingHistory sets the "smoking_history" field.
func (mrc *MedicalRecordCreate) SetSmokingHistory(i int32) *MedicalRecordCreate {
	mrc.mutation.SetSmokingHistory(i)
	return mrc
}

// SetNillableSmokingHistory sets the "smoking_history" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableSmokingHistory(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetSmokingHistory(*i)
	}
	return mrc
}

// SetDrinkingHistory sets the "drinking_history" field.
func (mrc *MedicalRecordCreate) SetDrinkingHistory(i int32) *MedicalRecordCreate {
	mrc.mutation.SetDrinkingHistory(i)
	return mrc
}

// SetNillableDrinkingHistory sets the "drinking_history" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableDrinkingHistory(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetDrinkingHistory(*i)
	}
	return mrc
}

// SetAllergyHistory sets the "allergy_history" field.
func (mrc *MedicalRecordCreate) SetAllergyHistory(i int32) *MedicalRecordCreate {
	mrc.mutation.SetAllergyHistory(i)
	return mrc
}

// SetNillableAllergyHistory sets the "allergy_history" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableAllergyHistory(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetAllergyHistory(*i)
	}
	return mrc
}

// SetHeartRate sets the "heart_rate" field.
func (mrc *MedicalRecordCreate) SetHeartRate(i int32) *MedicalRecordCreate {
	mrc.mutation.SetHeartRate(i)
	return mrc
}

// SetNillableHeartRate sets the "heart_rate" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableHeartRate(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetHeartRate(*i)
	}
	return mrc
}

// SetBloodPressure sets the "blood_pressure" field.
func (mrc *MedicalRecordCreate) SetBloodPressure(s string) *MedicalRecordCreate {
	mrc.mutation.SetBloodPressure(s)
	return mrc
}

// SetNillableBloodPressure sets the "blood_pressure" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableBloodPressure(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetBloodPressure(*s)
	}
	return mrc
}

// SetOxygenSaturation sets the "oxygen_saturation" field.
func (mrc *MedicalRecordCreate) SetOxygenSaturation(f float64) *MedicalRecordCreate {
	mrc.mutation.SetOxygenSaturation(f)
	return mrc
}

// SetNillableOxygenSaturation sets the "oxygen_saturation" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableOxygenSaturation(f *float64) *MedicalRecordCreate {
	if f != nil {
		mrc.SetOxygenSaturation(*f)
	}
	return mrc
}

// SetBloodGlucose sets the "blood_glucose" field.
func (mrc *MedicalRecordCreate) SetBloodGlucose(f float64) *MedicalRecordCreate {
	mrc.mutation.SetBloodGlucose(f)
	return mrc
}

// SetNillableBloodGlucose sets the "blood_glucose" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableBloodGlucose(f *float64) *MedicalRecordCreate {
	if f != nil {
		mrc.SetBloodGlucose(*f)
	}
	return mrc
}

// SetWeight sets the "weight" field.
func (mrc *MedicalRecordCreate) SetWeight(f float64) *MedicalRecordCreate {
	mrc.mutation.SetWeight(f)
	return mrc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableWeight(f *float64) *MedicalRecordCreate {
	if f != nil {
		mrc.SetWeight(*f)
	}
	return mrc
}

// SetWaistCircumference sets the "waist_circumference" field.
func (mrc *MedicalRecordCreate) SetWaistCircumference(f float64) *MedicalRecordCreate {
	mrc.mutation.SetWaistCircumference(f)
	return mrc
}

// SetNillableWaistCircumference sets the "waist_circumference" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableWaistCircumference(f *float64) *MedicalRecordCreate {
	if f != nil {
		mrc.SetWaistCircumference(*f)
	}
	return mrc
}

// SetBodyFat sets the "body_fat" field.
func (mrc *MedicalRecordCreate) SetBodyFat(f float64) *MedicalRecordCreate {
	mrc.mutation.SetBodyFat(f)
	return mrc
}

// SetNillableBodyFat sets the "body_fat" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableBodyFat(f *float64) *MedicalRecordCreate {
	if f != nil {
		mrc.SetBodyFat(*f)
	}
	return mrc
}

// SetDiagnosis sets the "diagnosis" field.
func (mrc *MedicalRecordCreate) SetDiagnosis(s string) *MedicalRecordCreate {
	mrc.mutation.SetDiagnosis(s)
	return mrc
}

// SetNillableDiagnosis sets the "diagnosis" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableDiagnosis(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetDiagnosis(*s)
	}
	return mrc
}

// SetDietTherapy sets the "diet_therapy" field.
func (mrc *MedicalRecordCreate) SetDietTherapy(i int32) *MedicalRecordCreate {
	mrc.mutation.SetDietTherapy(i)
	return mrc
}

// SetNillableDietTherapy sets the "diet_therapy" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableDietTherapy(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetDietTherapy(*i)
	}
	return mrc
}

// SetExerciseTherapy sets the "exercise_therapy" field.
func (mrc *MedicalRecordCreate) SetExerciseTherapy(i int32) *MedicalRecordCreate {
	mrc.mutation.SetExerciseTherapy(i)
	return mrc
}

// SetNillableExerciseTherapy sets the "exercise_therapy" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableExerciseTherapy(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetExerciseTherapy(*i)
	}
	return mrc
}

// SetMedicationTherapy sets the "medication_therapy" field.
func (mrc *MedicalRecordCreate) SetMedicationTherapy(i int32) *MedicalRecordCreate {
	mrc.mutation.SetMedicationTherapy(i)
	return mrc
}

// SetNillableMedicationTherapy sets the "medication_therapy" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableMedicationTherapy(i *int32) *MedicalRecordCreate {
	if i != nil {
		mrc.SetMedicationTherapy(*i)
	}
	return mrc
}

// SetTreatmentPlan sets the "treatment_plan" field.
func (mrc *MedicalRecordCreate) SetTreatmentPlan(s string) *MedicalRecordCreate {
	mrc.mutation.SetTreatmentPlan(s)
	return mrc
}

// SetNillableTreatmentPlan sets the "treatment_plan" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableTreatmentPlan(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetTreatmentPlan(*s)
	}
	return mrc
}

// SetDoctorID sets the "doctor_id" field.
func (mrc *MedicalRecordCreate) SetDoctorID(s string) *MedicalRecordCreate {
	mrc.mutation.SetDoctorID(s)
	return mrc
}

// SetNillableDoctorID sets the "doctor_id" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableDoctorID(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetDoctorID(*s)
	}
	return mrc
}

// SetAppointmentID sets the "appointment_id" field.
func (mrc *MedicalRecordCreate) SetAppointmentID(s string) *MedicalRecordCreate {
	mrc.mutation.SetAppointmentID(s)
	return mrc
}

// SetNillableAppointmentID sets the "appointment_id" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableAppointmentID(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetAppointmentID(*s)
	}
	return mrc
}

// SetRemarks sets the "remarks" field.
func (mrc *MedicalRecordCreate) SetRemarks(s string) *MedicalRecordCreate {
	mrc.mutation.SetRemarks(s)
	return mrc
}

// SetNillableRemarks sets the "remarks" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableRemarks(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetRemarks(*s)
	}
	return mrc
}

// SetUserID sets the "user_id" field.
func (mrc *MedicalRecordCreate) SetUserID(s string) *MedicalRecordCreate {
	mrc.mutation.SetUserID(s)
	return mrc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableUserID(s *string) *MedicalRecordCreate {
	if s != nil {
		mrc.SetUserID(*s)
	}
	return mrc
}

// SetID sets the "id" field.
func (mrc *MedicalRecordCreate) SetID(u uuid.UUID) *MedicalRecordCreate {
	mrc.mutation.SetID(u)
	return mrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mrc *MedicalRecordCreate) SetNillableID(u *uuid.UUID) *MedicalRecordCreate {
	if u != nil {
		mrc.SetID(*u)
	}
	return mrc
}

// Mutation returns the MedicalRecordMutation object of the builder.
func (mrc *MedicalRecordCreate) Mutation() *MedicalRecordMutation {
	return mrc.mutation
}

// Save creates the MedicalRecord in the database.
func (mrc *MedicalRecordCreate) Save(ctx context.Context) (*MedicalRecord, error) {
	mrc.defaults()
	return withHooks(ctx, mrc.sqlSave, mrc.mutation, mrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *MedicalRecordCreate) SaveX(ctx context.Context) *MedicalRecord {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *MedicalRecordCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *MedicalRecordCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mrc *MedicalRecordCreate) defaults() {
	if _, ok := mrc.mutation.CreatedAt(); !ok {
		v := medicalrecord.DefaultCreatedAt()
		mrc.mutation.SetCreatedAt(v)
	}
	if _, ok := mrc.mutation.UpdatedAt(); !ok {
		v := medicalrecord.DefaultUpdatedAt()
		mrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mrc.mutation.SmokingHistory(); !ok {
		v := medicalrecord.DefaultSmokingHistory
		mrc.mutation.SetSmokingHistory(v)
	}
	if _, ok := mrc.mutation.DrinkingHistory(); !ok {
		v := medicalrecord.DefaultDrinkingHistory
		mrc.mutation.SetDrinkingHistory(v)
	}
	if _, ok := mrc.mutation.AllergyHistory(); !ok {
		v := medicalrecord.DefaultAllergyHistory
		mrc.mutation.SetAllergyHistory(v)
	}
	if _, ok := mrc.mutation.ID(); !ok {
		v := medicalrecord.DefaultID()
		mrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *MedicalRecordCreate) check() error {
	if _, ok := mrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MedicalRecord.created_at"`)}
	}
	if _, ok := mrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MedicalRecord.updated_at"`)}
	}
	if _, ok := mrc.mutation.PatientName(); !ok {
		return &ValidationError{Name: "patient_name", err: errors.New(`ent: missing required field "MedicalRecord.patient_name"`)}
	}
	if _, ok := mrc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "MedicalRecord.gender"`)}
	}
	if _, ok := mrc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "MedicalRecord.age"`)}
	}
	if _, ok := mrc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "MedicalRecord.phone_number"`)}
	}
	return nil
}

func (mrc *MedicalRecordCreate) sqlSave(ctx context.Context) (*MedicalRecord, error) {
	if err := mrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
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
	mrc.mutation.id = &_node.ID
	mrc.mutation.done = true
	return _node, nil
}

func (mrc *MedicalRecordCreate) createSpec() (*MedicalRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &MedicalRecord{config: mrc.config}
		_spec = sqlgraph.NewCreateSpec(medicalrecord.Table, sqlgraph.NewFieldSpec(medicalrecord.FieldID, field.TypeUUID))
	)
	if id, ok := mrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mrc.mutation.CreatedAt(); ok {
		_spec.SetField(medicalrecord.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mrc.mutation.UpdatedAt(); ok {
		_spec.SetField(medicalrecord.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mrc.mutation.PatientName(); ok {
		_spec.SetField(medicalrecord.FieldPatientName, field.TypeString, value)
		_node.PatientName = value
	}
	if value, ok := mrc.mutation.Gender(); ok {
		_spec.SetField(medicalrecord.FieldGender, field.TypeInt32, value)
		_node.Gender = value
	}
	if value, ok := mrc.mutation.Age(); ok {
		_spec.SetField(medicalrecord.FieldAge, field.TypeInt32, value)
		_node.Age = value
	}
	if value, ok := mrc.mutation.IDCardNumber(); ok {
		_spec.SetField(medicalrecord.FieldIDCardNumber, field.TypeString, value)
		_node.IDCardNumber = value
	}
	if value, ok := mrc.mutation.PhoneNumber(); ok {
		_spec.SetField(medicalrecord.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := mrc.mutation.ChiefComplaint(); ok {
		_spec.SetField(medicalrecord.FieldChiefComplaint, field.TypeString, value)
		_node.ChiefComplaint = value
	}
	if value, ok := mrc.mutation.PresentIllness(); ok {
		_spec.SetField(medicalrecord.FieldPresentIllness, field.TypeString, value)
		_node.PresentIllness = value
	}
	if value, ok := mrc.mutation.PastHistory(); ok {
		_spec.SetField(medicalrecord.FieldPastHistory, field.TypeString, value)
		_node.PastHistory = value
	}
	if value, ok := mrc.mutation.SmokingHistory(); ok {
		_spec.SetField(medicalrecord.FieldSmokingHistory, field.TypeInt32, value)
		_node.SmokingHistory = value
	}
	if value, ok := mrc.mutation.DrinkingHistory(); ok {
		_spec.SetField(medicalrecord.FieldDrinkingHistory, field.TypeInt32, value)
		_node.DrinkingHistory = value
	}
	if value, ok := mrc.mutation.AllergyHistory(); ok {
		_spec.SetField(medicalrecord.FieldAllergyHistory, field.TypeInt32, value)
		_node.AllergyHistory = value
	}
	if value, ok := mrc.mutation.HeartRate(); ok {
		_spec.SetField(medicalrecord.FieldHeartRate, field.TypeInt32, value)
		_node.HeartRate = value
	}
	if value, ok := mrc.mutation.BloodPressure(); ok {
		_spec.SetField(medicalrecord.FieldBloodPressure, field.TypeString, value)
		_node.BloodPressure = value
	}
	if value, ok := mrc.mutation.OxygenSaturation(); ok {
		_spec.SetField(medicalrecord.FieldOxygenSaturation, field.TypeFloat64, value)
		_node.OxygenSaturation = value
	}
	if value, ok := mrc.mutation.BloodGlucose(); ok {
		_spec.SetField(medicalrecord.FieldBloodGlucose, field.TypeFloat64, value)
		_node.BloodGlucose = value
	}
	if value, ok := mrc.mutation.Weight(); ok {
		_spec.SetField(medicalrecord.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if value, ok := mrc.mutation.WaistCircumference(); ok {
		_spec.SetField(medicalrecord.FieldWaistCircumference, field.TypeFloat64, value)
		_node.WaistCircumference = value
	}
	if value, ok := mrc.mutation.BodyFat(); ok {
		_spec.SetField(medicalrecord.FieldBodyFat, field.TypeFloat64, value)
		_node.BodyFat = value
	}
	if value, ok := mrc.mutation.Diagnosis(); ok {
		_spec.SetField(medicalrecord.FieldDiagnosis, field.TypeString, value)
		_node.Diagnosis = value
	}
	if value, ok := mrc.mutation.DietTherapy(); ok {
		_spec.SetField(medicalrecord.FieldDietTherapy, field.TypeInt32, value)
		_node.DietTherapy = value
	}
	if value, ok := mrc.mutation.ExerciseTherapy(); ok {
		_spec.SetField(medicalrecord.FieldExerciseTherapy, field.TypeInt32, value)
		_node.ExerciseTherapy = value
	}
	if value, ok := mrc.mutation.MedicationTherapy(); ok {
		_spec.SetField(medicalrecord.FieldMedicationTherapy, field.TypeInt32, value)
		_node.MedicationTherapy = value
	}
	if value, ok := mrc.mutation.TreatmentPlan(); ok {
		_spec.SetField(medicalrecord.FieldTreatmentPlan, field.TypeString, value)
		_node.TreatmentPlan = value
	}
	if value, ok := mrc.mutation.DoctorID(); ok {
		_spec.SetField(medicalrecord.FieldDoctorID, field.TypeString, value)
		_node.DoctorID = value
	}
	if value, ok := mrc.mutation.AppointmentID(); ok {
		_spec.SetField(medicalrecord.FieldAppointmentID, field.TypeString, value)
		_node.AppointmentID = value
	}
	if value, ok := mrc.mutation.Remarks(); ok {
		_spec.SetField(medicalrecord.FieldRemarks, field.TypeString, value)
		_node.Remarks = value
	}
	if value, ok := mrc.mutation.UserID(); ok {
		_spec.SetField(medicalrecord.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	return _node, _spec
}

// MedicalRecordCreateBulk is the builder for creating many MedicalRecord entities in bulk.
type MedicalRecordCreateBulk struct {
	config
	err      error
	builders []*MedicalRecordCreate
}

// Save creates the MedicalRecord entities in the database.
func (mrcb *MedicalRecordCreateBulk) Save(ctx context.Context) ([]*MedicalRecord, error) {
	if mrcb.err != nil {
		return nil, mrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*MedicalRecord, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MedicalRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *MedicalRecordCreateBulk) SaveX(ctx context.Context) []*MedicalRecord {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *MedicalRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *MedicalRecordCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
