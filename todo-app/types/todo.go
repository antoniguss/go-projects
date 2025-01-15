package types

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID            int
	Description   string
	TimeAdded     time.Time
	TimeCompleted sql.NullTime
	IsCompleted   bool
}
