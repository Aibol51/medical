// Code generated by ent, DO NOT EDIT.

package dictionarydetail

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dictionarydetail type in the database.
	Label = "dictionary_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldDictionaryID holds the string denoting the dictionary_id field in the database.
	FieldDictionaryID = "dictionary_id"
	// EdgeDictionaries holds the string denoting the dictionaries edge name in mutations.
	EdgeDictionaries = "dictionaries"
	// Table holds the table name of the dictionarydetail in the database.
	Table = "sys_dictionary_details"
	// DictionariesTable is the table that holds the dictionaries relation/edge.
	DictionariesTable = "sys_dictionary_details"
	// DictionariesInverseTable is the table name for the Dictionary entity.
	// It exists in this package in order to avoid circular dependency with the "dictionary" package.
	DictionariesInverseTable = "sys_dictionaries"
	// DictionariesColumn is the table column denoting the dictionaries relation/edge.
	DictionariesColumn = "dictionary_id"
)

// Columns holds all SQL columns for dictionarydetail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldSort,
	FieldTitle,
	FieldKey,
	FieldValue,
	FieldDictionaryID,
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
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort uint32
)

// OrderOption defines the ordering options for the DictionaryDetail queries.
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

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByDictionaryID orders the results by the dictionary_id field.
func ByDictionaryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDictionaryID, opts...).ToFunc()
}

// ByDictionariesField orders the results by dictionaries field.
func ByDictionariesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDictionariesStep(), sql.OrderByField(field, opts...))
	}
}
func newDictionariesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DictionariesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DictionariesTable, DictionariesColumn),
	)
}
