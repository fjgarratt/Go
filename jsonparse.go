package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Ip   string
	Port string
}

var g_config Config

const input_filename = "src\\ip.json"
const outpt_filename = "src\\reip.json"

func main() {
	bytes, _ := ioutil.ReadFile(input_filename)
	err := json.Unmarshal(bytes, &g_config)

	if err != nil {
		fmt.Println(err)
	}

	rejson, err := json.Marshal(g_config)
	err = ioutil.WriteFile(outpt_filename, rejson, 0644)

	fmt.Println(g_config)
}
