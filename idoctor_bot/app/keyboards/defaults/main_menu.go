package keyboard

import "fmt"

const MainMenuKeyboard = `
	{
		"keyboard": [
		[
			{
				"text": "%s",
				"web_app": {
					"url": "https://dmi-staging.mehnat.uz/%s/telegram"
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

var MainMenuKeyboardMap = map[string]string{
	"uz": fmt.Sprintf(MainMenuKeyboard, AppealKeyboardSend["uz"], "uz", AppealKeyboardLanguage["uz"]),
	"ru": fmt.Sprintf(MainMenuKeyboard, AppealKeyboardSend["ru"], "ru", AppealKeyboardLanguage["ru"]),
	"en": fmt.Sprintf(MainMenuKeyboard, AppealKeyboardSend["en"], "en", AppealKeyboardLanguage["en"]),
}

var Back = map[string]string{
	"ru": "Назад ⬅️",
	"uz": "Ortga ⬅️",
	"en": "Back ⬅️",
}
