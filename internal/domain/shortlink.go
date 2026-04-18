package domain

import (
	"time"
)

type ShortLink struct {
	ID          string
	OriginalURL string
	Code        string
	CreatedAt   time.Time
	ExpiresAt   *time.Time // nil = no expira
	VisitCount  int
}

func (s *ShortLink) IsActive(now time.Time) bool { // Pasamos now como parámetro. Esto evita acoplar el dominio a time.Now() y lo hace testeable.
	if s.ExpiresAt == nil {
		return true
	}
	return now.Before(*s.ExpiresAt) // Devuelve true si now ocurre antes que ExpiresAt.
}

func (s *ShortLink) IncreaseVisits() {
	s.VisitCount++
}

func (s *ShortLink) CanBeResolved(now time.Time) bool {
	return s.IsActive(now)
}

func (s *ShortLink) Resolve(now time.Time) (string, error) {
	if !s.IsActive(now) {
		return "", ErrLinkExpired
	}
	s.VisitCount++

	return s.OriginalURL, nil
}
