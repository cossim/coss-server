package model

import (
	"encoding/json"
	"fmt"
	"github.com/cossim/coss-server/internal/live/api/dto"
)

type RoomType string

const (
	GroupRoomType = "group"
	UserRoomType  = "user"
)

type Room struct {
	Room string `json:"room"`
}

func (r *Room) ToJSONString() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type RoomInfo struct {
	Room            string                        `json:"room"`
	Type            RoomType                      `json:"type"`
	SenderID        string                        `json:"sender_id"`
	NumParticipants uint32                        `json:"num_participants"`
	MaxParticipants uint32                        `json:"max_participants"`
	Participants    map[string]*ActiveParticipant `json:"participants"`
	Option          dto.CallOption                `json:"option"`
}

func (r *RoomInfo) ToJSONString() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type UserRoomInfo struct {
	Room            string                        `json:"room"`
	SenderID        string                        `json:"sender_id"`
	RecipientID     string                        `json:"recipient_id"`
	NumParticipants uint32                        `json:"num_participants"`
	MaxParticipants uint32                        `json:"max_participants"`
	Participants    map[string]*ActiveParticipant `json:"participants"`
}

func (r *UserRoomInfo) ToMap() (map[string]string, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}

	resultMap := make(map[string]string)
	for key, value := range result {
		if key == "participants" {
			participantsJSON, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			resultMap[key] = string(participantsJSON)
		} else {
			resultMap[key] = fmt.Sprintf("%v", value)
		}
	}

	return resultMap, nil
}

func (r *UserRoomInfo) ToInterface() (map[string]interface{}, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRoomInfo) ToJSONString() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *UserRoomInfo) FromMap(data interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, r)
}

type GroupRoomInfo struct {
	Room            string                        `json:"room"`
	GroupID         uint32                        `json:"group_id"`
	SenderID        string                        `json:"sender_id"`
	Participants    map[string]*ActiveParticipant `json:"participants"`
	NumParticipants uint32                        `json:"num_participants"`
	MaxParticipants uint32                        `json:"max_participants"`
}

type ActiveParticipant struct {
	Connected bool `json:"connected"`
}

func (r *GroupRoomInfo) ToMap() (map[string]string, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}

	resultMap := make(map[string]string)
	for key, value := range result {
		if key == "participants" {
			participantsJSON, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			resultMap[key] = string(participantsJSON)
		} else {
			resultMap[key] = fmt.Sprintf("%v", value)
		}
	}

	return resultMap, nil
}

func (r *GroupRoomInfo) ToInterface() (map[string]interface{}, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *GroupRoomInfo) ToJSONString() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *GroupRoomInfo) FromMap(data interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, r)
}
