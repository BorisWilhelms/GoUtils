package BasicAuth

import (
	"testing"
)

func Test_CreateShouldReturnCorrectValue(t *testing.T) {
	expected := "dXNlcm5hbWU6cGFzc3dvcmQ=" //username:password
	auth := Create("username", "password")

	if expected != auth {
		t.Errorf("Expected %s but got %s", expected, auth)
	}
}
