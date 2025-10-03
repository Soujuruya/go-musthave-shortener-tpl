package repository

import "go-musthave-shortener-tpl/internal/model"

type URLRepository interface {
	Save(url model.URL) string
	Get(id string) (model.URL, bool)
}
