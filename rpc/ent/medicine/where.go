// Code generated by ent, DO NOT EDIT.

package medicine

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldSort, v))
}

// NameZh applies equality check predicate on the "name_zh" field. It's identical to NameZhEQ.
func NameZh(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameZh, v))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameEn, v))
}

// NameRu applies equality check predicate on the "name_ru" field. It's identical to NameRuEQ.
func NameRu(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameRu, v))
}

// NameKk applies equality check predicate on the "name_kk" field. It's identical to NameKkEQ.
func NameKk(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameKk, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldQuantity, v))
}

// DescriptionZh applies equality check predicate on the "description_zh" field. It's identical to DescriptionZhEQ.
func DescriptionZh(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionZh, v))
}

// DescriptionEn applies equality check predicate on the "description_en" field. It's identical to DescriptionEnEQ.
func DescriptionEn(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionEn, v))
}

// DescriptionRu applies equality check predicate on the "description_ru" field. It's identical to DescriptionRuEQ.
func DescriptionRu(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionRu, v))
}

// DescriptionKk applies equality check predicate on the "description_kk" field. It's identical to DescriptionKkEQ.
func DescriptionKk(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionKk, v))
}

// Remarks applies equality check predicate on the "remarks" field. It's identical to RemarksEQ.
func Remarks(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldRemarks, v))
}

// Images applies equality check predicate on the "images" field. It's identical to ImagesEQ.
func Images(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldImages, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldStatus))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldSort, v))
}

// NameZhEQ applies the EQ predicate on the "name_zh" field.
func NameZhEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameZh, v))
}

// NameZhNEQ applies the NEQ predicate on the "name_zh" field.
func NameZhNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldNameZh, v))
}

// NameZhIn applies the In predicate on the "name_zh" field.
func NameZhIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldNameZh, vs...))
}

// NameZhNotIn applies the NotIn predicate on the "name_zh" field.
func NameZhNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldNameZh, vs...))
}

// NameZhGT applies the GT predicate on the "name_zh" field.
func NameZhGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldNameZh, v))
}

// NameZhGTE applies the GTE predicate on the "name_zh" field.
func NameZhGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldNameZh, v))
}

// NameZhLT applies the LT predicate on the "name_zh" field.
func NameZhLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldNameZh, v))
}

// NameZhLTE applies the LTE predicate on the "name_zh" field.
func NameZhLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldNameZh, v))
}

// NameZhContains applies the Contains predicate on the "name_zh" field.
func NameZhContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldNameZh, v))
}

// NameZhHasPrefix applies the HasPrefix predicate on the "name_zh" field.
func NameZhHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldNameZh, v))
}

// NameZhHasSuffix applies the HasSuffix predicate on the "name_zh" field.
func NameZhHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldNameZh, v))
}

// NameZhEqualFold applies the EqualFold predicate on the "name_zh" field.
func NameZhEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldNameZh, v))
}

// NameZhContainsFold applies the ContainsFold predicate on the "name_zh" field.
func NameZhContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldNameZh, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldNameEn, v))
}

// NameRuEQ applies the EQ predicate on the "name_ru" field.
func NameRuEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameRu, v))
}

// NameRuNEQ applies the NEQ predicate on the "name_ru" field.
func NameRuNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldNameRu, v))
}

// NameRuIn applies the In predicate on the "name_ru" field.
func NameRuIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldNameRu, vs...))
}

// NameRuNotIn applies the NotIn predicate on the "name_ru" field.
func NameRuNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldNameRu, vs...))
}

// NameRuGT applies the GT predicate on the "name_ru" field.
func NameRuGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldNameRu, v))
}

// NameRuGTE applies the GTE predicate on the "name_ru" field.
func NameRuGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldNameRu, v))
}

// NameRuLT applies the LT predicate on the "name_ru" field.
func NameRuLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldNameRu, v))
}

// NameRuLTE applies the LTE predicate on the "name_ru" field.
func NameRuLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldNameRu, v))
}

// NameRuContains applies the Contains predicate on the "name_ru" field.
func NameRuContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldNameRu, v))
}

// NameRuHasPrefix applies the HasPrefix predicate on the "name_ru" field.
func NameRuHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldNameRu, v))
}

// NameRuHasSuffix applies the HasSuffix predicate on the "name_ru" field.
func NameRuHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldNameRu, v))
}

// NameRuEqualFold applies the EqualFold predicate on the "name_ru" field.
func NameRuEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldNameRu, v))
}

// NameRuContainsFold applies the ContainsFold predicate on the "name_ru" field.
func NameRuContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldNameRu, v))
}

// NameKkEQ applies the EQ predicate on the "name_kk" field.
func NameKkEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldNameKk, v))
}

// NameKkNEQ applies the NEQ predicate on the "name_kk" field.
func NameKkNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldNameKk, v))
}

// NameKkIn applies the In predicate on the "name_kk" field.
func NameKkIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldNameKk, vs...))
}

// NameKkNotIn applies the NotIn predicate on the "name_kk" field.
func NameKkNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldNameKk, vs...))
}

// NameKkGT applies the GT predicate on the "name_kk" field.
func NameKkGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldNameKk, v))
}

// NameKkGTE applies the GTE predicate on the "name_kk" field.
func NameKkGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldNameKk, v))
}

// NameKkLT applies the LT predicate on the "name_kk" field.
func NameKkLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldNameKk, v))
}

// NameKkLTE applies the LTE predicate on the "name_kk" field.
func NameKkLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldNameKk, v))
}

// NameKkContains applies the Contains predicate on the "name_kk" field.
func NameKkContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldNameKk, v))
}

// NameKkHasPrefix applies the HasPrefix predicate on the "name_kk" field.
func NameKkHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldNameKk, v))
}

// NameKkHasSuffix applies the HasSuffix predicate on the "name_kk" field.
func NameKkHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldNameKk, v))
}

// NameKkEqualFold applies the EqualFold predicate on the "name_kk" field.
func NameKkEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldNameKk, v))
}

// NameKkContainsFold applies the ContainsFold predicate on the "name_kk" field.
func NameKkContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldNameKk, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldQuantity, v))
}

// DescriptionZhEQ applies the EQ predicate on the "description_zh" field.
func DescriptionZhEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionZh, v))
}

// DescriptionZhNEQ applies the NEQ predicate on the "description_zh" field.
func DescriptionZhNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldDescriptionZh, v))
}

// DescriptionZhIn applies the In predicate on the "description_zh" field.
func DescriptionZhIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldDescriptionZh, vs...))
}

// DescriptionZhNotIn applies the NotIn predicate on the "description_zh" field.
func DescriptionZhNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldDescriptionZh, vs...))
}

// DescriptionZhGT applies the GT predicate on the "description_zh" field.
func DescriptionZhGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldDescriptionZh, v))
}

// DescriptionZhGTE applies the GTE predicate on the "description_zh" field.
func DescriptionZhGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldDescriptionZh, v))
}

// DescriptionZhLT applies the LT predicate on the "description_zh" field.
func DescriptionZhLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldDescriptionZh, v))
}

// DescriptionZhLTE applies the LTE predicate on the "description_zh" field.
func DescriptionZhLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldDescriptionZh, v))
}

// DescriptionZhContains applies the Contains predicate on the "description_zh" field.
func DescriptionZhContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldDescriptionZh, v))
}

// DescriptionZhHasPrefix applies the HasPrefix predicate on the "description_zh" field.
func DescriptionZhHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldDescriptionZh, v))
}

// DescriptionZhHasSuffix applies the HasSuffix predicate on the "description_zh" field.
func DescriptionZhHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldDescriptionZh, v))
}

// DescriptionZhIsNil applies the IsNil predicate on the "description_zh" field.
func DescriptionZhIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldDescriptionZh))
}

// DescriptionZhNotNil applies the NotNil predicate on the "description_zh" field.
func DescriptionZhNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldDescriptionZh))
}

// DescriptionZhEqualFold applies the EqualFold predicate on the "description_zh" field.
func DescriptionZhEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldDescriptionZh, v))
}

// DescriptionZhContainsFold applies the ContainsFold predicate on the "description_zh" field.
func DescriptionZhContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldDescriptionZh, v))
}

// DescriptionEnEQ applies the EQ predicate on the "description_en" field.
func DescriptionEnEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionEn, v))
}

// DescriptionEnNEQ applies the NEQ predicate on the "description_en" field.
func DescriptionEnNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldDescriptionEn, v))
}

// DescriptionEnIn applies the In predicate on the "description_en" field.
func DescriptionEnIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldDescriptionEn, vs...))
}

// DescriptionEnNotIn applies the NotIn predicate on the "description_en" field.
func DescriptionEnNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldDescriptionEn, vs...))
}

// DescriptionEnGT applies the GT predicate on the "description_en" field.
func DescriptionEnGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldDescriptionEn, v))
}

// DescriptionEnGTE applies the GTE predicate on the "description_en" field.
func DescriptionEnGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldDescriptionEn, v))
}

// DescriptionEnLT applies the LT predicate on the "description_en" field.
func DescriptionEnLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldDescriptionEn, v))
}

// DescriptionEnLTE applies the LTE predicate on the "description_en" field.
func DescriptionEnLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldDescriptionEn, v))
}

// DescriptionEnContains applies the Contains predicate on the "description_en" field.
func DescriptionEnContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldDescriptionEn, v))
}

// DescriptionEnHasPrefix applies the HasPrefix predicate on the "description_en" field.
func DescriptionEnHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldDescriptionEn, v))
}

// DescriptionEnHasSuffix applies the HasSuffix predicate on the "description_en" field.
func DescriptionEnHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldDescriptionEn, v))
}

// DescriptionEnIsNil applies the IsNil predicate on the "description_en" field.
func DescriptionEnIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldDescriptionEn))
}

// DescriptionEnNotNil applies the NotNil predicate on the "description_en" field.
func DescriptionEnNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldDescriptionEn))
}

// DescriptionEnEqualFold applies the EqualFold predicate on the "description_en" field.
func DescriptionEnEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldDescriptionEn, v))
}

// DescriptionEnContainsFold applies the ContainsFold predicate on the "description_en" field.
func DescriptionEnContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldDescriptionEn, v))
}

// DescriptionRuEQ applies the EQ predicate on the "description_ru" field.
func DescriptionRuEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionRu, v))
}

// DescriptionRuNEQ applies the NEQ predicate on the "description_ru" field.
func DescriptionRuNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldDescriptionRu, v))
}

// DescriptionRuIn applies the In predicate on the "description_ru" field.
func DescriptionRuIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldDescriptionRu, vs...))
}

// DescriptionRuNotIn applies the NotIn predicate on the "description_ru" field.
func DescriptionRuNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldDescriptionRu, vs...))
}

// DescriptionRuGT applies the GT predicate on the "description_ru" field.
func DescriptionRuGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldDescriptionRu, v))
}

// DescriptionRuGTE applies the GTE predicate on the "description_ru" field.
func DescriptionRuGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldDescriptionRu, v))
}

// DescriptionRuLT applies the LT predicate on the "description_ru" field.
func DescriptionRuLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldDescriptionRu, v))
}

// DescriptionRuLTE applies the LTE predicate on the "description_ru" field.
func DescriptionRuLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldDescriptionRu, v))
}

// DescriptionRuContains applies the Contains predicate on the "description_ru" field.
func DescriptionRuContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldDescriptionRu, v))
}

// DescriptionRuHasPrefix applies the HasPrefix predicate on the "description_ru" field.
func DescriptionRuHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldDescriptionRu, v))
}

// DescriptionRuHasSuffix applies the HasSuffix predicate on the "description_ru" field.
func DescriptionRuHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldDescriptionRu, v))
}

// DescriptionRuIsNil applies the IsNil predicate on the "description_ru" field.
func DescriptionRuIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldDescriptionRu))
}

// DescriptionRuNotNil applies the NotNil predicate on the "description_ru" field.
func DescriptionRuNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldDescriptionRu))
}

// DescriptionRuEqualFold applies the EqualFold predicate on the "description_ru" field.
func DescriptionRuEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldDescriptionRu, v))
}

// DescriptionRuContainsFold applies the ContainsFold predicate on the "description_ru" field.
func DescriptionRuContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldDescriptionRu, v))
}

// DescriptionKkEQ applies the EQ predicate on the "description_kk" field.
func DescriptionKkEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescriptionKk, v))
}

// DescriptionKkNEQ applies the NEQ predicate on the "description_kk" field.
func DescriptionKkNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldDescriptionKk, v))
}

// DescriptionKkIn applies the In predicate on the "description_kk" field.
func DescriptionKkIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldDescriptionKk, vs...))
}

// DescriptionKkNotIn applies the NotIn predicate on the "description_kk" field.
func DescriptionKkNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldDescriptionKk, vs...))
}

// DescriptionKkGT applies the GT predicate on the "description_kk" field.
func DescriptionKkGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldDescriptionKk, v))
}

// DescriptionKkGTE applies the GTE predicate on the "description_kk" field.
func DescriptionKkGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldDescriptionKk, v))
}

// DescriptionKkLT applies the LT predicate on the "description_kk" field.
func DescriptionKkLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldDescriptionKk, v))
}

// DescriptionKkLTE applies the LTE predicate on the "description_kk" field.
func DescriptionKkLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldDescriptionKk, v))
}

// DescriptionKkContains applies the Contains predicate on the "description_kk" field.
func DescriptionKkContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldDescriptionKk, v))
}

// DescriptionKkHasPrefix applies the HasPrefix predicate on the "description_kk" field.
func DescriptionKkHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldDescriptionKk, v))
}

// DescriptionKkHasSuffix applies the HasSuffix predicate on the "description_kk" field.
func DescriptionKkHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldDescriptionKk, v))
}

// DescriptionKkIsNil applies the IsNil predicate on the "description_kk" field.
func DescriptionKkIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldDescriptionKk))
}

// DescriptionKkNotNil applies the NotNil predicate on the "description_kk" field.
func DescriptionKkNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldDescriptionKk))
}

// DescriptionKkEqualFold applies the EqualFold predicate on the "description_kk" field.
func DescriptionKkEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldDescriptionKk, v))
}

// DescriptionKkContainsFold applies the ContainsFold predicate on the "description_kk" field.
func DescriptionKkContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldDescriptionKk, v))
}

// RemarksEQ applies the EQ predicate on the "remarks" field.
func RemarksEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldRemarks, v))
}

// RemarksNEQ applies the NEQ predicate on the "remarks" field.
func RemarksNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldRemarks, v))
}

// RemarksIn applies the In predicate on the "remarks" field.
func RemarksIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldRemarks, vs...))
}

// RemarksNotIn applies the NotIn predicate on the "remarks" field.
func RemarksNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldRemarks, vs...))
}

// RemarksGT applies the GT predicate on the "remarks" field.
func RemarksGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldRemarks, v))
}

// RemarksGTE applies the GTE predicate on the "remarks" field.
func RemarksGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldRemarks, v))
}

// RemarksLT applies the LT predicate on the "remarks" field.
func RemarksLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldRemarks, v))
}

// RemarksLTE applies the LTE predicate on the "remarks" field.
func RemarksLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldRemarks, v))
}

// RemarksContains applies the Contains predicate on the "remarks" field.
func RemarksContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldRemarks, v))
}

// RemarksHasPrefix applies the HasPrefix predicate on the "remarks" field.
func RemarksHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldRemarks, v))
}

// RemarksHasSuffix applies the HasSuffix predicate on the "remarks" field.
func RemarksHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldRemarks, v))
}

// RemarksIsNil applies the IsNil predicate on the "remarks" field.
func RemarksIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldRemarks))
}

// RemarksNotNil applies the NotNil predicate on the "remarks" field.
func RemarksNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldRemarks))
}

// RemarksEqualFold applies the EqualFold predicate on the "remarks" field.
func RemarksEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldRemarks, v))
}

// RemarksContainsFold applies the ContainsFold predicate on the "remarks" field.
func RemarksContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldRemarks, v))
}

// ImagesEQ applies the EQ predicate on the "images" field.
func ImagesEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldImages, v))
}

// ImagesNEQ applies the NEQ predicate on the "images" field.
func ImagesNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldImages, v))
}

// ImagesIn applies the In predicate on the "images" field.
func ImagesIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldImages, vs...))
}

// ImagesNotIn applies the NotIn predicate on the "images" field.
func ImagesNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldImages, vs...))
}

// ImagesGT applies the GT predicate on the "images" field.
func ImagesGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldImages, v))
}

// ImagesGTE applies the GTE predicate on the "images" field.
func ImagesGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldImages, v))
}

// ImagesLT applies the LT predicate on the "images" field.
func ImagesLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldImages, v))
}

// ImagesLTE applies the LTE predicate on the "images" field.
func ImagesLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldImages, v))
}

// ImagesContains applies the Contains predicate on the "images" field.
func ImagesContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldImages, v))
}

// ImagesHasPrefix applies the HasPrefix predicate on the "images" field.
func ImagesHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldImages, v))
}

// ImagesHasSuffix applies the HasSuffix predicate on the "images" field.
func ImagesHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldImages, v))
}

// ImagesIsNil applies the IsNil predicate on the "images" field.
func ImagesIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldImages))
}

// ImagesNotNil applies the NotNil predicate on the "images" field.
func ImagesNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldImages))
}

// ImagesEqualFold applies the EqualFold predicate on the "images" field.
func ImagesEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldImages, v))
}

// ImagesContainsFold applies the ContainsFold predicate on the "images" field.
func ImagesContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldImages, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Medicine) predicate.Medicine {
	return predicate.Medicine(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Medicine) predicate.Medicine {
	return predicate.Medicine(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Medicine) predicate.Medicine {
	return predicate.Medicine(sql.NotPredicates(p))
}
