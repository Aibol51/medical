package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Expert holds the schema definition for the Expert entity.
type Expert struct {
	ent.Schema
}

func (Expert) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_zh").Optional().Comment("Chinese name | 中文名称"),
		field.String("name_en").Optional().Comment("English name | 英文名称"),
		field.String("name_ru").Optional().Comment("Russian name | 俄语名称"),
		field.String("name_kk").Optional().Comment("Kazakh name | 哈萨克语名称"),

		field.Text("content_zh").Optional().Comment("Chinese content | 中文内容"),
		field.Text("content_en").Optional().Comment("English content | 英文内容"),
		field.Text("content_ru").Optional().Comment("Russian content | 俄语内容"),
		field.Text("content_kk").Optional().Comment("Kazakh content | 哈萨克语内容"),

		field.String("cover_url").Optional().Comment("Cover image URL | 封面图片URL"),
	}
}

func (Expert) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (Expert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "expert"},
	}
}
