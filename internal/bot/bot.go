package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/viper"
	"log"
	"tg-bot-youtube-parser/config"
)

func Run() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(viper.GetString("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	listenMessages(bot)
	//log.Printf("Authorized on account %s", bot.Self.UserName)

}

func listenMessages(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	chatId := viper.GetInt64("TELEGRAM_CHAT_ID")

	urls, err := GetUrls()
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		err := gocron.Every(1).Minute().Do(sendUrlMessages, urls, chatId, bot)
		if err != nil {
			log.Fatal(err)
			return
		}
		<-gocron.Start()
	}
}

func sendUrlMessages(urls [2]string, chatId int64, bot *tgbotapi.BotAPI) {
	for _, url := range urls {
		videoUrl, err := GetLastVideo(url)
		if err != nil {
			log.Fatal(err)
			return
		}

		msg := tgbotapi.NewMessage(chatId, videoUrl)

		_, err = bot.Send(msg)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
