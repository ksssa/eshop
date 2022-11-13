package util

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"sync"
	"time"
)

var node *snowflake.Node
var once sync.Once

func nodeID() int64 {
	var nodeIdS int64
	nodeIdS = rand.NewSource(time.Now().UnixNano()).Int63()
	return nodeIdS % 8
}

func GenerateID() int64 {
	once.Do(func() {
		var err error
		nId := nodeID()
		snowflake.NodeBits = 3
		snowflake.StepBits = 12
		node, err = snowflake.NewNode(nId)
		if err != nil {
			panic(err)
		}
	})
	id := node.Generate()
	return id.Int64()
}
