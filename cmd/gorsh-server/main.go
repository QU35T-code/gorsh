package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/audibleblink/gorsh/internal/dashboard"
	"github.com/briandowns/spinner"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var spin *spinner.Spinner

var opts struct {
	Iface  string `short:"i" long:"host" description:"Interface address on which to bind" default:"127.0.0.1" required:"true"`
	Port   string `short:"p" long:"port" description:"Port on which to bind" default:"9000" required:"true"`
	Socket string `short:"s" long:"socket" description:"Domain socket from which the program reads"`
}

func init() {
	_, err := flags.Parse(&opts)
	_ = err
	// the flags package returns an error when calling --help for
	// some reason so we look for that and exit gracefully
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		spin.Stop()
		os.Exit(1)
	}()
}

func main() {
	fmt.Println(fmt.Sprintf("Listening on : %s:%s", opts.Iface, opts.Port))
	dashboard.Menu()
}
