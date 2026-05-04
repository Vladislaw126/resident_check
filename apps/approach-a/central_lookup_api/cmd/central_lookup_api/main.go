package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/api"
	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/app"
	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/audit"
	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/client"
	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/mapper"
	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/router"
	"github.com/Vladislaw126/resident-check-poc/apps/approach-a/central_lookup_api/internal/validator"
)

func main() {
	log.Println("starting central_lookup_api...")

	// low-level helpers and dependencies.
	requestValidator := validator.New()
	authorityRouter := router.New()
	residentRecordClient := client.New("http://localhost:8081")
	responseMapper := mapper.New()
	auditLogger := audit.New()

	// the application/service layer that coordinates the use case.
	lookupService := app.NewLookupService(
		requestValidator,
		authorityRouter,
		residentRecordClient,
		responseMapper,
		auditLogger,
	)

	// the HTTP handler layer.
	handler := api.NewHandler(lookupService)

	// register routes.
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/lookup", handler.Lookup)

	// configure and start the HTTP server.
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("central_lookup_api listening on %s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
