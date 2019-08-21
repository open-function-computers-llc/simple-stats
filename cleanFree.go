package main

import (
	"strings"
)

func cleanFree(raw string) (map[string]string, map[string]string) {
	mem := make(map[string]string, 7)
	swap := make(map[string]string, 3)
	var keys []string

	words := strings.Fields(raw)
	currentlyProcessing := "keys"
	keyIndex := 0
	for _, word := range words {
		if word == "Mem:" {
			currentlyProcessing = "mem"
			keyIndex = 0
			continue
		}
		if word == "Swap:" {
			currentlyProcessing = "swap"
			keyIndex = 0
			continue
		}

		if currentlyProcessing == "keys" {
			keys = append(keys, word)
		}
		if currentlyProcessing == "mem" {
			if keyIndex >= len(keys) {
				continue
			}
			mem[keys[keyIndex]] = word
			keyIndex++
		}
		if currentlyProcessing == "swap" {
			swap[keys[keyIndex]] = word
			keyIndex++
		}
	}
	return mem, swap
}
