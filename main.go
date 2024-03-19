package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	temperatureMap := make(map[string]Reading)

	// read the file using scanner
	startTime := time.Now()
	f, err := os.Open("./measurements.txt")
	if err != nil {
		log.Fatalf("failed to read the measurements file, reason=%s", err)
		return
	}
	log.Printf("delay in reading file = %v", time.Since(startTime))

	defer f.Close()

	scanner := bufio.NewScanner(f)
	startTime = time.Now()
	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, ";")
		val, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatalf("failed to convert temperature val to float, reason=%s", err)
			return
		}

		reading := temperatureMap[parts[0]]
		if val < reading.Min {
			reading.Min = val
		} else if val > reading.Max {
			reading.Max = val
		}
		reading.Median = (reading.Median*reading.Count + val) / (reading.Count + 1)
		reading.Count++

		temperatureMap[parts[0]] = reading
	}
	log.Printf("delay in the loop = %v", time.Since(startTime))

	if scanner.Err() != nil {
		log.Fatalf("Failed to read in complete file, reason=%s", err)
		return
	}

	log.Printf("total distinct cities found = %d", len(temperatureMap))
}
