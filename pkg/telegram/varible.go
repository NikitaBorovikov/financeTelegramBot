package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var OperationKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("📉Добавить расход", "addExpence"),
		tgbotapi.NewInlineKeyboardButtonData("📈Добавить доход", "addIncome"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Список расходов", "getExpence"),
		tgbotapi.NewInlineKeyboardButtonData("Список доходов", "getIncome"),
	),
)

var ExpenceSelectKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("За неделю", "EXweek"),
		tgbotapi.NewInlineKeyboardButtonData("За месяц", "EXmonth"),
		tgbotapi.NewInlineKeyboardButtonData("За год", "EXyear"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("За все время", "EXallTime"),
	),
)

var IncomeSelectKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("За неделю", "INweek"),
		tgbotapi.NewInlineKeyboardButtonData("За месяц", "INmonth"),
		tgbotapi.NewInlineKeyboardButtonData("За год", "INyear"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("За все время", "INallTime"),
	),
)

var ReportMonthKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Январь", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Февраль", "2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Март", "3"),
		tgbotapi.NewInlineKeyboardButtonData("Апрель", "4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Май", "5"),
		tgbotapi.NewInlineKeyboardButtonData("Июнь", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Июль", "7"),
		tgbotapi.NewInlineKeyboardButtonData("Август", "8"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сентябрь", "9"),
		tgbotapi.NewInlineKeyboardButtonData("Октябрь", "10"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ноябрь", "11"),
		tgbotapi.NewInlineKeyboardButtonData("Декабрь", "12"),
	),
)
