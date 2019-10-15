package conf

import (
	"github.com/magiconair/properties/assert"
	"log"
	"testing"
)

func TestGetDbConf(t *testing.T) {
	Init("app.ini")
	port := GetConf("influxdb", "PORT")
	assert.Equal(t, port, "3242")

}

func TestGetDbConf1(t *testing.T) {
	Init("app.ini")
	portInt := GetConfReturnInt("influxdb", "PORT")
	log.Println("portInt:=", portInt)
	assert.Equal(t, portInt, int64(3242))

}

