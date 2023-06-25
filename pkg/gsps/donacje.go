// Package gsps interaces with GSPS site and APIs
package gsps

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Service contains basic service configuration
type Service struct {
	BaseURL string
}

// Count contains count of runs and bids
type Count struct {
	Runs   int `json:"runs"`
	Prizes int `json:"prizes"`
	Bids   int `json:"bids"`
	Donors int `json:"donors"`
}

// Agg contains aggregated donation onfi
type Agg struct {
	Amount  float64 `json:"amount,string"`
	Count   int     `json:"count"`
	Max     float64 `json:"max,string"`
	Average float64 `json:"avg"`
	Target  float64 `json:"target,string"`
}

// Donate contains GSPS tracker response
type Donate struct {
	Count Count `json:"count"`
	Agg   Agg   `json:"agg"`
}

// GetDonations retrieves info about an event donations
func (s Service) GetDonations(eventName string) (Donate, error) {
	combinedURL, err := url.Parse(s.BaseURL) //url.JoinPath(baseURL,"donacje")
	if err != nil {
		return Donate{}, err
	}
	combinedURL = combinedURL.JoinPath("donacje")
	combinedURL = combinedURL.JoinPath("index")
	if eventName != "" {
		combinedURL = combinedURL.JoinPath(eventName)
	}

	q := url.Values{}
	q.Add("json", "stunt_gp")

	combinedURL.RawQuery = q.Encode()

	//fmt.Println(combinedURL.String())
	resp, err := http.Get(combinedURL.String())
	if err != nil {
		return Donate{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Donate{}, err
	}
	err = resp.Body.Close()
	if err != nil {
		return Donate{}, err
	}

	var donate Donate

	err = json.Unmarshal(body, &donate)

	return donate, err
}

// {
//     "count": {
//         "runs": 573,
//         "prizes": 46,
//         "bids": 182,
//         "donors": 1083
//     },
//     "agg": {
//         "amount": "255057.67",
//         "count": 4702,
//         "max": "10000.00",
//         "avg": 54.244506592939175,
//         "target": null
//     }
// }
