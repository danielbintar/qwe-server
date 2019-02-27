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
		Position: model.Position {
			X: p.MinX,
			Y: p.MinY,
		},
	}

	repository.SpawnMonster(spawn)

	return spawn
}

func manageMonster(id uint) {
	for {
		time.Sleep(10000 * time.Millisecond)

		monster := repository.FindSpawnMonster(id)
		fmt.Println(monster.Position.X)
		fmt.Println(monster.Position.Y)
	}
}
