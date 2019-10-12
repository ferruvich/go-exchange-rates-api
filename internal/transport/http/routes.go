package http

import (
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// docs is generated by Swag CLI, wee need to import it.
	_ "github.com/ferruvich/go-exchange-rates-api/docs"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

// Routes returns our application routes
func Routes(
	address string,
	s service.Servicer,
) *gin.Engine {
	router := gin.Default()

	// For documentation purposes
	url := ginSwagger.URL(strings.Join([]string{address, "/swagger/doc.json"}, "/"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/rates", GetRatesHandler(s))
	router.GET("/value/:currency", GetEURValue(s))
	router.GET("/recommendation/:currency", Recommend(s))

	return router
}
