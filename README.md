# Internet Connectivity Checker

This is an educational project to periodically check internet connectivity and perform more in-depth tests when connectivity issues are detected. The program is written in Go and is designed to work on both Linux and macOS.

## Features
- Periodically checks internet connectivity by pinging Google's DNS server.
- If connectivity issues are detected, runs a speed test and traceroute.
- Prints the results to the console.

## Requirements
- Go (latest version)
- `speedtest-cli` (for speed tests)
- `traceroute` (for network tracing)

### Installation of Dependencies
Before running the Go program, you need to install `speedtest-cli` and `traceroute`.

```sh
# Install speedtest-cli
sudo apt-get install speedtest-cli   # On Ubuntu/Debian
brew install speedtest-cli           # On macOS

# Install traceroute
sudo apt-get install traceroute      # On Ubuntu/Debian
brew install traceroute              # On macOS
```

## How to Run

1. Clone the repository:
   ```sh
   git clone git@github.com:cerodriguez/go-netchecker.git
   ```

2. Navigate to the project directory:
   ```sh
   cd go-netchecker 
   ```

3. Run the program:
   ```sh
   go run main.go
   ```

## How to Test

1. Make sure you have the necessary dependencies installed.
2. Run the unit tests:
   ```sh
   go test
   ```

## Disclaimer

This project is for educational purposes only. It will connect to your network to check internet connectivity and perform tests.

