package main

type statsOutput struct {
	Uptime string `json:"uptime"`
	Discs  []disc `json:"discs"`
	RAM    struct {
		Mem  map[string]string `json:"mem"`
		Swap map[string]string `json:"swap"`
	} `json:"ram"`
	LoadAverages struct {
		OneMinute      string `json:"one-minute"`
		FiveMinutes    string `json:"five-minutes"`
		FifteenMinutes string `json:"fifteen-minutes"`
	} `json:"loadAverages"`
	AdditionalCommands []AdditionalCommand `json:"additional-commands"`
}

type disc struct {
	Size        string `json:"size"`
	MountPoint  string `json:"mountPoint"`
	Used        string `json:"used"`
	Free        string `json:"free"`
	UsedPercent string `json:"usedPercent"`
}
