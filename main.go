package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"regexp"
	"strings"
)

var (
	CharToMorse = make(map[string]string)
	MorseToChar = make(map[string]string)
)

func FlipMap(m map[string]string) map[string]string {
	ret := make(map[string]string, len(m))
	for index, object := range m {
		ret[object] = index
	}
	return ret
}

func init() {
	/*
		Eng to morse chars
	*/
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
	CharToMorse[" "] = "/" // Space

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
		ret += CharToMorse[s[i:i+1]] + " "
	}
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
	return ret
}

func main() {
	app := app.New()
	window := app.NewWindow("CPT - Brandon Plank")
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
