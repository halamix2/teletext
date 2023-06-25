// Package main of the updater tools, which regenerates tti files with current data
package main

import (
	"flag"
	"time"

	"github.com/halamix2/teletext/pkg/gsps"
	"github.com/halamix2/teletext/pkg/updater"
)

var (
	updateDelay  time.Duration
	baseURL      string
	currentEvent string
	inputDir     string
	outputDir    string
	u            updater.Updater
)

func setFlags() {
	flag.DurationVar(&updateDelay, "delay", 30*time.Second, "time when to update data")
	flag.StringVar(&baseURL, "baseURL", "https://gsps.pl", "base GSPS URL")
	flag.StringVar(&currentEvent, "currentEvent", "gspsdzieciom2023", "base GSPS URL")
	flag.StringVar(&inputDir, "input", "gspsdzieciom2023", "input dir")
	flag.StringVar(&outputDir, "output", "gspsdzieciom2023", "output dir")

	flag.Parse()
}

func main() {
	setFlags()

	ticker := time.NewTicker(updateDelay)

	gspsService := gsps.Service{BaseURL: baseURL}
	u = updater.Updater{
		BaseURL:      baseURL,
		CurrentEvent: currentEvent,
		InputDir:     inputDir,
		OutputDir:    outputDir,
		GSPSService:  gspsService,
	}

	u.UpdateGSPSData()
	u.Debug()
	for range ticker.C {
		u.UpdateGSPSData()
		u.Debug()
		u.UpdatePages()
	}
}
