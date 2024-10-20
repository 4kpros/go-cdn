package main

import (
	"cdn/cmd/api"
	"cdn/cmd/di"
	"cdn/common/helpers"
	"cdn/config"

	"go.uber.org/zap"
)

// Contains all errors during init() execution
var errInit error

func main() {
	// Check if there are any errors when initializing the app
	if errInit != nil {
		helpers.Logger.Warn(
			"There are some errors when initializing app!",
			zap.String("Error", "Please fix previous errors before."),
		)
		panic(errInit)
	}

	di.InjectDependencies()
	api.Start()
}

// Called before the main entry point. It's useful for setting up
// configurations before starting the application.
func init() {
	helpers.EnableLogger()

	// Load env
	errEnv := config.LoadEnv()
	if errEnv != nil {
		errInit = errEnv
		helpers.Logger.Error(
			"Failed to load env!",
			zap.String("Error", errEnv.Error()),
		)
	} else {
		helpers.Logger.Info("Env loaded!")
	}

	// Load OpenAPI templates
	errOpenAPITemplates := config.LoadOpenAPITemplates()
	if errOpenAPITemplates != nil {
		errInit = errOpenAPITemplates
		helpers.Logger.Error(
			"Failed to load OpenAPI templates!",
			zap.String("Error", errOpenAPITemplates.Error()),
		)
	} else {
		helpers.Logger.Info("OpenAPI templates loaded!")
	}
}
