package static

import (
	"fmt"

	appeal_dto "github.com/Mirsadikovv/idoctor_platform/src/module/appeal_service/dto"
	appeal_model "github.com/Mirsadikovv/idoctor_platform/src/module/appeal_service/model"
)

type MessageHandler = map[appeal_model.AppealWorkflowAction]MessageFunc
type MessageFunc func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string

var (
	LanguageFromIntToStr = map[int]string{
		1: "ru",
		2: "uz",
		3: "en",
	}
	LanguageFromStrToInt = map[string]int{
		"ru": 1,
		"uz": 2,
		"en": 3,
	}

	MessagesApeealCreate = map[appeal_model.AppealWorkflowAction]map[string]string{
		appeal_model.CREATED: {
			"en": "Dear user, your application has been accepted for processing.\n" +
				"You can check the detailed status of your application on the Online Application Platform of the Ministry of Employment and Labor Inspection,\n" +
				"as well as information about the location of inspectors.\n\n" +
				"Link to platform: <a href=\"https://dmi-staging.mehnat.uz/\">Online Application Platform</a>\n\n" +
				"Application number: <b>%06d</b>\n" +
				"Application password: <tg-spoiler><b>%s</b></tg-spoiler>",

			"ru": "Уважаемый пользователь, Ваша заявка принята на обработку.\n" +
				"Подробно о состоянии заявки Вы можете смотреть на Платформе для онлайн-заявок Министерства занятости и трудовой инспекции,\n" +
				"а также информацию о местонахождении инспекторов.\n\n" +
				"Ссылка на платформу: <a href=\"https://dmi-staging.mehnat.uz/\">Платформа для онлайн-заявок</a>\n\n" +
				"Номер заявки: <b>%06d</b>\n" +
				"Пароль от заявки: <tg-spoiler><b>%s</b></tg-spoiler>",

			"uz": "Hurmatli foydalanuvchi, Sizning murojaatingiz qabul qilindi.\n" +
				"Murojaatingizning batafsil holatini Mehnat va bandlik inspeksiyasi vazirligining Onlayn murojaatlar platformasida ko'rishingiz mumkin,\n" +
				"shuningdek, inspektorlarning joylashuvi haqida ma'lumot.\n\n" +
				"Platformaga havola: <a href=\"https://dmi-staging.mehnat.uz/\">Onlayn murojaatlar platformasi</a>\n\n" +
				"Murojaat raqami: <b>%06d</b>\n" +
				"Murojaat paroli: <tg-spoiler><b>%s</b></tg-spoiler>",
		},

		appeal_model.ACCEPT: {
			"en": "Dear user, your application has been accepted for review.\n" +
				"You can check the status of your application and the inspectors’ location on the Online Application Platform of the Ministry of Employment and Labor Inspection.\n\n" +
				"Link to platform: <a href=\"https://dmi-staging.mehnat.uz/\">Online Application Platform</a>\n\n" +
				"Application number: <b>%06d</b>\n" +
				"Application password: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.ACCEPT][3]) +
				"Accepted by inspector: <b>%s</b>", // CurrentInspectorName

			"ru": "Уважаемый пользователь, Ваша заявка принята на рассмотрение.\n" +
				"Статус заявки и местонахождение инспекторов Вы можете проверить на Платформе для онлайн-заявок Министерства занятости и трудовой инспекции.\n\n" +
				"Ссылка на платформу: <a href=\"https://dmi-staging.mehnat.uz/\">Платформа для онлайн-заявок</a>\n\n" +
				"Номер заявки: <b>%06d</b>\n" +
				"Пароль от заявки: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.ACCEPT][1]) +
				"Принявший инспектор: <b>%s</b>", // CurrentInspectorName

			"uz": "Hurmatli foydalanuvchi, arizangiz ko‘rib chiqish uchun qabul qilindi.\n" +
				"Ariza holati va inspektorlarning joylashuvini Bandlik va mehnat inspeksiyasi onlayn-ariza platformasida ko‘rishingiz mumkin.\n\n" +
				"Platformaga havola: <a href=\"https://dmi-staging.mehnat.uz/\">Onlayn murojaatlar platformasi</a>\n\n" +
				"Murojaat raqami: <b>%06d</b>\n" +
				"Murojaat paroli: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.ACCEPT][2]) +
				"Qabul qilgan inspektor: <b>%s</b>", // CurrentInspectorName
		},

		appeal_model.FORWARD: {
			"en": "Dear user, your application has been forwarded to another inspector.\n" +
				"You can check the status of your application and the inspectors’ location on the Online Application Platform of the Ministry of Employment and Labor Inspection.\n\n" +
				"Link to platform: <a href=\"https://dmi-staging.mehnat.uz/\">Online Application Platform</a>\n\n" +
				"Application number: <b>%06d</b>\n" +
				"Application password: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.FORWARD][3]) +
				"Sent by inspector: <b>%s</b>\n" + // OldInspectorName
				"Ceceived by inspector: <b>%s</b>", // CurrentInspectorName

			"ru": "Уважаемый пользователь, Ваша заявка была перенаправлена другому инспектору.\n" +
				"Статус заявки и местонахождение инспекторов Вы можете проверить на Платформе для онлайн-заявок Министерства занятости и трудовой инспекции.\n\n" +
				"Ссылка на платформу: <a href=\"https://dmi-staging.mehnat.uz/\">Платформа для онлайн-заявок</a>\n\n" +
				"Номер заявки: <b>%06d</b>\n" +
				"Пароль от заявки: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.FORWARD][1]) +
				"Направил инспектор: <b>%s</b>\n" + // OldInspectorName
				"Принял инспектор: <b>%s</b>", // CurrentInspectorName

			"uz": "Hurmatli foydalanuvchi, arizangiz boshqa inspektorga yo‘naltirildi.\n" +
				"Ariza holati va inspektorlarning joylashuvini Bandlik va mehnat inspeksiyasi onlayn-ariza platformasida ko‘rishingiz mumkin.\n\n" +
				"Platformaga havola: <a href=\"https://dmi-staging.mehnat.uz/\">Onlayn murojaatlar platformasi</a>\n\n" +
				"Murojaat raqami: <b>%06d</b>\n" +
				"Murojaat paroli: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.FORWARD][2]) +
				"Yuborgan inspektor: <b>%s</b>\n" + // OldInspectorName
				"Qabul qilgan inspektor: <b>%s</b>", // CurrentInspectorName
		},

		appeal_model.TRANSFER: {
			"en": "Dear user, your application has been transferred to another inspection office.\n" +
				"You can check the status of your application and the inspectors’ location on the Online Application Platform of the Ministry of Employment and Labor Inspection.\n\n" +
				"Link to platform: <a href=\"https://dmi-staging.mehnat.uz/\">Online Application Platform</a>\n\n" +
				"Application number: <b>%06d</b>\n" +
				"Application password: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.TRANSFER][3]) +
				"Sent from inspection office: <b>%s</b>\n" + // OldorganizationName
				"Received by inspection office: <b>%s</b>", // CurrentorganizationName

			"ru": "Уважаемый пользователь, Ваша заявка была перенаправлена в другую инспекцию.\n" +
				"Статус заявки и местонахождение инспекторов Вы можете проверить на Платформе для онлайн-заявок Министерства занятости и трудовой инспекции.\n\n" +
				"Ссылка на платформу: <a href=\"https://dmi-staging.mehnat.uz/\">Платформа для онлайн-заявок</a>\n\n" +
				"Номер заявки: <b>%06d</b>\n" +
				"Пароль от заявки: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.TRANSFER][1]) +
				"Отправившая инспекция: <b>%s</b>\n" + // OldorganizationName
				"Принявшая инспекция: <b>%s</b>", // CurrentorganizationName

			"uz": "Hurmatli foydalanuvchi, arizangiz boshqa inspeksiyaga yo‘naltirildi.\n" +
				"Ariza holati va inspektorlarning joylashuvini Bandlik va mehnat inspeksiyasi onlayn-ariza platformasida ko‘rishingiz mumkin.\n\n" +
				"Platformaga havola: <a href=\"https://dmi-staging.mehnat.uz/\">Onlayn murojaatlar platformasi</a>\n\n" +
				"Murojaat raqami: <b>%06d</b>\n" +
				"Murojaat paroli: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.TRANSFER][2]) +
				"Yuborgan inspeksiya: <b>%s</b>\n" + // OldorganizationName
				"Qabul qilgan inspeksiya: <b>%s</b>", // CurrentorganizationName
		},

		appeal_model.REJECT: {
			"en": "Dear user, your application has been rejected.\n" +
				"You can check the status of your application and the inspectors’ location on the Online Application Platform of the Ministry of Employment and Labor Inspection.\n\n" +
				"Link to platform: <a href=\"https://dmi-staging.mehnat.uz/\">Online Application Platform</a>\n\n" +
				"Application number: <b>%06d</b>\n" +
				"Application password: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.REJECT][3]) +
				"Application rejected by inspector: <b>%s</b>", // CurrentInspectorName

			"ru": "Уважаемый пользователь, Ваша заявка была отклонена.\n" +
				"Статус заявки и местонахождение инспекторов Вы можете проверить на Платформе для онлайн-заявок Министерства занятости и трудовой инспекции.\n\n" +
				"Ссылка на платформу: <a href=\"https://dmi-staging.mehnat.uz/\">Платформа для онлайн-заявок</a>\n\n" +
				"Номер заявки: <b>%06d</b>\n" +
				"Пароль от заявки: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.REJECT][1]) +
				"Заявку отклонил инспектор: <b>%s</b>", // CurrentInspectorName

			"uz": "Hurmatli foydalanuvchi, arizangiz rad etildi.\n" +
				"Ariza holati va inspektorlarning joylashuvini Bandlik va mehnat inspeksiyasi onlayn-ariza platformasida ko‘rishingiz mumkin.\n\n" +
				"Platformaga havola: <a href=\"https://dmi-staging.mehnat.uz/\">Onlayn murojaatlar platformasi</a>\n\n" +
				"Murojaat raqami: <b>%06d</b>\n" +
				"Murojaat paroli: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.REJECT][2]) +
				"Arizani rad qilgan inspektor: <b>%s</b>", // CurrentInspectorName
		},

		//Dear user, your application has been successfully closed.
		// You can view the application status and inspectors’ location on the Online Application Platform of the Employment and Labor Inspectorate.
		appeal_model.COMPLETE: {
			"en": "Dear user your application has been successfully closed.\n" +
				"You can check the status of your application and the inspectors’ location on the Online Application Platform of the Ministry of Employment and Labor Inspection.\n\n" +
				"Link to platform: <a href=\"https://dmi-staging.mehnat.uz/\">Online Application Platform</a>\n\n" +
				"Application number: <b>%06d</b>\n" +
				"Application password: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.COMPLETE][3]),
			"ru": "Уважаемый пользователь, ваша заявка успешно закрыта.\n" +
				"Статус заявки и местонахождение инспекторов Вы можете проверить на Платформе для онлайн-заявок Министерства занятости и трудовой инспекции.\n\n" +
				"Ссылка на платформу: <a href=\"https://dmi-staging.mehnat.uz/\">Платформа для онлайн-заявок</a>\n\n" +
				"Номер заявки: <b>%06d</b>\n" +
				"Пароль от заявки: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.COMPLETE][1]),
			"uz": "Hurmatli foydalanuvchi, arizangiz muvaffaqiyatli yopildi.\n" +
				"Ariza holati va inspektorlarning joylashuvini Bandlik va mehnat inspeksiyasi onlayn-ariza platformasida ko‘rishingiz mumkin.\n\n" +
				"Platformaga havola: <a href=\"https://dmi-staging.mehnat.uz/\">Onlayn murojaatlar platformasi</a>\n\n" +
				"Murojaat raqami: <b>%06d</b>\n" +
				"Murojaat paroli: <tg-spoiler><b>%s</b></tg-spoiler>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.COMPLETE][2]),
		},
	}

	Messages = map[appeal_model.AppealWorkflowAction]MessageFunc{
		appeal_model.CREATED: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			if msgTemplate, ok := MessagesApeealCreate[appeal_model.CREATED][language]; ok {
				return fmt.Sprintf(msgTemplate, appeal.Id, appeal.Password)
			} else {
				return fmt.Sprintf(MessagesApeealCreate[appeal_model.CREATED]["ru"], appeal.Id, appeal.Password)
			}
		},

		appeal_model.ACCEPT: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			var current string
			if latestWorkflow.CurrentInspectorName != nil {
				current = *latestWorkflow.CurrentInspectorName
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.ACCEPT][language]; ok {
				return fmt.Sprintf(msgTemplate,
					appeal.Id,
					appeal.Password,
					current,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.ACCEPT]["ru"],
					appeal.Id,
					appeal.Password,
					current,
				)
			}
		},

		appeal_model.FORWARD: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			var from, to string
			if latestWorkflow.FromInspectorName != nil {
				from = *latestWorkflow.FromInspectorName
			}

			if latestWorkflow.ToInspectorName != nil {
				to = *latestWorkflow.ToInspectorName
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.FORWARD][language]; ok {
				return fmt.Sprintf(
					msgTemplate,
					appeal.Id,
					appeal.Password,
					from,
					to,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.FORWARD]["ru"],
					appeal.Id,
					appeal.Password,
					from,
					to,
				)
			}
		},

		appeal_model.TRANSFER: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			var from, to string
			if latestWorkflow.FromOrgName != nil {
				from = *latestWorkflow.FromOrgName
			}

			if latestWorkflow.ToOrgName != nil {
				to = *latestWorkflow.ToOrgName
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.TRANSFER][language]; ok {
				return fmt.Sprintf(msgTemplate,
					appeal.Id,
					appeal.Password,
					from,
					to,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.TRANSFER]["ru"],
					appeal.Id,
					appeal.Password,
					from,
					to,
				)
			}
		},

		appeal_model.REJECT: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			var current string
			if latestWorkflow.CurrentInspectorName != nil {
				current = *latestWorkflow.CurrentInspectorName
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.REJECT][language]; ok {
				return fmt.Sprintf(msgTemplate,
					appeal.Id,
					appeal.Password,
					current,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.REJECT]["ru"],
					appeal.Id,
					appeal.Password,
					current,
				)
			}
		},

		appeal_model.COMPLETE: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			if msgTemplate, ok := MessagesApeealCreate[appeal_model.APPROVED][language]; ok {
				return fmt.Sprintf(
					msgTemplate,
					appeal.Id,
					appeal.Password,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.COMPLETE]["ru"],
					appeal.Id,
					appeal.Password,
				)
			}
		},
	}
)
