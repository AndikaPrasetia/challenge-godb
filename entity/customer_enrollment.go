package entity

import "time"

type CustomerEnrollment struct {
	Id        int
	Name      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
