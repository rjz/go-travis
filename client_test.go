package travis

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetRepository(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/repos/rjz/dingus", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "GET")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprintf(w, `{"repo":{"id":41}}`)
	})

	repo, err := tc.GetRepository("rjz", "dingus")
	if err != nil {
		t.Error(err)
	}

	expected := &Repository{ID: 41}
	if !reflect.DeepEqual(repo, expected) {
		t.Errorf("GetRepository returned %+v, expected %+v", repo, expected)
	}
}

func TestGetRepositorySettings(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/repos/42/settings", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "GET")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprint(w, `{"settings": {"builds_only_with_travis_yml":true}}`)
	})

	repoSettings, err := tc.GetRepositorySettings(42)
	if err != nil {
		t.Error(err)
	}

	expected := &RepositorySettings{BuildsOnlyWithTravisYml: Bool(true)}
	if !reflect.DeepEqual(repoSettings, expected) {
		t.Errorf("GetRepositorySettings returned %+v, expected %+v", repoSettings, expected)
	}
}

func TestUpdateRepositorySettings(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/repos/43/settings", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "PATCH")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprint(w, `{}`)
	})

	_, err := tc.UpdateRepositorySettings(43, &RepositorySettings{BuildPushes: Bool(true)})
	if err != nil {
		t.Error(err)
	}
}
