package app

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"log"
	"testing"
	"time"
)

func TestNewApp(t *testing.T) {
	httpSrv := http.NewServer()
	app := New(
		ID("SSSS"),
		Name("sss"),
		Server(httpSrv),
	)

	time.AfterFunc(30*time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
