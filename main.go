package main

import (
	"encoding/json"
	"sufz/sufz"
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Config struct {
    Threads int `json:"threads"`
    Depth   int `json:"depth"`
    Timeout int `json:"timeout"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParse := json.NewDecoder(configFile)
	err = jsonParse.Decode(&config)
	return config, err
}

func main()  {
	// Config
	config, _ := LoadConfig("config.json")
	fmt.Println(config)

	// FLAGS
	urlencode := flag.Bool("e", false, "Url Encode")
	urldecode := flag.Bool("d", false, "Url Decode")

	flag.Parse()
	scan := bufio.NewScanner(os.Stdin)

	if *urlencode {
		sufz.UrlEncode(*scan)
	}
	if *urldecode {
		sufz.UrlDecode(*scan)
	}
}