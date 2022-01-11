package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	CharToMorse[" "] = " "

	MorseToChar = FlipMap(CharToMorse)
	MorseToChar["   "] = " "
}

func CraftMorseFromString(s string) string {
	var ret string
	for i := 0; i < len(s); i++ {
		ret += CharToMorse[s[i:i+1]] + " "
	}
	return ret
}

func CraftStringFromMorse(m string) string {
	var ret string
	// Split the total string into sections
	codes := strings.Split(m, " ")

	for i := 0; i < len(codes); i++ {
		if codes[i] == " " {
			log.Println("e")
			ret += codes[i]
		}
		ret += MorseToChar[codes[i]]
	}
	ret = strings.ReplaceAll(ret, "   ", " ")
	return ret
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")
		text = strings.ToUpper(text)
		morse := CraftMorseFromString(text)
		fmt.Println(morse)
		back := CraftStringFromMorse(morse)
		fmt.Println(back)

	}
}
