package domain

import "errors"

var (
	ErrLinkExpired  = errors.New("short link expired")
	ErrLinkNotFound = errors.New("short link not found")
)
