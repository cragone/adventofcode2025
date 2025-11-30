package info

import (
	"fmt"
	"io"
	"net/http"
)

type Application struct {
	Session string `json:"session"`
}

func (a *Application) GetData(url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("could not create request")
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: a.Session})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not retrieve data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read api resp body")
	}

	return string(body), nil

}
