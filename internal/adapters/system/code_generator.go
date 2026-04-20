package system

import (
	"math/rand"
	"time"
	"url_shortener/internal/domain"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type codeGenerator struct {
	random *rand.Rand
}

func NewCodeGenerator() domain.CodeGenerator {
	source := rand.NewSource(time.Now().UnixNano()) // * rand.NewSource devuelve un objeto de tipo rand.Source, que es básicamente el generador interno de números aleatorios inicializado con esa semilla.
	// * time.Now().UnixNano() Obtiene el tiempo actual en nanosegundos. Es un valor que cambia constantemente.

	return &codeGenerator{
		random: rand.New(source), // * rand.New() Lo que hace es crear un generador de números aleatorios (*rand.Rand) usando una fuente (Source), que sí depende de la semilla.
	}
}

func (g *codeGenerator) Generate() (string, error) {
	length := 6

	b := make([]byte, length) // * Creamos un slice de tamaño length

	for i := range b {
		b[i] = charset[g.random.Intn(len(charset))]
	}

	return string(b), nil
}

// * Esto NO pertenece al dominio porque: cómo se genera el código ≠ regla de negocio
