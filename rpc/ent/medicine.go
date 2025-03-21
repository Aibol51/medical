// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/medicine"
)

// Medicine is the model entity for the Medicine schema.
type Medicine struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Sort Number | 排序编号
	Sort uint32 `json:"sort,omitempty"`
	// Medicine chinese name | 药品中文名称
	NameZh string `json:"name_zh,omitempty"`
	// Medicine english name | 药品英文名称
	NameEn string `json:"name_en,omitempty"`
	// Medicine russian name | 药品俄语名称
	NameRu string `json:"name_ru,omitempty"`
	// Medicine kazakh name | 药品哈萨克语名称
	NameKk string `json:"name_kk,omitempty"`
	// Quantity in stock | 库存数量
	Quantity uint32 `json:"quantity,omitempty"`
	// Description chinese | 药品中文描述
	DescriptionZh string `json:"description_zh,omitempty"`
	// Description english | 药品英文描述
	DescriptionEn string `json:"description_en,omitempty"`
	// Description russian | 药品俄语描述
	DescriptionRu string `json:"description_ru,omitempty"`
	// Description kazakh | 药品哈萨克语描述
	DescriptionKk string `json:"description_kk,omitempty"`
	// Remarks | 备注信息
	Remarks string `json:"remarks,omitempty"`
	// Images | 图片路径
	Images       string `json:"images,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Medicine) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case medicine.FieldID, medicine.FieldStatus, medicine.FieldSort, medicine.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case medicine.FieldNameZh, medicine.FieldNameEn, medicine.FieldNameRu, medicine.FieldNameKk, medicine.FieldDescriptionZh, medicine.FieldDescriptionEn, medicine.FieldDescriptionRu, medicine.FieldDescriptionKk, medicine.FieldRemarks, medicine.FieldImages:
			values[i] = new(sql.NullString)
		case medicine.FieldCreatedAt, medicine.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Medicine fields.
func (m *Medicine) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case medicine.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case medicine.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case medicine.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case medicine.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = uint8(value.Int64)
			}
		case medicine.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				m.Sort = uint32(value.Int64)
			}
		case medicine.FieldNameZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_zh", values[i])
			} else if value.Valid {
				m.NameZh = value.String
			}
		case medicine.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				m.NameEn = value.String
			}
		case medicine.FieldNameRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ru", values[i])
			} else if value.Valid {
				m.NameRu = value.String
			}
		case medicine.FieldNameKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_kk", values[i])
			} else if value.Valid {
				m.NameKk = value.String
			}
		case medicine.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				m.Quantity = uint32(value.Int64)
			}
		case medicine.FieldDescriptionZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_zh", values[i])
			} else if value.Valid {
				m.DescriptionZh = value.String
			}
		case medicine.FieldDescriptionEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_en", values[i])
			} else if value.Valid {
				m.DescriptionEn = value.String
			}
		case medicine.FieldDescriptionRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_ru", values[i])
			} else if value.Valid {
				m.DescriptionRu = value.String
			}
		case medicine.FieldDescriptionKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_kk", values[i])
			} else if value.Valid {
				m.DescriptionKk = value.String
			}
		case medicine.FieldRemarks:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remarks", values[i])
			} else if value.Valid {
				m.Remarks = value.String
			}
		case medicine.FieldImages:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value.Valid {
				m.Images = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Medicine.
// This includes values selected through modifiers, order, etc.
func (m *Medicine) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Medicine.
// Note that you need to call Medicine.Unwrap() before calling this method if this Medicine
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Medicine) Update() *MedicineUpdateOne {
	return NewMedicineClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Medicine entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Medicine) Unwrap() *Medicine {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Medicine is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Medicine) String() string {
	var builder strings.Builder
	builder.WriteString("Medicine(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", m.Sort))
	builder.WriteString(", ")
	builder.WriteString("name_zh=")
	builder.WriteString(m.NameZh)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(m.NameEn)
	builder.WriteString(", ")
	builder.WriteString("name_ru=")
	builder.WriteString(m.NameRu)
	builder.WriteString(", ")
	builder.WriteString("name_kk=")
	builder.WriteString(m.NameKk)
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", m.Quantity))
	builder.WriteString(", ")
	builder.WriteString("description_zh=")
	builder.WriteString(m.DescriptionZh)
	builder.WriteString(", ")
	builder.WriteString("description_en=")
	builder.WriteString(m.DescriptionEn)
	builder.WriteString(", ")
	builder.WriteString("description_ru=")
	builder.WriteString(m.DescriptionRu)
	builder.WriteString(", ")
	builder.WriteString("description_kk=")
	builder.WriteString(m.DescriptionKk)
	builder.WriteString(", ")
	builder.WriteString("remarks=")
	builder.WriteString(m.Remarks)
	builder.WriteString(", ")
	builder.WriteString("images=")
	builder.WriteString(m.Images)
	builder.WriteByte(')')
	return builder.String()
}

// Medicines is a parsable slice of Medicine.
type Medicines []*Medicine
