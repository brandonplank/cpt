/*

	This project was created during my 2022 school year for the Create Performance Task for
	AP Computer Science

	All rights are relinquished to the CPT management and staff for use for grading.

	CPT License, By Brandon Plank.
	Brandon Plank © 2022, All Rights reserved.


	* This code MAY be used for demonstration

	* This code may only be viewed by the CPT graders until after
	the grade has been entered

	* This code may NOT be published ander any name, but my own

	* This code may only be public AFTER the grading process

	* This code may NOT be used in any other public or private project

	By reviewing my code, you agree to this license.

*/

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"regexp"
	"strings"
)

var (
	CharToMorse = make(map[string]string)
	MorseToChar = make(map[string]string)
)

func FlipMap(m map[string]string) map[string]string {
	// Make a new map for the return with the length of m
	ret := make(map[string]string, len(m))
	for index, object := range m {
		ret[object] = index
	}
	return ret
}

// init Sets all the morse codes and flips it into another map
func init() {
	CharToMorse["A"] = ".-"
	CharToMorse["B"] = "-..."
	CharToMorse["C"] = "-.-."
	CharToMorse["D"] = "-.."
	CharToMorse["E"] = "."
	CharToMorse["F"] = "..-."
	CharToMorse["G"] = "--."
	CharToMorse["H"] = "...."
	CharToMorse["I"] = ".."
	CharToMorse["J"] = ".---"
	CharToMorse["K"] = "-.-"
	CharToMorse["L"] = ".-.."
	CharToMorse["M"] = "--"
	CharToMorse["N"] = "-."
	CharToMorse["O"] = "---"
	CharToMorse["P"] = ".--."
	CharToMorse["Q"] = "--.-"
	CharToMorse["R"] = ".-."
	CharToMorse["S"] = "..."
	CharToMorse["T"] = "-"
	CharToMorse["U"] = "..-"
	CharToMorse["V"] = "...-"
	CharToMorse["W"] = ".--"
	CharToMorse["X"] = "-..-"
	CharToMorse["Y"] = "-.--"
	CharToMorse["Z"] = "--.."

	CharToMorse["1"] = ".----"
	CharToMorse["2"] = "..---"
	CharToMorse["3"] = "...--"
	CharToMorse["4"] = "....-"
	CharToMorse["5"] = "....."
	CharToMorse["6"] = "-...."
	CharToMorse["7"] = "--..."
	CharToMorse["8"] = "---.."
	CharToMorse["9"] = "----."
	CharToMorse["0"] = "-----"

	CharToMorse["."] = ".-.-.-"
	CharToMorse[","] = "--..--"
	CharToMorse["?"] = "..--.."

	// Space, this is normalized in typed morse code
	CharToMorse[" "] = "/"

	MorseToChar = FlipMap(CharToMorse)
}

// IsMorseValidStageTwo Loops through all codes to make sure they match with the MAP
func IsMorseValidStageTwo(m string) bool {
	codes := strings.Split(m, " ")
	for i := 0; i < len(codes); i++ {
		if MorseToChar[codes[i]] == "" {
			return false
		}
	}
	return true
}

// For all regex, I used https://regex101.com/ to make build and test custom regex

// IsStringValid Makes sure that our translator can parse this string
func IsStringValid(m string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9.,?]+( [a-zA-Z0-9.,?]+)*$`).MatchString(m)
}

// IsMorseValidStageOne Loops through all codes to make sure they match with the MAP
func IsMorseValidStageOne(m string) bool {
	return regexp.MustCompile(`[.\-]+`).MatchString(m) && !regexp.MustCompile(`[a-zA-Z0-9]+`).MatchString(m)
}

// CraftMorseFromString Takes a message and converts all the characters into valid morse code
func CraftMorseFromString(s string) string {
	var ret string
	s = strings.ToUpper(s)
	for i := 0; i < len(s); i++ {
		ret += CharToMorse[s[i:i+1]]
		if i < len(s)-1 {
			ret += " "
		}
	}
	log.Println(fmt.Sprintf("[CONVERT] \"%s\" -> (%s)", s, ret))
	return ret
}

// CraftStringFromMorse Takes a morse code and converts it into human-readable text
func CraftStringFromMorse(m string) string {
	if !IsMorseValidStageTwo(m) {
		return "Invalid morse code"
	}
	var ret string
	// Split the total string into sections
	codes := strings.Split(m, " ")
	for i := 0; i < len(codes); i++ {
		ret += MorseToChar[codes[i]]
	}
	ret = strings.ReplaceAll(ret, "   ", " ")
	log.Println(fmt.Sprintf("[CONVERT] \"%s\" <- (%s)", ret, m))
	return ret
}

// main Provides an entry point for our application that contains the GUI init
func main() {
	log.Println("Starting GUI for CPT")
	// go background()
	cpt := app.New()
	window := cpt.NewWindow("CPT - Brandon Plank")
	window.Resize(fyne.Size{Width: 500, Height: 400})

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	output := widget.NewLabel("")

	window.SetContent(container.NewVBox(
		input,
		widget.NewButton("Translate", func() {
			if IsMorseValidStageOne(input.Text) {
				output.Text = CraftStringFromMorse(input.Text)
			} else {
				if IsStringValid(input.Text) {
					output.Text = CraftMorseFromString(input.Text)
				} else {
					output.Text = "Invalid string"
				}
			}
			output.Refresh()
		}),
		output,
	))
	window.ShowAndRun()
}
