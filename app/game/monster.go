package game

import (
	"fmt"
	"time"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/model"
)

func spawnMonster(id uint, monsterID uint, regionID uint, p model.RangePosition) model.MonsterSpawn {
	spawn := model.MonsterSpawn {
		ID: id,
		MonsterID: monsterID,
		RegionID: regionID,
	}

	repository.SpawnMonster(spawn)

	return spawn
}

func manageMonster(id uint) {
	for {
		time.Sleep(10000 * time.Millisecond)

		monster := repository.FindSpawnMonster(id)
		fmt.Println(monster.ID)
	}
}
