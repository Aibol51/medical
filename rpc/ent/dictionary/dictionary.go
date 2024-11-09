// Code generated by ent, DO NOT EDIT.

package dictionary

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dictionary type in the database.
	Label = "dictionary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeDictionaryDetails holds the string denoting the dictionary_details edge name in mutations.
	EdgeDictionaryDetails = "dictionary_details"
	// Table holds the table name of the dictionary in the database.
	Table = "sys_dictionaries"
	// DictionaryDetailsTable is the table that holds the dictionary_details relation/edge.
	DictionaryDetailsTable = "sys_dictionary_details"
	// DictionaryDetailsInverseTable is the table name for the DictionaryDetail entity.
	// It exists in this package in order to avoid circular dependency with the "dictionarydetail" package.
	DictionaryDetailsInverseTable = "sys_dictionary_details"
	// DictionaryDetailsColumn is the table column denoting the dictionary_details relation/edge.
	DictionaryDetailsColumn = "dictionary_id"
)

// Columns holds all SQL columns for dictionary fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldTitle,
	FieldName,
	FieldDesc,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
)

// OrderOption defines the ordering options for the Dictionary queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByDictionaryDetailsCount orders the results by dictionary_details count.
func ByDictionaryDetailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDictionaryDetailsStep(), opts...)
	}
}

// ByDictionaryDetails orders the results by dictionary_details terms.
func ByDictionaryDetails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDictionaryDetailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDictionaryDetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DictionaryDetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DictionaryDetailsTable, DictionaryDetailsColumn),
	)
}
