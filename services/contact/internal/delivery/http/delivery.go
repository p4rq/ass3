package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"architecture_go/services/contact/internal/useCase"
)

func init() {
	viper.SetDefault("HTTP_PORT", 80)
	viper.SetDefault("HTTP_HOST", "127.0.0.1")
	viper.SetDefault("IS_PRODUCTION", "false")
}

type Delivery struct {
	ucContact useCase.Contact
	ucGroup   useCase.Group
	router    *gin.Engine

	options Options
}

type Options struct{}

func New(ucContact useCase.Contact, ucGroup useCase.Group, options Options) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(options)

	d.router = d.initRouter()
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf("%s:%d", viper.GetString("HTTP_HOST"), uint16(viper.GetUint("HTTP_PORT"))))
}

func checkAuth(c *gin.Context) {
	c.Next()
}
