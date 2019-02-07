package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
	town_config "github.com/danielbintar/qwe-server/config/town"
	"github.com/danielbintar/qwe-server/model"
)

func townUsersKey(id uint) string {
	return "towns:" + strconv.FormatUint(uint64(id), 10) + ":users"
}

func GetTownUsers(id uint) []model.UserPosition {
	var users []model.UserPosition
	r, err := config.RedisInstance().HGetAll(townUsersKey(id)).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			panic(err)
		}
	} else {
		for k, v := range r {
			var user model.UserPosition
			json.Unmarshal([]byte(v), &user)
			u64, _ := strconv.ParseUint(k, 10, 32)
			user.ID = uint(u64)
			users = append(users, user)
		}
	}
	return users
}

func SetTownUser(townID uint, userID uint, x uint, y uint) {
	position := map[string]uint{
		"x": x,
		"y": y,
	}

	positionJson, _ := json.Marshal(position)

	err := config.RedisInstance().HSet(townUsersKey(townID), strconv.FormatUint(uint64(userID), 10), positionJson).Err()
	if err != nil { panic(err) }
}


func FindTown(id uint) *model.Town {
	for _, town := range town_config.Instance().Towns {
		if town.ID == id {
			return town
		}
	}

	return &model.Town{}
}
