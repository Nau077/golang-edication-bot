package repositories

import (
	"encoding/json"
	"golang-edication-bot/internal/services/events/telegram/models"
	"io/ioutil"
)

type InfoData interface {
	GetData(fileName string) (*models.GoInfoData, error)
}

type goInfoRepo struct {
}

func NewGoInfoRepo() *goInfoRepo {
	return &goInfoRepo{}
}

func (g *goInfoRepo) GetData(fileName string) (*models.GoInfoData, error) {
	file, err := ioutil.ReadFile("./go_base_data/" + fileName)
	if err != nil {
		return nil, err
	}
	data := models.GoInfoData{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
