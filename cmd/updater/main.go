// Package main of the updater tools, which regenerates tti files with current data
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/halamix2/teletext/pkg/gsps"
)

// Updater takes care of teletext page files dynamic update
type Updater struct {
	BaseURL      string
	CurrentEvent string
	InputDir     string
	OutputDir    string
	Total        gsps.Donate
	Current      gsps.Donate
	GSPSService  gsps.Service
}

var (
	updateDelay  time.Duration
	baseURL      string
	currentEvent string
	inputDir     string
	outputDir    string
	u            Updater
)

func main() {
	flag.DurationVar(&updateDelay, "delay", 5*time.Second, "time when to update data")
	flag.StringVar(&baseURL, "baseURL", "https://gsps.pl", "base GSPS URL")
	flag.StringVar(&currentEvent, "currentEvent", "gspsdzieciom2023", "base GSPS URL")
	flag.StringVar(&inputDir, "input", "gspsdzieciom2023", "input dir")
	flag.StringVar(&outputDir, "output", "gspsdzieciom2023", "output dir")

	flag.Parse()
	ticker := time.NewTicker(updateDelay)
	//quit := make(chan struct{})

	gspsService := gsps.Service{BaseURL: baseURL}
	//setup
	u = Updater{
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

// UpdateGSPSData updates GSPS tracker data
func (u *Updater) UpdateGSPSData() {
	//get money
	//fmt.Println("kekW")
	total, err := u.GSPSService.GetDonations("")
	if err != nil {
		fmt.Printf("Fatal get total donations: %v\n", err)
		os.Exit(1)
	}
	u.Total = total

	current, err := u.GSPSService.GetDonations(u.CurrentEvent)
	if err != nil {
		fmt.Printf("Fatal get 2023 donations: %v\n", err)
		os.Exit(1)
	}
	u.Current = current
}

// Debug prints some debug data
func (u *Updater) Debug() {
	fmt.Printf("total: %.2f\n", u.Total.Agg.Amount)
	fmt.Printf("total targ: %.2f\n", u.Total.Agg.Target)
	fmt.Printf("%s: %.2f\n", u.CurrentEvent, u.Current.Agg.Amount)
	fmt.Printf("%s targ: %.2f\n", u.CurrentEvent, u.Current.Agg.Target)
}

// UpdatePages updates tti files based on current data
func (u *Updater) UpdatePages() {

}
