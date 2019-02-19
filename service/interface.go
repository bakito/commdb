package service

import (
	"github.com/bakito/commdb/types"
)

type CommandService interface {
	GetAll(query string, page int, pageSize int) ([]types.Command, error)
	GetByID(id int64) (*types.Command, bool)
	DeleteByID(id int64) error
	Create(cmd *types.Command) (int64, error)
	Update(cmd *types.Command) error
}
