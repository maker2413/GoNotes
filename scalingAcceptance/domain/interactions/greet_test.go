package interactions_test

import (
	"testing"

	"github.com/maker2413/GoNotes/scalingAcceptance/domain/interactions"
	"github.com/maker2413/GoNotes/scalingAcceptance/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
