package main

import "strings"

func cleanUptime(raw string) (string, string, string, string) {
	parts := strings.Split(raw, ",")
	uptime := strings.Split(parts[0], "up")
	loadAverage1 := strings.Split(parts[len(parts)-3], ":")
	loadAverage5 := strings.TrimSpace(parts[len(parts)-2])
	loadAverage15 := strings.TrimSpace(parts[len(parts)-1])

	// return uptime, and then system load averages
	return strings.TrimSpace(uptime[1]), strings.TrimSpace(loadAverage1[1]), loadAverage5, loadAverage15
}
