package repositories

type Container struct {
	GoInfoRepo    InfoData
	GoTasksPgRepo TasksInfo
}

func NewContainer(goInfoRepo InfoData, goTasksPgRepo TasksInfo) *Container {
	return &Container{
		GoInfoRepo:    goInfoRepo,
		GoTasksPgRepo: goTasksPgRepo,
	}
}
