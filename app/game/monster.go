package game

import (
	// "fmt"
	// "time"
	"math/rand"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/model"
)

func spawnMonster(id uint, monsterID uint, regionID uint, bound model.RangePosition) model.MonsterSpawn {
	var x uint
	var y uint
	var p model.Position

	for {
		x = uint(rand.Intn(int(bound.MaxX - bound.MinX))) + bound.MinX
		y = uint(rand.Intn(int(bound.MaxY - bound.MinY))) + bound.MinY
		p = model.Position {
			X: x,
			Y: y,
		}

		occupy := repository.GetRegionOccupy(regionID, p)
		if occupy == nil {
			repository.SetRegionOccupy(regionID, p, id)
			break
		}
	}

	spawn := model.MonsterSpawn {
		ID: id,
		MonsterID: monsterID,
		RegionID: regionID,
		Position: p,
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
