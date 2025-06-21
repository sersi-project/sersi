package scaffold

import (
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
	_, err := s.client.SaveScaffold(types.ScaffoldRequest{
		Name:     config.Name,
		Scaffold: config.Scaffold,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ScaffoldService) GetAllScaffolds() ([]pkg.Config, error) {
	list, err := s.client.GetAllScaffolds()
	if err != nil {
		return nil, err
	}

	var configs []pkg.Config
	for _, scaffold := range list {
		configs = append(configs, pkg.Config{
			Name:     scaffold.Name,
			Scaffold: scaffold.Scaffold,
		})
	}

	return configs, nil
}

func (s *ScaffoldService) GetScaffold(name string) (*pkg.Config, error) {
	res, err := s.client.GetScaffold(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ScaffoldService) UpdateScaffold(config *pkg.Config) error {
	_, err := s.client.UpdateScaffold(types.ScaffoldRequest{
		Name:     config.Name,
		Scaffold: config.Scaffold,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ScaffoldService) DeleteScaffold(name string) error {
	_, err := s.client.DeleteScaffold(name)
	if err != nil {
		return err
	}

	return nil
}
