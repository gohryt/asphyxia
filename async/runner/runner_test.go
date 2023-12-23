package runner_test

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gohryt/asphyxia/async/runner"
)

func ServerStart(ctx context.Context, group *runner.Group, resource *http.Server) {
	err := resource.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		runner.Send(group, runner.ActionClose)

		log.Println(err)
	}
}

func ServerClose(ctx context.Context, group *runner.Group, resource *http.Server) {
	err := resource.Shutdown(nil)
	if err != nil {
		log.Println(err)
	}
}

func TestMain(t *testing.T) {
	group, shutdown := runner.New(context.Background(), os.Interrupt)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: time.Second,
		BaseContext: func(l net.Listener) context.Context {
			return shutdown
		},
	}

	runner.On(runner.ActionStart, group, server, ServerStart)
	runner.On(runner.ActionClose, group, server, ServerClose)

	runner.Wait(group, 4*time.Second)
}
