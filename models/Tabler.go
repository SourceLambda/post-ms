package models

// this interface enables changing the tablename (in plural by default) for all models
type Tabler interface {
	TableName() string
}