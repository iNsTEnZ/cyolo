package main

import (
	"cyolo-exercise/configuration"
	"cyolo-exercise/dao"
	"cyolo-exercise/output"
	"cyolo-exercise/rest"
	"cyolo-exercise/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Module interface {
	Start()
	Stop()
}

func main() {
	modules := make([]Module, 0)
	inMemory := dao.NewInMemory()
	srv := service.NewService(inMemory)
	cfg := configuration.GetRestConfiguration()
	modules = append(modules, rest.New(cfg, srv, output.NewPrinter()))

	RunModules(modules)
}

// RunModules runs each of the modules in a separate goroutine.
func RunModules(modules []Module) {
	if len(modules) > 0 {
		log.Println("Starting")

		for _, m := range modules {
			go m.Start()
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		for _, m := range modules {
			m.Stop()
		}

		log.Println("Stopping")
	}
}
