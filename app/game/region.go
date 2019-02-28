package game

import (
	"github.com/danielbintar/qwe-server/repository"
)

func manageRegion() {
	regions := repository.AllRegion()
	id := uint(1)
	for _, region := range regions {
		for _, regionMonster := range region.Monsters {
			for i := uint(1) ; i <= regionMonster.Total ; i++ {
				spawnMonster(id, regionMonster.ID, region.ID, regionMonster.Position)
				id += 1
			}
		}
	}
}
