package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kubikvest/api2/games"
	"gopkg.in/gcfg.v1"
)

func main() {
	app := &App{}
	config, err := app.readConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, _ := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	))

	app.DB = db

	gameM := games.GameMapper{
		DB: db,
	}

	game, err := gameM.GetGame("111")//app.Args.GameId)
	fmt.Println(game, err)
}

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

type Args struct {
	ConfPath string
	Command  string
	GameId   string
}

type App struct {
	DB   *sql.DB
	Args *Args
}

func (app *App) readConfig() (*Config, error) {
	config := &Config{}
	args, err := app.parseFlags(os.Args[1:])
	if err != nil {
		return config, err
	}

	app.Args = args

	err = gcfg.ReadFileInto(config, args.ConfPath)

	return config, err
}

func (app *App) parseFlags(args []string) (*Args, error) {
	usr, err := user.Current()

	fs := flag.NewFlagSet("main", flag.ContinueOnError)
	fs.Usage = usage

	confPath := fs.String("conf", fmt.Sprintf("%s%cconfig.ini", usr.HomeDir, os.PathSeparator), "Path to config")
	err = fs.Parse(args)

	if fs.NArg() != 2 {
		fs.Usage()
		err = errors.New("Not commands")
	}

	a := &Args{
		ConfPath: *confPath,
		Command:  fs.Arg(0),
		GameId:   fs.Arg(1),
	}

	return a, err
}

func usage() {
	fmt.Println(`Управление игрой

Использование:
	games [options] command gameID

Команды:
	get	Отображение игры

gameID:
	Идентификатор игры

Опции:
	-conf	Путь до конфигурационного файла
		По умолчанию файл config.ini из домашней папки
	`)
}
