// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/swiper"
)

// Swiper is the model entity for the Swiper schema.
type Swiper struct {
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
	// Chinese banner image | 中文轮播图
	BannerZh string `json:"banner_zh,omitempty"`
	// English banner image | 英文轮播图
	BannerEn string `json:"banner_en,omitempty"`
	// Russian banner image | 俄语轮播图
	BannerRu string `json:"banner_ru,omitempty"`
	// Kazakh banner image | 哈萨克语轮播图
	BannerKk string `json:"banner_kk,omitempty"`
	// Chinese content | 中文内容
	ContentZh string `json:"content_zh,omitempty"`
	// English content | 英文内容
	ContentEn string `json:"content_en,omitempty"`
	// Russian content | 俄语内容
	ContentRu string `json:"content_ru,omitempty"`
	// Kazakh content | 哈萨克语内容
	ContentKk string `json:"content_kk,omitempty"`
	// Jump Url | 跳转地址
	JumpURL      string `json:"jump_url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Swiper) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case swiper.FieldID, swiper.FieldStatus, swiper.FieldSort:
			values[i] = new(sql.NullInt64)
		case swiper.FieldTitleZh, swiper.FieldTitleEn, swiper.FieldTitleRu, swiper.FieldTitleKk, swiper.FieldBannerZh, swiper.FieldBannerEn, swiper.FieldBannerRu, swiper.FieldBannerKk, swiper.FieldContentZh, swiper.FieldContentEn, swiper.FieldContentRu, swiper.FieldContentKk, swiper.FieldJumpURL:
			values[i] = new(sql.NullString)
		case swiper.FieldCreatedAt, swiper.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Swiper fields.
func (s *Swiper) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case swiper.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case swiper.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case swiper.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case swiper.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = uint8(value.Int64)
			}
		case swiper.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				s.Sort = uint32(value.Int64)
			}
		case swiper.FieldTitleZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_zh", values[i])
			} else if value.Valid {
				s.TitleZh = value.String
			}
		case swiper.FieldTitleEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_en", values[i])
			} else if value.Valid {
				s.TitleEn = value.String
			}
		case swiper.FieldTitleRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_ru", values[i])
			} else if value.Valid {
				s.TitleRu = value.String
			}
		case swiper.FieldTitleKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_kk", values[i])
			} else if value.Valid {
				s.TitleKk = value.String
			}
		case swiper.FieldBannerZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_zh", values[i])
			} else if value.Valid {
				s.BannerZh = value.String
			}
		case swiper.FieldBannerEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_en", values[i])
			} else if value.Valid {
				s.BannerEn = value.String
			}
		case swiper.FieldBannerRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_ru", values[i])
			} else if value.Valid {
				s.BannerRu = value.String
			}
		case swiper.FieldBannerKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner_kk", values[i])
			} else if value.Valid {
				s.BannerKk = value.String
			}
		case swiper.FieldContentZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_zh", values[i])
			} else if value.Valid {
				s.ContentZh = value.String
			}
		case swiper.FieldContentEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_en", values[i])
			} else if value.Valid {
				s.ContentEn = value.String
			}
		case swiper.FieldContentRu:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_ru", values[i])
			} else if value.Valid {
				s.ContentRu = value.String
			}
		case swiper.FieldContentKk:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_kk", values[i])
			} else if value.Valid {
				s.ContentKk = value.String
			}
		case swiper.FieldJumpURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field jump_url", values[i])
			} else if value.Valid {
				s.JumpURL = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Swiper.
// This includes values selected through modifiers, order, etc.
func (s *Swiper) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Swiper.
// Note that you need to call Swiper.Unwrap() before calling this method if this Swiper
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Swiper) Update() *SwiperUpdateOne {
	return NewSwiperClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Swiper entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Swiper) Unwrap() *Swiper {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Swiper is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Swiper) String() string {
	var builder strings.Builder
	builder.WriteString("Swiper(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", s.Sort))
	builder.WriteString(", ")
	builder.WriteString("title_zh=")
	builder.WriteString(s.TitleZh)
	builder.WriteString(", ")
	builder.WriteString("title_en=")
	builder.WriteString(s.TitleEn)
	builder.WriteString(", ")
	builder.WriteString("title_ru=")
	builder.WriteString(s.TitleRu)
	builder.WriteString(", ")
	builder.WriteString("title_kk=")
	builder.WriteString(s.TitleKk)
	builder.WriteString(", ")
	builder.WriteString("banner_zh=")
	builder.WriteString(s.BannerZh)
	builder.WriteString(", ")
	builder.WriteString("banner_en=")
	builder.WriteString(s.BannerEn)
	builder.WriteString(", ")
	builder.WriteString("banner_ru=")
	builder.WriteString(s.BannerRu)
	builder.WriteString(", ")
	builder.WriteString("banner_kk=")
	builder.WriteString(s.BannerKk)
	builder.WriteString(", ")
	builder.WriteString("content_zh=")
	builder.WriteString(s.ContentZh)
	builder.WriteString(", ")
	builder.WriteString("content_en=")
	builder.WriteString(s.ContentEn)
	builder.WriteString(", ")
	builder.WriteString("content_ru=")
	builder.WriteString(s.ContentRu)
	builder.WriteString(", ")
	builder.WriteString("content_kk=")
	builder.WriteString(s.ContentKk)
	builder.WriteString(", ")
	builder.WriteString("jump_url=")
	builder.WriteString(s.JumpURL)
	builder.WriteByte(')')
	return builder.String()
}

// Swipers is a parsable slice of Swiper.
type Swipers []*Swiper
