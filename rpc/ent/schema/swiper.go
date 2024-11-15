package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Swiper holds the schema definition for the Swiper entity.
type Swiper struct {
	ent.Schema
}

func (Swiper) Fields() []ent.Field {
	return []ent.Field{
		field.String("title_zh").Optional().Comment("Chinese title | 中文标题"),
		field.String("title_en").Optional().Comment("English title | 英文标题"),
		field.String("title_ru").Optional().Comment("Russian title | 俄语标题"),
		field.String("title_kk").Optional().Comment("Kazakh title | 哈萨克语标题"),
		field.String("banner_zh").Optional().Comment("Chinese banner image | 中文轮播图"),
		field.String("banner_en").Optional().Comment("English banner image | 英文轮播图"),
		field.String("banner_ru").Optional().Comment("Russian banner image | 俄语轮播图"),
		field.String("banner_kk").Optional().Comment("Kazakh banner image | 哈萨克语轮播图"),

		field.Text("content_zh").Optional().Comment("Chinese content | 中文内容"),
		field.Text("content_en").Optional().Comment("English content | 英文内容"),
		field.Text("content_ru").Optional().Comment("Russian content | 俄语内容"),
		field.Text("content_kk").Optional().Comment("Kazakh content | 哈萨克语内容"),

		field.String("jump_url").Optional().Comment("Jump Url | 跳转地址"),
	}
}

func (Swiper) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (Swiper) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "Swiper"},
	}
}
