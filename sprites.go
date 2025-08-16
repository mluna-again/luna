package main

import _ "embed"

type asciiAnimation map[string][]string

type asciiPets map[string]asciiAnimation

//go:embed sprites/cat/idle-0.ascii
var catIdle0 string

//go:embed sprites/cat/idle-1.ascii
var catIdle1 string

//go:embed sprites/cat/idle-2.ascii
var catIdle2 string

//go:embed sprites/cat/idle-3.ascii
var catIdle3 string

//go:embed sprites/cat/idle-4.ascii
var catIdle4 string

//go:embed sprites/cat/idle-5.ascii
var catIdle5 string

//go:embed sprites/cat/idle-6.ascii
var catIdle6 string

//go:embed sprites/cat/sleeping-0.ascii
var catSleeping0 string

//go:embed sprites/cat/sleeping-1.ascii
var catSleeping1 string

//go:embed sprites/cat/sleeping-2.ascii
var catSleeping2 string

//go:embed sprites/cat/attacking-0.ascii
var catAttacking0 string

//go:embed sprites/cat/attacking-1.ascii
var catAttacking1 string

//go:embed sprites/cat/attacking-2.ascii
var catAttacking2 string

//go:embed sprites/cat/attacking-3.ascii
var catAttacking3 string

//go:embed sprites/cat/attacking-4.ascii
var catAttacking4 string

//go:embed sprites/cat/attacking-5.ascii
var catAttacking5 string

//go:embed sprites/cat/attacking-6.ascii
var catAttacking6 string

//go:embed sprites/cat/attacking-7.ascii
var catAttacking7 string

//go:embed sprites/cat/attacking-8.ascii
var catAttacking8 string

//go:embed sprites/turtle/idle-0.ascii
var turtleIdle0 string

//go:embed sprites/turtle/idle-1.ascii
var turtleIdle1 string

//go:embed sprites/turtle/idle-2.ascii
var turtleIdle2 string

//go:embed sprites/turtle/idle-3.ascii
var turtleIdle3 string

//go:embed sprites/turtle/idle-4.ascii
var turtleIdle4 string

//go:embed sprites/turtle/idle-5.ascii
var turtleIdle5 string

//go:embed sprites/turtle/idle-6.ascii
var turtleIdle6 string

//go:embed sprites/turtle/idle-7.ascii
var turtleIdle7 string

//go:embed sprites/turtle/sleeping-0.ascii
var turtleSleeping0 string

//go:embed sprites/turtle/sleeping-1.ascii
var turtleSleeping1 string

//go:embed sprites/turtle/sleeping-2.ascii
var turtleSleeping2 string

//go:embed sprites/turtle/sleeping-3.ascii
var turtleSleeping3 string

//go:embed sprites/turtle/sleeping-4.ascii
var turtleSleeping4 string

//go:embed sprites/turtle/sleeping-5.ascii
var turtleSleeping5 string

//go:embed sprites/turtle/sleeping-6.ascii
var turtleSleeping6 string

//go:embed sprites/turtle/sleeping-7.ascii
var turtleSleeping7 string

//go:embed sprites/turtle/sleeping-8.ascii
var turtleSleeping8 string

//go:embed sprites/turtle/sleeping-9.ascii
var turtleSleeping9 string

//go:embed sprites/turtle/sleeping-10.ascii
var turtleSleeping10 string

//go:embed sprites/turtle/sleeping-11.ascii
var turtleSleeping11 string

//go:embed sprites/turtle/attacking-0.ascii
var turtleAttacking0 string

//go:embed sprites/turtle/attacking-1.ascii
var turtleAttacking1 string

//go:embed sprites/turtle/attacking-2.ascii
var turtleAttacking2 string

//go:embed sprites/turtle/attacking-3.ascii
var turtleAttacking3 string

//go:embed sprites/turtle/attacking-4.ascii
var turtleAttacking4 string

//go:embed sprites/turtle/attacking-5.ascii
var turtleAttacking5 string

//go:embed sprites/turtle/attacking-6.ascii
var turtleAttacking6 string

//go:embed sprites/turtle/attacking-7.ascii
var turtleAttacking7 string

//go:embed sprites/turtle/attacking-8.ascii
var turtleAttacking8 string

//go:embed sprites/turtle/attacking-9.ascii
var turtleAttacking9 string

func getPets() asciiPets {
	return asciiPets{
		"turtle": asciiAnimation{
			"idle": []string{
				turtleIdle0,
				turtleIdle1,
				turtleIdle2,
				turtleIdle3,
				turtleIdle4,
				turtleIdle5,
				turtleIdle6,
				turtleIdle7,
			},
			"sleeping": []string{
				turtleSleeping0,
				turtleSleeping1,
				turtleSleeping2,
				turtleSleeping3,
				turtleSleeping4,
				turtleSleeping5,
				turtleSleeping6,
				turtleSleeping7,
				turtleSleeping8,
				turtleSleeping9,
				turtleSleeping10,
				turtleSleeping11,
			},
			"attacking": []string{
				turtleAttacking0,
				turtleAttacking1,
				turtleAttacking2,
				turtleAttacking3,
				turtleAttacking4,
				turtleAttacking5,
				turtleAttacking6,
				turtleAttacking7,
				turtleAttacking8,
				turtleAttacking9,
			},
		},
		"cat": asciiAnimation{
			"idle": []string{
				catIdle0,
				catIdle1,
				catIdle2,
				catIdle3,
				catIdle4,
				catIdle5,
				catIdle6,
			},
			"sleeping": []string{
				catSleeping0,
				catSleeping1,
				catSleeping2,
			},
			"attacking": []string{
				catAttacking0,
				catAttacking1,
				catAttacking2,
				catAttacking3,
				catAttacking4,
				catAttacking5,
				catAttacking6,
				catAttacking7,
				catAttacking8,
			},
		},
	}
}
