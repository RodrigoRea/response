package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

/*

 */
// Exit - Modelo de Resposta
type Exit struct {
	Dados    interface{}
	Mensagem string
	Erro     bool
}

/*

 */
// Response Interface
type Response interface {
	webJson(res ...interface{}) Exit
}

/*

 */
// webJson - Método de criação de retorno WEB
func (exit Exit) webJson(res ...interface{}) Exit {

	var erro error
	if len(res) > 0 {
		for _, t := range res {

			switch t.(type) { /*Inicio 1*/
			case nil:
				{
					continue
				}
			case error:
				erro = t.(error)
				if erro != nil {
					exit.Mensagem = string(erro.Error())
					break
				}
			default:

				switch reflect.TypeOf(t).Kind() { /*Inicio 2*/
				case reflect.Slice, reflect.Map:
					s := reflect.ValueOf(t)
					if s.Len() > 0 {
						exit.Dados = t
					} else {
						erro = errors.New("Nenhum registro encontrado")
						exit.Mensagem = string(erro.Error())
					}
				case reflect.Struct:
					s := int(reflect.TypeOf(t).Size())
					if s > 0 {
						exit.Dados = t
					} else {
						erro = errors.New("Nenhum registro encontrado")
						exit.Mensagem = string(erro.Error())
					}

				} /*Fim 2*/

			} /*Fim 1*/
		}
	}

	// Substituição da mensagem caso a mensagem seja recebida na função
	if len(res) > 0 {
		for _, t := range res {
			switch t.(type) {
			case string:
				exit.Mensagem = t.(string)
			}
		}
	}

	// Substituição da mensagem caso exista erro
	if len(res) > 0 {
		for _, t := range res {
			switch t.(type) {
			case error:
				erro = t.(error)
				if erro != nil {
					exit.Mensagem = string(erro.Error())
					break
				}
			}
		}
	}

	return Exit{
		Dados:    exit.Dados,
		Mensagem: exit.Mensagem,
		Erro:     (erro != nil),
	}
}

/*

 */
// Create - Cria a saida Json com base nos valores recebidos na função
//
// Ex Create( http.ResponseWriter, struct, string, err )
func Create(w http.ResponseWriter, params ...interface{}) {
	exit := Exit{}
	res := exit.webJson(params...)
	json.NewEncoder(w).Encode(res)
}

/*

 */
// JSON - Cria a saida Json com base nos valores recebidos na função
//
// Ex Create( http.ResponseWriter, struct, string, err )
func JSON(w http.ResponseWriter, params ...interface{}) {
	Create(w, params...)
}
