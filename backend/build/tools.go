//go:build tools
// +build tools

package tools

import (
	_ "github.com/LouisBrunner/lemmy-wails/cmd/lemmy"
	_ "github.com/wailsapp/wails/v2/cmd/wails"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
