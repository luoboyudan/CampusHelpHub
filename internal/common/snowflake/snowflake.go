package snowflake

import (
	"campushelphub/internal/config"
	"log"

	"github.com/bwmarrin/snowflake"
)

type IDgenarator interface {
	GenerateID() uint64
}

type SnowflakeIDGenerator struct {
	node *snowflake.Node
}

func NewSnowflakeIDGenerator(config *config.Config) IDgenarator {
	machineID := int64(config.MachineID)
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		log.Fatalf("Failed to create snowflake node: %v", err)
	}
	return &SnowflakeIDGenerator{node: node}
}

// 生成雪花ID
func (s *SnowflakeIDGenerator) GenerateID() uint64 {
	return uint64(s.node.Generate().Int64())
}
