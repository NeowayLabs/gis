// Program/script used to generated the csv test file.
// It queries the server of geohash.org for valid
// geohashes and print longitude, latitude and hash
// to stdout.
// Usage: go run gen.go > geohash.csv
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	timeout = 60 * time.Second
	urlfmt  = "http://geohash.org/?q=%s,%s&format=url&redirect=0"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
	os.Exit(1)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		randomHash()
	}
}

func randomHash() {
	lat, lon, url := geturl()

	fmt.Fprintf(os.Stderr, "Using url: %s\n", url)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		fatal(err)
	}
	defer cancel()

	if resp.StatusCode != 200 {
		fatal(fmt.Errorf("Status %d", resp.StatusCode))
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fatal(err)
	}
	parts := strings.Split(string(content), "/")
	if len(parts) != 4 {
		fatal(fmt.Errorf("Invalid url: %s", string(content)))
	}
	geohash := parts[3]
	fmt.Printf("%s,%s,%s\n", lon, lat, geohash)
}

func geturl() (string, string, string) {
	latrat, lonrat := rand.Float64(), rand.Float64()
	latint := 89 - rand.Intn(180)
	lonint := 179 - rand.Intn(360)
	lat := float64(latint) + latrat
	lon := float64(lonint) + lonrat
	latstr := strconv.FormatFloat(lat, 'f', 6, 64)
	lonstr := strconv.FormatFloat(lon, 'f', 6, 64)
	return latstr, lonstr, fmt.Sprintf(urlfmt, latstr, lonstr)
}
