// Package updater takes care of teletext page files dynamic update
package updater

import (
	"fmt"
	"os"

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
