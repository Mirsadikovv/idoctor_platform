package static

// Этот файл был закомментирован, так как он полностью зависит от appeal_service,
// который был удален из проекта. Если вам нужна эта функциональность,
// восстановите appeal_service или переработайте логику.

/*
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

	MessagesApeealCreate = map[appeal_model.AppealWorkflowAction]map[string]string{
		appeal_model.CREATED: {
			"en": "Dear user, your request has been successfully registered.\n\n" +
				"📌 Request ID: <b>%d</b>\n" +
				"🔐 Password: <b>%s</b>\n\n" +
				"Please save this information. You can track the status of your request using the ID and password.",

			"ru": "Уважаемый пользователь, ваша заявка успешно зарегистрирована.\n\n" +
				"📌 Номер заявки: <b>%d</b>\n" +
				"🔐 Пароль: <b>%s</b>\n\n" +
				"Пожалуйста, сохраните эту информацию. Вы можете отслеживать статус вашей заявки по номеру и паролю.",

			"uz": "Hurmatli foydalanuvchi, sizning arizangiz muvaffaqiyatli ro'yxatdan o'tkazildi.\n\n" +
				"📌 Ariza raqami: <b>%d</b>\n" +
				"🔐 Parol: <b>%s</b>\n\n" +
				"Iltimos, bu ma'lumotlarni saqlang. Siz ariza holatini raqam va parol orqali kuzatishingiz mumkin.",
		},
		appeal_model.ACCEPT: {
			"en": "Dear user,\n\n" +
				"📌 Request ID: <b>%d</b>\n" +
				"🔐 Password: <b>%s</b>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.ACCEPT][3]) +
				"\nYour request has been accepted for processing.",

			"ru": "Уважаемый пользователь,\n\n" +
				"📌 Номер заявки: <b>%d</b>\n" +
				"🔐 Пароль: <b>%s</b>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.ACCEPT][1]) +
				"\nВаша заявка принята в обработку.",

			"uz": "Hurmatli foydalanuvchi,\n\n" +
				"📌 Ariza raqami: <b>%d</b>\n" +
				"🔐 Parol: <b>%s</b>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.ACCEPT][2]) +
				"\nSizning arizangiz ko'rib chiqilmoqda.",
		},
		appeal_model.FORWARD: {
			"en": "Dear user,\n\n" +
				"📌 Request ID: <b>%d</b>\n" +
				"🔐 Password: <b>%s</b>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.FORWARD][3]) +
				"\nYour request has been forwarded to %s.",

			"ru": "Уважаемый пользователь,\n\n" +
				"📌 Номер заявки: <b>%d</b>\n" +
				"🔐 Пароль: <b>%s</b>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.FORWARD][1]) +
				"\nВаша заявка направлена в %s.",

			"uz": "Hurmatli foydalanuvchi,\n\n" +
				"📌 Ariza raqami: <b>%d</b>\n" +
				"🔐 Parol: <b>%s</b>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.FORWARD][2]) +
				"\nSizning arizangiz %s ga yo'naltirildi.",
		},
		appeal_model.TRANSFER: {
			"en": "Dear user,\n\n" +
				"📌 Request ID: <b>%d</b>\n" +
				"🔐 Password: <b>%s</b>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.TRANSFER][3]) +
				"\nYour request has been transferred to %s.",

			"ru": "Уважаемый пользователь,\n\n" +
				"📌 Номер заявки: <b>%d</b>\n" +
				"🔐 Пароль: <b>%s</b>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.TRANSFER][1]) +
				"\nВаша заявка передана в %s.",

			"uz": "Hurmatli foydalanuvchi,\n\n" +
				"📌 Ariza raqami: <b>%d</b>\n" +
				"🔐 Parol: <b>%s</b>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.TRANSFER][2]) +
				"\nSizning arizangiz %s ga o'tkazildi.",
		},
		appeal_model.REJECT: {
			"en": "Dear user,\n\n" +
				"📌 Request ID: <b>%d</b>\n" +
				"🔐 Password: <b>%s</b>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.REJECT][3]) +
				"\nReason: %s",

			"ru": "Уважаемый пользователь,\n\n" +
				"📌 Номер заявки: <b>%d</b>\n" +
				"🔐 Пароль: <b>%s</b>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.REJECT][1]) +
				"\nПричина: %s",

			"uz": "Hurmatli foydalanuvchi,\n\n" +
				"📌 Ariza raqami: <b>%d</b>\n" +
				"🔐 Parol: <b>%s</b>\n" +
				fmt.Sprintf("Holat: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.REJECT][2]) +
				"\nSabab: %s",
		},

		appeal_model.COMPLETE: {
			"en": "Dear user, your request has been completed.\n\n" +
				"📌 Request ID: <b>%d</b>\n" +
				"🔐 Password: <b>%s</b>\n" +
				fmt.Sprintf("Status: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.COMPLETE][3]),

			"ru": "Уважаемый пользователь, ваша заявка завершена.\n\n" +
				"📌 Номер заявки: <b>%d</b>\n" +
				"🔐 Пароль: <b>%s</b>\n" +
				fmt.Sprintf("Статус: <b>%s</b>\n", appeal_model.SystemActionComments[appeal_model.COMPLETE][1]),

			"uz": "Hurmatli foydalanuvchi, sizning arizangiz yakunlandi.\n\n" +
				"📌 Ariza raqami: <b>%d</b>\n" +
				"🔐 Parol: <b>%s</b>\n" +
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
			if latestWorkflow.Workflow.ToOrganization == nil {
				return ""
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.ACCEPT][language]; ok {
				return fmt.Sprintf(msgTemplate,
					appeal.Id,
					appeal.Password,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.ACCEPT]["ru"],
					appeal.Id,
					appeal.Password,
				)
			}
		},

		appeal_model.FORWARD: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			if latestWorkflow.Workflow.ToOrganization == nil {
				return ""
			}

			orgName := latestWorkflow.Workflow.ToOrganization.Name
			if latestWorkflow.Workflow.ToOrganization.Translations != nil {
				orgName = latestWorkflow.Workflow.ToOrganization.Translations[0].Translation
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.FORWARD][language]; ok {
				return fmt.Sprintf(
					msgTemplate,
					appeal.Id,
					appeal.Password,
					orgName,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.FORWARD]["ru"],
					appeal.Id,
					appeal.Password,
					orgName,
				)
			}
		},

		appeal_model.TRANSFER: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			if latestWorkflow.Workflow.ToOrganization == nil {
				return ""
			}

			orgName := latestWorkflow.Workflow.ToOrganization.Name
			if latestWorkflow.Workflow.ToOrganization.Translations != nil {
				orgName = latestWorkflow.Workflow.ToOrganization.Translations[0].Translation
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.TRANSFER][language]; ok {
				return fmt.Sprintf(msgTemplate,
					appeal.Id,
					appeal.Password,
					orgName,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.TRANSFER]["ru"],
					appeal.Id,
					appeal.Password,
					orgName,
				)
			}
		},

		appeal_model.REJECT: func(telegramId *int64, language string, latestWorkflow *appeal_dto.LatestWorkflow, appeal *appeal_dto.Appeal) string {
			reason := ""
			if latestWorkflow.Workflow.Comment != nil {
				reason = *latestWorkflow.Workflow.Comment
			}

			if msgTemplate, ok := MessagesApeealCreate[appeal_model.REJECT][language]; ok {
				return fmt.Sprintf(
					msgTemplate,
					appeal.Id,
					appeal.Password,
					reason,
				)
			} else {
				return fmt.Sprintf(
					MessagesApeealCreate[appeal_model.REJECT]["ru"],
					appeal.Id,
					appeal.Password,
					reason,
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
*/
