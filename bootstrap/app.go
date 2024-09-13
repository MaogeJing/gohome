package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env   *Env
	MySQL gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.MySQL = NewMySQLDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMySQLDBConnection(app.MySQL)
}
