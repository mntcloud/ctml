package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mntcloud/ctml/internal/parser"
	"github.com/mntcloud/watcher"
)

var template = `<html>
	<head>
		<script type="text/javascript" src="%s"></script>
	</head>
	%s
</html>`

func genHandleRoot(r io.Reader) func(*gin.Context) {
	return func(c *gin.Context) {
		p := parser.New(r)

		if err := p.Do(); err != nil {
			log.Printf("error on parsing: %s", err)
			return
		}

		c.Data(
			http.StatusOK,
			"text/html",
			[]byte(fmt.Sprintf(template, "assets/reload.js", p.AST.Print(0))),
		)
	}
}

func genHandleReload() func(*gin.Context) {
	var watch = watcher.New()

	var upgrader = websocket.Upgrader{}

	watch.Add("root.ctml")

	return func(c *gin.Context) {
		ws, upgradeErr := upgrader.Upgrade(c.Writer, c.Request, nil)
		originalClose := ws.CloseHandler()

		if upgradeErr != nil {
			log.Println("upgrade:", upgradeErr)
			return
		}

		ws.SetCloseHandler(func(code int, text string) error {
			watch.Close()

			return originalClose(code, text)
		})
		defer ws.Close()

		go func() {
			for {
				select {
				case event := <-watch.Event:
					f, _ := os.Open(event.Path)

					p := parser.New(f)
					if err := p.Do(); err != nil {
						log.Printf("watcher event: err parsing file: %s", err)
						continue
					}

					ws.WriteMessage(1, []byte(p.AST.Print(0)))
				case <-watch.Closed:
					return
				}
			}
		}()

		if err := watch.Start(time.Millisecond * 100); err != nil {
			log.Printf("watcher error: %s", err)
		}
	}
}
