package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
	Dbhost        string
	Password      string
	Dbuser        string
	Dbname        string
	Dbport        int

	Messages Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	ErrorDelDataMessage string `mapstructure:"errorDelDataMessage"`
	ErrorAddDataMessage string `mapstructure:"errorAddDataMessage"`
	ErrorEnterData      string `mapstructure:"errorEnterData"`
	ErrorNotExpence     string `mapstructure:"errorNotExpence"`
	ErrorNotIncome      string `mapstructure:"errorNotIncome"`
}

type Responses struct {
	SuccessDelMessage     string `mapstructure:"successDelMessage"`
	SuccessAddMessage     string `mapstructure:"successAddMessage"`
	UnknowCommandMessage  string `mapstructure:"unknowCommandMessage"`
	ChooseIntervalMessage string `mapstructure:"chooseIntervalMessage"`
	EnterExpenceMessage   string `mapstructure:"enterExpenceMessage"`
	EnterIncomeMessage    string `mapstructure:"enterIncomeMessage"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("token", &cfg.TelegramToken); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return err
	}
	token, exist := os.LookupEnv("TOKEN")
	if exist {
		cfg.TelegramToken = token
	}

	dbhost, exist := os.LookupEnv("DBHOST")
	if exist {
		cfg.Dbhost = dbhost
	}

	password, exist := os.LookupEnv("DBPASSWORD")
	if exist {
		cfg.Password = password
	}

	user, exist := os.LookupEnv("DBUSER")
	if exist {
		cfg.Dbuser = user
	}

	name, exist := os.LookupEnv("DBNAME")
	if exist {
		cfg.Dbname = name
	}

	port, exist := os.LookupEnv("DBPORT")
	if exist {
		cfg.Dbport, _ = strconv.Atoi(port)
	}
	return nil

}
