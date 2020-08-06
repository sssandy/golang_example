package main

const (
	win = 100 // he winning score in a game of Pig
)

type score struct {
	player int
	opponent int
	thisTurn int
}

//
type action func(current score) (result score, turnIsOver bool)

