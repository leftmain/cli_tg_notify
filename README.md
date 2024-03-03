# send2telegram

A fast and straightforward CLI tool for sending messages directly from your terminal to Telegram. Perfect for developers, system administrators, and automation enthusiasts who need to integrate instant notifications into their workflow.

## Table of Contents

- [Installation](#installation)
- [Setup](#setup)
- [Usage](#usage)


## Installation

### If you know what is GOPATH
```bash
git clone https://github.com/leftmain/send2telegram.git
cd send2telegram
go install
```

### If you want just binary
TODO

## Setup
You need to provide telegram bot token and chat_id.
To do it, run `send2telegram -s` and follow instructions.


## Usage

### Send messages to notify yourself
```bash
sleep 1h ; send2telegram 'Time to go!'
```

### Use different bots
```bash
send2telegram -c ~/tg_configs/work_bot 'Ready for deploy'
send2telegram -c ~/tg_configs/personal_bot 'Order snacks to be happy'
```

Add alices to make it useful:
```bash
# ~/.bashrc or ~/.zshrc or what_you_use config
alias work_bot="send2telegram -c ~/tg_configs/work_bot"
alias personal_bot="send2telegram -c ~/tg_configs/personal_bot"
```

