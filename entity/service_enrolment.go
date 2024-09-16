package entity

import "time"

type ServiceEnrollment struct {
	Id        int
	Name      string
	Unit      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
