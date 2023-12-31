package util

import (
	"github.com/charmbracelet/log"

	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func Update() {
	currentLocation, err := findCurrentLocation()
	if err != nil {
		log.Fatal("Could not find current installation location", "error", err)
	}
	log.Info("Current location is " + currentLocation)

	log.Info("Looking for binary", "os", runtime.GOOS, "arch", runtime.GOARCH)

	url := "https://frank-mayer.github.io/yab/yab-" + runtime.GOOS + "-" + runtime.GOARCH
	if runtime.GOOS == "windows" {
		url += ".exe"
	}

	log.Info("Downloading from " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not download new binary", "error", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Could not download new binary", "error", resp.Status)
	}

	err = install(currentLocation, resp.Body)
	if err != nil {
		log.Fatal("Could not install new binary", "error", err)
	}

	log.Info("Downloaded new binary")
}

func findCurrentLocation() (string, error) {
	binary, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return binary, nil
}
