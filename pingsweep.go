package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type IPResult struct {
	IP       string
	Hostname string
	Distance int // Distance from starting IP
}

func main() {
	var startIP, endIP, outputFile string

	// Define command-line flags
	flag.StringVar(&startIP, "start", "", "Starting IP address")
	flag.StringVar(&endIP, "end", "", "Ending IP address")
	flag.StringVar(&outputFile, "output", "", "Output CSV file name")
	flag.Parse()

	if startIP == "" || endIP == "" || outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Validate and parse IP addresses
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)
	if start == nil || end == nil {
		fmt.Println("Invalid start or end IP address")
		os.Exit(1)
	}

	if !isIPv4(start) || !isIPv4(end) {
		fmt.Println("Both start and end IP addresses must be IPv4")
		os.Exit(1)
	}

	// Channel to communicate results back
	results := make(chan IPResult)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	fmt.Println("Starting ping sweep...")
	startTime := time.Now() // Record the start time

	// Slice to collect results
	var allResults []IPResult

	// Perform ping sweeping from start to end IP addresses
	for ip := start; !ip.Equal(end); ip = nextIP(ip) {
		wg.Add(1)
		go func(ipStr string) {
			defer wg.Done()
			if ping(ipStr) {
				// If ping successful, perform NSLookup
				nslookupResult := nsLookup(ipStr)
				result := IPResult{IP: ipStr, Hostname: nslookupResult, Distance: distanceFrom(start, net.ParseIP(ipStr))}
				results <- result
			}
		}(ip.String())
	}
	wg.Add(1)
	go func(ipStr string) {
		defer wg.Done()
		if ping(endIP) {
			// If ping successful, perform NSLookup
			nslookupResult := nsLookup(endIP)
			result := IPResult{IP: endIP, Hostname: nslookupResult, Distance: distanceFrom(start, net.ParseIP(endIP))}
			results <- result
		}
	}(endIP)

	// Goroutine to collect results
	go func() {
		for result := range results {
			allResults = append(allResults, result)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the results channel
	close(results)

	// Sort results by distance
	sortByDistance(allResults)

	// Open output CSV file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"IP", "Hostname"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV:", err)
		os.Exit(1)
	}

	// Write sorted results to CSV
	for _, result := range allResults {
		record := []string{result.IP, result.Hostname}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing result to CSV:", err)
			// Optionally, handle the error or log it
		}
		fmt.Printf("Wrote to CSV: %s, %s\n", result.IP, result.Hostname)
	}

	// Calculate and print the elapsed time
	elapsed := time.Since(startTime)
	fmt.Println("Ping sweep completed in", elapsed)
}

func ping(ipStr string) bool {
	// Function to perform ICMP ping
	cmd := exec.Command("ping", "-c", "1", ipStr)
	err := cmd.Run()
	return err == nil
}

func nsLookup(ipStr string) string {
	// Function to perform NSLookup
	ips, err := net.LookupAddr(ipStr)
	if err != nil {
		return "NSLookup failed"
	}
	return strings.Join(ips, ", ")
}

func isIPv4(ip net.IP) bool {
	return ip.To4() != nil
}

func nextIP(ip net.IP) net.IP {
	next := make(net.IP, len(ip))
	copy(next, ip)
	for i := len(next) - 1; i >= 0; i-- {
		next[i]++
		if next[i] > 0 {
			break
		}
	}
	return next
}

func distanceFrom(start, ip net.IP) int {
	// Calculate distance from start IP
	startInt := ipToInt(start)
	ipInt := ipToInt(ip)
	return ipInt - startInt
}

func ipToInt(ip net.IP) int {
	// Convert IP address to integer for distance calculation
	ip = ip.To4()
	if ip == nil {
		return 0
	}
	return int(ip[0])<<24 | int(ip[1])<<16 | int(ip[2])<<8 | int(ip[3])
}

func sortByDistance(results []IPResult) {
	// Sort IPResults by Distance
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].Distance > results[j].Distance {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}


