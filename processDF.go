package main

import (
	"strings"
)

func processDF(raw string) []disc {
	var discs []disc
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		// continue on an empty line
		if len(words) == 0 {
			continue
		}
		// continue on a line that starts with a non "/dev" entry
		if len(words[0]) < 4 || words[0][0:4] != "/dev" {
			continue
		}

		// skip any snap mount points
		if len(words[5]) > 14 && words[5][0:14] == "/var/lib/snapd" {
			continue
		}

		disc := disc{
			Size:        words[1],
			Used:        words[2],
			Free:        words[3],
			UsedPercent: words[4],
			MountPoint:  words[5],
		}
		discs = append(discs, disc)
	}
	return discs
}
