package main

import (
	"flag"
	"log"
	"timeslot-service/internal/app"
)

func main() {
	//ctx := context.Background() обычно прокидывается дальше во все компоненты
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("could not create app: %s", err.Error())
		return
	}

	flag.Parse()
	if err := a.Run(); err != nil {
		log.Fatalf("could not run app: %s", err.Error())
	}
}
