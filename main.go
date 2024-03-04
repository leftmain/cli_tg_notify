package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"send2telegram/config"
	"send2telegram/telegram"
)

const DEFAULT_MESSAGE = "nil"
const MAX_INPUT_SIZE = 1024 * 1024

func readFromInput() string {
	file, err := os.Open("/dev/stdin")
	if err != nil {
		return DEFAULT_MESSAGE
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return DEFAULT_MESSAGE
	}

	if len(content) > MAX_INPUT_SIZE {
		content = content[:MAX_INPUT_SIZE]
	}

	return string(content)
}

func main() {
	configPathHelp := fmt.Sprintf("Config path (default: ~/%s)", config.DEFAULT_CONFIG_NAME)

	configPath := flag.String("c", "", configPathHelp)
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

	var msg string
	if len(flag.Args()) == 0 {
		msg = readFromInput()
	} else {
		msg = strings.Join(flag.Args(), "\n")
	}
	err = telegram.SendMessage(botConfig, msg)
	if err != nil {
		fmt.Printf("Cannot send message: %e", err)
	}
}
