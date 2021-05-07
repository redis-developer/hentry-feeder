package redis

import (
	"fmt"

	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/YashKumarVerma/hentry-feeder/internal/structure"
)

var client *redistimeseries.Client

// Init connection with redis timeseries database
func Init() {
	client = redistimeseries.NewClient("localhost:6379", "hentry", nil)
}

// Save a new entry into timeseries
func Save(data structure.PostBody) {

	var keyname = data.Id
	_, err := client.Info(keyname)
	if err != nil {
		fmt.Println("Key " + keyname + " does not exist")
		client.CreateKeyWithOptions(keyname, redistimeseries.DefaultCreateOptions)
	}

	response, err := client.AddAutoTs(keyname, float64(data.V))

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(response)
}
