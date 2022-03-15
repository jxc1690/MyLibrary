package install

import (
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
type Install struct {
}

/*
	名称，静止名字，说明
*/
func (p *Install) New(name, displayname, description string) bool {
	serConfig = &service.Config{
		Name:        name,
		DisplayName: displayname,
		Description: description,
	}
	pro = &program{}
	var err error
	s, err = service.New(pro, serConfig)
	if err != nil {
		return false
	}
	return true
}
func (p *Install) Install() {
	s.Install()
}
func (p *Install) Unistall() {
	s.Uninstall()
}
func (p *Install) Run(main func()) {
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
