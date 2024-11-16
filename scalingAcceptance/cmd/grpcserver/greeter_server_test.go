package main_test

import (
	"fmt"
	"testing"

	"github.com/maker2413/GoNotes/scalingAcceptance/adapters"
	"github.com/maker2413/GoNotes/scalingAcceptance/adapters/grpcserver"
	"github.com/maker2413/GoNotes/scalingAcceptance/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
