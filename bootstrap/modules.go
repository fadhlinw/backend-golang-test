package bootstrap

import (
	"github.com/fadhlinw/backend-golang-test/api/controllers"
	"github.com/fadhlinw/backend-golang-test/api/middlewares"
	"github.com/fadhlinw/backend-golang-test/api/routes"
	"github.com/fadhlinw/backend-golang-test/lib"
	"github.com/fadhlinw/backend-golang-test/repository"
	"github.com/fadhlinw/backend-golang-test/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
