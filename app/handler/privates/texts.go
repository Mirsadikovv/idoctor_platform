package handlers

var StartUserAdded = map[string]string{
	"uz": "Foydalanuvchi <a href=\"tg://user?id=%d\">%s</a> (@%v) botga qo'shildi",
	"ru": "Пользователь <a href=\"tg://user?id=%d\">%s</a> (@%v) был добавлен в бот",
	"en": "User <a href=\"tg://user?id=%d\">%s</a> (@%v) was added to the bot",
}

var UserEhoToAdmin = map[string]string{
	"uz": "<a href=\"tg://user?id=%d\">%s</a> (@%v) foydalanuvchidan botga xabar:\n\n%v",
	"ru": "Сообщение от пользователя <a href=\"tg://user?id=%d\">%s</a> (@%v):\n\n%v",
	"en": "New message from user <a href=\"tg://user?id=%d\">%s</a> (@%v):\n\n%v",
}

var StartWriteYourPhone = map[string]string{
	"uz": "Iltimos, telefon raqamingizni kiriting:",
	"ru": "Пожалуйста, введите свой номер телефона:",
	"en": "Please enter your phone number:",
}

var SendPhoneNumberMenu = map[string]string{
	"uz": "Menu:",
	"ru": "Меню:",
	"en": "Menu:",
}

var MenuText = map[string]string{
	"uz": "Xush kelibsiz, <a href=\"tg://user?id=%d\">%s</a>!",
	"ru": "Добро пожаловать, <a href=\"tg://user?id=%d\">%s</a>!",
	"en": "Welcome, <a href=\"tg://user?id=%d\">%s</a>!",
}

var HelpText = map[string]string{
	"uz": "Добро пожаловать в бота! \nКоманды:\n/start - Запустить бота\n/help - Помощь\n",
	"ru": "Добро пожаловать в бота! \nКоманды:\n/start - Запустить бота\n/help - Помощь\n",
	"en": "Welcome to the bot! \nCommands:\n/start - Start the bot\n/help - Help\n",
}

var BackChooseAction = map[string]string{
	"ru": "Выберите действие:",
	"uz": "Tanlang:",
	"en": "Choose action:",
}

var ChangeLanguageText = map[string]string{
	"ru": "Выберите язык:",
	"uz": "Tilni tanlang:",
	"en": "Choose language:",
}

var Language = map[string]string{
	"uz": "🇺🇿 O'zbek",
	"ru": "🇷🇺 Русский",
	"en": "🇬🇧 English",
}
