package travis

import (
	"encoding/json"
	"fmt"
)

type Repository struct {
	ID             int     `json:"id,omitempty"`
	Slug           *string `json:"slug,omitempty"`
	Description    *string `json:"description,omitempty"`
	GithubLanguage *string `json:"github_language,omitempty"`
	// Also build status info
}

type repositoryBody map[string]*Repository

type RepositorySettings struct {
	BuildsOnlyWithTravisYml *bool `json:"builds_only_with_travis_yml,omitempty"`
	BuildPushes             *bool `json:"build_pushes,omitempty"`
	BuildPullRequests       *bool `json:"build_pull_requests,omitempty"`
	MaximumNumberOfBuilds   *int  `json:"maximum_number_of_builds,omitempty"`
}

type repositorySettingsBody map[string]*RepositorySettings

func (tc *Client) GetRepository(owner, name string) (*Repository, error) {
	bytes, err := tc.Get(fmt.Sprintf("/repos/%s/%s", owner, name))
	if err != nil {
		return nil, err
	}

	resp := repositoryBody{"repo": nil}
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp["repo"], nil
}

func (tc *Client) GetRepositorySettings(repoId int) (settings *RepositorySettings, err error) {
	bytes, err := tc.Get(fmt.Sprintf("/repos/%d/settings", repoId))
	if err != nil {
		return nil, err
	}

	resp := repositorySettingsBody{"settings": nil}
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return resp["settings"], nil
}

func (tc *Client) UpdateRepositorySettings(repoId int, settings *RepositorySettings) (*RepositorySettings, error) {
	_, err := tc.Patch(fmt.Sprintf("/repos/%d/settings", repoId), repositorySettingsBody{"settings": settings})
	if err != nil {
		return nil, err
	}
	return settings, nil
}
