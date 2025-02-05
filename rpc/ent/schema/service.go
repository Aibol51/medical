package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Service schema for the ServiceInfo proto message
type Service struct {
	ent.Schema
}

// Fields defines the fields for the Service schema
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_zh").
			Comment("Service chinese name | 服务中文名称"),
		field.String("name_en").
			Comment("Service english name | 服务英文名称"),
		field.String("name_ru").
			Comment("Service russian name | 服务俄语名称"),
		field.String("name_kk").
			Comment("Service kazakh name | 服务哈萨克语名称"),
		field.Text("description_zh").Optional().
			Comment("Description chinese | 服务中文描述"),
		field.Text("description_en").Optional().
			Comment("Description english | 服务英文描述"),
		field.Text("description_ru").Optional().
			Comment("Description russian | 服务俄语描述"),
		field.Text("description_kk").Optional().
			Comment("Description kazakh | 服务哈萨克语描述"),
		field.String("cover_url").Optional().Comment("Cover image URL | 封面图片URL"),
	}
}

// Mixin defines the mixins for the Service schema
func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},     // Adds an ID field
		mixins.StatusMixin{}, // Adds a status field for active/inactive states
		mixins.SortMixin{},   // Adds a sort field for ordering
	}
}

// Annotations defines the annotations for the Service schema
func (Service) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "services"},
	}
}
