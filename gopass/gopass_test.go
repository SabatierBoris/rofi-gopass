package gopass

import (
	"os/exec"
	"testing"
)

func TestCreateOk(t *testing.T) {
	foo := func(string, ...string) *exec.Cmd {
		return nil
	}

	gopass, err := Create(foo)
	if err != nil {
		t.Errorf("err should be nil")
	}

	if gopass == nil {
		t.Errorf("gopass shouldn't be nil")
	}
}

func TestCreateKO(t *testing.T) {
	gopass, err := Create(nil)
	if err == nil {
		t.Errorf("err shouldn't be nil")
	}

	if gopass != nil {
		t.Errorf("gopass should be nil")
	}
}
