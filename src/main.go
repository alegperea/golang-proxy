package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	controllers "github.com/aperea/go-mlproxy/src/controllers"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	f, _ := os.OpenFile("/home/aperea/elk-filebeat/mylog/proxy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)

}

func LoggingMiddleware(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	t := time.Now()
	latency := time.Since(t)

	log.WithFields(log.Fields{
		"Host":    c.Request.Host,
		"Body":    body,
		"Method":  c.Request.Method,
		"Path":    c.Param("proxyPath"),
		"Latency": latency,
	}).Info("This is an info message.")
	c.Next()
}

func RulesMiddleware(c *gin.Context) {

	reader := strings.NewReader(`{"IP":"127.0.0.1"}`)
	request, err := http.NewRequest("GET", "http://localhost:9090/getrule", reader)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        c.Next()
    } else {
		c.JSON(http.StatusOK, gin.H{"Message": "Banned"})
		c.Abort()
		return
	}
	
	

}

func main() {
	r := gin.Default()
	r.Use(LoggingMiddleware)
	r.Use(RulesMiddleware)

	r.Any("/*proxyPath", controllers.Proxy)
	r.Run(":8080")
}
