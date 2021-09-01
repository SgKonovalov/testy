package vari

import (
	"github.com/adjust/redismq"
)

var TestQueue = redismq.CreateQueue("localhost", "6379", "", 9, "clicks")
