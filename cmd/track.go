package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track the IP with this command.",
	Long:  `Track the IP with this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ipAddress := range args {
				showData(ipAddress)
			}
		} else {
			fmt.Println("Please provide an IP Address to track.")
		}
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Loc      string `json:"loc"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func showData(ipAddress string) {
	url := "https://ipinfo.io/" + ipAddress + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response.")
	}

	fmt.Println("\nDATA FOUND :")

	fmt.Printf("IP: %s\nLAT & LON: %s\nCITY: %s\nREGION: %s\nCOUNTRY: %s\nISP: %s\nPOSTAL: %s\nTIMEZONE: %s\n", data.IP, data.Loc, data.City, data.Region, data.Country, data.Org, data.Postal, data.Timezone)

	fmt.Println("\n")
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response.")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
