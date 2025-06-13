package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	types "github.com/sersi-project/sersi/internal/types"
)

type API struct {
	Auth    types.APIAuth
	client  *http.Client
	baseURL string
}

func NewAPI() *API {
	url := os.Getenv("SERSI_API_URL")
	return &API{
		client:  &http.Client{},
		baseURL: url,
	}
}

func (a *API) Authenticate(email, password string) (*types.APIAuth, error) {
	body, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	req, err := http.NewRequest("POST", a.baseURL+"/authenticate", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := a.client.Do(req)

	defer req.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error logging in: %s", res.Status)
	}

	var auth types.APIAuthResponse
	respBody, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(respBody, &auth)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling auth config: %v", err)
	}
	return &auth.Data, nil
}

func (a *API) GetAllScaffolds() (*http.Response, error) {
	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", a.baseURL+"/scaffolds/"+config.UserID, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	res, err := a.client.Do(req)
	defer req.Body.Close() //nolint

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting all scaffolds: %s", res.Status)
	}
	return res, nil
}

func (a *API) GetScaffold(name string) (*http.Response, error) {
	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", a.baseURL+"/scaffolds/"+config.UserID+"/"+name, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	res, err := a.client.Do(req)
	defer req.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting scaffold: %s", res.Status)
	}
	return res, nil
}

func (a *API) SaveScaffold(scaffold types.ScaffoldRequest) (*http.Response, error) {
	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	sreq := types.SaveScaffoldRequest{
		Name:     scaffold.Name,
		UserID:   config.UserID,
		Scaffold: scaffold.Scaffold,
	}

	body, _ := json.Marshal(sreq)
	req, err := http.NewRequest("POST", a.baseURL+"/scaffolds", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	res, err := a.client.Do(req)
	defer req.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error saving scaffold: %s", res.Status)
	}
	return res, nil
}

func getAuthConfig() (*types.APIAuth, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(homeDir, ".sersi", "config.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ac := &types.APIAuth{}
	err = json.Unmarshal(data, ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}
