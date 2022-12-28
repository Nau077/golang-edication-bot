package repositories

import (
	"io/ioutil"
	"log"
	"os"
)

type InfoData interface {
	GetData(fileName string) ([]byte, error)
}

type goInfoRepo struct {
	staticPath string
}

func NewGoInfoRepo(staticPath string) *goInfoRepo {
	return &goInfoRepo{
		staticPath: staticPath,
	}
}

func (g *goInfoRepo) GetData(fileName string) ([]byte, error) {
	// file, err := ioutil.ReadFile(g.staticPath + "/go-base-data/" + fileName)
	file, err := os.Open(g.staticPath + "/go-base-data/" + fileName)
	if err != nil {
		return nil, err
	}
	// data := models.GoInfoData{}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := ioutil.ReadAll(file)
	// err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
