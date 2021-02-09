package main

import (
	"context"
	"errors"
	"os"

	"github.com/DanielTitkov/anomaly-detection-service/cmd/app/prepare"
	"github.com/DanielTitkov/anomaly-detection-service/internal/app"
	"github.com/DanielTitkov/anomaly-detection-service/internal/configs"
	"github.com/DanielTitkov/anomaly-detection-service/internal/logger"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent"
	mockAnalyzer "github.com/DanielTitkov/anomaly-detection-service/internal/service/analyzer/mock"
	randomDataset "github.com/DanielTitkov/anomaly-detection-service/internal/service/dataset/random"
	mockNotification "github.com/DanielTitkov/anomaly-detection-service/internal/service/notification/mock"

	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	logger.Info("starting service", "")

	args := os.Args[1:]
	if len(args) < 1 {
		logger.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	logger.Info("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		logger.Fatal("failed to load config", err)
	}
	logger.Info("loaded config", "")

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)
	dataset := randomDataset.NewService()
	notification := mockNotification.NewService()
	analyzer := mockAnalyzer.NewService()

	app := app.NewApp(cfg, logger, repo, dataset, notification, analyzer)

	server := prepare.NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
