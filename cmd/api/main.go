package main

import (
	"fmt"
	"log"
	"net/http"
	"url_shortener/internal/adapters/httpa"
	"url_shortener/internal/adapters/persistence/memory"
	"url_shortener/internal/adapters/system"
	"url_shortener/internal/application"

	"github.com/go-chi/chi/v5"
)

func main() {
	repo := memory.NewShortLinkRepository()

	clock := system.NewClock()

	codeGen := system.NewCodeGenerator()

	idGenerator := system.NewIDGenerator()

	createUC := application.NewCreateShortLink(repo, codeGen, clock, idGenerator)

	resolveUC := application.NewResolveShortLink(repo, clock)

	handler := httpa.NewHandler(resolveUC, createUC)

	fmt.Printf("%#v\n", createUC)

	r := chi.NewRouter()

	r.Post("/shortlinks", handler.Create)
	r.Get("/{code}", handler.Resolve)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
