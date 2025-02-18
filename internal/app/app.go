package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"time"
	"todo/cmd/config"
	handler "todo/internal/controllers/rest/handlers/task"
	"todo/internal/controllers/rest/routers"
	service "todo/internal/service/task"
	storage "todo/internal/storage/task"
	"todo/pkg/db"
	"todo/pkg/logger"
)

type App struct {
	Config   *config.Config
	FiberApp *fiber.App
	LogFile  *os.File
	Log      *slog.Logger
}

func NewApp(cfg *config.Config) (*App, error) {

	dbCfg := db.ConfigDB{
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		DBName:   cfg.DB.DBName,
	}

	pgx, err := db.NewConnectDB(dbCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %w", err)
	}

	logFile, err := logger.InitLogger(cfg.Env, cfg.Log.LogFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to init logFile: %w", err)
	}

	l := slog.Default()
	fiberApp := fiber.New()

	taskStorage := storage.NewTaskStorage(pgx, l)
	taskService := service.NewTaskService(taskStorage)
	taskHandler := handler.NewTaskHandler(taskService, l)

	routers.RegisterRoutes(fiberApp, taskHandler)

	return &App{
		Config:   cfg,
		FiberApp: fiberApp,
		LogFile:  logFile,
		Log:      l,
	}, nil
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.Config.Server.Port)
	a.Log.Info("Server is running on port", slog.String("port", addr))
	return a.FiberApp.Listen(addr)
}

func (a *App) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.FiberApp.ShutdownWithContext(ctx); err != nil {
		a.Log.Error("Fiber Shutdown Failed", slog.String("details", fmt.Sprintf("%v", err)))
	}

	if a.LogFile != nil {
		a.LogFile.Close()
	}

	return nil
}
