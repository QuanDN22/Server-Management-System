package domain

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Server_Name   string `json:"server_name" gorm:"column:server_name; not null; unique"`
	Server_IPv4   string `json:"server_ipv4" gorm:"column:server_ipv4; not null; unique"`
	Server_Status bool   `json:"server_status" gorm:"column:server_status; not null"`
}
