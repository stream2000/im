/*
@Time : 2020/2/19 11:33
@Author : Minus4
*/
package snowflake

import (
	"fmt"
	"testing"
)

var worker *Worker

func init() {
	worker, _ = NewWorker(1)
}

func TestTwo(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id, _ := worker.Generate()
		redisId := worker.RedisScoreMapping(id)
		fmt.Println(redisId)
	}
}
