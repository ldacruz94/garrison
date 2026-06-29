package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"garrison/internal/handlers"
	"garrison/internal/stores"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	mux := http.NewServeMux()
	connStr := os.Getenv("DATABASE_URL")
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connStr)

	if err != nil {
		log.Fatal("Failure setting up DB")
	}

	// Stores & Handlers
	missionStore := stores.NewMissionStore(pool)
	assetStore := stores.NewAssetStore(pool)
	personnelStore := stores.NewPersonnelStore(pool)

	missionHandler := handlers.NewMissionHandler(missionStore)
	assetHandler := handlers.NewAssetHandler(assetStore)
	personnelHandler := handlers.NewPersonnelHandler(personnelStore)

	// Missions Routes
	mux.HandleFunc("GET /missions/{id}", missionHandler.GetMissionByID)
	mux.HandleFunc("GET /missions", missionHandler.GetAllMissions)
	mux.HandleFunc("POST /missions", missionHandler.CreateMission)
	mux.HandleFunc("DELETE /missions/{id}", missionHandler.DeleteMission)
	mux.HandleFunc("PUT /missions/{id}", missionHandler.UpdateMission)

	// Asset Routes
	mux.HandleFunc("GET /assets/{id}", assetHandler.GetAssetByID)
	mux.HandleFunc("GET /assets", assetHandler.GetAllAssets)
	mux.HandleFunc("POST /assets", assetHandler.CreateAsset)
	mux.HandleFunc("DELETE /assets/{id}", assetHandler.DeleteAsset)
	mux.HandleFunc("PUT /assets/{id}", assetHandler.UpdateAsset)

	// Personnel Routes
	mux.HandleFunc("GET /personnel/{id}", personnelHandler.GetPersonnelByID)
	mux.HandleFunc("GET /personnel", personnelHandler.GetAllPersonnel)
	mux.HandleFunc("POST /personnel", personnelHandler.CreatePersonnel)
	mux.HandleFunc("DELETE /personnel/{id}", personnelHandler.DeletePersonnel)
	mux.HandleFunc("PUT /personnel/{id}", personnelHandler.UpdatePersonnel)

	port := 8080
	var server *http.Server

	if os.Getenv("MTLS_ENABLED") == "true" {
		port = 8443

		certPool, err := loadCaPool()

		if err != nil {
			log.Fatal("Error loading cert pool")
		}

		tlsConfig := &tls.Config{
			ClientCAs:  certPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			MinVersion: tls.VersionTLS13,
		}

		server = &http.Server{
			Addr:      fmt.Sprintf(":%d", port),
			Handler:   mux,
			TLSConfig: tlsConfig,
		}

		log.Println(fmt.Sprintf("Listening on port :%d", port))
		log.Fatal(server.ListenAndServeTLS("certs/server.crt", "certs/server.key"))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
			log.Fatal(err)
		}
	} else {
		server = &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		}

		log.Println(fmt.Sprintf("Listening on port :%d", port))
		log.Fatal(server.ListenAndServe())
	}

}

func loadCaPool() (*x509.CertPool, error) {
	caCert, err := os.ReadFile("certs/ca.crt")

	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	return certPool, nil
}
