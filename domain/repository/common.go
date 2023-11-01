package repository

type Repo interface {
	Init() error
	Search(name string) (string, error)
	Add(name string, pass string) error
}

func GetRepoInterface(repoName string) Repo {
	switch repoName {
	case "MySQL":
		return &MySQL{}
	case "AssiociativeArray":
		return &AssiociativeArray{}
	default:
		return nil
	}
}
