package app

import (
	grpcapp "github.com/NorthDice/AuthGRPC/internal/app/grpc"
	"github.com/NorthDice/AuthGRPC/internal/services/auth"
	"github.com/NorthDice/AuthGRPC/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTLL time.Duration,
) *App {

	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Error("Failed to initialize storage")
		panic(err)
	}
	authService := auth.New(log, storage, storage, storage, tokenTLL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
