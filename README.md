# IP Ping Sweep

This Go program performs a sweep of IP addresses in a given range and outputs the results to the console or to a CSV file. The program will ping each IP address in the specified range and, if successful, will attempt to retrieve the hostname associated with that IP address using `nslookup`. It also calculates the distance of each IP address from the starting IP in the range.

## Features

- Ping sweep across a given range of IPv4 addresses.
- NSLookup for resolving hostnames associated with the IP addresses.
- Optionally outputs results to a CSV file.
- Displays the results on the console with IP, Hostname, and Distance.
- Sorts results based on distance from the starting IP.

## Requirements

- Go 1.18 or later.
- Dependencies: `github.com/spf13/pflag` for command-line flag parsing.

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/AndrusK/pingsweep_go.git
    ```

2. Navigate into the project directory:
    ```bash
    cd pingsweep_go
    ```

3. Install required dependencies:
    ```bash
    go mod tidy
    ```

## Usage

### Command-line Flags

- `-s`, `--start`: Starting IP address (Required).
- `-e`, `--end`: Ending IP address (Required).
- `-o`, `--output`: Output file name for the CSV (Optional).

### Example Usage

#### Ping Sweep without output file

```bash
go run main.go --start 192.168.1.1 --end 192.168.1.10
