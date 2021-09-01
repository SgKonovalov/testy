package vari

import (
	"github.com/adjust/redismq"
)

var TestQueue = redismq.CreateQueue("127.0.0.1:", "6379", "", 9, "clicks")
