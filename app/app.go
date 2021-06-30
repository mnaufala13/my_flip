package app

import (
	"database/sql"
	bigflipPsql "flip/biglip/repository/postgres"
	bigflipSvc "flip/biglip/repository/svc"
	bigflipUC "flip/biglip/usecase"
	"flip/config"
	withdrawPsql "flip/withdraw/repository/postgres"
	withdrawUC "flip/withdraw/usecase"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	DB         *sql.DB
	Fiber      *fiber.App
	HttpClient *http.Client
	Repository *RepositoryContainer
	Usecase    *UsecaseContainer
}

type RepositoryContainer struct {
	BigflipPsql  bigflipPsql.BigFlipPsqlRepository
	BigflipSvc   bigflipSvc.BigFlipSvcRepository
	WithdrawPsql withdrawPsql.WithdrawRepository
}

type UsecaseContainer struct {
	BigflipUC  bigflipUC.BigFlipUsecase
	WithdrawUC withdrawUC.WithdrawUsecase
}

func NewApp(cfg config.Config) *App {
	db, err := newDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		panic(err)
	}

	hc := &http.Client{}
	rContainer := &RepositoryContainer{}
	rContainer.BigflipPsql = bigflipPsql.NewFlipper(db)
	rContainer.WithdrawPsql = withdrawPsql.NewWithdrawer(db)
	rContainer.BigflipSvc = bigflipSvc.NewFlipper(hc, db, rContainer.BigflipPsql)

	ucContainer := &UsecaseContainer{}
	ucContainer.BigflipUC = bigflipUC.NewBigflipUC(db, rContainer.BigflipSvc, rContainer.BigflipPsql)
	ucContainer.WithdrawUC = withdrawUC.NewWithdrawUC(db, rContainer.WithdrawPsql, ucContainer.BigflipUC)
	fiber := fiber.New()

	return &App{
		DB:         db,
		HttpClient: hc,
		Fiber:      fiber,
		Repository: rContainer,
		Usecase:    ucContainer,
	}
}

func newDB(host, port, user, password, dbname string) (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)
	// open database
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, errors.Wrap(err, "error connect to database")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "error checking connection to database")
	}
	log.Printf("connected to database %s:%s", host, port)
	return db, nil
}
