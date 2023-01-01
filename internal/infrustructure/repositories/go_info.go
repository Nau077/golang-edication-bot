package repositories

import (
	"context"
	"golang-edication-bot/internal/infrustructure/libs/db"
	"golang-edication-bot/internal/services/events/telegram/models"
	"io/ioutil"
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
)

const (
	GO_INFO = "go_info"
)

type InfoData interface {
	GetData(ctx context.Context, name string) (*models.GoInfo, error)
	GetDataListType(ctx context.Context, dataType string) ([]*models.GoInfo, error)
}

type goInfoRepo struct {
	staticPath string
}

type goInfoPgRepo struct {
	staticPath string
	client     db.Client
}

func NewGoInfoRepo(staticPath string, db db.Client) *goInfoRepo {
	return &goInfoRepo{
		staticPath: staticPath,
	}
}

func NewGoInfoPgRepo(staticPath string, db db.Client) *goInfoPgRepo {
	return &goInfoPgRepo{
		staticPath: staticPath,
		client:     db,
	}
}

func (g *goInfoRepo) GetData(_ context.Context, fileName string) (*models.GoInfo, error) {
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

	var goInfo = new(models.GoInfo)
	goInfo.Text = string(data)
	// err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, err
	}

	return goInfo, nil
}

func (g *goInfoPgRepo) GetData(ctx context.Context, name string) (*models.GoInfo, error) {
	builder := sq.Select("id, title, text, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(GO_INFO).
		Where(sq.Eq{"title": name}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "GetGoInfo",
		QueryRaw: query,
	}
	var goInfo = new(models.GoInfo)

	err = g.client.DB().GetContext(ctx, goInfo, q, args...)
	if err != nil {
		return nil, err
	}

	return goInfo, nil
}

func (g *goInfoPgRepo) GetDataListType(ctx context.Context, dataType string) ([]*models.GoInfo, error) {
	builder := sq.Select("id, title, text, text_type, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(GO_INFO).
		Where(sq.Eq{"text_type": dataType})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "GetGoInfoList",
		QueryRaw: query,
	}
	var records []*models.GoInfo

	err = g.client.DB().SelectContext(ctx, &records, q, args...)
	if err != nil {
		return nil, err
	}

	return records, nil
}
