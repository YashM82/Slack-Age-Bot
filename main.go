package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)

		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-8034411857650-80*****48183201-eRhz8SBdW******viEfT2XED")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0812TKBABW-*******-18b0d49968b2c7255c3e348e0c4f7****888ebf6195d4a8b9ebaf5f54281b7331b")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	ctx, cancel := context.WithCancel(context.Background())

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"g"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2024 - yob
			r := fmt.Sprintf("age in years is %d", age)
			response.Reply(r)
		},
	})

	bot.Command("end", &slacker.CommandDefinition{
		Description: "End the bot process",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Recieved End")
			response.Reply("Bot closing!")
			cancel()

		},
	})

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot closed")
}
