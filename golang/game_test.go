package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func Test_no_rolls_returns_0_score(t *testing.T) {
	g := game{}

	assert.Equal(t, 0, g.score())
}

func Test_first_roll_returns_score_equal_to_pins_knocked_down(t *testing.T) {
	g := game{}

	require.NoError(t, g.roll(1))

	assert.Equal(t, 1, g.score())
}

func Test_two_rolls_returns_score_equal_to_pins_knocked_down(t *testing.T) {
	g := game{}

	require.NoError(t, g.roll(1))
	require.NoError(t, g.roll(1))

	assert.Equal(t, 2, g.score())
}

func Test_roll_rejects_more_than_10_pins(t *testing.T) {
	g := game{}

	assert.Error(t, g.roll(11))
}

func Test_spare_is_calculated_with_third_roll_bonus(t *testing.T) {
	g := game{}

	require.NoError(t, g.roll(5))
	require.NoError(t, g.roll(5))
	require.NoError(t, g.roll(1))

	assert.Equal(t, 12, g.score())
}


