package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε κόσμε",    // Greek
	"en": "Hello world",      // English
	"fr": "Bonjour le monde", // French
	"de": "Hallo Welt",       // German
	"he": "שלום עולם",        // Hebrew
	"ur": "ہیلو دنیا",        // Urdu
	"es": "Hola mundo",       // Spanish
	"it": "Ciao mondo",       // Italian
	"ja": "こんにちは世界",          // Japanese
	"pt": "Olá mundo",        // Portuguese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
