package main

import (
	"flag"
	"fmt"
	"strings"

	"send2telegram/config"
	"send2telegram/telegram"
)

const DEFAULT_MESSAGE = "nil"

func main() {
	configPathHelp := fmt.Sprintf("Config path (default: ~/%s)", config.DEFAULT_CONFIG_NAME)

	configPath := flag.String("c", "", configPathHelp)
	message := flag.String("m", "", "Message (optional)")
	setupMode := flag.Bool("s", false, "Setup tg bot")

	flag.Parse()

	if *setupMode {
		config.SetupConfig()
		return
	}

	botConfig, err := config.LoadConfig(*configPath)
	if err != nil {
		return
	}

	msg := *message
	if msg == "" {
		if len(flag.Args()) == 0 {
			msg = DEFAULT_MESSAGE
		} else {
			msg = strings.Join(flag.Args(), "\n")
		}
	}
	err = telegram.SendMessage(botConfig, msg)
	if err != nil {
		fmt.Printf("Cannot send message: %e", err)
	}
}
