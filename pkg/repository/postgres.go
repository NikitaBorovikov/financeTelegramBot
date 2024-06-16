package repository

import (
	"database/sql"
	"financeBot/pkg/config"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type userSelect struct {
	summ     float64
	category string
}

type PostgresDB struct {
	db       *sql.DB
	messages *config.Messages
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

func (p *PostgresDB) AddExpence(user_id int64, sum float64, category string, date string) error {
	result, err := p.db.Exec("insert into expence (user_id, summ, category, expence_date) values ($1, $2, $3, $4)", user_id, sum, category, date)
	if err != nil {
		log.Printf("Insert request error user_id: %d", user_id)
		return err
	}
	log.Printf("successful addition of data user_id: %d", user_id)
	fmt.Println(result.RowsAffected())
	return nil
}

func (p *PostgresDB) AddIncome(user_id int64, sum float64, category string, date string) error {
	result, err := p.db.Exec("insert into income (user_id, summ, category, income_date) values ($1, $2, $3,$4)", user_id, sum, category, date)
	if err != nil {
		log.Printf("Insert request error user_id: %d", user_id)
		return err
	}
	log.Printf("successful addition of data user_id: %d", user_id)
	fmt.Println(result.RowsAffected())
	return nil
}

func (p *PostgresDB) DelExpence(user_id int64) error {
	result, err := p.db.Exec("DELETE FROM expence WHERE user_id = $1", user_id)
	if err != nil {
		log.Printf("Delete request error user_id: %d", user_id)
		return err
	}
	log.Printf("successful delete of data user_id: %d", user_id)
	fmt.Println(result.RowsAffected())
	return nil
}

func (p *PostgresDB) DelIncome(user_id int64) error {
	result, err := p.db.Exec("DELETE FROM income WHERE user_id = $1", user_id)
	if err != nil {
		log.Printf("Delete request error user_id: %d", user_id)
		return err
	}
	log.Printf("successful delete of data user_id: %d", user_id)
	fmt.Println(result.RowsAffected())
	return nil
}

func (p *PostgresDB) GetExpence(user_id int64, startDateInterval, endDateIntervale string) []userSelect {
	rows, err := p.db.Query("SELECT summ, category FROM expence WHERE user_id = $1 AND expence_date >= $2 AND expence_date <= $3", user_id, startDateInterval, endDateIntervale)
	if err != nil {
		panic(err)
	}
	log.Printf("successful select form expence table user_id: %d", user_id)
	userExpence := []userSelect{}

	for rows.Next() {
		p := userSelect{}
		err := rows.Scan(&p.summ, &p.category)
		if err != nil {
			log.Print(err)
			continue
		}
		userExpence = append(userExpence, p)
	}
	return userExpence

}

func (p *PostgresDB) GetIncome(user_id int64, startDateInterval, endDateIntervale string) []userSelect {
	rows, err := p.db.Query("SELECT summ, category FROM income WHERE user_id = $1 AND income_date >= $2 AND income_date <= $3", user_id, startDateInterval, endDateIntervale)
	if err != nil {
		panic(err)
	}
	log.Printf("successful select form income table user_id: %d", user_id)

	userIncome := []userSelect{}

	for rows.Next() {
		p := userSelect{}
		err := rows.Scan(&p.summ, &p.category)
		if err != nil {
			log.Print(err)
			continue
		}
		userIncome = append(userIncome, p)
	}
	return userIncome
}

func (r *PostgresDB) DelLastEspence(user_id int64) error {
	result, err := r.db.Exec("DELETE FROM expence WHERE user_id = $1 AND expence_id = (SELECT expence_id FROM expence ORDER BY expence_id DESC LIMIT 1)", user_id)
	if err != nil {
		log.Printf("error delete last expence user_id %d", user_id)
		return err
	}
	log.Printf("successful delete last expence user_id: %d", user_id)
	fmt.Println(result.RowsAffected())
	return nil

}

func (r *PostgresDB) DelLastIncome(user_id int64) error {
	result, err := r.db.Exec("DELETE FROM income WHERE user_id = $1 AND income_id = (SELECT income_id FROM income ORDER BY income_id DESC LIMIT 1)", user_id)
	if err != nil {
		log.Printf("error delete last income user_id: %d", user_id)
		return err
	}
	log.Printf("successful delete last income user_id %d", user_id)
	fmt.Println(result.RowsAffected())
	return nil
}

func (r *PostgresDB) MakeExpenceString(user_id int64, startDateInterval, endDateIntervale string) string {
	userExpence := r.GetExpence(user_id, startDateInterval, endDateIntervale)
	if len(userExpence) == 0 {
		return r.messages.ErrorNotExpence
	}
	resultMessage := ("Ваши расходы:\n")
	resultSumm := 0.0
	for _, p := range userExpence {
		expenceRow := "• " + p.category + "-" + strconv.FormatFloat(p.summ, 'f', -1, 64) + "\n"
		resultMessage += expenceRow
		resultSumm += p.summ
	}
	resultMessage += "📍Итог: " + strconv.FormatFloat(resultSumm, 'f', -1, 64)
	return resultMessage
}

func (r *PostgresDB) MakeIncomeString(user_id int64, startDateInterval, endDateIntervale string) string {
	userIncome := r.GetIncome(user_id, startDateInterval, endDateIntervale)
	if len(userIncome) == 0 {
		return r.messages.ErrorNotIncome
	}
	resultMessage := ("Ваши расходы:\n")
	resultSumm := 0.0
	for _, p := range userIncome {
		expenceRow := "• " + p.category + "-" + strconv.FormatFloat(p.summ, 'f', -1, 64) + "\n"
		resultMessage += expenceRow
		resultSumm += p.summ
	}
	resultMessage += "📍Итог: " + strconv.FormatFloat(resultSumm, 'f', -1, 64)
	return resultMessage
}

func (r *PostgresDB) MakeReportString(user_id int64, startDateInterval, endDateIntervale string) string {
	userExpence := r.GetExpence(user_id, startDateInterval, endDateIntervale)
	userIncome := r.GetIncome(user_id, startDateInterval, endDateIntervale)
	expenceSum := 0.0
	incomeSum := 0.0
	categoryExpenceSum := map[string]float64{}
	categoryIncomeSum := map[string]float64{}
	for _, p := range userExpence {
		expenceSum += p.summ
		categoryExpenceSum[p.category] += p.summ
	}
	for _, p := range userIncome {
		incomeSum += p.summ
		categoryIncomeSum[p.category] += p.summ
	}
	maxExpenceCategory := ""
	maxExpenceSum := 0.0

	for key, value := range categoryExpenceSum {
		if value > maxExpenceSum {
			maxExpenceSum = value
			maxExpenceCategory = key
		}
	}

	maxIncomeCategory := ""
	maxIncomeSum := 0.0

	for key, value := range categoryIncomeSum {
		if value > maxIncomeSum {
			maxIncomeSum = value
			maxIncomeCategory = key
		}
	}

	resultMessage := ""
	if incomeSum > expenceSum {
		diff := incomeSum - expenceSum
		resultMessage = fmt.Sprintf("•Ваш доход за выбранный месяц: %.0f \n•Ваши раходы за выбранный месяц: %.0f \n•Вы заработали на %.0f больше, чем потратили\n•Категория, на которую вы больше всего потратили: %s - %.0f\n•Основной источник дохода: %s - %.0f", incomeSum, expenceSum, diff, maxExpenceCategory, maxExpenceSum, maxIncomeCategory, maxIncomeSum)
	} else {
		diff := expenceSum - incomeSum
		resultMessage = fmt.Sprintf("•Ваш доход за выбранный месяц: %.0f \n•Ваши раходы за выбранный месяц: %.0f \n•Вы заработали на %.0f меньше, чем потратили\n•Категория, на которую вы больше всего потратили: %s - %.0f\n•Основной источник дохода: %s - %.0f", incomeSum, expenceSum, diff, maxExpenceCategory, maxExpenceSum, maxIncomeCategory, maxIncomeSum)
	}
	return resultMessage
}
