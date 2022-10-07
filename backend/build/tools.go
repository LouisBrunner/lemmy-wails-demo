//go:build tools
// +build tools

package tools

import (
	_ "github.com/LouisBrunner/lemmy-wails/cmd/lemmy"
  _ "honnef.co/go/tools/cmd/staticcheck"
)
