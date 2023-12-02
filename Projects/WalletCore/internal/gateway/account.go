package gateway

import "github.com.br/Nesbox314/fullcycle-course-notes-and-projects/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
