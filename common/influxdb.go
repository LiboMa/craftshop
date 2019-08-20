package common

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func InfluxQuery(query string) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	//q := client.NewQuery("SELECT count(value) FROM cpu_load", "mydb", "")
	q := client.NewQuery("SELECT count(value) FROM cpu_load", "markets.huobi", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}
