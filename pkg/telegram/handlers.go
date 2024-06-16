package telegram

import (
	"financeBot/pkg/logic"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var typeOfOperation = map[int64]string{}

type AddedDb struct {
	user_id  int64
	sum      float64
	category string
	date     string
}

func NewAddedDb(user_id int64, sum float64, category string, date string) AddedDb {
	return AddedDb{
		user_id:  user_id,
		sum:      sum,
		category: category,
		date:     date,
	}
}

// обработчик команд
func (b *Bot) handlerCommands(update tgbotapi.Update) error {

	log.Printf("Command message:[%s] %s", update.Message.From.UserName, update.Message.Text)

	switch update.Message.Command() {
	case "start":
		return b.handleStartCommand(update)
	case "add_expence":
		return b.handleAddExpenceCommand(update)
	case "add_income":
		return b.handleAddIncomeCommand(update)
	case "get_expence":
		return b.handleRequestExpenceCommand(update)
	case "get_income":
		return b.handleRequestIncomeCommand(update)
	case "get_report":
		return b.handleRequestReport(update)
	case "del_expence_history":
		return b.handleDelExpenceCommand(update)
	case "del_income_history":
		return b.handleDelIncomeCommand(update)
	case "del_last_expence":
		return b.handleDelLastExpenceCommand(update)
	case "del_last_income":
		return b.handleDelLastIncomeCommand(update)
	default:
		return b.handleUnknownCommand(update)
	}
}

// обработчик сообщений
func (b *Bot) handleMessage(update tgbotapi.Update) error {
	log.Printf("Text message:[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message != nil {

		category, sum, err := logic.IsCorrectEnter(update.Message.Text)

		if err == nil {
			currenctTime := logic.GetTime()
			if typeOfOperation[update.Message.Chat.ID] == "addExpence" {
				b.handleAddExpence(update, sum, category, currenctTime)
			} else if typeOfOperation[update.Message.Chat.ID] == "addIncome" {
				b.handleAddIncome(update, sum, category, currenctTime)
			}
		} else {
			err = errEnterData
			b.handleErrors(update.Message.Chat.ID, err)
		}
	}
	return nil
}

// обработчик коллбэков
func (b *Bot) handleCallback(update tgbotapi.Update) error {

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := b.bot.Request(callback); err != nil {
		panic(err)
	}
	switch update.CallbackQuery.Data {
	case "addExpence":
		return b.handleAddExpenceCallback(update)
	case "addIncome":
		return b.handleAddIncomeCallback(update)
	case "getExpence":
		return b.handleRequestExpence(update)
	case "getIncome":
		return b.handleRequestIncome(update)
	case "EXweek":
		return b.handleGetExpence(update, "week")
	case "EXmonth":
		return b.handleGetExpence(update, "month")
	case "EXyear":
		return b.handleGetExpence(update, "year")
	case "EXallTime":
		return b.handleGetExpence(update, "allTime")
	case "INweek":
		return b.handleGetIncome(update, "week")
	case "INmonth":
		return b.handleGetIncome(update, "month")
	case "INyear":
		return b.handleGetIncome(update, "year")
	case "INallTime":
		return b.handleGetIncome(update, "allTime")
	case "1":
		return b.handleGetReport(update, "Jan")
	case "2":
		return b.handleGetReport(update, "Feb")
	case "3":
		return b.handleGetReport(update, "Mar")
	case "4":
		return b.handleGetReport(update, "Apr")
	case "5":
		return b.handleGetReport(update, "May")
	case "6":
		return b.handleGetReport(update, "Jun")
	case "7":
		return b.handleGetReport(update, "Jul")
	case "8":
		return b.handleGetReport(update, "Avg")
	case "9":
		return b.handleGetReport(update, "Sep")
	case "10":
		return b.handleGetReport(update, "Oct")
	case "11":
		return b.handleGetReport(update, "Nov")
	case "12":
		return b.handleGetReport(update, "Dec")
	default:
		return b.handleUnknownCallback(update)
	}

}

func (b *Bot) handleStartCommand(update tgbotapi.Update) error {
	typeOfOperation[update.Message.Chat.ID] = "addExpence" //по умолчанию
	text := fmt.Sprintf("👋Привет, %s!\nВыбери тип операции:", update.Message.From.UserName)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = OperationKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddExpenceCommand(update tgbotapi.Update) error {
	typeOfOperation[update.Message.Chat.ID] = "addExpence"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.EnterExpenceMessage)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddIncomeCommand(update tgbotapi.Update) error {
	typeOfOperation[update.Message.Chat.ID] = "addIncome"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.EnterIncomeMessage)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleRequestExpenceCommand(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.ChooseIntervalMessage)
	msg.ReplyMarkup = ExpenceSelectKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleRequestIncomeCommand(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.ChooseIntervalMessage)
	msg.ReplyMarkup = IncomeSelectKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleDelExpenceCommand(update tgbotapi.Update) error {
	err := b.postgresDb.DelExpence(update.Message.Chat.ID)
	if err != nil {
		b.handleErrors(update.Message.Chat.ID, err)
		return errDelData
	}
	//b.sendMessage(update, successDelMessage)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.SuccessDelMessage)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleDelIncomeCommand(update tgbotapi.Update) error {
	err := b.postgresDb.DelIncome(update.Message.Chat.ID)
	if err != nil {
		b.handleErrors(update.Message.Chat.ID, err)
		return errDelData
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.SuccessDelMessage)
	_, err = b.bot.Send(msg)
	return err

}

func (b *Bot) handleDelLastExpenceCommand(update tgbotapi.Update) error {
	err := b.postgresDb.DelLastEspence(update.Message.Chat.ID)
	if err != nil {
		b.handleErrors(update.Message.Chat.ID, err)
		return errDelData
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.SuccessDelMessage)
	_, err = b.bot.Send(msg)
	return err

}

func (b *Bot) handleDelLastIncomeCommand(update tgbotapi.Update) error {
	err := b.postgresDb.DelLastIncome(update.Message.Chat.ID)
	if err != nil {
		b.handleErrors(update.Message.Chat.ID, err)
		return errDelData
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.SuccessDelMessage)
	_, err = b.bot.Send(msg)
	return err

}

func (b *Bot) handleUnknownCommand(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.UnknowCommandMessage)
	_, err := b.bot.Send(msg)
	return err
}

// добавление в бд
func (b *Bot) handleAddExpence(update tgbotapi.Update, sum float64, category, currenctTime string) error {
	newrequest := NewAddedDb(update.Message.Chat.ID, sum, category, currenctTime)
	err := b.postgresDb.AddExpence(newrequest.user_id, newrequest.sum, newrequest.category, newrequest.date)

	if err != nil {
		b.handleErrors(update.Message.Chat.ID, err)
		return errAddData
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.SuccessAddMessage)
	_, err = b.bot.Send(msg)
	return err

}

// добавление в бд
func (b *Bot) handleAddIncome(update tgbotapi.Update, sum float64, category, currenctTime string) error {
	newrequest := NewAddedDb(update.Message.Chat.ID, sum, category, currenctTime)
	err := b.postgresDb.AddIncome(newrequest.user_id, newrequest.sum, newrequest.category, newrequest.date)
	if err != nil {
		b.handleErrors(update.Message.Chat.ID, err)
		return errAddData
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, b.messages.SuccessAddMessage)
	_, err = b.bot.Send(msg)
	return err

}

// сообщение с клавиатурой выбота временного итнервала
func (b *Bot) handleRequestExpence(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, b.messages.ChooseIntervalMessage)
	msg.ReplyMarkup = ExpenceSelectKeyboard
	_, err := b.bot.Send(msg)
	return err
}

// сообщение с клавиатурой выбора временного итнервала
func (b *Bot) handleRequestIncome(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, b.messages.ChooseIntervalMessage)
	msg.ReplyMarkup = IncomeSelectKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAddExpenceCallback(update tgbotapi.Update) error {
	typeOfOperation[update.CallbackQuery.Message.Chat.ID] = "addExpence"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, b.messages.EnterExpenceMessage)
	_, err := b.bot.Send(msg)
	return err

}

func (b *Bot) handleAddIncomeCallback(update tgbotapi.Update) error {
	typeOfOperation[update.CallbackQuery.Message.Chat.ID] = "addIncome"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, b.messages.EnterIncomeMessage)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCallback(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, b.messages.UnknowCommandMessage)
	_, err := b.bot.Send(msg)
	return err
}

// запрос в бд
func (b *Bot) handleGetExpence(update tgbotapi.Update, interval string) error {

	startDateInterval := logic.HandleTime(interval)
	endDateInterval := logic.GetTime()
	text := b.postgresDb.MakeExpenceString(update.CallbackQuery.Message.Chat.ID, startDateInterval, endDateInterval)
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleGetIncome(update tgbotapi.Update, interval string) error {
	startDateInterval := logic.HandleTime(interval)
	endDateInterval := logic.GetTime()
	text := b.postgresDb.MakeIncomeString(update.CallbackQuery.Message.Chat.ID, startDateInterval, endDateInterval)
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleRequestReport(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери месяц, за который нужно получить отчет")
	msg.ReplyMarkup = ReportMonthKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleGetReport(update tgbotapi.Update, month string) error {
	startDateInterval, endDateInterval := logic.GetDateMonth(month)
	text := b.postgresDb.MakeReportString(update.CallbackQuery.Message.Chat.ID, startDateInterval, endDateInterval)
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err

}
