package utils_test

import (
	. "github.com/feranwq/ppt/pkg/utils"
	"os"
	"testing"
)

func TestRemoveRemoveComment(t *testing.T) {
	mockComment := `// getPodMapForDeployment returns the Pods managed by a Deployment.
//
// It returns a map from ReplicaSet UID to a list of Pods controlled by that RS,
// according to the Pod's ControllerRef.`
	expected := ` getPodMapForDeployment returns the Pods managed by a Deployment.
It returns a map from ReplicaSet UID to a list of Pods controlled by that RS,according to the Pod's ControllerRef.`

	actual, err := RemoveComment(mockComment)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestGetEnvValue(t *testing.T) {
	err := os.Setenv("MOCK_TEST_GET_ENV_VALUE", "ok")
	defer os.Unsetenv("MOCK_TEST_GET_ENV_VALUE")
	if err != nil {
		t.Fatal(err)
	}
	expected := "ok"

	actual, err := GetEnvValue("MOCK_TEST_GET_ENV_VALUE")
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}
