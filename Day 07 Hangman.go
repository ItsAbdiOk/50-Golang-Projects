package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var stages = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`,
}

var wordList = []string{
	"abruptly",
	"absurd",
	"abyss",
	"affix",
	"askew",
	"avenue",
	"awkward",
	"axiom",
	"azure",
	"bagpipes",
	"bandwagon",
	"banjo",
	"bayou",
	"beekeeper",
	"bikini",
	"blitz",
	"blizzard",
	"boggle",
	"bookworm",
	"boxcar",
	"boxful",
	"buckaroo",
	"buffalo",
	"buffoon",
	"buxom",
	"buzzard",
	"buzzing",
	"buzzwords",
	"caliph",
	"cobweb",
	"cockiness",
	"croquet",
	"crypt",
	"curacao",
	"cycle",
	"daiquiri",
	"dirndl",
	"disavow",
	"dizzying",
	"duplex",
	"dwarves",
	"embezzle",
	"equip",
	"espionage",
	"euouae",
	"exodus",
	"faking",
	"fishhook",
	"fixable",
	"fjord",
	"flapjack",
	"flopping",
	"fluffiness",
	"flyby",
	"foxglove",
	"frazzled",
	"frizzled",
	"fuchsia",
	"funny",
	"gabby",
	"galaxy",
	"galvanize",
	"gazebo",
	"giaour",
	"gizmo",
	"glowworm",
	"glyph",
	"gnarly",
	"gnostic",
	"gossip",
	"grogginess",
	"haiku",
	"haphazard",
	"hyphen",
	"iatrogenic",
	"icebox",
	"injury",
	"ivory",
	"ivy",
	"jackpot",
	"jaundice",
	"jawbreaker",
	"jaywalk",
	"jazziest",
	"jazzy",
	"jelly",
	"jigsaw",
	"jinx",
	"jiujitsu",
	"jockey",
	"jogging",
	"joking",
	"jovial",
	"joyful",
	"juicy",
	"jukebox",
	"jumbo",
	"kayak",
	"kazoo",
	"keyhole",
	"khaki",
	"kilobyte",
	"kiosk",
	"kitsch",
	"kiwifruit",
	"klutz",
	"knapsack",
	"larynx",
	"lengths",
	"lucky",
	"luxury",
	"lymph",
	"marquis",
	"matrix",
	"megahertz",
	"microwave",
	"mnemonic",
	"mystify",
	"naphtha",
	"nightclub",
	"nowadays",
	"numbskull",
	"nymph",
	"onyx",
	"ovary",
	"oxidize",
	"oxygen",
	"pajama",
	"peekaboo",
	"phlegm",
	"pixel",
	"pizazz",
	"pneumonia",
	"polka",
	"pshaw",
	"psyche",
	"puppy",
	"puzzling",
	"quartz",
	"queue",
	"quips",
	"quixotic",
	"quiz",
	"quizzes",
	"quorum",
	"razzmatazz",
	"rhubarb",
	"rhythm",
	"rickshaw",
	"schnapps",
	"scratch",
	"shiv",
	"snazzy",
	"sphinx",
	"spritz",
	"squawk",
	"staff",
	"strength",
	"strengths",
	"stretch",
	"stronghold",
	"stymied",
	"subway",
	"swivel",
	"syndrome",
	"thriftless",
	"thumbscrew",
	"topaz",
	"transcript",
	"transgress",
	"transplant",
	"triphthong",
	"twelfth",
	"twelfths",
	"unknown",
	"unworthy",
	"unzip",
	"uptown",
	"vaporize",
	"vixen",
	"vodka",
	"voodoo",
	"vortex",
	"voyeurism",
	"walkway",
	"waltz",
	"wave",
	"wavy",
	"waxy",
	"wellspring",
	"wheezy",
	"whiskey",
	"whizzing",
	"whomever",
	"wimpy",
	"witchcraft",
	"wizard",
	"woozy",
	"wristwatch",
	"wyvern",
	"xylophone",
	"yachtsman",
	"yippee",
	"yoked",
	"youthful",
	"yummy",
	"zephyr",
	"zigzag",
	"zigzagging",
	"zilch",
	"zipper",
	"zodiac",
	"zombie",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	gameComplete := false
	chosenWord := wordList[rand.Intn(len(wordList))]

	lettersGuessed := len(chosenWord)
	blankChosenWord := strings.Repeat("_", len(chosenWord))

	lives := 6

	fmt.Println("H A N G M A N")
	for !gameComplete {
		var guess string
		fmt.Print("Guess a letter: ")
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)

		temp := lettersGuessed
		for i := 0; i < len(chosenWord); i++ {
			if guess == string(chosenWord[i]) {
				blankChosenWord = blankChosenWord[:i] + guess + blankChosenWord[i+1:]
				lettersGuessed--
			}
		}

		if temp != lettersGuessed {
			fmt.Println("Correct guess")
		} else {
			fmt.Println("Incorrect guess")
			lives--
			fmt.Println(stages[lives])
		}

		fmt.Println(blankChosenWord)

		if lettersGuessed == 0 {
			gameComplete = true
			fmt.Println("You WIN")
		} else if lives == 0 {
			gameComplete = true
			fmt.Printf("You Lose, the word was %s\n", chosenWord)
		}
	}
}
