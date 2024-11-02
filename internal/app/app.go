package app

import (
	"timeslot-service/internal/app/handler"
	"timeslot-service/internal/db"
	"timeslot-service/internal/repository"
	"timeslot-service/internal/service/reservation"
	"timeslot-service/internal/service/timeslot"
	"timeslot-service/internal/transport"
)

type App struct {
	timeslotService    *timeslot.Service
	reservationService *reservation.Service
	server             *transport.Server
	handler            *handler.Handler
	repository         *repository.Repository
	db                 *db.DB
}

func NewApp() (*App, error) {
	app := &App{}
	app.initDeps()

	return app, nil
}

func (a *App) Run() error {
	return a.server.Run()
}

func (a *App) initDeps() {
	inits := []func(){
		a.initDB,
		a.initRepository,
		a.initService,
		a.initHandler,
		a.initServer,
	}

	for _, fn := range inits {
		fn()
	}
}

func (a *App) initDB() {
	a.db = db.NewDB()
}

func (a *App) initRepository() {
	a.repository = repository.NewRepository(a.db)
}

func (a *App) initService() {
	a.timeslotService = timeslot.NewService(a.repository)
	a.reservationService = reservation.NewService(a.repository)
}

func (a *App) initHandler() {
	a.handler = handler.NewHandler(a.timeslotService)
}

func (a *App) initServer() {
	opts := transport.Options{
		HttpPort: ":8080",
	}
	a.server = transport.NewServer(opts, a.handler)
}
