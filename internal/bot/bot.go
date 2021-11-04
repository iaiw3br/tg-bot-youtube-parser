package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/viper"
	"log"
	"tg-bot-youtube-parser/config"
	"tg-bot-youtube-parser/internal/database"
)

func Run() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bot, err := tgbotapi.NewBotAPI(viper.GetString("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	listenMessages(bot, db)
	//log.Printf("Authorized on account %s", bot.Self.UserName)

}

func listenMessages(bot *tgbotapi.BotAPI, db *pgxpool.Pool) {
	chatId := viper.GetInt64("TELEGRAM_CHAT_ID")

	err := gocron.Every(1).Minute().Do(sendUrlMessages, chatId, bot)
	if err != nil {
		log.Fatal(err)
		return
	}

	<-gocron.Start()
}

func sendUrlMessages(chatId int64, bot *tgbotapi.BotAPI, db *pgxpool.Pool) {
	urls, err := database.GetAllURLs(db)
	if err != nil {
		log.Fatal(err)
	}

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
