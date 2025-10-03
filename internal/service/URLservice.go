package service

import (
	"math/rand"

	"go-musthave-shortener-tpl/internal/model"
	"go-musthave-shortener-tpl/internal/repository"
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) Shortner(original string) string {
	id := s.generateUniqueID(6)
	url := model.URL{ID: id, Original: original}
	s.repo.Save(url)
	return id
}

func (s *URLService) GetOriginal(id string) (string, bool) {
	u, ok := s.repo.Get(id)
	if !ok {
		return "", false
	}
	return u.Original, true
}
func (s *URLService) generateUniqueID(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		b := make([]byte, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		id := string(b)
		if _, exists := s.repo.Get(id); !exists {
			return id
		}
	}
}
