package runner_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/gohryt/asphyxia/async/runner"
)

func TestMain(t *testing.T) {
	group, _ := runner.Prepare(context.Background(), os.Interrupt)

	server := &http.Server{
		Addr: ":8080",
	}

	runner.On(runner.ActionStart, group, server, func(ctx context.Context, group *runner.Group, resource *http.Server) error {
		err := resource.ListenAndServe()

		if err == http.ErrServerClosed {
			return nil
		}

		return err
	})

	runner.On(runner.ActionClose, group, server, func(ctx context.Context, group *runner.Group, resource *http.Server) error {
		return server.Shutdown(ctx)
	})

	group.Wait()
}
