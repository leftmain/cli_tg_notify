# send2telegram

A fast and straightforward CLI tool designed for sending messages directly from your terminal to a specific Telegram bot chat. Ideal for developers, system administrators, and automation enthusiasts looking to integrate instant notifications into their workflow. Please note, `send2telegram` is tailored for individual bot chats and does not support sending messages to group chats (yet).

## Table of Contents

- [Installation](#installation)
- [Setup](#setup)
- [Usage](#usage)


## Installation

### If you have go
```bash
git clone https://github.com/leftmain/send2telegram.git
cd send2telegram
go build
```

### If you want just the binary
TODO

## Setup
You need to provide a Telegram bot token and chat_id.
To do this, run command and follow the instructions:
```bash
send2telegram -s
```

## Usage

### Send messages to notify yourself
```bash
sleep 1h ; send2telegram 'Time to go!'
```

### One argument - one line
```bash
send2telegram 'Line 1' 'Line 2'
```

### Without message arguments you can write to stdin
```bash
echo 'blablabla' | send2telegram -c ~/my_lovely_config
```
> **Notice:** Arguments have priority, meaning if arguments are provided, stdin input will be ignored.

### Use different bots
```bash
send2telegram -c ~/tg_configs/work_bot 'Ready for deploy'
send2telegram -c ~/tg_configs/personal_bot 'Order snacks to be happy'
```

Add alices to make it useful:
```bash
# Add these lines to ~/.bashrc, ~/.zshrc, or your preferred shell configuration file
alias work_bot="send2telegram -c ~/tg_configs/work_bot"
alias personal_bot="send2telegram -c ~/tg_configs/personal_bot"
...
```

