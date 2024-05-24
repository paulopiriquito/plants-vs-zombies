package game

import "plants-vs-zombies/lawn"

type Game interface {
	Start() error
	Lawn() lawn.Lawn
}

type game struct {
	lawn lawn.Lawn
}

func (g game) Lawn() lawn.Lawn {
	return g.lawn
}

func (g game) Start() error {
	return nil
}

func NewGame(lawnMap []string) Game {
	return game{lawn: lawn.New(lawnMap)}
}