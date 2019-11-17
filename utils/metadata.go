package utils

import (
	"time"
)

/**
 * META é responsável por manter as informações de cada collection e seus campos
 * LastUpdateDate serve de controle pra quando META foi alterado ou não.
 */
type META struct {
	Collections    map[string][]string `json:"collections"`
	LastUpdateDate time.Time           `json:"lastUpdateDate"`
}

func (meta *META) FindMetadataByName(paramName string) (result bool) {
	result = false

	for i := 0; i < len(meta.Collections); i++ {
		if meta.Collections[paramName] != nil {
			result = true
		}
	}
	return
}

func (meta *META) NotContainsFieldInCollection(fieldName string, collection string) (result bool) {
	result = false

	for field := range meta.Collections[collection] {
		if fieldName == meta.Collections[collection][field] {
			result = true
		}
	}

	return
}
