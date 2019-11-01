package main

import (
	"github.com/castrob/valaritas"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/**
 * A função main é responsável por inicializar o servidor e criar o listener para conexões na porta especificada
 * O Pacote labstack/echo é uma biblioteca/framework para facilitar funções de REST para a API.
 * Já está implementada para usar um JWT Token para autenticação (gerar um no jwt.io com a chave especificada)
 * a forma de declarar cada endpoit é o caminho e qual a função que irá lidar com essas conexões.
 */
func main() {
	e := echo.New()
	apiGroup := e.Group("/api")
	apiGroup.Use(middleware.JWT([]byte("b3@W4ry")))

	//Endpoits para cada função
	apiGroup.GET("/", valaritas.Root)
	apiGroup.GET("/:collection/_search", valaritas.Retrieve)
	apiGroup.POST("/:collection/_create", valaritas.Create)
	apiGroup.PUT("/:collection/_update", valaritas.Update)
	apiGroup.DELETE("/:collection/_delete", valaritas.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
