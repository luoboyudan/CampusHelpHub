package cache

import (
	"campushelphub/internal/config"
)

type KeyBuilder struct {
	Prefix string
}

func NewKeyBuilder(config *config.Config) *KeyBuilder {
	return &KeyBuilder{Prefix: config.Redis.Prefix}
}

func (k *KeyBuilder) BuildKey(key string) string {
	return k.Prefix + ":" + key
}
