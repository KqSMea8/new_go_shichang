package main

import (
	"fmt"
	"git.code.oa.com/fip-team/fiorm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"xinxin/service/util"
	log "github.com/sirupsen/logrus"
)

//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=basic_daily_data
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=basic_daily_history
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=goods
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=report_file
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=report_notify
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=user
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=vendor
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=vendor_price
//go:generate db2struct.exe -host=119.29.87.223 -port=3306 -dbname=shichang -user=root -password=dceb66c7d2408b87f9cb5bcbd16316f59c1ddf06e4837297b9c2ffe08fa1f5504dc4821d4628 -key=shichang -table=vendor_price_history
func main(){
	// Creates a router without any middleware by default
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	UseRouters(r)

	initLog()
	initConfig()
	initORM()

	r.Run(":8081")
}

func initConfig(){
	viper.SetEnvPrefix("config")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading config file:%s", err))
		os.Exit(1)
	}
}

func initORM() {
	dbSettings := &fiorm.DbSettings{}
	util.ReadKey("db", dbSettings)
	fiorm.InitDB(dbSettings)
}

func initLog(){
	var filename string = "./logs/log.txt"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	}else{
		log.SetOutput(f)
	}
}