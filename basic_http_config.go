package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !")
}

type Config struct {
	Port    string `json:"port"`
	BaseDir string `json:"base_dir"`
}

func (c Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%#v", c)
}

func main() {
	fileBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}

	var config Config

	if err := json.Unmarshal(fileBytes, &config); err != nil {
		fmt.Println(err)
	}

	http.Handle("/", config)

	fmt.Println("Serving from port :", config.Port)
	e := http.ListenAndServe("localhost:"+config.Port, nil)
	if e != nil {
		fmt.Println("err", e)
	}
}
