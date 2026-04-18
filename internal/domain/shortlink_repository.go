package domain

type ShortLinkRepository interface {
	FindByCode(code string) (*ShortLink, error)
	Save(link *ShortLink) error
} // Esto define que necesita el dominio del exterior
