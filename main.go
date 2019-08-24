package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/LiboMa/otcmarket/common"
	"github.com/LiboMa/otcmarket/conf"
	"github.com/LiboMa/otcmarket/markets"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*CORS middleware settings*/
// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }

func main() {

	// get config from config file

	configfile := flag.String("conf", "config.json", "specify json conf file for init server")
	flag.Parse()
	if configfile == nil {
		log.Fatal("a json format file needed")
	}

	appconfig := conf.LoadConfig(configfile)

	fmt.Println(appconfig.RedisDSN)
	//fmt.Println(appconfig.RequestLog)

	// init db connection
	if appconfig.ENV == "dev-local" {
		appconfig.MysqlDSN = os.Getenv("MYSQL_DSN")
		appconfig.RedisDSN = os.Getenv("REDIS_DSN")
	}
	db := common.InitDB(appconfig.MysqlDSN)
	defer db.Close()
	// init redis connection
	cacheclient := common.InitCache(appconfig.RedisDSN)
	defer cacheclient.Close()

	// init log operator
	// _file := filepath.Base(appconfig.RequestLog)
	_ = common.GetOrCreateDir(filepath.Dir(appconfig.RequestLog)) // check dir exist or not and created

	logfile, err := os.OpenFile(appconfig.RequestLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) //logfile, err := os.Create(appconfig.LogPath)
	//logfile, err := os.OpenFile("./request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644) //logfile, err := os.Create(appconfig.LogPath)
	if err != nil {
		panic(err)
	}
	// normal log writter
	gin.DefaultWriter = logfile // a shortcut of gin.DefaultWriter = io.MultiWriter(logfile)
	// error log writter
	gin.DefaultErrorWriter = logfile
	log.SetOutput(gin.DefaultWriter)

	// start gin
	r := gin.Default()
	// init routers
	//
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "x-requested-with"}
	config.AllowMethods = []string{"*"}
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

	// Schedule Tasks
	go markets.TaskRunner(30)

	if appconfig.ServicePort == 0 {

		r.Run(":8080")
	} else {
		r.Run(fmt.Sprintf("%s:%d", appconfig.ServiceHost, appconfig.ServicePort))
		//r.Run(strconv.Itoa(appconfig.ServicePort))

	}
}
