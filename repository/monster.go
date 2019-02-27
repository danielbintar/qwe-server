package repository

import (
	"strconv"
	"encoding/json"

	"github.com/danielbintar/qwe-server/config"
	monster_config "github.com/danielbintar/qwe-server/config/monster"
	"github.com/danielbintar/qwe-server/model"
)

func spawnMonsterKey() string {
	return "spawn_monster"
}

func SpawnMonster(monster model.MonsterSpawn) {
	monsterJson, _ := json.Marshal(monster)
	err := config.RedisInstance().HSet(spawnMonsterKey(), strconv.FormatUint(uint64(monster.ID), 10), monsterJson).Err()
	if err != nil { panic(err) }
}

func FindSpawnMonster(id uint) *model.MonsterSpawn {
	var monster model.MonsterSpawn

	r, err := config.RedisInstance().HGet(spawnMonsterKey(), strconv.FormatUint(uint64(id), 10)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		} else {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(r), &monster)
	}

	return &monster
}

func FindMonster(id uint) *model.Monster {
	for _, monster := range monster_config.Instance().Monsters {
		if monster.ID == id {
			return monster
		}
	}

	return nil
}
