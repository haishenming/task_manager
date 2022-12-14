// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"task_manager/internal/biz"
	"task_manager/internal/conf"
	"task_manager/internal/data"
	"task_manager/internal/server"
	"task_manager/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client, err := data.NewClient(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, logger, client)
	if err != nil {
		return nil, nil, err
	}
	taskRepo := data.NewTaskRepo(dataData, logger)
	employeeRepo := data.NewEmployeeRepo(dataData, logger)
	hospitalRepo := data.NewHospitalRepo(dataData, logger)
	taskUsecase := biz.NewTaskUsecase(taskRepo, employeeRepo, hospitalRepo, logger)
	taskService := service.NewTaskService(taskUsecase)
	grpcServer := server.NewGRPCServer(confServer, taskService, logger)
	httpServer := server.NewHTTPServer(confServer, taskService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
