package domain

import "time"

type Server struct {
	// gorm.Model
	ID uint `json:"server_id" gorm:"column:server_id; primarykey"`

	Server_Name   string `json:"server_name" gorm:"column:server_name; not null; unique"`
	Server_IPv4   string `json:"server_ipv4" gorm:"column:server_ipv4; not null; unique"`
	Server_Status string `json:"server_status" gorm:"column:server_status; not null"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; not null;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at; not null;"`
}
