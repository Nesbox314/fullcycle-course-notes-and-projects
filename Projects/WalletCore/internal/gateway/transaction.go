package gateway

import "github.com.br/Nesbox314/fullcycle-course-notes-and-projects/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
