package repository

import (
	"encoding/json"
	"strconv"

	"github.com/danielbintar/qwe-server/config"
	town_config "github.com/danielbintar/qwe-server/config/town"
	"github.com/danielbintar/qwe-server/model"
)

func townUsersKey(id int) string {
	return "towns:" + strconv.Itoa(id) + ":users"
}

func GetTownUsers(id int) []model.UserPosition {
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
			user.Id, _ = strconv.Atoi(k)
			users = append(users, user)
		}
	}
	return users
}

func SetTownUser(townId int, userId int, x int, y int) {
	position := map[string]int{
		"x": x,
		"y": y,
	}

	positionJson, _ := json.Marshal(position)

	err := config.RedisInstance().HSet(townUsersKey(townId), strconv.Itoa(userId), positionJson).Err()
	if err != nil { panic(err) }
}


func FindTown(id int) *model.Town {
	for _, town := range town_config.Instance().Towns {
		if town.Id == id {
			return town
		}
	}

	return &model.Town{}
}
