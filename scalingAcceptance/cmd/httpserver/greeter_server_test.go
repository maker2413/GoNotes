package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/maker2413/GoNotes/scalingAcceptance/adapters"
	"github.com/maker2413/GoNotes/scalingAcceptance/adapters/httpserver"
	"github.com/maker2413/GoNotes/scalingAcceptance/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
