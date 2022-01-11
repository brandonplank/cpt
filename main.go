package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	CharToMorse[" "] = "/"

	MorseToChar = FlipMap(CharToMorse)
}

func IsMorseValid(m string) bool {
	codes := strings.Split(m, " ")
	for i := 0; i < len(codes); i++ {
		if MorseToChar[codes[i]] == "" {
			return false
		}
	}
	return true
}

func CraftMorseFromString(s string) string {
	var ret string
	s = strings.ToUpper(s)
	for i := 0; i < len(s); i++ {
		ret += CharToMorse[s[i:i+1]] + " "
	}
	return ret
}

func CraftStringFromMorse(m string) string {
	if !IsMorseValid(m) {
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

	modeLabel := widget.NewLabel("Text to Morse")

	output := widget.NewLabel("")

	textToMorse := false

	window.SetContent(container.NewVBox(
		input,
		widget.NewButton("Translate", func() {
			if textToMorse {
				output.Text = CraftMorseFromString(input.Text)
				output.Refresh()
			} else {
				output.Text = CraftStringFromMorse(input.Text)
				output.Refresh()
			}
		}),
		container.NewHBox(
			widget.NewCheck("", func(isOn bool) {
				textToMorse = !isOn
				if !isOn {
					modeLabel.Text = "Text To Morse"
				} else {
					modeLabel.Text = "Morse to Text"
				}
				modeLabel.Refresh()
			}),
			modeLabel,
		),
		output,
	))
	window.ShowAndRun()
}
