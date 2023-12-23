package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	app := &cli.App{
		Name: "ctml",
		Usage: "build html pages in a convenient way",
		Action: func(*cli.Context) error {
			fmt.Println("I don't know this command. Refer to help command for the full list of commands")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "publish",
				Usage: "build page from ctml to html and publish it in the output folder",
				Action: func(ctx *cli.Context) error {
					fmt.Println("starting build...")
					return nil
				},
			},
			{
				Name: "server",
				Usage: "run a web server with hot reloading capabilities",
				Action: func(ctx *cli.Context) error {
					gin.SetMode(gin.ReleaseMode)

					r := gin.New()

					r.GET("/ping", func(c *gin.Context) {
						c.JSON(200, gin.H{
							"message": "pong",
						})
					})

					return r.Run()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}	
}