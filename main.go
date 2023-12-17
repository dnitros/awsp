package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"regexp"
)

func getConfigFileLocation() string {
	configLocation := os.Getenv("HOME") + "/.aws/config"
	if value, flag := os.LookupEnv("AWS_CONFIG_FILE"); flag {
		configLocation = value
	}
	return configLocation
}

func main() {
	// check the config file using AWS_CONFIG_FILE or ~/.aws/config
	configLocation := getConfigFileLocation()

	file, err := os.Open(configLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	regex, err := regexp.Compile("(\\[profile )|(\\[)|(])")
	var matches []string

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if regex.MatchString(scanner.Text()) {
			match := regex.ReplaceAllString(scanner.Text(), "")
			matches = append(matches, match)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// prompt for profile choice
	prompt := promptui.Select{
		Label: "Select profile",
		Items: matches,
	}

	_, selectedProfile, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", selectedProfile)

	// write the choice to .awsp file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	location := fmt.Sprintf("%s/.awsp", homeDir)
	awspFile, err := os.OpenFile(location, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = awspFile.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	awspWriter := bufio.NewWriter(awspFile)
	_, err = fmt.Fprint(awspWriter, selectedProfile)
	if err != nil {
		log.Fatal(err)
	}
	err = awspWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
