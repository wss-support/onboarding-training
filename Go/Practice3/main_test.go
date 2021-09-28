package main

import (
	"github.com/hashicorp/go-version"
	"testing"
)

func TestReadVersionFile(t *testing.T) {
	_, _ = version.NewVersion("20.5.7")
}
