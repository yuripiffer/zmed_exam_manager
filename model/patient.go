package model

import (
	"cloud.google.com/go/civil"
	"time"
)

type Patient struct {
	Id                  string                 `json:"id"`
	Document            string                 `json:"document"`
	Name                string                 `json:"name"`
	BirthDate           *civil.Date            `json:"birth_date"`
	Status              PatientStatus          `json:"status"`
	CreatedAt           *time.Time             `json:"created_at,omitempty"`
	UpdatedAt           *time.Time             `json:"updated_at,omitempty"`
	DeletedAt           *time.Time             `json:"deleted_at"`
	Cellphone           string                 `json:"cellphone"`
	CellphoneVerifiedAt *time.Time             `json:"cellphone_verified_at,omitempty"`
	Email               string                 `json:"email"`
	EmailVerifiedAt     *time.Time             `json:"email_verified_at,omitempty"`
	Data                map[string]interface{} `json:"data,omitempty"`
}

type PatientStatus string

const (
	StatusPending  PatientStatus = "pending"
	StatusActive                 = "active"
	StatusInactive               = "inactive"
)
