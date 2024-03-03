package config

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupConfig() {
	var config SendConfig
	var bot *tgbotapi.BotAPI
	var configPath string

	defaultConfigPath := getDefaultConfigPath()
	fmt.Printf("Provide path to config file (%s): ", defaultConfigPath)
	n, err := fmt.Scanf("%s", &configPath)
	if n == 0 || err != nil || configPath == "" {
		configPath = defaultConfigPath
	}

	for config.TelegramApiToken == "" {
		fmt.Printf("Provide telegram bot token (https://t.me/BotFather): ")

		n, err = fmt.Scanf("%s", &config.TelegramApiToken)
		if n == 0 || err != nil {
			fmt.Println("Wrong token format")
			continue
		}

		bot, err = tgbotapi.NewBotAPI(config.TelegramApiToken)
		if err != nil {
			fmt.Printf("Wrong token: %v\n", err)
			config.TelegramApiToken = ""
			continue
		}
	}

	fmt.Printf("Authorized on bot https://t.me/%s\n", bot.Self.UserName)
	fmt.Println("Send any message to the bot to get chat id")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			var answer string

			fmt.Printf("Are you %s? (y/n): ", update.Message.From.UserName)
			fmt.Scanf("%s", &answer)

			if answer == "y" {
				config.ChatId = update.Message.Chat.ID
				fmt.Printf("Chat id: %d\n", config.ChatId)
				break
			} else {
				fmt.Println("Send any message to the bot to get chat id")
			}
		}
	}

	err = SaveConfig(config, configPath)
	if err != nil {
		fmt.Printf("Cannot save config: %v\n", err)
		return
	}

	fmt.Printf("Config saved to %s\n", configPath)
}
