package install

import (
	"errors"
	"github.com/kardianos/service"
)

var (
	mains     func()
	logger    service.Logger
	s         service.Service
	serConfig *service.Config
	pro       *program
)

type program struct{}
type install struct {
}

/*
	名称，静止名字，说明
*/
func New(name, displayname, description string) (*install, error) {
	serConfig = &service.Config{
		Name:        name,
		DisplayName: displayname,
		Description: description,
	}
	pro = &program{}
	var err error
	s, err = service.New(pro, serConfig)
	if err != nil {
		return nil, errors.New("错误")
	}
	p := install{}
	return &p, nil
}

func (p *install) Install() {
	s.Install()
}
func (p *install) Unistall() {
	s.Uninstall()
}
func (p *install) Run(main func()) {
	mains = main
	s.Run()
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	mains()
}

func (p *program) Stop(s service.Service) error {
	return nil
}
