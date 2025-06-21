package authorization

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sersi-project/sersi/internal/api"
)

type AuthConfig struct {
	Token        string `json:"access_token"`
	UserID       string `json:"user_id"`
	Email        string `json:"email"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func CheckStatus() (userID string, ok bool) {
	configPath, err := getConfigPath()
	if err != nil {
		ok = false
		return
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		ok = false
		return
	}

	authConfig := &AuthConfig{}
	err = json.Unmarshal(data, &authConfig)
	if err != nil {
		ok = false
		return
	}

	err = authConfig.Validate()
	if err != nil {
		ok = false
		return
	}

	err = authConfig.AttemptRefresh()
	if err != nil {
		ok = false
		return
	}

	userID = authConfig.UserID
	ok = true
	return
}

func Login(email, password string) error {
	res, err := api.NewAPI().Authenticate(email, password)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Unix() + res.ExpiresIn

	authConfig := &AuthConfig{
		Token:        res.Token,
		UserID:       res.UserID,
		Email:        email,
		ExpiresIn:    expiresAt,
		RefreshToken: res.RefreshToken,
	}

	err = authConfig.Validate()
	if err != nil {
		return err
	}

	err = authConfig.writeToConfig()
	if err != nil {
		return err
	}

	return nil
}

func (ac *AuthConfig) Validate() error {
	if ac.Token == "" || ac.RefreshToken == "" || ac.UserID == "" || ac.Email == "" || ac.ExpiresIn == 0 {
		return fmt.Errorf("config is not valid")
	}

	if ac.ExpiresIn+40800 < time.Now().Unix() {
		return fmt.Errorf("config is not valid")
	}

	return nil
}

func (ac *AuthConfig) AttemptRefresh() error {
	if ac.ExpiresIn < time.Now().Unix() {
		if ac.ExpiresIn+40800 < time.Now().Unix() {
			return fmt.Errorf("config is not valid")
		}
		token, err := api.NewAPI().RefreshToken()
		if err != nil {
			return fmt.Errorf("error refreshing token: %v", err)
		}
		ac.Token = token.Token
		ac.ExpiresIn = time.Now().Unix() + token.ExpiresIn
		ac.RefreshToken = token.RefreshToken

		err = ac.writeToConfig()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ac *AuthConfig) writeToConfig() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	acBytes, err := json.Marshal(ac)
	if err != nil {
		return err
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0o755); err != nil {
			return err
		}
	}

	if err := os.WriteFile(filepath.Join(configDir, "config.json"), acBytes, 0o644); err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	targetDir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(targetDir, "config.json"), nil
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(homeDir, ".sersi")
	return path, nil
}
