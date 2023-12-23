package runner_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gohryt/asphyxia/async/context"
	"github.com/gohryt/asphyxia/async/runner"
)

func ServerStart(ctx *context.Context, group *runner.Group, resource *http.Server) {
	err := resource.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		runner.Send(group, runner.ActionClose)

		log.Println(err)
	}
}

func ServerClose(ctx *context.Context, group *runner.Group, resource *http.Server) {
	err := resource.Shutdown(nil)
	if err != nil {
		log.Println(err)
	}
}

func TestMain(t *testing.T) {
	group, _ := runner.New(context.Background(), os.Interrupt)

	server := &http.Server{
		Addr: ":8080",
	}

	runner.On(runner.ActionStart, group, server, ServerStart)
	runner.On(runner.ActionClose, group, server, ServerClose)

	runner.Wait(group, 4*time.Second)
}
