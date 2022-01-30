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

func generatedID() int64 {
	return snowflk.Generate().Int64()
}
