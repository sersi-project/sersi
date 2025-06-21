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
	"github.com/sersi-project/sersi/pkg"
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

	req, err := http.NewRequest("POST", a.baseURL+"/auth/sign-in", bytes.NewReader(body))
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

func (a *API) RefreshToken() (*types.APIAuth, error) {
	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	refreshRequest := map[string]string{
		"refresh_token": config.RefreshToken,
		"user_id":       config.UserID,
	}

	body, _ := json.Marshal(refreshRequest)
	req, err := http.NewRequest("POST", a.baseURL+"/auth/refresh", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	res, err := a.client.Do(req)
	defer res.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error refreshing token: %s", res.Status)
	}

	var auth types.APIAuthResponse
	respBody, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(respBody, &auth)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling auth config: %v", err)
	}
	return &auth.Data, nil
}

func (a *API) GetAllScaffolds() ([]types.GetScaffoldResponse, error) {
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
	defer res.Body.Close() //nolint

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting all scaffolds: %s", res.Status)
	}

	var list types.APIScaffoldsResponse
	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return list.Data, nil
}

func (a *API) GetScaffold(name string) (*pkg.Config, error) {
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
	defer res.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting scaffold: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var scaffold pkg.Config
	err = json.Unmarshal(body, &scaffold)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return &scaffold, nil
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
	defer res.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 500 {
		return nil, fmt.Errorf("client failed to save scaffold: %s", res.Status)
	}

	if res.StatusCode == 409 {
		return nil, fmt.Errorf("scaffold name already exists in store")
	}

	return res, nil
}

func (a *API) DeleteScaffold(name string) (*http.Response, error) {
	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", a.baseURL+"/scaffolds/"+config.UserID+"/"+name, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	res, err := a.client.Do(req)
	defer res.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("client failed to delete scaffold: %s", res.Status)
	}
	return res, nil
}

func (a *API) UpdateScaffold(scaffold types.ScaffoldRequest) (*http.Response, error) {
	config, err := getAuthConfig()
	if err != nil {
		return nil, err
	}

	sreq := types.SaveScaffoldRequest{
		Name:     scaffold.Name,
		Scaffold: scaffold.Scaffold,
	}

	body, _ := json.Marshal(sreq)
	req, err := http.NewRequest("PUT", a.baseURL+"/scaffolds/"+config.UserID+"/"+scaffold.Name, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.Token)
	res, err := a.client.Do(req)
	defer res.Body.Close() //nolint
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("client failed to update scaffold: %s", res.Status)
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
