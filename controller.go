package valaritas

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"sync"

	"github.com/castrob/valaritas/utils"
	"github.com/labstack/echo"
)

var (
	metadata = &utils.META{
		Collections: make(map[string][]string),
	}
	lockedResources = &LOCK{}
	mux sync.Mutex
)

/**
 * Root para acessos em /api/
 */
func Root(ctx echo.Context) error {
	// Exemplo de como usar o metadata (podemos mudar isso dps)
	log.Printf(" Antes de atualizar: %+v", metadata)
	metadata.Collections["users"] = []string{"a", "b", "c"}
	metadata.LastUpdateDate = time.Now()
	log.Printf(" Depois de atualizar: %+v", metadata)
	log.Printf("%+v", lockedResources)
	return ctx.JSON(http.StatusOK, "Authors: Felipe Megale\nGuilherme Galvão\nJoão Castro\n Natália Miranda\nPUC Minas 2019")
}

/**
 * Tratar os inserts em uma collection
 */
func Create(ctx echo.Context) error {
	// lock arquivo de dados
	mux.Lock()
	var request = echo.Map{}
	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
	}
	fmt.Println(request)

	if request["collection"] != nil {
		collection := fmt.Sprintf("%v", request["collection"])

		// todos os campos que estao chegando
		fields := []string{}
		for key := range request {
			if key != "collection" {
				fields = append(fields, key)
			}
		}

		if metadata.FindMetadataByName(collection) {
			// verificar todos os campos que existem na collection
			// e inserir os novos
			for field := range fields {
				if metadata.NotContainsFieldInCollection(fields[field], collection) {
					metadata.Collections[collection] = append(metadata.Collections[collection], fields[field])
				}
			}
		} else {
			// inserir nova chave com seus valores
			metadata.Collections[collection] = fields
		}
		log.Printf(" Depois de criar/inserir: %+v", metadata)
	}
	// fmt.Println(paramName)

	// unlock arquivo de dados
	mux.Unlock()
	return ctx.JSON(http.StatusOK, fmt.Sprintf("Collection %s created successfully!"))
}

/**
 * Tratar os buscas em uma collection
 */
func Retrieve(ctx echo.Context) error {

	mux.Lock()
	//var request = echo.Map{}
	var paramName = ctx.ParamValues()[0]
	fmt.Println("values %+v", paramName)
	if metadata.FindMetadataByName(paramName) {// encontrou os dados
		//log.Printf("Collections paranName %+v", metadata.Collections[paramName])
		//return metadata.Collections["user"]
	}else{ // caso nao encontrou os dados criar um novo
		//log.Printf("Collections paramName %+v", metadata.Collections[paramName])
		Create(ctx)
	}

	mux.Unlock()
	return ctx.JSON(http.StatusOK, "Search Working")
}

/**
 * Tratar os updates em uma collection
 */
func Update(ctx echo.Context) error {

	mux.Lock()
	var paramName = ctx.ParamValues()[0]
	fmt.Println(paramName)

	mux.Unlock()
	return ctx.JSON(http.StatusOK, "Update Working")
}

/**
 * Tratar os deletes em uma collection
 */
func Delete(ctx echo.Context) error {
	mux.Lock()
	var paramName = ctx.ParamValues()[0]
	fmt.Println(paramName)
	mux.Unlock()
	return ctx.JSON(http.StatusOK, "Delete Working")
}
