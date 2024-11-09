package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Medicine schema for the MedicineInfo proto message
type Medicine struct {
	ent.Schema
}

// Fields defines the fields for the Medicine schema
func (Medicine) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("Medicine name | 药品名称"),
		field.Uint32("quantity").
			Comment("Quantity in stock | 库存数量"),
		field.String("description").Optional().
			Comment("Description | 药品描述"),
		field.String("remarks").Optional().
			Comment("Remarks | 备注信息"),
		field.String("images").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("Images | 图片路径"),
	}
}

// Mixin defines the mixins for the Medicine schema
func (Medicine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},     // Adds an ID field
		mixins.StatusMixin{}, // Adds a status field for active/inactive states
		mixins.SortMixin{},   // Adds a sort field for ordering
	}
}

// Annotations defines the annotations for the Medicine schema
func (Medicine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "medicines"},
	}
}
