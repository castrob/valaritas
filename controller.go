package valaritas

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	// Collections é responsável por tratar de forma concorrente quais as collections existentes
	collections = sync.Map{}
	// Channel de Commandos a serem executados pelas rotinas workers
	jobs = make(chan COMMAND)
	// Channel de Result no qual cada Rotina emite seu resultado
	results = make (chan string)
)

// Funcao Worker responsável por rodar em go routines usando um channel de jobs e um channel de results
func worker(id string, jobs <- chan COMMAND, result <- chan string) {
	// Itera por cada Job e o Executa
	for job := range jobs {
		fmt.Println("Worker", id, "started job", job.Command, "Collection", job.Collection)
		switch job.Command {
		case "CREATE": // Create é o mais simples, é tratado no próprio metodo
			fileWrite(job.Collection, job.Data)
			results <- job.CommandID
		case "UPDATE": // Update primeiro tenta encontrar o documento, quando encontra atualiza seu valor inteiro
			resultMessage := " Document Not Found"
			content, err := fileRead(job.Collection)
			if err != nil {
				resultMessage =  "Error: " + err.Error()
			} else {
				// Removendo os {} para conseguir dar o find
				searchValue := strings.ReplaceAll(string(job.Search), "}", "")
				searchValue = strings.ReplaceAll(searchValue, "{", "")
				for key, _ := range content {
					// Caso seja igual, atualiza com o novo data
					if strings.Contains(content[key], searchValue) {
						content[key] = string(job.Data)
						fmt.Println(content)
						updatedCollection, err := json.Marshal(content)
						if err != nil {
							resultMessage =  "Error: " + err.Error()
						}else{
							updatedCollectionString := string(updatedCollection)
							fileWrite(job.Collection, []byte(updatedCollectionString))
							resultMessage = " Document Updated!"
							break
						}
					}
				}
			}
			// Coloca no channel o resultado desse comando
			results <- "Operation " + job.CommandID + " " + resultMessage
		case "DELETE": // Delete realiza uma busca para encontrar o item e em seguida o remove da collection
			resultMessage := " Document Not Found"
			content, err := fileRead(job.Collection)
			if err != nil {
				resultMessage =  "Error: " + err.Error()
			} else {
				searchValue := strings.ReplaceAll(string(job.Search), "}", "")
				searchValue = strings.ReplaceAll(searchValue, "{", "")
				for key, _ := range content {
					fmt.Println(content[key], " ", searchValue)
					if strings.Contains(content[key], searchValue) {
						delete(content, key)
						fmt.Println(content)
						updatedCollection, err := json.Marshal(content)
						if err != nil {
							resultMessage =  "Error: " + err.Error()
						}else{
							updatedCollectionString := string(updatedCollection)
							fileWrite(job.Collection, []byte(updatedCollectionString))
							resultMessage = " Document Deleted!"
							break
						}
					}
				}
			}
			results <- "Operation " + job.CommandID + " " + resultMessage
		case "RETRIEVE": // Retrieve itera sobre todos os documentos da collection e em seguida o retorna
			resultMessage := " Document Not Found"
			content, err := fileRead(job.Collection)
			if err != nil {
				resultMessage = "Error: " + err.Error()
			} else {
				searchValue := strings.ReplaceAll(string(job.Search), "}", "")
				searchValue = strings.ReplaceAll(searchValue, "{", "")
				for key, value := range content {
					fmt.Println(content[key], " ", searchValue)
					if strings.Contains(content[key], searchValue) {
						resultMessage = value
						break
					}
				}
			}
			results <- "Operation " + job.CommandID + " " + resultMessage
		default:
			fmt.Println("Default Switch")
		}
	}
}

// Funcao responsavel por escrita ao arquivo
// Parametros: Collection String -> Qual o Nome da Collection e consequentemente do arquivo
//				data [] byte -> Qual o dado a ser escrito no arquivo
// Essa funcao utiliza a biblioteca Log e a funcao log.Output que é responsável pela concorrencia na escrita e leitura
func fileWrite(collection string, data []byte){
	f, err := os.OpenFile(collection + ".valaritasdb", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(f, "", 0)
	logger.Output(2, string(data))
}

// Funcao responsavel por leitura ao arquivo
// Parametros: Collection String -> Qual o Nome da Collection e consequentemente do arquivo
// Retorno: Collection inteira
func fileRead(collection string) (map[string]string, error){
	// Open our jsonFile
	jsonFile, err := os.OpenFile(collection + ".valaritasdb", os.O_RDONLY, 0644)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Open: ",err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]string
	err = json.Unmarshal([]byte(byteValue), &result)
	return result, err
}

/**
 * Root para acessos em /api/
 */
func Root(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Authors: Felipe Megale\nGuilherme Galvão\nJoão Castro\n Natália Miranda\nPUC Minas 2019")
}

/**
 * Tratar os inserts em uma collection
 */
func Create(ctx echo.Context) error {
	// lock arquivo de dados
	var request = echo.Map{}
	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
	}
	collection := fmt.Sprintf("%v", request["collection"])
	delete(request, "collection")
	requestData, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	requestJsonString := string(requestData)
	var command COMMAND
	// check if collection already exist or not
	if _, ok := collections.Load(collection); ok {
		result, err := fileRead(collection)
		if err != nil {
			fmt.Println(err)
		}
		documentId := guuid.New().String()
		result[documentId] = requestJsonString
		newDocumentBytes, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		newDocumentString := string(newDocumentBytes)
		fmt.Printf("New Document on Collection %s", newDocumentString)
		command = COMMAND{
			CommandID:  guuid.New().String(),
			Command:    "CREATE",
			Collection: collection,
			Search:     nil,
			Data:       []byte(newDocumentString),
		}
	} else {
		collections.Store(collection, true)
		newDocuments := make(map[string]string)
		documentId := guuid.New().String()
		newDocuments[documentId] = requestJsonString
		newDocumentBytes, err := json.Marshal(newDocuments)
		if err != nil {
			fmt.Println(err)
		}
		newDocumentString := string(newDocumentBytes)
		fmt.Printf("New Collection %s", newDocumentString)
		command = COMMAND{
			CommandID:  documentId,
			Command:    "CREATE",
			Collection: collection,
			Search:     nil,
			Data:       []byte(newDocumentString),
		}
	}
	go worker(command.CommandID, jobs, results)
	jobs <- command
	var result string
	for result = <-results ;result != command.CommandID; {
		continue
	}
	return ctx.JSON(http.StatusOK, fmt.Sprintf("Create Success!"))
}

/**
 * Tratar os buscas em uma collection
 */
func Retrieve(ctx echo.Context) error {
	var request = echo.Map{}
	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
	}
	collection := fmt.Sprintf("%v", request["collection"])
	delete(request, "collection")
	requestData, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	requestJsonString := string(requestData)
	var command COMMAND
	if _, ok := collections.Load(collection); ok {
		command = COMMAND{
			CommandID:  guuid.New().String(),
			Command:    "RETRIEVE",
			Collection: collection,
			Search:     []byte(requestJsonString),
			Data:       nil,
		}
	} else {
		return ctx.JSON(http.StatusBadRequest, "Collection not found!")
	}
	go worker(command.CommandID, jobs, results)
	jobs <- command
	var result string
	for result = <- results ; !strings.Contains(result, command.CommandID); {
		continue
	}
	return ctx.JSON(http.StatusOK, result)
}

/**
 * Tratar os updates em uma collection
 */
func Update(ctx echo.Context) error {
	var request = echo.Map{}
	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
	}
	collection := fmt.Sprintf("%v", request["collection"])
	delete(request, "collection")
	searchValue := request["search"]
	dataValue := request["data"]

	searchValueBytes, err := json.Marshal(searchValue)
	if err != nil {
		fmt.Println(err)
	}
	dataValueBytes, err := json.Marshal(dataValue)
	if err != nil {
		fmt.Println(err)
	}

	searchValueString := string(searchValueBytes)
	dataValueString := string(dataValueBytes)

	var command COMMAND
	if _, ok := collections.Load(collection); ok {
		command = COMMAND{
			CommandID:  guuid.New().String(),
			Command:    "UPDATE",
			Collection: collection,
			Search:     []byte(searchValueString),
			Data:       []byte(dataValueString),
		}
	} else {
		return ctx.JSON(http.StatusBadRequest, "Collection not found!")
	}
	go worker(command.CommandID, jobs, results)
	jobs <- command
	var result string
	for result = <- results ; !strings.Contains(result, command.CommandID); {
		fmt.Println(result, " ", command.CommandID)
		continue
	}
	return ctx.JSON(http.StatusOK, result)
}

/**
 * Tratar os deletes em uma collection
 */
func Delete(ctx echo.Context) error {
	var request = echo.Map{}
	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
	}
	collection := fmt.Sprintf("%v", request["collection"])
	delete(request, "collection")
	requestData, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	requestJsonString := string(requestData)
	var command COMMAND
	if _, ok := collections.Load(collection); ok {
		command = COMMAND{
			CommandID:  guuid.New().String(),
			Command:    "DELETE",
			Collection: collection,
			Search:     []byte(requestJsonString),
			Data:       nil,
		}
	} else {
		return ctx.JSON(http.StatusBadRequest, "Collection not found!")
	}
	go worker(command.CommandID, jobs, results)
	jobs <- command
	var result string
	for result = <- results ; !strings.Contains(result, command.CommandID); {
		fmt.Println(result, " ", command.CommandID)
		continue
	}
	return ctx.JSON(http.StatusOK, result)
}
