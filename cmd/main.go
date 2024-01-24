package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"go.uber.org/automaxprocs/maxprocs"
	"p3ld3v.dev/template/app"
	"p3ld3v.dev/template/cmd/env"
)

func setThreadCount(dep app.Dependencies) {
	// Want to see what maxprocs reports.
	opt := maxprocs.Logger(func(format string, args ...interface{}) {
		dep.Logger.Info(fmt.Sprintf(format, args...))
	})
	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(opt); err != nil {
		dep.Logger.Error("maxprocs: %w", err)
	}
	dep.Logger.Info("GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}

// The entrypoint for the application.
func main() {
	dep, err := app.NewDependencies(env.LoadConfig())
	if err != nil {
		dep.Logger.Panic("Error creating dependencies", err)
	}
	setThreadCount(*dep)
	// Start the server
	server := app.StartServer(dep)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		addr := fmt.Sprintf("%s:%s", dep.Config.Host, dep.Config.Port)
		dep.Logger.Info(fmt.Sprintf("Statring server on %s. PID: %d", addr, os.Getpid()))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			dep.Logger.Error(fmt.Sprintf("Error: %v\n", err))
		}
	}()

	// Wait for an interrupt signal
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		dep.Logger.Info("Server is gracefully shutdown: %v\n", err)
	} else {
		server.Close()
		dep.Logger.Info("Server gracefully stopped")
	}
}
