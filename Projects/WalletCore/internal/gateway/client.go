package gateway

import "github.com.br/Nesbox314/fullcycle-course-notes-and-projects/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
