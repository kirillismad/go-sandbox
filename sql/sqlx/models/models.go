package models

type TableMetadata[Cols any] struct {
	TableName string
	Columns   Cols
}
