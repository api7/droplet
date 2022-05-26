package main

import (
	"github.com/gin-gonic/gin"
	"github.com/api7/droplet"
	"github.com/api7/droplet/wrapper"
	ginwrap "github.com/api7/droplet/wrapper/gin"
	"reflect"
)

func main() {
	r := gin.Default()
	r.POST("/json_input/:id", ginwrap.Wraps(JsonInputDo, wrapper.InputType(reflect.TypeOf(&JsonInput{}))))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type JsonInput struct {
	ID    string   `auto_read:"id,path" json:"id"`
	User  string   `auto_read:"user,header" json:"user"`
	IPs   []string `json:"ips"`
	Count int      `json:"count"`
	Body  []byte   `auto_read:"@body"`
}

func JsonInputDo(ctx droplet.Context) (interface{}, error) {
	input := ctx.Input().(*JsonInput)

	return input, nil
}
