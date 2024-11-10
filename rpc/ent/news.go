// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/news"
)

// News is the model entity for the News schema.
type News struct {
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
	// Chinese title | 中文标题
	TitleZh string `json:"title_zh,omitempty"`
	// English title | 英文标题
	TitleEn string `json:"title_en,omitempty"`
	// Russian title | 俄语标题
	TitleRu string `json:"title_ru,omitempty"`
	// Kazakh title | 哈萨克语标题
	TitleKk string `json:"title_kk,omitempty"`
	// Chinese content | 中文内容
	ContentZh string `json:"content_zh,omitempty"`
	// English content | 英文内容
	ContentEn string `json:"content_en,omitempty"`
	// Russian content | 俄语内容
	ContentRu string `json:"content_ru,omitempty"`
	// Kazakh content | 哈萨克语内容
	ContentKk string `json:"content_kk,omitempty"`
	// Cover image URL | 封面图片URL
	CoverURL     string `json:"cover_url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*News) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case news.FieldID, news.FieldStatus, news.FieldSort:
			values[i] = new(sql.NullInt64)
		case news.FieldTitleZh, news.FieldTitleEn, news.FieldTitleRu, news.FieldTitleKk, news.FieldContentZh, news.FieldContentEn, news.FieldContentRu, news.FieldContentKk, news.FieldCoverURL:
			values[i] = new(sql.NullString)
		case news.FieldCreatedAt, news.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the News fields.
func (n *News) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case news.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = uint64(value.Int64)
		case news.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case news.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case news.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				n.Status = uint8(value.Int64)
			}
		case news.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				n.Sort = uint32(value.Int64)
			}
		case news.FieldTitleZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_zh", values[i])
			} else if value.Valid {
				n.TitleZh = value.String
			}
		case news.FieldTitleEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_en", values[i])
			} else if value.Valid {
				n.TitleEn = value.String
			}
		case news.FieldTitleRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_ru", values[i])
			} else if value.Valid {
				n.TitleRu = value.String
			}
		case news.FieldTitleKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_kk", values[i])
			} else if value.Valid {
				n.TitleKk = value.String
			}
		case news.FieldContentZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_zh", values[i])
			} else if value.Valid {
				n.ContentZh = value.String
			}
		case news.FieldContentEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_en", values[i])
			} else if value.Valid {
				n.ContentEn = value.String
			}
		case news.FieldContentRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_ru", values[i])
			} else if value.Valid {
				n.ContentRu = value.String
			}
		case news.FieldContentKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_kk", values[i])
			} else if value.Valid {
				n.ContentKk = value.String
			}
		case news.FieldCoverURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_url", values[i])
			} else if value.Valid {
				n.CoverURL = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the News.
// This includes values selected through modifiers, order, etc.
func (n *News) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// Update returns a builder for updating this News.
// Note that you need to call News.Unwrap() before calling this method if this News
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *News) Update() *NewsUpdateOne {
	return NewNewsClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the News entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *News) Unwrap() *News {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: News is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *News) String() string {
	var builder strings.Builder
	builder.WriteString("News(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", n.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", n.Sort))
	builder.WriteString(", ")
	builder.WriteString("title_zh=")
	builder.WriteString(n.TitleZh)
	builder.WriteString(", ")
	builder.WriteString("title_en=")
	builder.WriteString(n.TitleEn)
	builder.WriteString(", ")
	builder.WriteString("title_ru=")
	builder.WriteString(n.TitleRu)
	builder.WriteString(", ")
	builder.WriteString("title_kk=")
	builder.WriteString(n.TitleKk)
	builder.WriteString(", ")
	builder.WriteString("content_zh=")
	builder.WriteString(n.ContentZh)
	builder.WriteString(", ")
	builder.WriteString("content_en=")
	builder.WriteString(n.ContentEn)
	builder.WriteString(", ")
	builder.WriteString("content_ru=")
	builder.WriteString(n.ContentRu)
	builder.WriteString(", ")
	builder.WriteString("content_kk=")
	builder.WriteString(n.ContentKk)
	builder.WriteString(", ")
	builder.WriteString("cover_url=")
	builder.WriteString(n.CoverURL)
	builder.WriteByte(')')
	return builder.String()
}

// NewsSlice is a parsable slice of News.
type NewsSlice []*News
