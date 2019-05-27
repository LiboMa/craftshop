package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/LiboMa/otcmarket/common"
	"github.com/LiboMa/otcmarket/conf"
	"github.com/LiboMa/otcmarket/markets"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// get config from config file

	configfile := flag.String("conf", "config.json", "specify json conf file for init server")
	flag.Parse()

	if configfile == nil {
		log.Fatal("a json format file needed")
	}

	appconfig := conf.LoadConfig(configfile)

	//fmt.Println(appconfig.RedisDSN)
	// fmt.Println(appconfig.RequestLog)

	// init db connection
	db := common.InitDB()
	defer db.Close()
	// init redis connection
	cacheclient := common.InitCache(appconfig.RedisDSN)
	defer cacheclient.Close()

	// init log files
	// logfile, _ := os.OpenFile(appconfig.RequestLog, os.O_CREATE|os.O_APPEND, 0644)
	// gin.DefaultWriter = io.MultiWriter(logfile)

	// errlogfile, _ := os.OpenFile(appconfig.ErrorLog, os.O_CREATE|os.O_APPEND, 0644)
	// gin.DefaultErrorWriter = io.MultiWriter(errlogfile)

	// start gin
	r := gin.Default()
	// init routers
	//
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	v1 := r.Group("/api")

	markets.MarketsRegister(v1.Group("/markets"))
	markets.MarketsAnonymousRegister(v1.Group("/markets"))

	testMock := r.Group("/api/ping")

	testMock.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// db operationss
	go markets.TaskRunner(30)

	if appconfig.ServicePort == 0 {

		r.Run(":8080")
	} else {

		r.Run(fmt.Sprintf(":%d", appconfig.ServicePort))
		//r.Run(strconv.Itoa(appconfig.ServicePort))

	}
}
