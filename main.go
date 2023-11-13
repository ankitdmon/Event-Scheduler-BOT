package main

import (
	"fmt"

	"github.com/ankitdmon/Event-Scheduler-BOT/bot"
	"github.com/ankitdmon/Event-Scheduler-BOT/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
