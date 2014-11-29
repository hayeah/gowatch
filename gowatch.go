package main

import (
	"encoding/json"
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
)

type OutputEvent struct {
	Event string `json:"event"`
	Path  string `json:"path"`
}

func main() {
	var err error
	dirs := os.Args[1:]
	if len(dirs) == 0 {
		log.Fatalln("no directory to watch")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}

	// watch all given directories
	for _, dir := range dirs {
		log.Printf("watching: %s\n", dir)
		err = watcher.Add(dir)
		if err != nil {
			log.Fatalln(err)
		}
	}

	output := os.Stdout

	// stream fsnotify events
	var outputEvent OutputEvent
	for {
		event := <-watcher.Events

		outputEvent.Event = opName(event.Op)
		outputEvent.Path = event.Name

		encoder := json.NewEncoder(output)
		err = encoder.Encode(outputEvent)
		if err != nil {
			log.Println(err)
		}

		// fmt.Printf("%v %s\n", tag, event.Name)
	}
}

func opName(op fsnotify.Op) string {
	switch op {
	case fsnotify.Create:
		return "create"
	case fsnotify.Write:
		return "write"
	case fsnotify.Remove:
		return "remove"
	case fsnotify.Rename:
		return "rename"
	case fsnotify.Chmod:
		return "chmod"
	default:
		return "wtf"
	}
}
