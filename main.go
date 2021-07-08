package main

import (
	"notion-integration-bot/cron"
	"notion-integration-bot/discordbot"
)


func main() {
	discordbot.CreateChattingBot()
	cron.CreateCron()
	
	<-make(chan struct{})
	return
}

