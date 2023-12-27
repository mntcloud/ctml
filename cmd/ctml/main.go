package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/mntcloud/ctml/internal/parser"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ctml",
		Usage: "build html pages in a convenient way",
		Action: func(*cli.Context) error {
			fmt.Println("I don't know this command. Refer to help command for the full list of commands")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "publish",
				Usage: "build page from ctml to html and publish it to _output folder",
				Action: func(ctx *cli.Context) error {
					f, err := os.Open("root.ctml")

					if err != nil {
						return fmt.Errorf("Open root file: %s", err)
					}

					if err := os.Mkdir("_output", 0660); err != nil {
						if os.IsExist(err) {
							os.Remove("_output")
							os.Mkdir("_output", 0660)
						} else {
							return fmt.Errorf("Make output dir: %s", err)
						}
					}

					p := parser.New(f)
					if err := p.Do(); err != nil {
						return fmt.Errorf("Parser error: %s", err)
					}

					writeFile, _ := os.Create("./_output/index.html")

					writeFile.WriteString(fmt.Sprintf(template, "", p.AST.Print(0)))

					writeFile.Close()

					return nil
				},
			},
			{
				Name: "test",
				Action: func(ctx *cli.Context) error {
					p, _ := os.Executable()

					fmt.Printf("path: %s", p)

					return nil
				},
			},
			{
				Name:  "server",
				Usage: "run a web server with hot reloading capabilities",
				Action: func(ctx *cli.Context) error {
					f, err := os.Open("root.ctml")

					if err != nil {
						return fmt.Errorf("Open root file: %s", err)
					}

					if err := os.Mkdir("_output", 0660); err != nil {
						return fmt.Errorf("Make output dir: %s", err)
					}

					currentExec, _ := os.Executable()

					// Server setup
					gin.SetMode(gin.ReleaseMode)

					r := gin.New()

					r.StaticFile(path.Dir(currentExec)+"/assets/reload.js", "assets/reload.js")
					r.GET("/", genHandleRoot(f))

					r.GET("/reload", genHandleReload())

					return r.Run()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
