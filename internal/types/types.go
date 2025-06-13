package structs

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

type SaveScaffoldRequest struct {
	Name     string             `json:"name"`
	UserID   string             `json:"user_id"`
	Scaffold pkg.ScaffoldConfig `json:"scaffold"`
}

type ScaffoldRequest struct {
	Name     string             `json:"name"`
	Scaffold pkg.ScaffoldConfig `json:"scaffold"`
}
