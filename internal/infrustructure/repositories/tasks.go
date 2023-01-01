package repositories

import (
	"context"
	"golang-edication-bot/internal/infrustructure/libs/db"
	"golang-edication-bot/internal/services/events/telegram/models"
	"math/rand"
	"time"

	sq "github.com/Masterminds/squirrel"
)

const (
	TASKS = "tasks"
)

type TasksInfo interface {
	GetTask(ctx context.Context) (*models.Task, error)
	GetTasksSolution(ctx context.Context, id int64) (string, error)
}

type goTasksPgRepo struct {
	staticPath string
	client     db.Client
}

func NewTasksPgRepo(staticPath string, db db.Client) *goTasksPgRepo {
	return &goTasksPgRepo{
		staticPath: staticPath,
		client:     db,
	}
}

func (g *goTasksPgRepo) GetTask(ctx context.Context) (*models.Task, error) {
	builder := sq.Select("id, title, tasks_text, tasks_solution, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(TASKS)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "GetGoInfoList",
		QueryRaw: query,
	}
	var records []*models.Task

	err = g.client.DB().SelectContext(ctx, &records, q, args...)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().Unix())

	return records[rand.Intn(len(records))], nil
}

func (g *goTasksPgRepo) GetTasksSolution(ctx context.Context, id int64) (string, error) {
	builder := sq.Select("id, title, tasks_solution").
		PlaceholderFormat(sq.Dollar).
		From(TASKS).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "GetTaskSolution",
		QueryRaw: query,
	}
	var task = new(models.Task)

	err = g.client.DB().GetContext(ctx, task, q, args...)
	if err != nil {
		return "", err
	}

	return task.TasksSolution, nil
}
