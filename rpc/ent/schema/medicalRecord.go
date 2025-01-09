package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// MedicalRecord holds the schema definition for the MedicalRecord entity.
type MedicalRecord struct {
	ent.Schema
}

func (MedicalRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("patient_name").Optional().
			Comment("Patient name | 患者姓名"),

		field.String("phone_number").Optional().
			Comment("Phone number | 联系电话"),

		field.Int32("gender").Optional().
			Comment("Gender 1:male 2:female | 性别 1:男 2:女"),

		field.Int32("age").Optional().
			Comment("Age | 年龄"),

		field.Int64("visit_time").Optional().
			Comment("Visit time | 就诊时间"),

		field.String("diagnosis").Optional().
			Comment("Diagnosis | 诊断"),

		field.String("treatment_plan").Optional().
			Comment("Treatment plan | 治疗方案"),

		field.String("prescription").Optional().
			Comment("Prescription | 处方"),

		field.String("examination_results").Optional().
			Comment("Examination results | 检查结果"),

		field.String("doctor_advice").Optional().
			Comment("Doctor's advice | 医嘱"),

		field.String("doctor_id").Optional().
			Comment("Doctor ID | 医生ID"),

		field.String("department").Optional().
			Comment("Department | 科室"),

		field.String("appointment_id").Optional().
			Comment("Related appointment ID | 关联预约ID"),

		field.String("remarks").Optional().
			Comment("Remarks | 备注信息"),

		field.String("user_id").Optional().
			Comment("User ID | 用户ID"),
	}
}

func (MedicalRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{}, // Using UUID as primary key | 使用UUID作为主键
	}
}

func (MedicalRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "medical_records"},
	}
}
