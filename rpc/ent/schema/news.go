package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

func (News) Fields() []ent.Field {
	return []ent.Field{
		field.String("title_zh").Optional().Comment("Chinese title | 中文标题"),
		field.String("title_en").Optional().Comment("English title | 英文标题"),
		field.String("title_ru").Optional().Comment("Russian title | 俄语标题"),
		field.String("title_kk").Optional().Comment("Kazakh title | 哈萨克语标题"),

		field.Text("content_zh").Optional().Comment("Chinese content | 中文内容"),
		field.Text("content_en").Optional().Comment("English content | 英文内容"),
		field.Text("content_ru").Optional().Comment("Russian content | 俄语内容"),
		field.Text("content_kk").Optional().Comment("Kazakh content | 哈萨克语内容"),

		field.String("cover_url").Optional().Comment("Cover image URL | 封面图片URL"),
	}
}

func (News) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (News) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "news"},
	}
}
