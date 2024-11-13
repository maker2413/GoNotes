package scaling_acceptance_test

import (
	"testing"

	scaling_acceptance "github.com/maker2413/GoNotes/scalingAcceptance"
	"github.com/maker2413/GoNotes/scalingAcceptance/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t,
		specifications.GreetAdapter(scaling_acceptance.Greet),
	)
}
