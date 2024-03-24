package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/bhavik402/reminders-api-go/api-rest/pkg/app"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Reminders Service: Setting up")

	// crete configuration
	var cfg app.Config
	var env string
	flag.IntVar(&cfg.Port, "port", 8452, "Server Port")
	flag.StringVar(&env, "env", "dev", "Environment: dev|prod")
	flag.StringVar(&cfg.Db.DriverName, "driverName", "sqlite", "Driver Name: sqlite|postgres|mysql")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "reminders.db", "DB DSN")
	flag.IntVar(&cfg.Db.MaxOpenConns, "db-max-open-conns", 25, "max open connections")
	flag.IntVar(&cfg.Db.MaxIdleConns, "db-max-idle-conns", 25, "max idle connections")
	flag.StringVar(&cfg.Db.MaxIdleTime, "db-max-idle-time", "15m", "max connection idle time")
	flag.Parse()
	cfg.Env = app.ToEnvType(env)

	// create app
	app, err := app.NewApp(cfg)
	if err != nil {
		fmt.Println(err.Error())
		app.Logger.Fatal("Failed to create server",
			zap.Int("port", app.Cfg.Port),
			zap.String("env", env),
		)
	}
	defer app.Logger.Sync()
	fmt.Println(app.ToString())

	// create router
	var handler http.Handler = app.WebApiRouter()

	addr := fmt.Sprintf(":%d", app.Cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("Started Server...")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
		app.Logger.Fatal("Failed to start server")
	}
}
