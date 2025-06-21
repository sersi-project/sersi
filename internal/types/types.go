package types

import (
	"github.com/sersi-project/sersi/pkg"
)

type APIAuth struct {
	UserID       string `json:"user_id"`
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type APIAuthResponse struct {
	Message string  `json:"message"`
	Data    APIAuth `json:"data"`
}

type APIScaffoldsResponse struct {
	Message string                `json:"message"`
	Data    []GetScaffoldResponse `json:"data"`
}

type SaveScaffoldRequest struct {
	Name     string             `json:"name"`
	UserID   string             `json:"user_id"`
	Scaffold pkg.ScaffoldConfig `json:"scaffold"`
}

type ScaffoldRequest struct {
	Name     string             `json:"name"`
	Scaffold pkg.ScaffoldConfig `json:"scaffold"`
}

type GetScaffoldResponse struct {
	ID        int64              `json:"id"`
	Name      string             `json:"name"`
	Scaffold  pkg.ScaffoldConfig `json:"scaffold"`
	UserID    string             `json:"user_id"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
}
