package repository

import (
	"strconv"

	"github.com/danielbintar/qwe-server/config"
)

func loginCharacterKey() string {
	return "character_login"
}

func IsLoginCharacter(id uint) bool {
	_, err := config.RedisInstance().HGet(loginCharacterKey(), strconv.FormatUint(uint64(id), 10)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return false
		} else {
			panic(err)
		}
	}

	return true
}

func SetLoginCharacter(id uint) {
	err := config.RedisInstance().HSet(loginCharacterKey(), strconv.FormatUint(uint64(id), 10), "true").Err()
	if err != nil { panic(err) }
}

func UnsetLoginCharacter(id uint) {
	err := config.RedisInstance().HDel(loginCharacterKey(), strconv.FormatUint(uint64(id), 10)).Err()
	if err != nil { panic(err) }
}
