package valaritas

import "time"

/**
 * Exemplo de um Comando que deve ser recebido via request.
 * @Commands: Insert, Search, Update, Delete
 * @Search: Em casos de Search, Update e Delete, é utilizado para primeiro buscar o dado
 * @Data: é o campo no qual deve ser utilizado cada chave:valor para inserir na base, em casos de update alterar usando eles
 * @Params: Upsert, ReturnId
 */
type COMMAND struct {
	Command string      `json:"command"`
	Search  interface{} `json:"search,omitempty"`
	Data    interface{} `json:"data"`
	Params  interface{} `json:"params,omitempty"`
}

/**
 * LOCK é responsável por manter as informações de cada collection.document que está com lock
 * LastUpdateDate serve de controle pra quando LOCK foi alterado ou não
 */
type LOCK struct {
	Resources      []interface{} `json:"resources"`
	LastUpdateDate time.Time     `json:"lastUpdateDate"`
}

/**
 * o modelo de cada Documento é esse
 */
type DOCUMENT struct {
	Id             string      `json:"_id"`
	Content        interface{} `json:"content"`
	CreateDate     time.Time   `json:"createDate"`
	LastUpdateDate time.Time   `json:"lastUpdateDate"`
}
