package storage

type StorageType string

const (
	SQLite StorageType = "sqlite"
	CSV    StorageType = "csv"
)

func NewStorage(storageType StorageType, path string) (Storage, error) {
	switch storageType {
	case SQLite:
		return NewSQLiteStorage(path)
	case CSV:
		return NewCSVStorage(path)
	default:
		return NewSQLiteStorage(path)
	}
}
