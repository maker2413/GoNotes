package main

import (
	"database/sql"
	"encoding/gob"
	"final-project/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"

func main() {
	// Connect to database
	db := initDB()

	// Create sessions
	session := initSession()

	// Create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create some channels

	// Create waitgroup
	wg := sync.WaitGroup{}

	// Set up the application config
	app := Config{
		Session:       session,
		DB:            db,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Wait:          &wg,
		Models:        data.New(db),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	// Set up mail
	app.Mailer = app.createMail()
	go app.listenForMail()

	// Listen for signals
	go app.listenForShutdown()

	// Listen for erros
	go app.listenForErrors()

	// Listen for web connections
	app.serve()
}

func (app *Config) listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.ErrorChanDone:
			return
		}
	}
}

func (app *Config) serve() {
	// Start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}

	return conn
}

// connectToDB tries to connect to postgres, and backs off until a connection
// is made, or we have not connected after 10 tries
func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Println("connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("Backing off for 1 second...")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

// openDB opens a connection to Postgres, using the DSN read from the
// environment variable DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// initSession sets up a session, using Redis for session store
func initSession() *scs.SessionManager {
	gob.Register(data.User{})

	// set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

// initRedis returns a pool of connections to Redis using the environment
// variable REDIS
func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// Perform any cleanup tasks
	app.InfoLog.Println("would run cleanup tasks...")

	// Block until WaitGroup is empty
	app.Wait.Wait()

	app.Mailer.DoneChan <- true
	app.ErrorChanDone <- true

	app.InfoLog.Println("closing channels and shutting down application...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
	close(app.ErrorChan)
	close(app.ErrorChanDone)
}

func (app *Config) createMail() Mail {
	// Create channels
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	m := Mail{
		Domain:      "127.0.0.1",
		Host:        "127.0.0.1",
		Port:        1025,
		Encryption:  "none",
		FromName:    "Info",
		FromAddress: "info@mycompany.com",
		Wait:        app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
	}

	return m
}
