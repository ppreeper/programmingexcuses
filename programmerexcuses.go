package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"gopkg.in/antchfx/htmlquery.v1"
)

// Constants for URL and User-Agent header
const (
	baseURL   = "http://programmingexcuses.com/"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36"
)

// getExcuse fetches the excuse from the website
func getExcuse() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch excuse, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	body, err := getExcuse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	doc, err := htmlquery.Parse(strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	list := htmlquery.FindOne(doc, "//div[@class='wrapper']")
	a := htmlquery.FindOne(list, "//a")
	fmt.Println(htmlquery.InnerText(a))
}
