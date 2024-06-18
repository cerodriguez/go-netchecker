package main

import (
	"testing"
)

// TestCheckConnectivity tests the checkConnectivity function.
func TestCheckConnectivity(t *testing.T) {
	if !checkConnectivity() {
		t.Error("Expected connectivity to be fine, but it was not")
	}
}

// Since runSpeedTest and runTraceroute functions just execute external commands,
// they are more like integration tests. However, we'll still create simple tests to ensure they run without errors.

func TestRunSpeedTest(t *testing.T) {
	// Assuming speedtest-cli is installed
	runSpeedTest()
	// Manual inspection of output is needed as we are not capturing output in the function
}

func TestRunTraceroute(t *testing.T) {
	// Assuming traceroute is installed
	runTraceroute()
	// Manual inspection of output is needed as we are not capturing output in the function
}
