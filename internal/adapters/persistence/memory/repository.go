package memory

import (
	"sync"
	"url_shortener/internal/domain"
)

type repository struct {
	data map[string]*domain.ShortLink // * aquí declaramos el tipo; aún no creamos el map
	mu   sync.RWMutex
}

func NewShortLinkRepository() domain.ShortLinkRepository {
	return &repository{
		data: make(map[string]*domain.ShortLink), // * Con el make sí crea el map en memoria
	}
}

func (r *repository) Save(link *domain.ShortLink) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[link.Code] = link

	return nil
}

func (r *repository) FindByCode(code string) (*domain.ShortLink, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	link, ok := r.data[code]

	if !ok {
		return nil, domain.ErrLinkNotFound
	}

	return link, nil
}
