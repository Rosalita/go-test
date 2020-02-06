package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToRomanNumeral(t *testing.T) {

	tests := []struct {
		number       int
		romanNumeral string
	}{
		
	}

	for _, test := range tests {
		romanNumeral := toRomanNumeral(test.number)
		assert.Equal(t, test.romanNumeral, romanNumeral)
	}

}
