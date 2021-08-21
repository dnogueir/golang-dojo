package solution_test

import (
	"testing"

	"github.com/dnogueir/golang-dojo/solution"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	sum := solution.Multi_threaded()
	require.Equal(t, 319600, sum)
}
