package service

import (
	"github.com/bakito/commdb/types"
	"github.com/go-xorm/xorm"
)

func NewCommandService(orm *xorm.Engine) CommandService {
	return &commandService{
		orm: orm,
	}
}

type commandService struct {
	orm *xorm.Engine
}

func (s *commandService) GetAll(query string, page int, pageSize int) ([]types.Command, error) {

	commands := []types.Command{}
	var err error

	if query != "" {
		err = s.orm.Where("command like ?", "%"+query+"%").Or("keywords like ?", "%"+query+"%").Limit(pageSize, page).Find(&commands)
	} else {
		err = s.orm.Limit(pageSize, page).Find(&commands)
	}

	return commands, err
}

func (s *commandService) GetByID(id int64) (*types.Command, bool) {
	command := &types.Command{ID: id}
	if ok, _ := s.orm.Get(command); ok {
		return command, true
	}
	return nil, false
}

func (s *commandService) DeleteByID(id int64) error {
	command := &types.Command{}
	_, err := s.orm.ID(id).Delete(command)

	return err
}

func (s *commandService) Create(cmd *types.Command) (int64, error) {
	return s.orm.Insert(cmd)
}

func (s *commandService) Update(cmd *types.Command) error {
	_, err := s.orm.ID(cmd.ID).Update(cmd)
	return err
}
