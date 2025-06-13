package scaffold

import (
	"encoding/json"
	"fmt"
	"io"

	types "github.com/sersi-project/sersi/internal/types"

	"github.com/sersi-project/sersi/internal/api"
	"github.com/sersi-project/sersi/pkg"
)

type Scaffold interface {
	Generate() error
}

type ScaffoldBuilder interface {
	Build() Scaffold
}

type ScaffoldService struct {
	client *api.API
}

func NewScaffoldService() *ScaffoldService {
	return &ScaffoldService{
		client: api.NewAPI(),
	}
}

func (s *ScaffoldService) SaveScaffold(config *pkg.Config) error {
	res, err := s.client.SaveScaffold(types.ScaffoldRequest{
		Name:     config.Name,
		Scaffold: config.Scaffold,
	})
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("error saving scaffold: %s", res.Status)
	}

	fmt.Println("Scaffold saved successfully")
	return nil
}

func (s *ScaffoldService) GetAllScaffolds() error {
	res, err := s.client.GetAllScaffolds()
	if err != nil {
		return err
	}

	var list []pkg.ScaffoldConfig
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &list)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScaffoldService) GetScaffold(name string) (*pkg.Config, error) {
	res, err := s.client.GetScaffold(name)
	if err != nil {
		return nil, err
	}

	var scaffold pkg.ScaffoldConfig
	body, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &scaffold)
	if err != nil {
		return nil, err
	}

	config := &pkg.Config{
		Scaffold: scaffold,
		Name:     name,
	}

	return config, nil
}
