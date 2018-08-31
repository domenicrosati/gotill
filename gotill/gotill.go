package gotill

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	tillSendURL    = "https://platform.tillmobile.com/api/send?username=%s&api_key=%s"
	tillResultsURL = "https://platform.tillmobile.com/api/results"
	tillResultURL  = "https://platform.tillmobile.com/api/results/%s/"
	tillStatsURL   = "https://platform.tillmobile.com/api/stats/"
	tillStatURL    = "https://platform.tillmobile.com/api/stats/%s/"
)

var (
	username = os.Getenv("USERNAME")
	apikey   = os.Getenv("API_KEY")
)

type SendResp struct {
	ResultsURL        string `json:"results_url,omitempty"`
	StatsURL          string `json:"stats_url,omitempty"`
	ProjectLaunchGUID string `json:"project_launch_guid,omitempty"`
}

func SendMessage(number string, text string) (string, error) {
	jsonStr := []byte(fmt.Sprintf(`{"phone":["%s"], "text": "%s"}`, number, text))
	resp, err := http.Post(fmt.Sprintf(tillSendURL, username, apikey), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return "", nil
}

func MakeQuestion() {}

func SendQuestion() {}

func GetResults() {}
func GetResult()  {}
func GetStats()   {}
func GetStat()    {}
