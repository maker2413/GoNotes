package specifications

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Maker")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Maker")
}
