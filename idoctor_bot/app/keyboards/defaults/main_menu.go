package keyboard

import "fmt"

const MainMenuKeyboard = `
	{
		"keyboard": [
		[
			{
				"text": "%s",
				"web_app": {
					"url": "https://mirsadikovv.github.io?user_id=%d"
				}
			}
		],
		[
			{
				"text": "%s",
				"callback_data": "language"
			}
		]
		],
		"resize_keyboard": true,
		"one_time_keyboard": true
	}`

var (
	AppealKeyboardSend = map[string]string{
		"ru": "Подать заявку📄",
		"uz": "Murojaat yo'llash📄",
		"en": "Send an application📄",
	}

	AppealKeyboardLanguage = map[string]string{
		"ru": "Поменять язык🌐",
		"uz": "Tilni o'zgartirish🌐",
		"en": "Change language🌐",
	}
)

func GetMainMenuKeyboard(lang string, userID int64) string {
	return fmt.Sprintf(
		MainMenuKeyboard,
		lang,
		userID,
		AppealKeyboardLanguage[lang],
	)
}

var Back = map[string]string{
	"ru": "Назад ⬅️",
	"uz": "Ortga ⬅️",
	"en": "Back ⬅️",
}
