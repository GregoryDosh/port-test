package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	cli "github.com/urfave/cli"
	"github.com/valyala/fasthttp"
)

var (
	welcomeMessage string
)

func main() {
	app := cli.NewApp()
	app.Name = "port-test"
	app.Usage = "opens ports for connectivity testing purposes"
	app.Version = "0.1"
	app.Action = httpListen
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "port,p",
			Usage:  "TCP `port` to listen on.",
			EnvVar: "TCP_PORT",
			Value:  "80",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	if _, err := fmt.Fprint(ctx, welcomeMessage); err != nil {
		log.Error(err)
	}
}

func serverStart(httpPort int) {
	if err := fasthttp.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", httpPort), requestHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

func httpListen(c *cli.Context) {
	httpPort := c.Int("port")
	welcomeMessage = fmt.Sprintf("Port %d open!", httpPort)

	log.Infof("starting server on port %d", httpPort)

	go serverStart(httpPort)
	select {}
}
