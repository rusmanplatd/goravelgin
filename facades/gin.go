package facades

import (
	"log"

	gin "github.com/rusmanplatd/goravelgin"

	"github.com/rusmanplatd/goravelframework/contracts/route"
)

func Route(driver string) route.Route {
	instance, err := gin.App.MakeWith(gin.BindingRoute, map[string]any{
		"driver": driver,
	})

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*gin.Route)
}
