package middleware

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func saveLog(name string, s ...string) {
	// open file
	path := "../appmiddleware/log/" + name + ".log"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, items := range s {
		if _, err := file.WriteString(items + " "); err != nil {
			panic(err)
		}
	}

	if _, err := file.WriteString("\n"); err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = file
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// print time
		t := time.Now()
		name := t.Format("2006-01-02")
		started := t.Format("2006-01-02 15:04:05")

		// executes pending handlers in the chain inside the calling handler.
		ctx.Next()
		// after request
		latency := time.Since(t)

		// print status we are sending
		status := ctx.Writer.Status()

		// print method used
		method := ctx.Request.Method

		// save to log file
		go func() {
			saveLog(name, started, method, ctx.Request.URL.Path, strconv.Itoa(status), latency.String())
		}()
		time.Sleep(time.Millisecond * 10)
	}
}
