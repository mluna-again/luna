package luna

import _ "embed"

type asciiAnimation map[LunaVariant]map[LunaSize]map[LunaAnimation][]string

type asciiPets map[LunaPet]asciiAnimation

//go:embed sprites/amogus/amogus0.ascii
var amogus0 string

//go:embed sprites/amogus/amogus1.ascii
var amogus1 string

//go:embed sprites/amogus/amogus2.ascii
var amogus2 string

//go:embed sprites/amogus/amogus3.ascii
var amogus3 string

//go:embed sprites/amogus/amogus4.ascii
var amogus4 string

//go:embed sprites/amogus/amogus5.ascii
var amogus5 string

//go:embed sprites/bunny/idle0_sm.ascii
var bunnyIdle0SM string

//go:embed sprites/bunny/idle0_md.ascii
var bunnyIdle0MD string

//go:embed sprites/bunny/idle0_xl.ascii
var bunnyIdle0XL string

//go:embed sprites/bunny/idle1_sm.ascii
var bunnyIdle1SM string

//go:embed sprites/bunny/idle1_md.ascii
var bunnyIdle1MD string

//go:embed sprites/bunny/idle1_xl.ascii
var bunnyIdle1XL string

//go:embed sprites/bunny/idle2_sm.ascii
var bunnyIdle2SM string

//go:embed sprites/bunny/idle2_md.ascii
var bunnyIdle2MD string

//go:embed sprites/bunny/idle2_xl.ascii
var bunnyIdle2XL string

//go:embed sprites/bunny/idle3_sm.ascii
var bunnyIdle3SM string

//go:embed sprites/bunny/idle3_md.ascii
var bunnyIdle3MD string

//go:embed sprites/bunny/idle3_xl.ascii
var bunnyIdle3XL string

//go:embed sprites/bunny/idle4_sm.ascii
var bunnyIdle4SM string

//go:embed sprites/bunny/idle4_md.ascii
var bunnyIdle4MD string

//go:embed sprites/bunny/idle4_xl.ascii
var bunnyIdle4XL string

//go:embed sprites/bunny/idle5_sm.ascii
var bunnyIdle5SM string

//go:embed sprites/bunny/idle5_md.ascii
var bunnyIdle5MD string

//go:embed sprites/bunny/idle5_xl.ascii
var bunnyIdle5XL string

//go:embed sprites/bunny/idle6_sm.ascii
var bunnyIdle6SM string

//go:embed sprites/bunny/idle6_md.ascii
var bunnyIdle6MD string

//go:embed sprites/bunny/idle6_xl.ascii
var bunnyIdle6XL string

//go:embed sprites/bunny/idle7_sm.ascii
var bunnyIdle7SM string

//go:embed sprites/bunny/idle7_md.ascii
var bunnyIdle7MD string

//go:embed sprites/bunny/idle7_xl.ascii
var bunnyIdle7XL string

//go:embed sprites/blackcat/idle0_sm.ascii
var blackcatIdle0SM string

//go:embed sprites/blackcat/idle0_md.ascii
var blackcatIdle0MD string

//go:embed sprites/blackcat/idle0_xl.ascii
var blackcatIdle0XL string

//go:embed sprites/blackcat/idle1_sm.ascii
var blackcatIdle1SM string

//go:embed sprites/blackcat/idle1_md.ascii
var blackcatIdle1MD string

//go:embed sprites/blackcat/idle1_xl.ascii
var blackcatIdle1XL string

//go:embed sprites/blackcat/idle2_sm.ascii
var blackcatIdle2SM string

//go:embed sprites/blackcat/idle2_md.ascii
var blackcatIdle2MD string

//go:embed sprites/blackcat/idle2_xl.ascii
var blackcatIdle2XL string

//go:embed sprites/blackcat/idle3_sm.ascii
var blackcatIdle3SM string

//go:embed sprites/blackcat/idle3_md.ascii
var blackcatIdle3MD string

//go:embed sprites/blackcat/idle3_xl.ascii
var blackcatIdle3XL string

//go:embed sprites/blackcat/idle4_sm.ascii
var blackcatIdle4SM string

//go:embed sprites/blackcat/idle4_md.ascii
var blackcatIdle4MD string

//go:embed sprites/blackcat/idle4_xl.ascii
var blackcatIdle4XL string

//go:embed sprites/blackcat/idle5_sm.ascii
var blackcatIdle5SM string

//go:embed sprites/blackcat/idle5_md.ascii
var blackcatIdle5MD string

//go:embed sprites/blackcat/idle5_xl.ascii
var blackcatIdle5XL string

//go:embed sprites/blackcat/idle6_sm.ascii
var blackcatIdle6SM string

//go:embed sprites/blackcat/idle6_md.ascii
var blackcatIdle6MD string

//go:embed sprites/blackcat/idle6_xl.ascii
var blackcatIdle6XL string

//go:embed sprites/blackcat/sleeping0_sm.ascii
var blackcatSleeping0SM string

//go:embed sprites/blackcat/sleeping0_md.ascii
var blackcatSleeping0MD string

//go:embed sprites/blackcat/sleeping0_xl.ascii
var blackcatSleeping0XL string

//go:embed sprites/blackcat/sleeping1_sm.ascii
var blackcatSleeping1SM string

//go:embed sprites/blackcat/sleeping1_md.ascii
var blackcatSleeping1MD string

//go:embed sprites/blackcat/sleeping1_xl.ascii
var blackcatSleeping1XL string

//go:embed sprites/blackcat/sleeping2_sm.ascii
var blackcatSleeping2SM string

//go:embed sprites/blackcat/sleeping2_md.ascii
var blackcatSleeping2MD string

//go:embed sprites/blackcat/sleeping2_xl.ascii
var blackcatSleeping2XL string

//go:embed sprites/blackcat/attacking0_sm.ascii
var blackcatAttacking0SM string

//go:embed sprites/blackcat/attacking0_md.ascii
var blackcatAttacking0MD string

//go:embed sprites/blackcat/attacking0_xl.ascii
var blackcatAttacking0XL string

//go:embed sprites/blackcat/attacking1_sm.ascii
var blackcatAttacking1SM string

//go:embed sprites/blackcat/attacking1_md.ascii
var blackcatAttacking1MD string

//go:embed sprites/blackcat/attacking1_xl.ascii
var blackcatAttacking1XL string

//go:embed sprites/blackcat/attacking2_sm.ascii
var blackcatAttacking2SM string

//go:embed sprites/blackcat/attacking2_md.ascii
var blackcatAttacking2MD string

//go:embed sprites/blackcat/attacking2_xl.ascii
var blackcatAttacking2XL string

//go:embed sprites/blackcat/attacking3_sm.ascii
var blackcatAttacking3SM string

//go:embed sprites/blackcat/attacking3_md.ascii
var blackcatAttacking3MD string

//go:embed sprites/blackcat/attacking3_xl.ascii
var blackcatAttacking3XL string

//go:embed sprites/blackcat/attacking4_sm.ascii
var blackcatAttacking4SM string

//go:embed sprites/blackcat/attacking4_md.ascii
var blackcatAttacking4MD string

//go:embed sprites/blackcat/attacking4_xl.ascii
var blackcatAttacking4XL string

//go:embed sprites/blackcat/attacking5_sm.ascii
var blackcatAttacking5SM string

//go:embed sprites/blackcat/attacking5_md.ascii
var blackcatAttacking5MD string

//go:embed sprites/blackcat/attacking5_xl.ascii
var blackcatAttacking5XL string

//go:embed sprites/blackcat/attacking6_sm.ascii
var blackcatAttacking6SM string

//go:embed sprites/blackcat/attacking6_md.ascii
var blackcatAttacking6MD string

//go:embed sprites/blackcat/attacking6_xl.ascii
var blackcatAttacking6XL string

//go:embed sprites/blackcat/attacking7_sm.ascii
var blackcatAttacking7SM string

//go:embed sprites/blackcat/attacking7_md.ascii
var blackcatAttacking7MD string

//go:embed sprites/blackcat/attacking7_xl.ascii
var blackcatAttacking7XL string

//go:embed sprites/blackcat/attacking8_sm.ascii
var blackcatAttacking8SM string

//go:embed sprites/blackcat/attacking8_md.ascii
var blackcatAttacking8MD string

//go:embed sprites/blackcat/attacking8_xl.ascii
var blackcatAttacking8XL string

//go:embed sprites/cat/idle0_sm.ascii
var catIdle0SM string

//go:embed sprites/cat/idle0_md.ascii
var catIdle0MD string

//go:embed sprites/cat/idle0_xl.ascii
var catIdle0XL string

//go:embed sprites/cat/idle1_sm.ascii
var catIdle1SM string

//go:embed sprites/cat/idle1_md.ascii
var catIdle1MD string

//go:embed sprites/cat/idle1_xl.ascii
var catIdle1XL string

//go:embed sprites/cat/idle2_sm.ascii
var catIdle2SM string

//go:embed sprites/cat/idle2_md.ascii
var catIdle2MD string

//go:embed sprites/cat/idle2_xl.ascii
var catIdle2XL string

//go:embed sprites/cat/idle3_sm.ascii
var catIdle3SM string

//go:embed sprites/cat/idle3_md.ascii
var catIdle3MD string

//go:embed sprites/cat/idle3_xl.ascii
var catIdle3XL string

//go:embed sprites/cat/idle4_sm.ascii
var catIdle4SM string

//go:embed sprites/cat/idle4_md.ascii
var catIdle4MD string

//go:embed sprites/cat/idle4_xl.ascii
var catIdle4XL string

//go:embed sprites/cat/idle5_sm.ascii
var catIdle5SM string

//go:embed sprites/cat/idle5_md.ascii
var catIdle5MD string

//go:embed sprites/cat/idle5_xl.ascii
var catIdle5XL string

//go:embed sprites/cat/idle6_sm.ascii
var catIdle6SM string

//go:embed sprites/cat/idle6_md.ascii
var catIdle6MD string

//go:embed sprites/cat/idle6_xl.ascii
var catIdle6XL string

//go:embed sprites/turtle/idle0_sm.ascii
var turtleIdle0SM string

//go:embed sprites/turtle/idle0_md.ascii
var turtleIdle0MD string

//go:embed sprites/turtle/idle0_xl.ascii
var turtleIdle0XL string

//go:embed sprites/turtle/idle1_sm.ascii
var turtleIdle1SM string

//go:embed sprites/turtle/idle1_md.ascii
var turtleIdle1MD string

//go:embed sprites/turtle/idle1_xl.ascii
var turtleIdle1XL string

//go:embed sprites/turtle/idle2_sm.ascii
var turtleIdle2SM string

//go:embed sprites/turtle/idle2_md.ascii
var turtleIdle2MD string

//go:embed sprites/turtle/idle2_xl.ascii
var turtleIdle2XL string

//go:embed sprites/turtle/idle3_sm.ascii
var turtleIdle3SM string

//go:embed sprites/turtle/idle3_md.ascii
var turtleIdle3MD string

//go:embed sprites/turtle/idle3_xl.ascii
var turtleIdle3XL string

//go:embed sprites/turtle/idle4_sm.ascii
var turtleIdle4SM string

//go:embed sprites/turtle/idle4_md.ascii
var turtleIdle4MD string

//go:embed sprites/turtle/idle4_xl.ascii
var turtleIdle4XL string

//go:embed sprites/turtle/idle5_sm.ascii
var turtleIdle5SM string

//go:embed sprites/turtle/idle5_md.ascii
var turtleIdle5MD string

//go:embed sprites/turtle/idle5_xl.ascii
var turtleIdle5XL string

//go:embed sprites/turtle/idle6_sm.ascii
var turtleIdle6SM string

//go:embed sprites/turtle/idle6_md.ascii
var turtleIdle6MD string

//go:embed sprites/turtle/idle6_xl.ascii
var turtleIdle6XL string

//go:embed sprites/turtle/idle7_sm.ascii
var turtleIdle7SM string

//go:embed sprites/turtle/idle7_md.ascii
var turtleIdle7MD string

//go:embed sprites/turtle/idle7_xl.ascii
var turtleIdle7XL string

//go:embed sprites/bunny/sleeping0_sm.ascii
var bunnySleeping0SM string

//go:embed sprites/bunny/sleeping0_md.ascii
var bunnySleeping0MD string

//go:embed sprites/bunny/sleeping0_xl.ascii
var bunnySleeping0XL string

//go:embed sprites/bunny/sleeping1_sm.ascii
var bunnySleeping1SM string

//go:embed sprites/bunny/sleeping1_md.ascii
var bunnySleeping1MD string

//go:embed sprites/bunny/sleeping1_xl.ascii
var bunnySleeping1XL string

//go:embed sprites/cat/sleeping0_sm.ascii
var catSleeping0SM string

//go:embed sprites/cat/sleeping0_md.ascii
var catSleeping0MD string

//go:embed sprites/cat/sleeping0_xl.ascii
var catSleeping0XL string

//go:embed sprites/cat/sleeping1_sm.ascii
var catSleeping1SM string

//go:embed sprites/cat/sleeping1_md.ascii
var catSleeping1MD string

//go:embed sprites/cat/sleeping1_xl.ascii
var catSleeping1XL string

//go:embed sprites/cat/sleeping2_sm.ascii
var catSleeping2SM string

//go:embed sprites/cat/sleeping2_md.ascii
var catSleeping2MD string

//go:embed sprites/cat/sleeping2_xl.ascii
var catSleeping2XL string

//go:embed sprites/turtle/sleeping0_sm.ascii
var turtleSleeping0SM string

//go:embed sprites/turtle/sleeping0_md.ascii
var turtleSleeping0MD string

//go:embed sprites/turtle/sleeping0_xl.ascii
var turtleSleeping0XL string

//go:embed sprites/turtle/sleeping1_sm.ascii
var turtleSleeping1SM string

//go:embed sprites/turtle/sleeping1_md.ascii
var turtleSleeping1MD string

//go:embed sprites/turtle/sleeping1_xl.ascii
var turtleSleeping1XL string

//go:embed sprites/turtle/sleeping2_sm.ascii
var turtleSleeping2SM string

//go:embed sprites/turtle/sleeping2_md.ascii
var turtleSleeping2MD string

//go:embed sprites/turtle/sleeping2_xl.ascii
var turtleSleeping2XL string

//go:embed sprites/turtle/sleeping3_sm.ascii
var turtleSleeping3SM string

//go:embed sprites/turtle/sleeping3_md.ascii
var turtleSleeping3MD string

//go:embed sprites/turtle/sleeping3_xl.ascii
var turtleSleeping3XL string

//go:embed sprites/turtle/sleeping4_sm.ascii
var turtleSleeping4SM string

//go:embed sprites/turtle/sleeping4_md.ascii
var turtleSleeping4MD string

//go:embed sprites/turtle/sleeping4_xl.ascii
var turtleSleeping4XL string

//go:embed sprites/turtle/sleeping5_sm.ascii
var turtleSleeping5SM string

//go:embed sprites/turtle/sleeping5_md.ascii
var turtleSleeping5MD string

//go:embed sprites/turtle/sleeping5_xl.ascii
var turtleSleeping5XL string

//go:embed sprites/turtle/sleeping6_sm.ascii
var turtleSleeping6SM string

//go:embed sprites/turtle/sleeping6_md.ascii
var turtleSleeping6MD string

//go:embed sprites/turtle/sleeping6_xl.ascii
var turtleSleeping6XL string

//go:embed sprites/turtle/sleeping7_sm.ascii
var turtleSleeping7SM string

//go:embed sprites/turtle/sleeping7_md.ascii
var turtleSleeping7MD string

//go:embed sprites/turtle/sleeping7_xl.ascii
var turtleSleeping7XL string

//go:embed sprites/turtle/sleeping8_sm.ascii
var turtleSleeping8SM string

//go:embed sprites/turtle/sleeping8_md.ascii
var turtleSleeping8MD string

//go:embed sprites/turtle/sleeping8_xl.ascii
var turtleSleeping8XL string

//go:embed sprites/turtle/sleeping9_sm.ascii
var turtleSleeping9SM string

//go:embed sprites/turtle/sleeping9_md.ascii
var turtleSleeping9MD string

//go:embed sprites/turtle/sleeping9_xl.ascii
var turtleSleeping9XL string

//go:embed sprites/turtle/sleeping10_sm.ascii
var turtleSleeping10SM string

//go:embed sprites/turtle/sleeping10_md.ascii
var turtleSleeping10MD string

//go:embed sprites/turtle/sleeping10_xl.ascii
var turtleSleeping10XL string

//go:embed sprites/turtle/sleeping11_sm.ascii
var turtleSleeping11SM string

//go:embed sprites/turtle/sleeping11_md.ascii
var turtleSleeping11MD string

//go:embed sprites/turtle/sleeping11_xl.ascii
var turtleSleeping11XL string

//go:embed sprites/bunny/attacking0_sm.ascii
var bunnyAttacking0SM string

//go:embed sprites/bunny/attacking0_md.ascii
var bunnyAttacking0MD string

//go:embed sprites/bunny/attacking0_xl.ascii
var bunnyAttacking0XL string

//go:embed sprites/bunny/attacking1_sm.ascii
var bunnyAttacking1SM string

//go:embed sprites/bunny/attacking1_md.ascii
var bunnyAttacking1MD string

//go:embed sprites/bunny/attacking1_xl.ascii
var bunnyAttacking1XL string

//go:embed sprites/bunny/attacking2_sm.ascii
var bunnyAttacking2SM string

//go:embed sprites/bunny/attacking2_md.ascii
var bunnyAttacking2MD string

//go:embed sprites/bunny/attacking2_xl.ascii
var bunnyAttacking2XL string

//go:embed sprites/bunny/attacking3_sm.ascii
var bunnyAttacking3SM string

//go:embed sprites/bunny/attacking3_md.ascii
var bunnyAttacking3MD string

//go:embed sprites/bunny/attacking3_xl.ascii
var bunnyAttacking3XL string

//go:embed sprites/bunny/attacking4_sm.ascii
var bunnyAttacking4SM string

//go:embed sprites/bunny/attacking4_md.ascii
var bunnyAttacking4MD string

//go:embed sprites/bunny/attacking4_xl.ascii
var bunnyAttacking4XL string

//go:embed sprites/bunny/attacking5_sm.ascii
var bunnyAttacking5SM string

//go:embed sprites/bunny/attacking5_md.ascii
var bunnyAttacking5MD string

//go:embed sprites/bunny/attacking5_xl.ascii
var bunnyAttacking5XL string

//go:embed sprites/bunny/attacking6_sm.ascii
var bunnyAttacking6SM string

//go:embed sprites/bunny/attacking6_md.ascii
var bunnyAttacking6MD string

//go:embed sprites/bunny/attacking6_xl.ascii
var bunnyAttacking6XL string

//go:embed sprites/cat/attacking0_sm.ascii
var catAttacking0SM string

//go:embed sprites/cat/attacking0_md.ascii
var catAttacking0MD string

//go:embed sprites/cat/attacking0_xl.ascii
var catAttacking0XL string

//go:embed sprites/cat/attacking1_sm.ascii
var catAttacking1SM string

//go:embed sprites/cat/attacking1_md.ascii
var catAttacking1MD string

//go:embed sprites/cat/attacking1_xl.ascii
var catAttacking1XL string

//go:embed sprites/cat/attacking2_sm.ascii
var catAttacking2SM string

//go:embed sprites/cat/attacking2_md.ascii
var catAttacking2MD string

//go:embed sprites/cat/attacking2_xl.ascii
var catAttacking2XL string

//go:embed sprites/cat/attacking3_sm.ascii
var catAttacking3SM string

//go:embed sprites/cat/attacking3_md.ascii
var catAttacking3MD string

//go:embed sprites/cat/attacking3_xl.ascii
var catAttacking3XL string

//go:embed sprites/cat/attacking4_sm.ascii
var catAttacking4SM string

//go:embed sprites/cat/attacking4_md.ascii
var catAttacking4MD string

//go:embed sprites/cat/attacking4_xl.ascii
var catAttacking4XL string

//go:embed sprites/cat/attacking5_sm.ascii
var catAttacking5SM string

//go:embed sprites/cat/attacking5_md.ascii
var catAttacking5MD string

//go:embed sprites/cat/attacking5_xl.ascii
var catAttacking5XL string

//go:embed sprites/cat/attacking6_sm.ascii
var catAttacking6SM string

//go:embed sprites/cat/attacking6_md.ascii
var catAttacking6MD string

//go:embed sprites/cat/attacking6_xl.ascii
var catAttacking6XL string

//go:embed sprites/cat/attacking7_sm.ascii
var catAttacking7SM string

//go:embed sprites/cat/attacking7_md.ascii
var catAttacking7MD string

//go:embed sprites/cat/attacking7_xl.ascii
var catAttacking7XL string

//go:embed sprites/cat/attacking8_sm.ascii
var catAttacking8SM string

//go:embed sprites/cat/attacking8_md.ascii
var catAttacking8MD string

//go:embed sprites/cat/attacking8_xl.ascii
var catAttacking8XL string

//go:embed sprites/turtle/attacking0_sm.ascii
var turtleAttacking0SM string

//go:embed sprites/turtle/attacking0_md.ascii
var turtleAttacking0MD string

//go:embed sprites/turtle/attacking0_xl.ascii
var turtleAttacking0XL string

//go:embed sprites/turtle/attacking1_sm.ascii
var turtleAttacking1SM string

//go:embed sprites/turtle/attacking1_md.ascii
var turtleAttacking1MD string

//go:embed sprites/turtle/attacking1_xl.ascii
var turtleAttacking1XL string

//go:embed sprites/turtle/attacking2_sm.ascii
var turtleAttacking2SM string

//go:embed sprites/turtle/attacking2_md.ascii
var turtleAttacking2MD string

//go:embed sprites/turtle/attacking2_xl.ascii
var turtleAttacking2XL string

//go:embed sprites/turtle/attacking3_sm.ascii
var turtleAttacking3SM string

//go:embed sprites/turtle/attacking3_md.ascii
var turtleAttacking3MD string

//go:embed sprites/turtle/attacking3_xl.ascii
var turtleAttacking3XL string

//go:embed sprites/turtle/attacking4_sm.ascii
var turtleAttacking4SM string

//go:embed sprites/turtle/attacking4_md.ascii
var turtleAttacking4MD string

//go:embed sprites/turtle/attacking4_xl.ascii
var turtleAttacking4XL string

//go:embed sprites/turtle/attacking5_sm.ascii
var turtleAttacking5SM string

//go:embed sprites/turtle/attacking5_md.ascii
var turtleAttacking5MD string

//go:embed sprites/turtle/attacking5_xl.ascii
var turtleAttacking5XL string

//go:embed sprites/turtle/attacking6_sm.ascii
var turtleAttacking6SM string

//go:embed sprites/turtle/attacking6_md.ascii
var turtleAttacking6MD string

//go:embed sprites/turtle/attacking6_xl.ascii
var turtleAttacking6XL string

//go:embed sprites/turtle/attacking7_sm.ascii
var turtleAttacking7SM string

//go:embed sprites/turtle/attacking7_md.ascii
var turtleAttacking7MD string

//go:embed sprites/turtle/attacking7_xl.ascii
var turtleAttacking7XL string

//go:embed sprites/turtle/attacking8_sm.ascii
var turtleAttacking8SM string

//go:embed sprites/turtle/attacking8_md.ascii
var turtleAttacking8MD string

//go:embed sprites/turtle/attacking8_xl.ascii
var turtleAttacking8XL string

//go:embed sprites/turtle/attacking9_sm.ascii
var turtleAttacking9SM string

//go:embed sprites/turtle/attacking9_md.ascii
var turtleAttacking9MD string

//go:embed sprites/turtle/attacking9_xl.ascii
var turtleAttacking9XL string

// Lord have mercy
func getPets() asciiPets {
	return asciiPets{
		"amogus": asciiAnimation{
			"default": map[LunaSize]map[LunaAnimation][]string{
				"sm": map[LunaAnimation][]string{
					"idle": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
					"sleeping": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
					"attacking": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
				},
				"md": map[LunaAnimation][]string{
					"idle": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
					"sleeping": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
					"attacking": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
				},
				"xl": map[LunaAnimation][]string{
					"idle": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
					"sleeping": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
					"attacking": []string{
						amogus0,
						amogus1,
						amogus2,
						amogus3,
						amogus4,
						amogus5,
					},
				},
			},
		},
		"bunny": asciiAnimation{
			"default": map[LunaSize]map[LunaAnimation][]string{
				"sm": map[LunaAnimation][]string{
					"idle": []string{
						bunnyIdle0SM,
						bunnyIdle1SM,
						bunnyIdle2SM,
						bunnyIdle3SM,
						bunnyIdle5SM,
						bunnyIdle6SM,
						bunnyIdle7SM,
					},
					"sleeping": []string{
						bunnySleeping0SM,
						bunnySleeping1SM,
					},
					"attacking": []string{
						bunnyAttacking0SM,
						bunnyAttacking1SM,
						bunnyAttacking2SM,
						bunnyAttacking3SM,
						bunnyAttacking4SM,
						bunnyAttacking5SM,
						bunnyAttacking6SM,
					},
				},
				"md": map[LunaAnimation][]string{
					"idle": []string{
						bunnyIdle0MD,
						bunnyIdle1MD,
						bunnyIdle2MD,
						bunnyIdle3MD,
						bunnyIdle5MD,
						bunnyIdle6MD,
						bunnyIdle7MD,
					},
					"sleeping": []string{
						bunnySleeping0MD,
						bunnySleeping1MD,
					},
					"attacking": []string{
						bunnyAttacking0MD,
						bunnyAttacking1MD,
						bunnyAttacking2MD,
						bunnyAttacking3MD,
						bunnyAttacking4MD,
						bunnyAttacking5MD,
						bunnyAttacking6MD,
					},
				},
				"xl": map[LunaAnimation][]string{
					"idle": []string{
						bunnyIdle0XL,
						bunnyIdle1XL,
						bunnyIdle2XL,
						bunnyIdle3XL,
						bunnyIdle5XL,
						bunnyIdle6XL,
						bunnyIdle7XL,
					},
					"sleeping": []string{
						bunnySleeping0XL,
						bunnySleeping1XL,
					},
					"attacking": []string{
						bunnyAttacking0XL,
						bunnyAttacking1XL,
						bunnyAttacking2XL,
						bunnyAttacking3XL,
						bunnyAttacking4XL,
						bunnyAttacking5XL,
						bunnyAttacking6XL,
					},
				},
			},
		},
		"turtle": asciiAnimation{
			"default": map[LunaSize]map[LunaAnimation][]string{
				"sm": map[LunaAnimation][]string{
					"idle": []string{
						turtleIdle0SM,
						turtleIdle1SM,
						turtleIdle2SM,
						turtleIdle3SM,
						turtleIdle4SM,
						turtleIdle5SM,
						turtleIdle6SM,
						turtleIdle7SM,
					},
					"sleeping": []string{
						turtleSleeping0SM,
						turtleSleeping1SM,
						turtleSleeping2SM,
						turtleSleeping3SM,
						turtleSleeping4SM,
						turtleSleeping5SM,
						turtleSleeping6SM,
						turtleSleeping7SM,
						turtleSleeping8SM,
						turtleSleeping9SM,
						turtleSleeping10SM,
						turtleSleeping11SM,
					},
					"attacking": []string{
						turtleAttacking0SM,
						turtleAttacking1SM,
						turtleAttacking2SM,
						turtleAttacking3SM,
						turtleAttacking4SM,
						turtleAttacking5SM,
						turtleAttacking6SM,
						turtleAttacking7SM,
						turtleAttacking8SM,
						turtleAttacking9SM,
					},
				},
				"md": map[LunaAnimation][]string{
					"idle": []string{
						turtleIdle0MD,
						turtleIdle1MD,
						turtleIdle2MD,
						turtleIdle3MD,
						turtleIdle4MD,
						turtleIdle5MD,
						turtleIdle6MD,
						turtleIdle7MD,
					},
					"sleeping": []string{
						turtleSleeping0MD,
						turtleSleeping1MD,
						turtleSleeping2MD,
						turtleSleeping3MD,
						turtleSleeping4MD,
						turtleSleeping5MD,
						turtleSleeping6MD,
						turtleSleeping7MD,
						turtleSleeping8MD,
						turtleSleeping9MD,
						turtleSleeping10MD,
						turtleSleeping11MD,
					},
					"attacking": []string{
						turtleAttacking0MD,
						turtleAttacking1MD,
						turtleAttacking2MD,
						turtleAttacking3MD,
						turtleAttacking4MD,
						turtleAttacking5MD,
						turtleAttacking6MD,
						turtleAttacking7MD,
						turtleAttacking8MD,
						turtleAttacking9MD,
					},
				},
				"xl": map[LunaAnimation][]string{
					"idle": []string{
						turtleIdle0XL,
						turtleIdle1XL,
						turtleIdle2XL,
						turtleIdle3XL,
						turtleIdle4XL,
						turtleIdle5XL,
						turtleIdle6XL,
						turtleIdle7XL,
					},
					"sleeping": []string{
						turtleSleeping0XL,
						turtleSleeping1XL,
						turtleSleeping2XL,
						turtleSleeping3XL,
						turtleSleeping4XL,
						turtleSleeping5XL,
						turtleSleeping6XL,
						turtleSleeping7XL,
						turtleSleeping8XL,
						turtleSleeping9XL,
						turtleSleeping10XL,
						turtleSleeping11XL,
					},
					"attacking": []string{
						turtleAttacking0XL,
						turtleAttacking1XL,
						turtleAttacking2XL,
						turtleAttacking3XL,
						turtleAttacking4XL,
						turtleAttacking5XL,
						turtleAttacking6XL,
						turtleAttacking7XL,
						turtleAttacking8XL,
						turtleAttacking9XL,
					},
				},
			},
		},
		"cat": asciiAnimation{
			"ragdoll": map[LunaSize]map[LunaAnimation][]string{
				"sm": map[LunaAnimation][]string{
					"idle": []string{
						catIdle0SM,
						catIdle1SM,
						catIdle2SM,
						catIdle3SM,
						catIdle4SM,
						catIdle5SM,
						catIdle6SM,
					},
					"sleeping": []string{
						catSleeping0SM,
						catSleeping1SM,
						catSleeping2SM,
					},
					"attacking": []string{
						catAttacking0SM,
						catAttacking1SM,
						catAttacking2SM,
						catAttacking3SM,
						catAttacking4SM,
						catAttacking5SM,
						catAttacking6SM,
						catAttacking7SM,
						catAttacking8SM,
					},
				},
				"md": map[LunaAnimation][]string{
					"idle": []string{
						catIdle0MD,
						catIdle1MD,
						catIdle2MD,
						catIdle3MD,
						catIdle4MD,
						catIdle5MD,
						catIdle6MD,
					},
					"sleeping": []string{
						catSleeping0MD,
						catSleeping1MD,
						catSleeping2MD,
					},
					"attacking": []string{
						catAttacking0MD,
						catAttacking1MD,
						catAttacking2MD,
						catAttacking3MD,
						catAttacking4MD,
						catAttacking5MD,
						catAttacking6MD,
						catAttacking7MD,
						catAttacking8MD,
					},
				},
				"xl": map[LunaAnimation][]string{
					"idle": []string{
						catIdle0XL,
						catIdle1XL,
						catIdle2XL,
						catIdle3XL,
						catIdle4XL,
						catIdle5XL,
						catIdle6XL,
					},
					"sleeping": []string{
						catSleeping0XL,
						catSleeping1XL,
						catSleeping2XL,
					},
					"attacking": []string{
						catAttacking0XL,
						catAttacking1XL,
						catAttacking2XL,
						catAttacking3XL,
						catAttacking4XL,
						catAttacking5XL,
						catAttacking6XL,
						catAttacking7XL,
						catAttacking8XL,
					},
				},
			},
			"black": map[LunaSize]map[LunaAnimation][]string{
				"sm": map[LunaAnimation][]string{
					"idle": []string{
						blackcatIdle0SM,
						blackcatIdle1SM,
						blackcatIdle2SM,
						blackcatIdle3SM,
						blackcatIdle4SM,
						blackcatIdle5SM,
						blackcatIdle6SM,
					},
					"sleeping": []string{
						blackcatSleeping0SM,
						blackcatSleeping1SM,
						blackcatSleeping2SM,
					},
					"attacking": []string{
						blackcatAttacking0SM,
						blackcatAttacking1SM,
						blackcatAttacking2SM,
						blackcatAttacking3SM,
						blackcatAttacking4SM,
						blackcatAttacking5SM,
						blackcatAttacking6SM,
						blackcatAttacking7SM,
						blackcatAttacking8SM,
					},
				},
				"md": map[LunaAnimation][]string{
					"idle": []string{
						blackcatIdle0MD,
						blackcatIdle1MD,
						blackcatIdle2MD,
						blackcatIdle3MD,
						blackcatIdle4MD,
						blackcatIdle5MD,
						blackcatIdle6MD,
					},
					"sleeping": []string{
						blackcatSleeping0MD,
						blackcatSleeping1MD,
						blackcatSleeping2MD,
					},
					"attacking": []string{
						blackcatAttacking0MD,
						blackcatAttacking1MD,
						blackcatAttacking2MD,
						blackcatAttacking3MD,
						blackcatAttacking4MD,
						blackcatAttacking5MD,
						blackcatAttacking6MD,
						blackcatAttacking7MD,
						blackcatAttacking8MD,
					},
				},
				"xl": map[LunaAnimation][]string{
					"idle": []string{
						blackcatIdle0XL,
						blackcatIdle1XL,
						blackcatIdle2XL,
						blackcatIdle3XL,
						blackcatIdle4XL,
						blackcatIdle5XL,
						blackcatIdle6XL,
					},
					"sleeping": []string{
						blackcatSleeping0XL,
						blackcatSleeping1XL,
						blackcatSleeping2XL,
					},
					"attacking": []string{
						blackcatAttacking0XL,
						blackcatAttacking1XL,
						blackcatAttacking2XL,
						blackcatAttacking3XL,
						blackcatAttacking4XL,
						blackcatAttacking5XL,
						blackcatAttacking6XL,
						blackcatAttacking7XL,
						blackcatAttacking8XL,
					},
				},
			},
		},
	}
}
