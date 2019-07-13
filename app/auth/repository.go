package auth

type Repository struct {}

func NewRepository() *Repository {
	return new(Repository)
} 