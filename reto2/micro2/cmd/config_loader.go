package main

import (
	"bufio"
	"os"
	"strings"
)

type Configuration struct {
	IP        string
	Port      string
	Directory string
	QName     string
}

func loadConfig(directory string) (*Configuration, error) {
	file, err := os.Open(directory + "/.env")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Configuration{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "IP":
			config.IP = parts[1]
		case "PORT":
			config.Port = parts[1]
		case "DIRECTORY":
			config.Directory = parts[1]
		case "RABBIT_QUEUE":
			config.QName = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
