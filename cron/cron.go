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

		developeIsChange, developeDetail := notionbot.CheckDepartment(developementUrl)
		designIsChange, designDetail := notionbot.CheckDepartment(designUrl)
		financeIsChange, financeDetail := notionbot.CheckDepartment(financeUrl)
		marketingIsChange, marketingDetail := notionbot.CheckDepartment(marketingUrl)
		if(developeIsChange){
			for _, v := range developeDetail{
				m := discordbot.DepartmentMessageTransfer(v, 4388240)
				discordbot.SendMessageEmbed(config.DevelopmentChannelId, m)
			}
		}
		if(designIsChange){
			for _, v := range designDetail{
				m := discordbot.DepartmentMessageTransfer(v, 15093467)
				discordbot.SendMessageEmbed(config.DesignChannelId, m)
			}
		}
		if(financeIsChange){
			for _, v := range financeDetail{
				m := discordbot.DepartmentMessageTransfer(v, 15424552)
				discordbot.SendMessageEmbed(config.FinanceChannelId, m)
			}
		}
		if(marketingIsChange){
			for _, v := range marketingDetail{
				m := discordbot.DepartmentMessageTransfer(v, 3466199)
				discordbot.SendMessageEmbed(config.MarketingChannelId, m)
			}
		}
	})
	cronManager.Start()
}