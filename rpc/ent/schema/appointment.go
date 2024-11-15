package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Appointment holds the schema definition for the Appointment entity.
type Appointment struct {
	ent.Schema
}

func (Appointment) Fields() []ent.Field {
	return []ent.Field{
		field.String("patient_name").Optional().
			Comment("Patient name | 患者姓名"),

		field.String("phone_number").Optional().
			Comment("Phone number | 联系电话"),

		field.String("id_card").Optional().
			Comment("ID card number | 身份证号"),

		field.Int32("gender").Optional().
			Comment("Gender 1:male 2:female | 性别 1:男 2:女"),

		field.Int32("age").Optional().
			Comment("Age | 年龄"),

		field.Int64("appointment_time").Optional().
			Comment("Appointment time | 预约时间"),

		field.String("symptoms").Optional().
			Comment("Symptoms description | 症状描述"),

		field.Int32("status").Optional().Default(1).
			Comment("Status 1:pending 2:confirmed 3:completed 4:cancelled 5:expired | 状态 1:待确认 2:已确认 3:已完成 4:已取消 5:已过期"),

		field.String("remarks").Optional().
			Comment("Remarks | 备注信息"),

		field.String("user_id").Optional().
			Comment("User ID | 用户ID"),
	}
}

func (Appointment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{}, // 使用 UUID 作为主键
	}
}

func (Appointment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "appointments"},
	}
}
