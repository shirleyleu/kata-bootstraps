package main

import "errors"

type game struct {
	frames     []frame
	scoreCount int
	spare      bool
}

type frame struct {
	rolls []int
}

func (g *game) roll(n int) error {
	if n > 10 {
		return errors.New("cannot knock down more than 10 pins")
	}

	if g.frames == nil {
		g.frames = make([]frame, 1)
	}

	lastFrameIndex := len(g.frames) - 1

	if g.spare {
		g.scoreCount += n
		g.spare = false
	}

	if len(g.frames[lastFrameIndex].rolls) == 2 {
		g.frames = append(g.frames, frame{rolls: []int{n}})
	} else {
		g.frames[lastFrameIndex].rolls = append(g.frames[lastFrameIndex].rolls, n)
		var totalScore int
		for _, v := range g.frames[lastFrameIndex].rolls {
			totalScore += v
		}
		if totalScore == 10 {
			g.spare = true
		}
	}

	g.scoreCount += n

	return nil
}

func (g *game) score() int {
	return g.scoreCount
}
