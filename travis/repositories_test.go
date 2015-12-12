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

	expected := &Repository{ID: 41}
	if repo, _ := tc.GetRepository("rjz", "dingus"); !reflect.DeepEqual(repo, expected) {
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

	expected := &RepositorySettings{BuildsOnlyWithTravisYml: Bool(true)}
	if repoSettings, _ := tc.GetRepositorySettings(42); !reflect.DeepEqual(repoSettings, expected) {
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

	if _, err := tc.UpdateRepositorySettings(43, &RepositorySettings{BuildPushes: Bool(true)}); err != nil {
		t.Error(err)
	}
}
