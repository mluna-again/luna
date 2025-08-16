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

func getPets() asciiPets {
	return asciiPets{
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
