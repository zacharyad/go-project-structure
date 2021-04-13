package main

import (
	"testing"

	a "github.com/stretchr/testify/assert"
	math "github.com/zacharyad/go-project-structure/pkg/maths"
)

func TestThatWeHaveTestifyWorking(t *testing.T) {
	a.Equal(t, 10, 10, "they should be equal")
}

func TestSubFunc(t *testing.T) {

	a.Equal(t, math.Sub(5, 5), 0, "IT should be 5 - 5 which equals 0")
}

func TestAddFunc(t *testing.T) {
	a.Equal(t, math.Add(5, 5), 10, "they should be equal")
}
