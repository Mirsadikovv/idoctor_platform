package bot_service

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Mirsadikovv/idoctor_platform/src/common/utils"
	botMsg_dto "github.com/Mirsadikovv/idoctor_platform/src/module/bot_service/dto"
	botMsg_model "github.com/Mirsadikovv/idoctor_platform/src/module/bot_service/model"
	"github.com/Mirsadikovv/shared/logger"
	"github.com/Mirsadikovv/shared/pg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type BotMessageService interface {
	Create(*botMsg_dto.BotMessageCreate) error
	Update(*botMsg_dto.BotMessageUpdate, pg.Filter) error
	Find(pg.Filter) ([]botMsg_dto.BotMessageDto, error)
	SendMessage(botMsg_dto.SendMessageRequest) error
	Init()
}

type botMessageService struct {
	db       *gorm.DB
	log      logger.Logger
	bot      *tgbotapi.BotAPI
	mu       sync.Mutex
	lastSent map[int64]time.Time
}

func NewBotMessageService(db *gorm.DB, log logger.Logger) BotMessageService {
	var token = utils.GetEnv("BOT_TOKEN", "")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = false

	return &botMessageService{
		db:       db,
		log:      log,
		bot:      bot,
		lastSent: make(map[int64]time.Time),
		mu:       sync.Mutex{},
	}
}

func (s *botMessageService) Create(input *botMsg_dto.BotMessageCreate) error {
	var model = &botMsg_model.BotMessage{
		AppealId:   input.AppealId,
		TelegramId: input.TelegramId,
		Message:    input.Message,
	}

	return pg.Create(s.db, model, "id")
}

func (s *botMessageService) Update(input *botMsg_dto.BotMessageUpdate, filter pg.Filter) error {
	if _, err := pg.Update[botMsg_model.BotMessage](s.db, input, filter); err != nil {
		return err
	}

	return nil
}

func (s *botMessageService) Find(filter pg.Filter) ([]botMsg_dto.BotMessageDto, error) {
	return pg.FindWithScan[botMsg_model.BotMessage, botMsg_dto.BotMessageDto](s.db, filter)
}

func (s *botMessageService) SendMessage(input botMsg_dto.SendMessageRequest) error {
	var msg = tgbotapi.NewMessage(input.TelegramId, input.Text)
	msg.ParseMode = "HTML"

	_, err := s.bot.Send(msg)
	return err
}

func (s *botMessageService) Init() {
	for {
		var filter = func(tx *gorm.DB) *gorm.DB {
			return tx.
				Where("is_blocked = ? AND sent_at IS NULL", false).
				Order("created_at ASC").
				Limit(100)
		}

		items, err := s.Find(filter)
		if err != nil {
			log.Error("Failed to fetch messages:", err.Error())
			time.Sleep(time.Minute)
			continue
		}

		// If no messages to send, sleep for 1 minute
		{
			if len(items) == 0 {
				time.Sleep(time.Minute)
				continue
			}
		}

		for _, item := range items {
			go func(msg botMsg_dto.BotMessageDto) {
				if err := s.processMessageWithRateLimit(msg); err != nil {
					s.log.Error("Failed to process message:", err.Error())
				}
			}(item)
		}

		time.Sleep(time.Minute)
	}
}

func (s *botMessageService) processMessageWithRateLimit(item botMsg_dto.BotMessageDto) error {
	s.mu.Lock()

	lastTime, exists := s.lastSent[item.TelegramId]
	if exists {
		var wait = time.Since(lastTime)

		if wait < time.Second {
			time.Sleep(time.Second - wait + time.Millisecond*100)
		}
	}

	s.lastSent[item.TelegramId] = time.Now()
	s.mu.Unlock()

	err := s.SendMessage(
		botMsg_dto.SendMessageRequest{
			TelegramId: item.TelegramId,
			Text:       item.Message,
		},
	)

	if err != nil {
		if strings.Contains(err.Error(), "Too Many Requests") {
			if retryAfter := parseRetryAfter(err.Error()); retryAfter > 0 {
				time.Sleep(time.Duration(retryAfter) * time.Second)
				return s.processMessageWithRateLimit(item)
			}
		}

		if strings.Contains(err.Error(), "bot was blocked by the user") {
			return s.markAsBlocked(item.Id)
		}

		return fmt.Errorf("send error (id=%d): %w", item.Id, err)
	}

	return s.markAsSent(item)
}

func (s *botMessageService) markAsSent(item botMsg_dto.BotMessageDto) error {
	var (
		filter = func(tx *gorm.DB) *gorm.DB { return tx.Where("id = ?", item.Id) }
		now    = time.Now()
	)

	return s.Update(
		&botMsg_dto.BotMessageUpdate{
			TelegramId: item.TelegramId,
			Message:    item.Message,
			SentAt:     now,
			IsBlocked:  false,
		},
		filter,
	)
}

func (s *botMessageService) markAsBlocked(id int64) error {
	var filter = func(tx *gorm.DB) *gorm.DB { return tx.Where("id = ?", id) }
	return s.Update(&botMsg_dto.BotMessageUpdate{IsBlocked: true}, filter)
}

func parseRetryAfter(errMsg string) int {
	var (
		re    = regexp.MustCompile(`retry after (\d+)`)
		match = re.FindStringSubmatch(errMsg)
	)

	{
		if len(match) == 2 {
			sec, err := strconv.Atoi(match[1])
			{
				if err == nil {
					return sec
				}
			}
		}
	}

	return 0
}
