package utility

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	expected := 15
	actual := Sum(1, 2, 3, 4, 5)

	if expected != actual {
		t.Fatalf("expected %v, but got %v", expected, actual)
		// t.Fatal("errorrr")
	}

	log.Println("done")
}

func TestSumTestify(t *testing.T) {
	type tabletest struct {
		title    string
		expected int
		actual   int
	}

	var tabletests = []tabletest{
		{
			title:    "no send data",
			expected: -1,
			actual:   Sum(),
		},
		{
			title:    "send data with zero value",
			expected: -1,
			actual:   Sum(1, 0, 2),
		},
		{
			title:    "send data with minus value",
			expected: -1,
			actual:   Sum(1, -1, 2),
		},
		{
			title:    "success",
			expected: 15,
			actual:   Sum(1, 2, 3, 4, 5),
		},
	}

	for _, test := range tabletests {
		t.Run(test.title, func(t *testing.T) {
			require.Equal(t, test.expected, test.actual)
		})
	}

}
