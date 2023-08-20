// Package updater takes care of teletext page files dynamic update
package updater

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
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
	Time         time.Time
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
	fmt.Println("looping pages")
	walkFunc := u.getWalkFunc()
	err := filepath.WalkDir(u.InputDir, walkFunc)
	if err != nil {
		fmt.Printf("Could not walk %s: %v", u.InputDir, err)
	}
}

func (u *Updater) getWalkFunc() fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// skip dir entries
			return nil
		}
		if filepath.Ext(path) != ".tti" {
			return nil
		}

		fmt.Printf("Found %s\n", path)

		err = u.parsePage(path)
		return err
	}
}

func (u *Updater) parsePage(path string) error {
	// TODO Do I need sprig here?
	templateData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// additionalFuncs := template.FuncMap{ "string": func(){} }
	t, err := template.New(path).Funcs(sprig.TxtFuncMap()).Parse(string(templateData))
	if err != nil {
		return err
	}
	justFilename := filepath.Base(path)
	outputPath := filepath.Join(u.OutputDir, justFilename)
	w, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	// TODO provide data
	u.Time = time.Now()
	err = t.Execute(w, u)
	if err != nil {
		return err
	}
	err = w.Close()
	return err
}
