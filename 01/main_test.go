package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {

	tests := []struct {
		name             string
		expectedGreeting string
	}{
		{"Atelier", "Hello Atelier!"},
	}

	for _, test := range tests {
		greeting := greet(test.name)
		assert.Equal(t, test.expectedGreeting, greeting)
	}

}
