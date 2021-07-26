package discordbot

import (
	"fmt"
	"log"
	"notion-integration-bot/config"
	"notion-integration-bot/model"

	"github.com/bwmarrin/discordgo"
)

var BotId string

func CreateBot() *discordgo.Session {
	discordClient, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	if err = discordClient.Open();err != nil {
		log.Fatal(err)
	}

	return discordClient
}

func DepartmentMessageTransfer(detail model.DepartmentDetail, color int) (message discordgo.MessageEmbed){
	description := ""

	if(detail.Action == "Create"){
		description = "Notion has something created !"
	} else {
		description = "Notion has something updated !"
	}

	var fd []*discordgo.MessageEmbedField
	for i := 0;i < len(detail.FieldSet);i++ {
		var e discordgo.MessageEmbedField
		e.Name = detail.FieldSet[i].Key
		e.Value = detail.FieldSet[i].Value
		fd = append(fd, &e)
	}

	message = discordgo.MessageEmbed{
		Title:       detail.Title,
		Description: description,
		Color:       color,
		Fields:      fd,
	}
	return
}

func SendMessageEmbedToUser(userID string, message discordgo.MessageEmbed){
	discordClient := CreateBot()

	st, err := discordClient.UserChannelCreate(userID)
	if err != nil {
		fmt.Print(err)
		return
	}

	_, err = discordClient.ChannelMessageSendEmbed(st.ID, &message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func DistributeMessage(detailSet *[]model.DepartmentDetail) {
	for _, detail := range *detailSet{
		m := DepartmentMessageTransfer(detail, 4388240)

		if(detail.AssigneeEmail == config.YENTINGCHEN_EMAIL){
			SendMessageEmbedToUser(config.YENTINGCHEN, m)
		} else if (detail.AssigneeEmail == config.TADHSUEH_EMAIL) {
			SendMessageEmbedToUser(config.TADHSUEH, m)
		} else if (detail.AssigneeEmail == config.YUANLIN_EMAIL){
			SendMessageEmbedToUser(config.YUANLIN, m)
		} else if (detail.AssigneeEmail == config.YUTUNG_EMAIL){
			SendMessageEmbedToUser(config.YUTUNG, m)
		} else if (detail.AssigneeEmail == config.FANGFANG_EMAIL){
			SendMessageEmbedToUser(config.FANGFANG, m)
		} else if (detail.AssigneeEmail == config.WINNIEK_EMAIL){
			SendMessageEmbedToUser(config.WINNIEK, m)
		} else if (detail.AssigneeEmail == config.MARYCHOO_EMAIL){
			SendMessageEmbedToUser(config.MARYCHOO, m)
		}
	}
}

func CreateChattingBot() {
	discordClient, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	user, err := discordClient.User("@me")
	if err != nil {
		log.Fatal(err)
	}

	BotId = user.ID
	discordClient.AddHandler(ChatHandler)

	if err = discordClient.Open();err != nil {
		log.Fatal(err)
	}

	return
}

func ChatHandler(client *discordgo.Session, discordMessage *discordgo.MessageCreate) {
	if discordMessage.Author.ID == BotId {
		return
	}

	if discordMessage.Content == "Yen's Notion Integration Bot! 請自我介紹!" {
		var fd []*discordgo.MessageEmbedField
		var e = &discordgo.MessageEmbedField{}
		e.Name = "#關注創建資訊"
		e.Value = ":rocket: 如果notion上面創建了新的page,我會幫大家發個通知喔~"
		fd = append(fd, e)
		e = &discordgo.MessageEmbedField{}
		e.Name = "#關注更新資訊"
		e.Value = ":100: 如果notion上面有page更新了，我也會通知大家喔!"
		fd = append(fd, e)
		e = &discordgo.MessageEmbedField{}
		e.Name = "#跟我聊聊天(開發中)"
		e.Value = ":yum: 之後我會跟大家聊天喔!"
		fd = append(fd, e)
		e = &discordgo.MessageEmbedField{}
		e.Name = "#發送開會資訊"
		e.Value = ":+1: 之後新的會議議程新增在notion的會議室頻道，都會由我正式公告!"
		fd = append(fd, e)

		message := discordgo.MessageEmbed{
			Title:       ":fire: 大家好! 我是由陳彥廷獨立開發的整合機器人!",
			Description: "我是專門為workfe設計的lol~ 之後我將會跟大家一起努力!!!",
			Color:       16775936,
			Fields:      fd,
		}

		_, err := client.ChannelMessageSendEmbed(config.MeetingRoomChannelId, &message)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else if (discordMessage.Content == "Yen's Notion Integration Bot 版本更新!"){
		var fd []*discordgo.MessageEmbedField
		var e = &discordgo.MessageEmbedField{}
		e.Name = "#私訊團隊成員 提醒相關issue"
		e.Value = ":rocket: notion上面的issue有提及成員 都會收到貼心提醒喔~"
		fd = append(fd, e)
		e = &discordgo.MessageEmbedField{}
		e.Name = "#多個提及成員"
		e.Value = ":rocket: 若是有多個提及的成員在同一個issue中 全部都會收到通知喔! 大家可以多多善用這提及多人的功能!"
		fd = append(fd, e)
		e = &discordgo.MessageEmbedField{}
		e.Name = "#提供頁面連結"
		e.Value = ":rocket: 新版本的通知內會有頁面連結 點近去就可以進到notion了喔~"
		fd = append(fd, e)

		message := discordgo.MessageEmbed{
			Title:       ":fire: 大家好! 我的新版本已經上線了!!",
			Description: "我是專門為workfe設計的整合機器人~ 之後我將會跟大家一起努力!!!",
			Color:       16775936,
			Fields:      fd,
		}

		_, err := client.ChannelMessageSendEmbed(config.MeetingRoomChannelId, &message)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}



