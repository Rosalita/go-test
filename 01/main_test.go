package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T){

	tests := []struct {
		name  string
		expectedGreeting string
	}{
		{"Atelier", "Hello Atelier!"},
		{"Gwen", "Hello Gwen!"},
	}

	for _, test := range tests {
		result := Greet(test.name)
		assert.Equal(t, test.expectedGreeting, result)
	}

}