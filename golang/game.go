package main

import (
	"errors"
)

type game struct {
	frames      []frame // 10 frames in each game
	scoreCount  int
	spare       bool
	strike      bool
	strikeRolls []int
}

type frame struct {
	rolls []int // [1, 2] (max de 3)
}

func (g *game) roll(pinsKnockedDown int) error {
	if pinsKnockedDown > 10 {
		return errors.New("cannot knock down more than 10 pins")
	}

	if g.frames == nil {
		g.frames = make([]frame, 1, 10)
	}

	lastFrameIndex := len(g.frames) - 1

	if g.spare {
		g.scoreCount += pinsKnockedDown
		g.spare = false
	}

	if g.strike == true {
		g.strikeRolls = append(g.strikeRolls, pinsKnockedDown)
		if len(g.strikeRolls) == 2 {
			for _, v := range g.strikeRolls {
				g.scoreCount += v
			}
			g.strike = false
			g.strikeRolls = []int{}
		}
	}

	if len(g.frames[lastFrameIndex].rolls) == 2 || lastFrameIndex == 0 {
		g.frames = append(g.frames, frame{rolls: []int{pinsKnockedDown}})

		if pinsKnockedDown == 10 {
			g.strike = true
		}
	} else {
		g.frames[lastFrameIndex].rolls = append(g.frames[lastFrameIndex].rolls,
			pinsKnockedDown)
		var totalScore int
		for _, v := range g.frames[lastFrameIndex].rolls {
			totalScore += v
		}
		if totalScore == 10 {
			g.spare = true
		}
	}

	g.scoreCount += pinsKnockedDown

	return nil
}

func (g *game) score() int {
	return g.scoreCount
}
