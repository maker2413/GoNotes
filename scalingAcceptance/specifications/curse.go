package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meany MeanGreeter) {
	got, err := meany.Curse("Maker")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, Maker!")
}
