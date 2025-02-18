package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"todo/cmd/config"
	_ "todo/docs"
	"todo/internal/app"
)

// @version         1.0
// @description     API для работа с задачами.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:5555
// @BasePath
func main() {
	cfg := config.MustConfig()

	appInstance, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer appInstance.Close()

	go func() {
		if err := appInstance.Run(); err != nil {
			log.Printf("failed to start HTTP server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := appInstance.Close(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}
