package mosh

import (
	"gitlab.hive.thyth.com/chronostruct/go-mosh/pkg/mosh"
	"strings"

	"testing"
)

func TestInternalVersion(t *testing.T) {
	version := mosh.GetVersion()
	if !strings.HasPrefix(version, "mosh") {
		t.Errorf("Unexpected mosh version: " + version)
	}
}
