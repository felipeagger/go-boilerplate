package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var snowflk *snowflake.Node

func init() {

	var err error
	snowflk, err = snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}

}

func GeneratedID() int64 {
	return snowflk.Generate().Int64()
}
