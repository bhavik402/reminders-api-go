package app

import (
	"fmt"
	"log"

	"github.com/bhavik402/reminders-api-go/api-rest/internal/data/reminders"
	"go.uber.org/zap"
)

const (
	OPEN_DB_ERR_MSG = "failed to open %q database"
)

type Config struct {
	Port int
	Env  EnvironmentType
	Db   struct {
		DriverName   string
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

type Models struct {
	Reminders *reminders.ReminderModel
}

type App struct {
	Cfg    Config
	Logger *zap.Logger
	Models *Models
}

func NewApp(cfg Config) (*App, error) {

	l, err := getZapLogger(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}

	rms, err := getRemindersModels(cfg.Db.DriverName, cfg.Db.Dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to get reminders models setup %w", err)
	}

	return &App{
		Cfg:    cfg,
		Logger: l,
		Models: &Models{
			Reminders: rms,
		},
	}, nil
}

func getZapLogger(e EnvironmentType) (*zap.Logger, error) {
	switch e {
	case PROD:
		return zap.NewProduction()
	default:
		return zap.NewDevelopment()
	}
}

func (app *App) ToString() string {
	result := fmt.Sprintf("Server Port: %d \nEnvironment: %s", app.Cfg.Port, app.Cfg.Env.ToString())
	localUrl := fmt.Sprintf("\nhttp://localhost:%d", app.Cfg.Port)
	return result + localUrl
}

func getRemindersModels(drv string, dsn string) (*reminders.ReminderModel, error) {
	rm, err := reminders.New(drv, dsn)
	if err != nil {
		return nil, err
	}
	return rm, nil
}
