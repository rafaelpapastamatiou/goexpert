package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var dbConfig DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	MarshalConfig("config.json")

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				fmt.Println("Event:", event)

				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				fmt.Println("FSNotify Watcher Error:", err)
			}
		}
	}()

	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &dbConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("Config loaded:", dbConfig)
}
