package utils_test

import (
	. "github.com/feranwq/ppt/pkg/utils"
	"os"
	"testing"
)

func TestGoComment(t *testing.T) {
	mockComment := `// getPodMapForDeployment returns the Pods managed by a Deployment.
//
// It returns a map from ReplicaSet UID to a list of Pods controlled by that RS,
// according to the Pod's ControllerRef.`
	expected := ` getPodMapForDeployment returns the Pods managed by a Deployment. It returns a map from ReplicaSet UID to a list of Pods controlled by that RS,according to the Pod's ControllerRef.`

	actual, err := FormatComment(mockComment)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestLinuxMan(t *testing.T) {
	mockComment := `Cron  is  started  from  /etc/rc.d/init.d or /etc/init.d when classical sysvinit scripts are used. In case systemd is enabled, then unit file is installed into /lib/systemd/sys‐
	tem/crond.service and daemon is started by systemctl start crond.service command. It returns immediately, thus, there is no need to need to start it with the '&' parameter.`
	expected := `Cron  is  started  from  /etc/rc.d/init.d or /etc/init.d when classical sysvinit scripts are used. In case systemd is enabled, then unit file is installed into /lib/systemd/system/crond.service and daemon is started by systemctl start crond.service command. It returns immediately, thus, there is no need to need to start it with the '&' parameter.`

	actual, err := FormatComment(mockComment)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestLinuxDoc(t *testing.T) {
	mockComment := `By default, Linux follows an optimistic memory allocation strategy.
This means that when malloc() returns non-NULL there is  no  guarantee
that  the  memory  really is available.  This is a really bad bug.`
	expected := `By default, Linux follows an optimistic memory allocation strategy. This means that when malloc() returns non-NULL there is  no  guarantee that  the  memory  really is available.  This is a really bad bug.`

	actual, err := FormatComment(mockComment)
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
