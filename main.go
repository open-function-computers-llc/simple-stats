package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/dchest/uniuri"
)

func main() {
	port := os.Getenv("OFCOSTATSPORT")
	if port == "" {
		port = "48832"
	}

	token := os.Getenv("OFCOSTATSTOKEN")
	if token == "" {
		token = uniuri.NewLen(32)
		fmt.Println("token not found in system ENV, randomly set to: " + token)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/json")

		// check for matching token
		if r.FormValue("token") != token {
			sendError(w, "token mismatch")
			return
		}

		stats := statsOutput{}
		var result bytes.Buffer

		// get server uptime
		upTimeCommand := exec.Command("uptime")
		upTimeCommand.Stdout = &result
		upTimeCommand.Run()
		stats.Uptime, stats.LoadAverages.OneMinute, stats.LoadAverages.FiveMinutes, stats.LoadAverages.FifteenMinutes = cleanUptime(result.String())

		// get available server RAM
		result.Reset()
		freeRAMCommand := exec.Command("free")
		freeRAMCommand.Stdout = &result
		freeRAMCommand.Run()
		stats.RAM.Mem, stats.RAM.Swap = cleanFree(result.String())

		// get available disc space
		result.Reset()
		dfCommand := exec.Command("df")
		dfCommand.Stdout = &result
		dfCommand.Run()
		stats.Discs = processDF(result.String())

		output, _ := json.Marshal(stats)
		w.Write(output)
	})
	fmt.Println("system listening on port: " + port)
	http.ListenAndServe(":"+port, nil)
}
