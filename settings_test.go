package travis

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListEnvironmentVariables(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/settings/env_vars", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "GET")
		expectQueryParam(t, r, "repository_id", "61")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprintf(w, `{"env_vars":[{"id":"1234","name":"foo","value":"bar"}]}`)
	})

	envVars, err := tc.ListEnvironmentVariables(61)
	if err != nil {
		t.Error(err)
	}

	expected := []EnvironmentVariable{{ID: String("1234"), Name: String("foo"), Value: String("bar")}}
	if !reflect.DeepEqual(envVars, expected) {
		t.Errorf("ListEnvironmentVariables returned %+v, expected %+v", envVars, expected)
	}
}

func TestCreateEnvironmentVariable(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/settings/env_vars", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "POST")
		expectQueryParam(t, r, "repository_id", "62")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprintf(w, `{}`)
	})

	_, err := tc.CreateEnvironmentVariable(62, &EnvironmentVariable{ID: String("62abc")})
	if err != nil {
		t.Error(err)
	}
}

func UpdateEnvironmentVariable(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/settings/env_vars/62abc", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "PATCH")
		expectQueryParam(t, r, "repository_id", "62")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprintf(w, `{}`)
	})

	_, err := tc.UpdateEnvironmentVariable(62, "62abc", &EnvironmentVariable{Name: String("SOMEVAR")})
	if err != nil {
		t.Error(err)
	}
}

func DestroyEnvironmentVariable(t *testing.T) {
	setup("test-token")
	defer teardown()

	mux.HandleFunc("/settings/env_vars/62abc", func(w http.ResponseWriter, r *http.Request) {
		expectMethod(t, r, "PATCH")
		expectQueryParam(t, r, "repository_id", "62")
		expectTravisHeaders(t, r, "test-token")
		fmt.Fprintf(w, `{}`)
	})

	if err := tc.DestroyEnvironmentVariable(62, "62abc"); err != nil {
		t.Error(err)
	}
}
