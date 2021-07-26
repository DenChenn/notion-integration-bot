package cron

import (
	"notion-integration-bot/config"
	"notion-integration-bot/discordbot"
	"notion-integration-bot/notionbot"

	"github.com/robfig/cron"
)

func CreateCron() {
	cronManager := cron.New()
	cronManager.AddFunc("*/10 * * * * *", func(){
		developementUrl := "https://api.notion.com/v1/databases/" + config.DepartmentDatabaseId + "/query"
		designUrl := "https://api.notion.com/v1/databases/" + config.DesignDatabaseChannelId + "/query"
		financeUrl := "https://api.notion.com/v1/databases/" + config.FinanceDatabaseId + "/query"
		marketingUrl := "https://api.notion.com/v1/databases/" + config.MarketingDatabaseId + "/query"

		developeIsChange, developeDetailSet := notionbot.CheckDepartment(developementUrl)
		designIsChange, designDetailSet := notionbot.CheckDepartment(designUrl)
		financeIsChange, financeDetailSet := notionbot.CheckDepartment(financeUrl)
		marketingIsChange, marketingDetailSet := notionbot.CheckDepartment(marketingUrl)

		if(developeIsChange){
			discordbot.DistributeMessage(&developeDetailSet)
		}
		if(designIsChange){
			discordbot.DistributeMessage(&designDetailSet)
		}
		if(financeIsChange){
			discordbot.DistributeMessage(&financeDetailSet)
		}
		if(marketingIsChange){
			discordbot.DistributeMessage(&marketingDetailSet)
		}
	})
	cronManager.Start()
}