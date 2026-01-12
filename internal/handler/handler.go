package handler

import (
	"github.com/TelitsynNikita/test_telegram_bot/internal/service"
	"github.com/yanzay/tbot/v2"
)

type Handler struct {
	Server  *tbot.Server
	Client  *tbot.Client
	Service service.Service
}

func NewHandler(token string, service service.Service) *Handler {
	return &Handler{
		Server:  tbot.New(token),
		Service: service,
	}
}

func (h *Handler) InitHandlers() {
	h.Client = h.Server.Client()

	h.Server.HandleMessage("/start", func(m *tbot.Message) {
		btnRock := tbot.InlineKeyboardButton{
			Text:         "Rock",
			CallbackData: "rock",
		}

		btnPaper := tbot.InlineKeyboardButton{
			Text:         "Paper",
			CallbackData: "paper",
		}

		btnScissors := tbot.InlineKeyboardButton{
			Text:         "Scissors",
			CallbackData: "scissors",
		}

		h.Server.Client().SendMessage(m.Chat.ID, "Pick an option", tbot.OptInlineKeyboardMarkup(&tbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]tbot.InlineKeyboardButton{
				{btnRock, btnPaper, btnScissors},
			},
		}))
	})

	h.Server.HandleCallback(func(query *tbot.CallbackQuery) {
		h.Server.Client().SendMessage(query.Message.Chat.ID, query.Data)
	})
}
