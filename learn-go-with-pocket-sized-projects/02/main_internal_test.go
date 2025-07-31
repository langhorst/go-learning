package main

import "testing"

func Example() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"Greek": {
			lang: "el",
			want: "Χαίρετε κόσμε",
		},
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"German": {
			lang: "de",
			want: "Hallo Welt",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום עולם",
		},
		"Urdu": {
			lang: "ur",
			want: "ہیلو دنیا",
		},
		"Spanish": {
			lang: "es",
			want: "Hola mundo",
		},
		"Italian": {
			lang: "it",
			want: "Ciao mondo",
		},
		"Japanese": {
			lang: "ja",
			want: "こんにちは世界",
		},
		"Portuguese": {
			lang: "pt",
			want: "Olá mundo",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
	}

	// range over all the scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
