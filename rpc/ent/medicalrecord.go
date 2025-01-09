// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/medicalrecord"
)

// MedicalRecord is the model entity for the MedicalRecord schema.
type MedicalRecord struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Patient name | 患者姓名
	PatientName string `json:"patient_name,omitempty"`
	// Phone number | 联系电话
	PhoneNumber string `json:"phone_number,omitempty"`
	// Gender 1:male 2:female | 性别 1:男 2:女
	Gender int32 `json:"gender,omitempty"`
	// Age | 年龄
	Age int32 `json:"age,omitempty"`
	// Visit time | 就诊时间
	VisitTime int64 `json:"visit_time,omitempty"`
	// Diagnosis | 诊断
	Diagnosis string `json:"diagnosis,omitempty"`
	// Treatment plan | 治疗方案
	TreatmentPlan string `json:"treatment_plan,omitempty"`
	// Prescription | 处方
	Prescription string `json:"prescription,omitempty"`
	// Examination results | 检查结果
	ExaminationResults string `json:"examination_results,omitempty"`
	// Doctor's advice | 医嘱
	DoctorAdvice string `json:"doctor_advice,omitempty"`
	// Doctor ID | 医生ID
	DoctorID string `json:"doctor_id,omitempty"`
	// Department | 科室
	Department string `json:"department,omitempty"`
	// Related appointment ID | 关联预约ID
	AppointmentID string `json:"appointment_id,omitempty"`
	// Remarks | 备注信息
	Remarks string `json:"remarks,omitempty"`
	// User ID | 用户ID
	UserID       string `json:"user_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MedicalRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case medicalrecord.FieldGender, medicalrecord.FieldAge, medicalrecord.FieldVisitTime:
			values[i] = new(sql.NullInt64)
		case medicalrecord.FieldPatientName, medicalrecord.FieldPhoneNumber, medicalrecord.FieldDiagnosis, medicalrecord.FieldTreatmentPlan, medicalrecord.FieldPrescription, medicalrecord.FieldExaminationResults, medicalrecord.FieldDoctorAdvice, medicalrecord.FieldDoctorID, medicalrecord.FieldDepartment, medicalrecord.FieldAppointmentID, medicalrecord.FieldRemarks, medicalrecord.FieldUserID:
			values[i] = new(sql.NullString)
		case medicalrecord.FieldCreatedAt, medicalrecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case medicalrecord.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MedicalRecord fields.
func (mr *MedicalRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case medicalrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				mr.ID = *value
			}
		case medicalrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mr.CreatedAt = value.Time
			}
		case medicalrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mr.UpdatedAt = value.Time
			}
		case medicalrecord.FieldPatientName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field patient_name", values[i])
			} else if value.Valid {
				mr.PatientName = value.String
			}
		case medicalrecord.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				mr.PhoneNumber = value.String
			}
		case medicalrecord.FieldGender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				mr.Gender = int32(value.Int64)
			}
		case medicalrecord.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				mr.Age = int32(value.Int64)
			}
		case medicalrecord.FieldVisitTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field visit_time", values[i])
			} else if value.Valid {
				mr.VisitTime = value.Int64
			}
		case medicalrecord.FieldDiagnosis:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field diagnosis", values[i])
			} else if value.Valid {
				mr.Diagnosis = value.String
			}
		case medicalrecord.FieldTreatmentPlan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field treatment_plan", values[i])
			} else if value.Valid {
				mr.TreatmentPlan = value.String
			}
		case medicalrecord.FieldPrescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prescription", values[i])
			} else if value.Valid {
				mr.Prescription = value.String
			}
		case medicalrecord.FieldExaminationResults:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field examination_results", values[i])
			} else if value.Valid {
				mr.ExaminationResults = value.String
			}
		case medicalrecord.FieldDoctorAdvice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field doctor_advice", values[i])
			} else if value.Valid {
				mr.DoctorAdvice = value.String
			}
		case medicalrecord.FieldDoctorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field doctor_id", values[i])
			} else if value.Valid {
				mr.DoctorID = value.String
			}
		case medicalrecord.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field department", values[i])
			} else if value.Valid {
				mr.Department = value.String
			}
		case medicalrecord.FieldAppointmentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field appointment_id", values[i])
			} else if value.Valid {
				mr.AppointmentID = value.String
			}
		case medicalrecord.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				mr.Remarks = value.String
			}
		case medicalrecord.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				mr.UserID = value.String
			}
		default:
			mr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MedicalRecord.
// This includes values selected through modifiers, order, etc.
func (mr *MedicalRecord) Value(name string) (ent.Value, error) {
	return mr.selectValues.Get(name)
}

// Update returns a builder for updating this MedicalRecord.
// Note that you need to call MedicalRecord.Unwrap() before calling this method if this MedicalRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (mr *MedicalRecord) Update() *MedicalRecordUpdateOne {
	return NewMedicalRecordClient(mr.config).UpdateOne(mr)
}

// Unwrap unwraps the MedicalRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mr *MedicalRecord) Unwrap() *MedicalRecord {
	_tx, ok := mr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MedicalRecord is not a transactional entity")
	}
	mr.config.driver = _tx.drv
	return mr
}

// String implements the fmt.Stringer.
func (mr *MedicalRecord) String() string {
	var builder strings.Builder
	builder.WriteString("MedicalRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("patient_name=")
	builder.WriteString(mr.PatientName)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(mr.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", mr.Gender))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", mr.Age))
	builder.WriteString(", ")
	builder.WriteString("visit_time=")
	builder.WriteString(fmt.Sprintf("%v", mr.VisitTime))
	builder.WriteString(", ")
	builder.WriteString("diagnosis=")
	builder.WriteString(mr.Diagnosis)
	builder.WriteString(", ")
	builder.WriteString("treatment_plan=")
	builder.WriteString(mr.TreatmentPlan)
	builder.WriteString(", ")
	builder.WriteString("prescription=")
	builder.WriteString(mr.Prescription)
	builder.WriteString(", ")
	builder.WriteString("examination_results=")
	builder.WriteString(mr.ExaminationResults)
	builder.WriteString(", ")
	builder.WriteString("doctor_advice=")
	builder.WriteString(mr.DoctorAdvice)
	builder.WriteString(", ")
	builder.WriteString("doctor_id=")
	builder.WriteString(mr.DoctorID)
	builder.WriteString(", ")
	builder.WriteString("department=")
	builder.WriteString(mr.Department)
	builder.WriteString(", ")
	builder.WriteString("appointment_id=")
	builder.WriteString(mr.AppointmentID)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(mr.Remarks)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(mr.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// MedicalRecords is a parsable slice of MedicalRecord.
type MedicalRecords []*MedicalRecord
