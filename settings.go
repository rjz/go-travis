package travis

import (
	"encoding/json"
	"fmt"
)

type EnvironmentVariable struct {
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Value        *string `json:"value,omitempty"`
	Public       *bool   `json:"public,omitempty"`
	RepositoryID int     `json:"repository_id,omitempty"`
}

type environmentVariableBody map[string]*EnvironmentVariable

type environmentVariableListBody map[string][]EnvironmentVariable

func (tc *Client) ListEnvironmentVariables(repoId int) ([]EnvironmentVariable, error) {
	bytes, err := tc.Get(fmt.Sprintf("/settings/env_vars?repository_id=%d", repoId))
	if err != nil {
		return nil, err
	}

	resp := environmentVariableListBody{"env_vars": nil}
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp["env_vars"], nil
}

func (tc *Client) CreateEnvironmentVariable(repoId int, e *EnvironmentVariable) (*EnvironmentVariable, error) {
	body := environmentVariableBody{"env_var": e}
	_, err := tc.Post(fmt.Sprintf("/settings/env_vars?repository_id=%d", repoId), body)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (tc *Client) UpdateEnvironmentVariable(repoId int, envVarId string, e *EnvironmentVariable) (*EnvironmentVariable, error) {
	body := environmentVariableBody{"env_var": e}
	_, err := tc.Patch(fmt.Sprintf("/settings/env_vars/%s?repository_id=%d", envVarId, repoId), body)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (tc *Client) DestroyEnvironmentVariable(repoId int, envVarId string) error {
	_, err := tc.Delete(fmt.Sprintf("/settings/env_vars/%s?repository_id=%d", envVarId, repoId))
	return err
}
