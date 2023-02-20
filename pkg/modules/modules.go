package modules

import "gorm.io/gorm"

type Modules struct {
	Author AuthorModule
	Book   BookModule
}

func NewModules(db *gorm.DB) Modules {
	return Modules{
		Author: NewAuthorModule(db),
		Book:   NewBookModule(db),
	}
}
