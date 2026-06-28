package main

import (
	"context"
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

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
