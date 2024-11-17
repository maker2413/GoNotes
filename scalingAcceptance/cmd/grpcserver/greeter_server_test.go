package main_test

import (
	"fmt"
	"testing"

	"github.com/maker2413/GoNotes/scalingAcceptance/adapters"
	"github.com/maker2413/GoNotes/scalingAcceptance/adapters/grpcserver"
	"github.com/maker2413/GoNotes/scalingAcceptance/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
