package main

import "os"

type Environment struct {
	GoogleFontsApiKey string
}

var env Environment

func init() {
	// Load the Google Fonts API key from environment variable
	googleFontsApiKey, exists := os.LookupEnv("GOOGLE_FONTS_API_KEY")
	if !exists || googleFontsApiKey == "" {
		panic("GOOGLE_FONTS_API_KEY environment variable is not set")
	}

	env = Environment{
		GoogleFontsApiKey: googleFontsApiKey,
	}
}
