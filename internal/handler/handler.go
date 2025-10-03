package handler

import "go-musthave-shortener-tpl/internal/service"

var urlService *service.URLService

func SetURLService(s *service.URLService) {
	urlService = s
}
