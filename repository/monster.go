package repository

import (
	"strconv"
	"encoding/json"

	"github.com/danielbintar/qwe-server/config"
	monster_config "github.com/danielbintar/qwe-server/config/monster"
	"github.com/danielbintar/qwe-server/model"
)

func spawnMonsterKey(regionID uint) string {
	return "spawn_monster:" + strconv.FormatUint(uint64(regionID), 10)
}

func SpawnMonster(monster model.MonsterSpawn) {
	monsterJson, _ := json.Marshal(monster)
	err := config.RedisInstance().HSet(spawnMonsterKey(monster.RegionID), strconv.FormatUint(uint64(monster.ID), 10), monsterJson).Err()
	if err != nil { panic(err) }
}

func AllSpawnMonster(regionID uint) []*model.MonsterSpawn {
	monsters := []*model.MonsterSpawn{}

	r, err := config.RedisInstance().HGetAll(spawnMonsterKey(regionID)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		} else {
			panic(err)
		}
	} else {
		for _, v := range r {
			var monster *model.MonsterSpawn
			json.Unmarshal([]byte(v), &monster)
			monsters = append(monsters, monster)
		}
	}

	return monsters
}

func FindMonster(id uint) *model.Monster {
	for _, monster := range monster_config.Instance().Monsters {
		if monster.ID == id {
			return monster
		}
	}

	return nil
}
