package tg

import (
	"context"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hardstylez72/cry/internal/log"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Config struct {
	Token          string
	UserRepository repository.UserRepository
}

type Bot struct {
	driver *tgbotapi.BotAPI
	c      *Config
}

func NewBot(ctx context.Context, c *Config) (*Bot, error) {
	// get secret key from keyring

	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		return nil, errors.Wrap(err, "tgbotapi.NewBotAPI")
	}

	log.Log.Info("tg bot running")

	b := &Bot{
		driver: bot,
		c:      c,
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Log.Info("tg bot stop")
				return
			case update, ok := <-bot.GetUpdatesChan(tgbotapi.UpdateConfig{
				Timeout:        int(time.Second * 30),
				AllowedUpdates: nil,
			}):
				if !ok {
					return
				}

				if update.Message == nil {
					continue
				}
				var (
					updateChatID = update.Message.Chat.ID
				)

				if update.Message.IsCommand() {
					url := "drop-hunter.pro"
					if config.CFG.Env != config.Prod {
						url = "cry.monster"
					}

					var text string
					switch update.Message.Command() {
					case "start":
						text = "Send your email address from " + url + " to start receiving alerts"
					default:
						text = "Unknown command. Send /help to see a list of commands."
					}

					if _, err := bot.Send(tgbotapi.NewMessage(updateChatID, text)); err != nil {
						log.Log.Error(zap.Error(errors.Wrap(err, "Send")))
					}
				} else {
					msgCtx, cancel := context.WithTimeout(ctx, time.Second*10)
					if err := b.c.UserRepository.SubscribeAlerts(msgCtx, update.Message.Text, strconv.Itoa(int(updateChatID))); err != nil {
						log.Log.Error(zap.Error(errors.Wrap(err, "SubscribeAlerts")))
						if _, err := bot.Send(tgbotapi.NewMessage(updateChatID, "email not found")); err != nil {
							log.Log.Error(zap.Error(errors.Wrap(err, "Send")))
						}
					} else {
						if _, err := bot.Send(tgbotapi.NewMessage(updateChatID, "subscribed")); err != nil {
							log.Log.Error(zap.Error(errors.Wrap(err, "Send")))
						}
					}
					cancel()

				}

			}
		}

	}()

	return b, err
}

func (b *Bot) ProcessStarted(chatId, processId string) error {
	chatInt, err := strconv.Atoi(chatId)
	if err != nil {
		return err
	}

	domain := "https://cry.monster"
	if config.CFG.Env == config.Prod {
		domain = "https://drop-hunter.pro"
	}
	out := "[process](" + domain + "/process/" + processId + ") started"
	m := tgbotapi.NewMessage(int64(chatInt), out)
	m.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.driver.Send(m); err != nil {
		return errors.Wrap(err, "b.driver.Send")
	}
	return nil
}

func (b *Bot) ProcessFinished(chatId, processId string) error {
	chatInt, err := strconv.Atoi(chatId)
	if err != nil {
		return err
	}

	domain := "https://cry.monster"
	if config.CFG.Env == config.Prod {
		domain = "https://drop-hunter.pro"
	}
	out := "[process](" + domain + "/process/" + processId + ") finished"
	m := tgbotapi.NewMessage(int64(chatInt), out)
	m.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.driver.Send(m); err != nil {
		return errors.Wrap(err, "b.driver.Send")
	}
	return nil
}

func (b *Bot) SendAlert(chatId, processId, text string) error {
	chatInt, err := strconv.Atoi(chatId)
	if err != nil {
		return err
	}

	domain := "https://cry.monster"
	if config.CFG.Env == config.Prod {
		domain = "https://drop-hunter.pro"
	}
	out := "[error](" + domain + "/process/" + processId + ") "
	text = tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text)
	out = out + text
	m := tgbotapi.NewMessage(int64(chatInt), out)
	m.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.driver.Send(m); err != nil {
		return errors.Wrap(err, "b.driver.Send")
	}
	return nil
}
