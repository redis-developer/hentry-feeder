package redis

import (
	"fmt"
	"strconv"

	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/YashKumarVerma/hentry-feeder/internal/config"
	"github.com/YashKumarVerma/hentry-feeder/internal/structure"
)

var client *redistimeseries.Client

// Init connection with redis timeseries database
func Init() {
	client = redistimeseries.NewClient(config.Load.RedisURL, "hentry", nil)
	fmt.Println("redis.connected")
}

// Save a new entry into timeseries
func Save(prefix string, data structure.PostBody) {
	var keyname = prefix + ":" + data.Id
	_, err := client.Info(keyname)
	if err != nil {
		fmt.Println("Key " + keyname + " does not exist")
		client.CreateKeyWithOptions(keyname, redistimeseries.DefaultCreateOptions)
	}

	parsedFloat, err := strconv.ParseFloat(data.V, 64)
	if err != nil {
		fmt.Println("Error converting timestamp to float64")
		fmt.Println(err)
	}
	response, err := client.AddAutoTs(keyname, parsedFloat)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(prefix, response)
}
