package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
	region_config "github.com/danielbintar/qwe-server/config/region"
	"github.com/danielbintar/qwe-server/model"
)

func regionUsersKey(id uint) string {
	return "regions:" + strconv.FormatUint(uint64(id), 10) + ":users"
}

func UnsetRegionCharacterPosition(characterID uint, regionID uint) {
	err := config.RedisInstance().HDel(regionUsersKey(regionID), strconv.FormatUint(uint64(characterID), 10)).Err()
	if err != nil { panic(err) }
}

func SetRegionCharacterPosition(id uint, position *model.CharacterPosition) {
	coordinate := map[string]uint{
		"x": position.X,
		"y": position.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)
	err := config.RedisInstance().HSet(regionUsersKey(id), strconv.FormatUint(uint64(position.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
}

func GetRegionCharactersPosition(id uint) []*model.CharacterPosition {
	var positions []*model.CharacterPosition
	r, err := config.RedisInstance().HGetAll(regionUsersKey(id)).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			panic(err)
		}
	} else {
		for k, v := range r {
			var position *model.CharacterPosition
			json.Unmarshal([]byte(v), &position)
			u64, _ := strconv.ParseUint(k, 10, 32)
			position.ID = uint(u64)
			positions = append(positions, position)
		}
	}
	return positions
}

func getRegionCharacterPosition(id uint, characterID uint) *model.CharacterPosition {
	var position *model.CharacterPosition
	r, err := config.RedisInstance().HGet(regionUsersKey(id), strconv.FormatUint(uint64(characterID), 10)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		} else {
			panic(err)
		}
	} else {
		json.Unmarshal([]byte(r), &position)
		position.ID = characterID
	}
	return position
}

func MovingCharacterInRegion(regionID uint, movement *model.CharacterMovement) *model.CharacterPosition {
	position := getRegionCharacterPosition(regionID, movement.ID)

	if int(position.X) + movement.X > 0 {
		position.X = uint(int(position.X) + movement.X)
	} else {
		position.X = 0
	}

	if int(position.Y) + movement.Y > 0 {
		position.Y = uint(int(position.Y) + movement.Y)
	} else {
		position.Y = 0
	}

	coordinate := map[string]uint{
		"x": position.X,
		"y": position.Y,
	}

	coordinateJson, _ := json.Marshal(coordinate)

	err := config.RedisInstance().HSet(regionUsersKey(regionID), strconv.FormatUint(uint64(position.ID), 10), coordinateJson).Err()
	if err != nil { panic(err) }
	return position
}


func FindRegion(id uint) *model.Region {
	for _, region := range region_config.Instance().Regions {
		if region.ID == id {
			return region
		}
	}

	return &model.Region{}
}
