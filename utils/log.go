package utils

import "log"

func Log(text string) {
	log.Printf("%v: %v", text, text)
}
