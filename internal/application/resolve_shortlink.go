package application

import (
	"url_shortener/internal/domain"
)

type resolveShortLink struct {
	repo  domain.ShortLinkRepository
	clock domain.Clock
}

type ResolveShortLinkI interface {
	Execute(code string) (string, error)
}

func NewResolveShortLink(repo domain.ShortLinkRepository, clock domain.Clock) *resolveShortLink {
	return &resolveShortLink{
		repo:  repo,
		clock: clock,
	}
}

func (uc *resolveShortLink) Execute(code string) (string, error) {

	link, err := uc.repo.FindByCode(code)
	if err != nil {
		return "", err
	}

	// Lógica del dominio
	url, err := link.Resolve(uc.clock.Now())
	if err != nil {
		return "", err
	}

	// Persistir cambio
	err = uc.repo.Save(link)
	if err != nil {
		return "", err
	}

	return url, nil
}

/*
Mira lo elegante que queda:

	-> El caso de uso no sabe reglas
	-> La entidad no sabe DB
	-> El repo no sabe negocio

Hexagonal puro.
*/
