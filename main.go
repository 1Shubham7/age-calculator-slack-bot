package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
	"golang.org/x/tools/go/analysis"
)

func printCommandEvents(analysisChannel <-chan *slacker.CommandEvent){
	for event := range analysisChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Command)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5567624998498-5591373117392-KILAcyz0bzA7XWKOc6Bo4OiE")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05GLLRBY5T-5567640932018-3bcb4b009f3fdb481ff6eed5f4852d281622e6b0fdd98646a295a8c2b3052af1")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())


	bot.Command("my yob is <year>", &slacker.CommandDefinition){
		Description: "yob calculator",
		Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter ){
			rear := request.Param("year")

		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}