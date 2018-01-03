package main

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
	log "gopkg.in/Sirupsen/logrus.v0"
	cli "gopkg.in/urfave/cli.v1"
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
			EnvVar: "HTTP_PORT",
			Value:  "80",
		},
	}
	app.Run(os.Args)
}

func httpListen(c *cli.Context) {
	httpPort := c.Int("port")

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json; charset=utf8")
		fmt.Fprint(ctx, fmt.Sprintf("Port %d open!", httpPort))
	}

	log.Infof("starting server on port %d", httpPort)

	go func() {
		if err := fasthttp.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", httpPort), requestHandler); err != nil {
			log.Fatalf("error in ListenAndServe: %s", err)
		}
	}()
	select {}
}
