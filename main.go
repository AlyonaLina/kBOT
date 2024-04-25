/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"
	"time"
    "os"
	tele "gopkg.in/telebot.v3"
)


func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TELE_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 20 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
