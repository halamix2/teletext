// Package main of the updater tools, which regenerates tti files with current data
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/halamix2/teletext/pkg/gsps"
	"github.com/halamix2/teletext/pkg/updater"
)

var (
	updateGSPSDelay time.Duration
	updateDelay     time.Duration
	baseURL         string
	currentEvent    string
	inputDir        string
	outputDir       string
	u               updater.Updater
)

func setFlags() {
	flag.DurationVar(&updateGSPSDelay, "delayGSPS", 30*time.Second, "time when to update GSPS data")
	flag.DurationVar(&updateDelay, "delay", 30*time.Second, "time when to update pages")
	flag.StringVar(&baseURL, "baseURL", "https://gsps.pl", "base GSPS URL")
	flag.StringVar(&currentEvent, "currentEvent", "gsps2023", "base GSPS URL")
	flag.StringVar(&inputDir, "input", "prezentacja_input", "input dir")
	flag.StringVar(&outputDir, "output", "prezentacja_wst", "output dir")

	flag.Parse()
}

func main() {
	setFlags()

	tickerGSPS := time.NewTicker(updateGSPSDelay)
	tickerPages := time.NewTicker(updateDelay)

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
	for {
		select {
		case <-tickerGSPS.C:
			fmt.Println("Updating GSPS")
			u.UpdateGSPSData()
			u.Debug()
			u.UpdatePages()
		case <-tickerPages.C:
			// u.UpdateGSPSData()
			// u.Debug()
			u.UpdatePages()
		}

	}

}
