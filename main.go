package main

import (
	"fmt"
	gui "github.com/AllenDang/giu" // Golang maps for imgui
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

var morseCodeHelp []string

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
	CharToMorse["!"] = "-.-.--"
	CharToMorse["("] = "-.--."
	CharToMorse[")"] = "-.--.-"
	CharToMorse[":"] = "---..."

	// Space, this is normalized in typed morse code
	CharToMorse[" "] = "/"

	// For mose to english conversion
	MorseToChar = FlipMap(CharToMorse)

	// Add all morse codes to a list for autocomplete
	for _, object := range CharToMorse {
		morseCodeHelp = append(morseCodeHelp, object)
	}
}

// IsMorseValidStageTwo Loops through all codes to make sure they match with the MAP
func IsMorseValidStageTwo(s string) bool {
	codes := strings.Split(s, " ")
	for i := 0; i < len(codes); i++ {
		if MorseToChar[codes[i]] == "" {
			return false
		}
	}
	return true
}

// For all regex, I used https://regex101.com/ to make build and test custom regex

// IsStringValid Makes sure that our translator can parse this string
func IsStringValid(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z\d.,?!(): ]+( [a-zA-Z\d.,?!(): ]+)*$`).MatchString(s)
}

// IsMorseValidStageOne Loops through all codes to make sure they match with the MAP
func IsMorseValidStageOne(s string) bool {
	return regexp.MustCompile(`[.\-]+`).MatchString(s) && !regexp.MustCompile(`[a-zA-Z\d]+`).MatchString(s)
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
func CraftStringFromMorse(s string) string {
	if !IsMorseValidStageTwo(s) {
		return "Invalid morse code"
	}
	var ret string
	// Split the total string into sections
	codes := strings.Split(s, " ")
	for i := 0; i < len(codes); i++ {
		ret += MorseToChar[codes[i]]
	}
	ret = strings.ReplaceAll(ret, "   ", " ")
	log.Println(fmt.Sprintf("[CONVERT] \"%s\" <- (%s)", ret, s))
	return ret
}

var text string
var output string
var autoTranslate bool

// DetectAndTranslate gets text from the GUI and checks if its morse or english, then converts accordingly
func DetectAndTranslate() {
	if IsMorseValidStageOne(text) {
		output = CraftStringFromMorse(text)
	} else {
		if IsStringValid(text) {
			output = CraftMorseFromString(text)
		} else {
			output = "Invalid string"
		}
	}
	gui.Update()
}

// GuiLoop provides the main GUI loop for our app.
func GuiLoop() {
	gui.SingleWindow().Layout(
		gui.Align(gui.AlignCenter).To(
			gui.Label("Morse Code Translator"),
			gui.InputText(&text).AutoComplete(morseCodeHelp).OnChange(func() {
				if autoTranslate {
					DetectAndTranslate()
				}
			}),
			gui.Row(
				gui.Button("Translate").OnClick(DetectAndTranslate),
				gui.Button("Clear").OnClick(func() {
					text = ""
					output = ""
					gui.Update()
				}),
			),
			gui.Checkbox("Auto-Translate", &autoTranslate),
			gui.Label(output).Wrapped(true),
		),
	)
}

// main Provides an entry point for our application that contains the GUI init
func main() {
	log.Println("Starting GUI")
	win := gui.NewMasterWindow("Morse Code Translator", 500, 300, gui.MasterWindowFlagsNotResizable)
	win.Run(GuiLoop)
}
