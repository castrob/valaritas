package valaritas

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

var (
	metadata = &META{}
	lockedResources = &LOCK{}
)

/**
 * Root para acessos em /api/
 */

func Root(ctx echo.Context) error {
	// Exemplo de como usar o metadata (podemos mudar isso dps)
	userCollection := map[string]interface{}{
		"Name": "user",
		"Fields": []interface{}{
			"nome",
			"sexo",
			"idade",
		},
	}
	log.Printf(" Antes de atualizar: %+v", metadata)
	metadata.Collections = append(metadata.Collections, userCollection)
	metadata.LastUpdateDate = time.Now()
	log.Printf(" Depois de atualizar: %+v", metadata)
	log.Printf("%+v", lockedResources)
	return ctx.JSON(http.StatusOK, "Authors: Felipe Megale\nGuilherme Galvão\nJoão Castro\n Natália Miranda\nPUC Minas 2019")
}


/**
 * Tratar os inserts em uma collection
 */
func Create(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Create Working")
}

/**
 * Tratar os buscas em uma collection
 */
func Retrieve(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Search Working")
}

/**
 * Tratar os updates em uma collection
 */
func Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Update Working")
}

/**
 * Tratar os deletes em uma collection
 */
func Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Delete Working")
}


