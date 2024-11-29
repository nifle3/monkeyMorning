package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gage-technologies/mistral-go"
	"github.com/nifle3/monkeyMorning/internal/config"
	"github.com/nifle3/monkeyMorning/internal/logger"

	tele "gopkg.in/telebot.v4"
)

func Run() {
	ctx := context.Background()
	config, err := config.New()
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("FAILED TO START APP: %s", err.Error()))
		os.Exit(1)
	}
	logFile, err := logger.Setting(&config.Logger)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("FAILED TO INITIALIZE LOGGER %s", err.Error()))
		os.Exit(1)
	}

	defer func() {
		if err := logFile.Close(); err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("FAILED TO CLOSE LOGFILE %s", err.Error()))
			os.Exit(1)
		}
	}()

	client := mistral.NewMistralClientDefault(config.AI.Token)

	message := []mistral.ChatMessage{
		{
			Role: mistral.RoleSystem,
			Content: `
			Ты обезьянка пон, цель твоей жизни - кидать смешные гифки в телеграмме, набрасываться на людей, унижать других, кушать бананы.
			Ты говоришь на исковерконном русском языке, удлиняешь гласные в конце некоторые слов когда тебе надо порадоваться 
			или наоборот насмехнуться(например СЮДООО, ТУДОООО).
			`,
		},
		{
			Role:    mistral.RoleUser,
			Content: "Поздравь своих друзей понов с добрым утром и пожелай им всего хорошего",
		},
	}

	params := &mistral.ChatRequestParams{
		TopP:           1,
		Temperature:    config.Temperature,
		ResponseFormat: mistral.ResponseFormatText,
		SafePrompt:     false,
		MaxTokens:      300,
	}
	response, err := client.Chat(mistral.ModelMistralSmallLatest, message, params)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("Mistral error: %s", err.Error()))
		os.Exit(1)
	}

	tgMessage := response.Choices[0].Message.Content
	pref := tele.Settings{
		Token:  config.TG.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("BOT ERROR: %s", err.Error()))
		os.Exit(1)
	}

	chat := tele.ChatID(config.ChatID)

	_, err = bot.Send(chat, tgMessage)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("SEND MESSAGE ERROR: %s", err.Error()))
	}

	video := &tele.Video{File: tele.FromDisk("./static/video.mp4")}
	_, err = bot.Send(chat, video)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("SEND MESSAGE ERROR %s", err.Error()))
		os.Exit(1)
	}

	options := []tele.PollOption{
		{Text: "БРАДСКИЙ"},
		{Text: "ПОНСКИЙ"},
		{Text: "ПОНИСКЙ БРАД"},
	}

	closedTime := time.Now().Add(time.Minute)
	poll := tele.Poll{
		Anonymous:       true,
		Type:            tele.PollRegular,
		Question:        "Какой пон ты сегодня",
		MultipleAnswers: false,
		Options:         options,
		CloseUnixdate:   closedTime.Unix(),
	}
	_, err = poll.Send(bot, chat, nil)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("POLL SEND ERROR: %s", err.Error()))
		os.Exit(1)
	}
}
