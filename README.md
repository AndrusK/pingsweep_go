IP Ping Sweep

This Go program performs a sweep of IP addresses in a given range and outputs the results to the console or to a CSV file. The program will ping each IP address in the specified range and, if successful, will attempt to retrieve the hostname associated with that IP address using nslookup. It also calculates the distance of each IP address from the starting IP in the range.
Features

    Ping sweep across a given range of IPv4 addresses.
    NSLookup for resolving hostnames associated with the IP addresses.
    Optionally outputs results to a CSV file.
    Displays the results on the console with IP, Hostname, and Distance.
    Sorts results based on distance from the starting IP.

Requirements

    Go 1.18 or later.
    Dependencies: github.com/spf13/pflag for command-line flag parsing.

Installation

    Clone this repository:

git clone https://github.com/your-repo/ip-ping-sweep.git

Navigate into the project directory:

cd ip-ping-sweep

Install required dependencies:

    go mod tidy

Usage
Command-line Flags

    -s, --start: Starting IP address (Required).
    -e, --end: Ending IP address (Required).
    -o, --output: Output file name for the CSV (Optional).

Example Usage
Ping Sweep without output file

go run main.go --start 192.168.1.1 --end 192.168.1.10

This command will ping all IPs from 192.168.1.1 to 192.168.1.10 and display the results (IP, Hostname) on the console.
Ping Sweep with output to CSV

go run main.go --start 192.168.1.1 --end 192.168.1.10 --output results.csv

This command will save the results of the ping sweep in a CSV file named results.csv. The CSV file will contain two columns: IP and Hostname.
Sample Output (Console)

IP, Hostname
192.168.1.1, router.local
192.168.1.3, device-1.local
192.168.1.5, device-2.local

Sample Output (CSV)

IP, Hostname
192.168.1.1, router.local
192.168.1.3, device-1.local
192.168.1.5, device-2.local

Functions

    ping: Performs a single ICMP ping to the given IP address and returns whether the ping was successful.
    nsLookup: Resolves the hostname of a given IP address using nslookup.
    isIPv4: Checks whether the given IP address is an IPv4 address.
    nextIP: Returns the next IP address in the sequence (used for iterating over the IP range).
    distanceFrom: Calculates the distance of an IP address from the starting IP address (in terms of the integer value of the IP address).
    ipToInt: Converts an IPv4 address to an integer to facilitate distance calculations.
    sortByDistance: Sorts the results based on the distance from the starting IP address.

Notes

    The program assumes the IP range is valid and that the start IP is less than or equal to the end IP.
    The ping functionality is platform-dependent; it uses ping -n on Windows and ping -c on other systems like Linux or macOS.
    NSLookup may fail if the IP address does not have a corresponding hostname.
