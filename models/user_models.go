package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	RoleSuperAdmin = "superadmin"
	RoleAdmin      = "admin"
	RoleSales      = "sales"
	RoleSupport    = "support"
	RoleCustomer   = "customer"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"` // Excluded from JSON responses
	Role      string         `gorm:"size:50;not null;index" json:"role"`
	Status    string         `gorm:"size:20;not null;default:'active'" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete support
}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=superadmin admin sales support customer"`
	Status   string `json:"status" binding:"omitempty,oneof=active inactive"`
}


