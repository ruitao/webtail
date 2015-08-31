package main

import (
	"bufio"
	"flag"
	"github.com/dustin/go-broadcast"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	broadcaster broadcast.Broadcaster
	addr        = flag.String("addr", "0.0.0.0:2333", "run on given address")
)

func runWebServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetHTMLTemplate(html)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "webtail", nil)
	})
	r.GET("/log", func(c *gin.Context) {
		listener := make(chan interface{})
		broadcaster.Register(listener)
		defer broadcaster.Unregister(listener)

		c.Stream(func(w io.Writer) bool {
			c.SSEvent("message", <-listener)
			return true
		})
	})
	r.Run(*addr)
}

func main() {
	broadcaster = broadcast.NewBroadcaster(10)
	scanner := bufio.NewScanner(os.Stdin)
	go runWebServer()
	for scanner.Scan() {
		broadcaster.Submit(scanner.Text())
	}
}
