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
		field.String("patient_name").
			Comment("Patient name | 姓名"),

		field.Int32("gender").
			Comment("Gender 1:male 2:female | 性别 1:男 2:女"),

		field.Int32("age").
			Comment("Age | 年龄"),

		field.String("id_card_number").Optional().
			Comment("ID card number | 身份证号"),

		field.String("phone_number").
			Comment("Phone number | 电话"),

		field.String("chief_complaint").Optional().
			Comment("Chief complaint | 主诉"),

		field.String("present_illness").Optional().
			Comment("Present illness history | 现病史"),

		field.String("past_history").Optional().
			Comment("Past history | 既往史"),

		field.Int32("smoking_history").Optional().Default(1).
			Comment("Smoking history | 吸烟史: 1-无, 2-有, 3-一般, 4-多, 5-已戒"),

		field.Int32("drinking_history").Optional().Default(1).
			Comment("Drinking history | 饮酒史: 1-无, 2-有, 3-一般, 4-多, 5-已戒"),

		field.Int32("allergy_history").Optional().Default(1).
			Comment("Allergy history | 过敏史: 1-无, 2-有"),

		// 查体信息
		field.Int32("heart_rate").Optional().
			Comment("Heart rate (beats/min) | 心率(次/分)"),

		field.String("blood_pressure").Optional().
			Comment("Blood pressure (mmHg) | 血压(mmHg)"),

		field.Float("oxygen_saturation").Optional().
			Comment("Oxygen saturation (%) | 指脉氧(%)"),

		field.Float("blood_glucose").Optional().
			Comment("Blood glucose (mmol/L) | 血糖(mmol/L)"),

		field.Float("weight").Optional().
			Comment("Weight (kg) | 体重(kg)"),

		field.Float("waist_circumference").Optional().
			Comment("Waist circumference (cm) | 腰围(cm)"),

		field.Float("body_fat").Optional().
			Comment("Body fat percentage (%) | 体脂数(%)"),

		field.String("diagnosis").Optional().
			Comment("Diagnosis | 诊断"),

		// 治疗方案
		field.Int32("diet_therapy").Optional().
			Comment("Diet therapy | 饮食治疗"),

		field.Int32("exercise_therapy").Optional().
			Comment("Exercise therapy | 运动治疗"),

		field.Int32("medication_therapy").Optional().
			Comment("Medication therapy | 药物治疗"),

		field.String("treatment_plan").Optional().
			Comment("Overall treatment plan | 综合治疗方案"),

		field.String("doctor_id").Optional().
			Comment("Doctor ID | 医生ID"),

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
