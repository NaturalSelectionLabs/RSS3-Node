// Code generated by ent, DO NOT EDIT.

package feed

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the feed type in the database.
	Label = "feed"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChain holds the string denoting the chain field in the database.
	FieldChain = "chain"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldTotalActions holds the string denoting the total_actions field in the database.
	FieldTotalActions = "total_actions"
	// FieldActions holds the string denoting the actions field in the database.
	FieldActions = "actions"
	// FieldFeeValue holds the string denoting the fee_value field in the database.
	FieldFeeValue = "fee_value"
	// FieldFeeToken holds the string denoting the fee_token field in the database.
	FieldFeeToken = "fee_token"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the feed in the database.
	Table = "feeds"
)

// Columns holds all SQL columns for feed fields.
var Columns = []string{
	FieldID,
	FieldChain,
	FieldPlatform,
	FieldFrom,
	FieldTo,
	FieldTag,
	FieldType,
	FieldStatus,
	FieldIndex,
	FieldTotalActions,
	FieldActions,
	FieldFeeValue,
	FieldFeeToken,
	FieldTimestamp,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// ChainValidator is a validator for the "chain" field. It is called by the builders before save.
	ChainValidator func(string) error
	// FromValidator is a validator for the "from" field. It is called by the builders before save.
	FromValidator func(string) error
	// ToValidator is a validator for the "to" field. It is called by the builders before save.
	ToValidator func(string) error
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
	// DefaultIndex holds the default value on creation for the "index" field.
	DefaultIndex uint
	// DefaultTotalActions holds the default value on creation for the "total_actions" field.
	DefaultTotalActions uint
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Feed queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChain orders the results by the chain field.
func ByChain(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChain, opts...).ToFunc()
}

// ByPlatform orders the results by the platform field.
func ByPlatform(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatform, opts...).ToFunc()
}

// ByFrom orders the results by the from field.
func ByFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrom, opts...).ToFunc()
}

// ByTo orders the results by the to field.
func ByTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTo, opts...).ToFunc()
}

// ByTag orders the results by the tag field.
func ByTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTag, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}

// ByTotalActions orders the results by the total_actions field.
func ByTotalActions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalActions, opts...).ToFunc()
}

// ByFeeValue orders the results by the fee_value field.
func ByFeeValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeeValue, opts...).ToFunc()
}

// ByFeeToken orders the results by the fee_token field.
func ByFeeToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeeToken, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
