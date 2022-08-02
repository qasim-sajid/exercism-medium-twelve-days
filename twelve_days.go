package twelve

import (
	"fmt"
)

var christmasDays = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var christmasGifts = []string{"Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds", "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking", "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming"}
var christmasGiftsCount = []string{"a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}

func Verse(verseIndex int) string {
	verse := ""

	verseIndex--
	if verseIndex < 0 || verseIndex >= len(christmasDays) {
		return verse
	}

	verse += fmt.Sprintf("On the %s day of Christmas my true love gave to me:", christmasDays[verseIndex])
	for i := verseIndex; i >= 0; i-- {
		if i == 0 && verseIndex >= 1 {
			verse += " and"
		}
		verse += fmt.Sprintf(" %s %s", christmasGiftsCount[i], christmasGifts[i])
		if i == 0 {
			verse += "."
		} else {
			verse += ","
		}
	}

	return verse
}

func Song() string {
	song := ""
	for i := 1; i <= len(christmasDays); i++ {
		song += Verse(i)
		if i < len(christmasDays) {
			song += "\n"
		}
	}

	return song
}
