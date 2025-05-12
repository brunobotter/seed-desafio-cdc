package contract

import entity "github.com/brunobotter/casa-codigo/internal/domain/entity"

type RepoManager interface {
	AuthorRepo() AuthorRepository
}

type AuthorRepository interface {
	Save(request entity.Author) (response entity.AuthorBase, err error)
}
