package main

import (
	"fmt"

	"github.com/macedo/category_service-go/config"
	"github.com/macedo/category_service-go/models"
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

	p.Task("default", D{"hello"})

	p.Task("hello", func(c *Context) {
		name := c.Args.ZeroString("name", "n")
		if name == "" {
			Bash("echo Hello $USER!")
		} else {
			fmt.Println("Hello", name)
		}
	})

	p.Task("db_create", func(c *Context) {
		Bash("createdb -U category_service-go category_service-go")
	})

	p.Task("db_migrate", func(c *Context) {
		db := config.PostgresSession()
		db.CreateTable(&models.Category{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Category{})
	})

	p.Task("db_setup", D{"db_create", "db_migrate"})
}

func main() {
	Godo(tasks)
}
