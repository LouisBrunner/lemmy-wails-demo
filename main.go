package main

import (
	"embed"

	"github.com/LouisBrunner/lemmy-wails-demo/backend/application"
	"github.com/LouisBrunner/lemmy-wails-demo/backend/contracts"
	"github.com/LouisBrunner/lemmy-wails/backend"
	"github.com/LouisBrunner/lemmy-wails/backend/api"
	"github.com/sirupsen/logrus"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

//go:embed wails.json
var rawConfig []byte

func main() {
	backend.Start(api.Options[contracts.Bindings, application.Config]{
		Assets:    assets,
		Icon:      icon,
		WailsJSON: rawConfig,
		AppMaker: func(logger *logrus.Logger, cfg api.ConfigManager[application.Config]) (api.App[contracts.Bindings], error) {
			return application.New(logger, cfg), nil
		},
	})
}
