package entity

import (
	"database/sql"
	"time"
)

type OrderEnrollment struct {
	Id             int
	CustomerId     int
	OrderDate      time.Time
	CompletionDate sql.NullTime
	ReceivedBy     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
