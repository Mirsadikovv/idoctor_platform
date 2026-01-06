package keyboard

import "fmt"

const MainMenuKeyboard = `
	{
		"keyboard": [
		[
			{
				"text": "%s",
				"web_app": {
					"url": "https://idoctor.tvgo.uz"
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
		"ru": "Панель управления📄",
		"uz": "Panelga kirish📄",
		"en": "Admin panel📄",
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
		AppealKeyboardSend[lang],
		AppealKeyboardLanguage[lang],
	)
}

var Back = map[string]string{
	"ru": "Назад ⬅️",
	"uz": "Ortga ⬅️",
	"en": "Back ⬅️",
}
