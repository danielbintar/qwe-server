package game

import (
	// "fmt"
	// "time"
	"math/rand"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/model"
)

func spawnMonster(id uint, monsterID uint, regionID uint, p model.RangePosition) model.MonsterSpawn {
	x := uint(rand.Intn(int(p.MaxX - p.MinX))) + p.MinX
	y := uint(rand.Intn(int(p.MaxY - p.MinY))) + p.MinY

	spawn := model.MonsterSpawn {
		ID: id,
		MonsterID: monsterID,
		RegionID: regionID,
		Position: model.Position {
			X: x,
			Y: y,
		},
	}

	repository.SpawnMonster(spawn)

	return spawn
}

func manageMonster(id uint) {
	// for {
		// time.Sleep(10000 * time.Millisecond)

		// monster := repository.FindSpawnMonster(id)
		// fmt.Println(monster.Position.X)
		// fmt.Println(monster.Position.Y)
	// }
}
