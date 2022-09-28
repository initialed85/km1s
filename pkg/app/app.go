package app

import "log"

type App struct{}

func New() (*App, error) {
	a := App{}

	return &a, nil
}

func (a *App) Run() error {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)

	return nil
}
