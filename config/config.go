package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type message struct {
	Content string
	Time    string
}

//todo allow for the creation of numerous unique ID's and IP's to foster multicast functionality
type Config struct {
	MinD int
	MaxD int
	ID   int
	IP   string
	Port string
}

//extract delay values and populate the structure for future reference
func getDelayVals(firstLine []string) [2]int { //to tidy up ReadFile fxn; for the first line with delay values
	if len(firstLine) != 2 {
		panic(1)
	}

	var Delays [2]int
	var err error
	Delays[0], err = strconv.Atoi(firstLine[0]) //max delay
	if err != nil {
		panic(err)
	}
	Delays[1], err = strconv.Atoi(firstLine[1]) //min delay
	if err != nil {
		panic(err)
	}
	return Delays
}

func ReadFile(FileName string) []Config {

	var configs []Config

	//open and read config file
	ConfigFile, err := os.Open(FileName) //assuming this will be file name
	if err != nil {
		fmt.Println(err)
	}
	defer ConfigFile.Close()

	//read the files contents
	scanner := bufio.NewScanner(ConfigFile) //read the file line by line
	//parse out file line by line
	scanner.Scan()
	DVals := getDelayVals(strings.Split(scanner.Text(), " ")) //delay values, 0th index = min, 1st = max
	//move out of the first line

	for scanner.Scan() { //go line by line through the file
		temp := strings.Split(scanner.Text(), " ")
		id, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		conf := Config{MinD: DVals[0], MaxD: DVals[1], ID: id, IP: temp[1], Port: temp[2]} //populate fields
		configs = append(configs, conf)                                                    // append to slice
	}
	return configs
}
