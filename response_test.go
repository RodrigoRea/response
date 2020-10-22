package response

import (
	"reflect"
	"testing"
)

func TestWebJson(t *testing.T) {
	// Criando uma estrutura qualquer
	var Exemplo struct {
		Saida string
	}

	// Adicionando um valor qualquer
	Exemplo.Saida = "Exemplo de Saída"

	exit := Exit{}
	exitTeste := exit.webJson(Exemplo)

	// verificando se o retorno é uma estrutura
	val := reflect.ValueOf(exitTeste.Dados)
	if val.Kind() != reflect.Struct {
		//structType := val.Type()
		//tableName := structType.Name()
		//if tableName != "Exemplo" {
		t.Errorf("\nErro val.Kind() != reflect.Struct: %+v\n", exitTeste.Dados)
		//}
		//t.Errorf("\nErro em tableName: %+v\n", val.Field(1))
	}

}
