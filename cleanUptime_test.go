package main

import "testing"

func TestUptimeCommandCanBeParsedCorrectly(t *testing.T) {
	output := " 12:11:27 up  1:35,  1 user,  load average: 2.02, 2.17, 2.18\n"
	uptime, average1, average5, average15 := cleanUptime(output)
	if uptime != "1:35" || average1 != "2.02" || average5 != "2.17" || average15 != "2.18" {
		t.Error("Uptime command parsed incorrectly")
	}
}
