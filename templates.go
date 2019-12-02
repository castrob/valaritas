package valaritas

/**
 * Exemplo de um Comando que deve ser recebido via request.
 * @Commands: Insert, Search, Update, Delete
 * @Search: Em casos de Search, Update e Delete, é utilizado para primeiro buscar o dado
 * @Data: é o campo no qual deve ser utilizado cada chave:valor para inserir na base, em casos de update alterar usando eles
 */
type COMMAND struct {
	CommandID string	 	`json:"_id"`
	Command string      `json:"command"`
	Collection string	`json:"collection"`
	Search  []byte `json:"search,omitempty"`
	Data    []byte `json:"data"`
}