package repository

import (
	"go-musthave-shortener-tpl/internal/model"
)

type MemoryURLRepo struct {
	DB map[string]model.URL
}

func NewMemoryURLRepo() *MemoryURLRepo {
	return &MemoryURLRepo{DB: make(map[string]model.URL)}
}

func (r *MemoryURLRepo) Save(url model.URL) string {
	r.DB[url.ID] = url
	return url.ID
}

func (r *MemoryURLRepo) Get(id string) (model.URL, bool) {
	u, ok := r.DB[id]
	return u, ok
}
