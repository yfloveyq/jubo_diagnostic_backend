package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Method int

const (
	GET Method = iota + 1
	PUT
	POST
	DELETE
)

type RestfulController struct {
	r *gin.Engine
}

func NewRestfulController() *RestfulController {

	r := gin.Default()
	r.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		}
	}(),
	)
	return &RestfulController{
		r,
	}
}

func (controller *RestfulController) Start() {
	controller.r.Run() // listen and serve on 0.0.0.0:8080
}

func (controller *RestfulController) Bind(method Method, relativePath string, handler ...gin.HandlerFunc) gin.IRoutes {
	switch method {
	case GET:
		return controller.r.GET(relativePath, handler...)
	case PUT:
		return controller.r.PUT(relativePath, handler...)
	case POST:
		return controller.r.POST(relativePath, handler...)
	case DELETE:
		return controller.r.DELETE(relativePath, handler...)
	default:
		panic("Invalid HTTP method:" + strconv.Itoa(int(method)))
	}
}
