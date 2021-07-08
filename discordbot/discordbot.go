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
	fmt.Println(detail)
	if(detail.Action == "Create"){
		description = "Notion has something created !"
	} else {
		description = "Notion has something updated !"
	}

	var fd []*discordgo.MessageEmbedField
	for i := 1;i < 4;i++{
		var e discordgo.MessageEmbedField
		e.Name = detail.FieldSet[i].Key
		e.Value = detail.FieldSet[i].Value
		fd = append(fd, &e)
	}

	message = discordgo.MessageEmbed{
		Title:       detail.FieldSet[0].Value,
		Description: description,
		Color:       color,
		Fields:      fd,
	}
	return
}

func SendMessageEmbed(channelId string, message discordgo.MessageEmbed){
	discordClient := CreateBot()
	_, err := discordClient.ChannelMessageSendEmbed(channelId, &message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
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
	}
}