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

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldName, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v uint32) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldQuantity, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescription, v))
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

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldName, v))
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

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Medicine {
	return predicate.Medicine(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Medicine {
	return predicate.Medicine(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Medicine {
	return predicate.Medicine(sql.FieldContainsFold(FieldDescription, v))
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
