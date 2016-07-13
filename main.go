package main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var token = flag.String("token", "", "Telegram Bot API token")
var debug = flag.Bool("debug", false, "Bot debug")
var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type proverbs []string

func (p proverbs) String() string {
	var b bytes.Buffer
	for _, proverb := range p {
		b.WriteString(proverb)
		b.WriteByte('\n')
	}
	return b.String()
}

var all = proverbs{
	"Don't communicate by sharing memory, share memory by communicating",
	"Concurrency is not parallelism",
	"Channels orchestrate, mutexes serialize",
	"The bigger the interface, the weaker the abstraction",
	"Make the zero value useful",
	"interface{} says nothing",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite",
	"A little copying is better than a little dependency",
	"Syscall must always be guarded with build tags",
	"Cgo must always be guarded with build tags",
	"Cgo is not Go",
	"With the unsafe package, there are no guarantees",
	"Clear is better than clever",
	"Reflection is never clear",
	"Errors are values",
	"Don't just check errors, handle them gracefully",
	"Design the architecture, name the components, document the details",
	"Documentation is for users",
	"Don't panic",
}

func allProverbs() string {
	return all.String()
}

func randomProverb() string {
	return all[r.Intn(len(all))]
}

func main() {
	flag.Parse()
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = *debug
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		m := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			var command string
			if update.Message.Entities != nil {
				for _, entity := range *update.Message.Entities {
					if entity.Type == "bot_command" {
						command = update.Message.Text[entity.Offset:entity.Length]
						break
					}
				}
			}
			if command == "/all" {
				m.Text = allProverbs()
			}
			if command == "/random" {
				m.Text = randomProverb()
			}
		}

		m.ReplyToMessageID = update.Message.MessageID

		if len(m.Text) > 0 {
			bot.Send(m)
		}
	}
}
