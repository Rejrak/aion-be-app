package main

import (
	servConfig "be/internal/config"
	"be/internal/database/db"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"goa.design/clue/log"
)

const (
	fromDocsPath = "./gen/http/openapi3.yaml"
	docsPath     = "./static/openapi3.yaml"
)

func main() {
	srvConf := servConfig.LoadServerConfig()
	ctx := setupLoggingContext(srvConf.Debug)
	var wg sync.WaitGroup
	errc := make(chan error)

	go handleSignals(errc) // Start goroutine to listen for OS signals (e.g., SIGINT, SIGTERM)

	if err := moveFile(fromDocsPath, docsPath); err != nil {
		log.Debugf(ctx, "error: %v", err)
	}

	// Create a cancellable context to manage server shutdown.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() //Ensure the context is canceled to release resources

	switch srvConf.Domain {
	case "development":
		setupEnvironment(ctx, srvConf, &wg, errc) // Start the HTTP server for development
	case "production":
		setupEnvironment(ctx, srvConf, &wg, errc) // Start the HTTP server for production
	default:
		log.Fatal(ctx, fmt.Errorf("invalid host argument: %q (valid hosts: development|production)", srvConf.Domain)) // Fatal error for invalid domain
	}

	// Wait for an error or signal to exit.
	log.Printf(ctx, "exiting (%v)", <-errc)
	cancel()                  // Cancel context to begin shutdown process
	wg.Wait()                 // Wait for all goroutines to complete
	log.Printf(ctx, "exited") // Log when the application has fully exited
}

// handleSignals listens for OS signals and sends them to the error channel.
// This function enables graceful shutdown on system signals (e.g., SIGINT, SIGTERM).
func handleSignals(errc chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // Notify on interrupt or terminate signals
	errc <- fmt.Errorf("%s", <-c)                     // Send the received signal to the error channel as a formatted error
}

func setupEnvironment(ctx context.Context, srvConf *servConfig.ServerConfig, wg *sync.WaitGroup, errc chan error) {
	db.ConnectDb()
	u := srvConf.BuildServerURL(srvConf, ctx)         // Build server URL based on configuration
	HandleHttpServer(ctx, u, wg, errc, srvConf.Debug) // Start the HTTP server for development
}

func setupLoggingContext(debug bool) context.Context {
	format := log.FormatJSON

	if log.IsTerminal() {
		format = log.FormatTerminal
	}
	ctx := log.Context(context.Background(), log.WithFormat(format))
	if debug {
		ctx = log.Context(ctx, log.WithDebug())
		log.Debugf(ctx, "debug logs enabled")
	}
	return ctx
}

// moveFile moves a file from src to dst. If the destination file exists, it will be overwritten.
func moveFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read source file or already moved: %w", err)
	}

	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		return fmt.Errorf("failed to write destination file: %w", err)
	}

	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed to remove source file: %w", err)
	}

	return nil
}
