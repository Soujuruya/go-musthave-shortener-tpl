package app

import (
	"go-musthave-shortener-tpl/internal/handler"
	"go-musthave-shortener-tpl/internal/repository"
	"go-musthave-shortener-tpl/internal/router"
	"go-musthave-shortener-tpl/internal/service"
)

type App struct {
	route *router.Router
	dbase *repository.MemoryURLRepo
}

func New() *App {

	urlRepo := repository.NewMemoryURLRepo()

	urlService := service.NewURLService(urlRepo)

	handler.SetURLService(urlService)

	route := router.New()

	return &App{route: route}
}

func (a *App) Run() error {
	return a.route.Run()
}
