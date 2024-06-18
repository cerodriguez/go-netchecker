package main

import (
	"fmt"
	"os/exec"
	"time"
)

// checkConnectivity pings Google's DNS server to check connectivity.
func checkConnectivity() bool {
	_, err := exec.Command("ping", "-c", "3", "8.8.8.8").Output()
	if err != nil {
		fmt.Println("Error pinging:", err)
		return false
	}
	return true
}

// runSpeedTest runs a speed test using speedtest-cli.
func runSpeedTest() {
	out, err := exec.Command("speedtest-cli").Output()
	if err != nil {
		fmt.Println("Error running speed test:", err)
		return
	}
	fmt.Println("Speed Test Result:\n", string(out))
}

// runTraceroute runs a traceroute to a reliable host.
func runTraceroute() {
	out, err := exec.Command("traceroute", "8.8.8.8").Output()
	if err != nil {
		fmt.Println("Error running traceroute:", err)
		return
	}
	fmt.Println("Traceroute Result:\n", string(out))
}

func main() {
	// Set the interval for connectivity checks
	interval := time.Minute * 5

	for {
		if !checkConnectivity() {
			fmt.Println("Internet connectivity issue detected.")
			runSpeedTest()
			runTraceroute()
		} else {
			fmt.Println("Internet connectivity is fine.")
		}

		// Wait for the next interval
		time.Sleep(interval)
	}
}

