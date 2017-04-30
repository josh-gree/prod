package main

import (
	"math/rand"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/josh-gree/comm"
	"github.com/labstack/echo"
)

var j = comm.JobMessage{}
var r = comm.ResMessage{}
var public = false // read from cml

func Prod(data []float64, id int) {
	log.Info("Doing computation")
	sum := 1.0
	for _, d := range data {
		sum *= d
	}
	time.Sleep(time.Duration(rand.Int31n(10000)) * time.Millisecond)

	resmsg := comm.ResMessage{Id: id, Result: sum}

	resmsg.Send("public:8000/res")
}

func main() {
	e := echo.New()

	e.POST("/job", j.Recieve(public, Prod))

	e.Start(":8000")
}
