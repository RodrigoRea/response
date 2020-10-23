# response

install:

go get github.com/RodrigoRea/response

# Descrição 

O Pacote foi criado para facilitar e padronizar o formato de resposta em json, sempre respondendo em um objeto 

<pre>{
    "Dados":{},
    "Erro": false,
    "Mensagem":""
}
</pre> 

Dados: (struct ou map)

Erro: (bool) - (Caso 'Dados' seja vazio, erro irá conter 'true')

é possível enviar o erro para a função e a própria função ira tratar o erro com a mensagem: EX: func respostaIV()

Mensagem: (string)


Utilize a função 'response.JSON(w, ?, ?)'

Podendo ser enviada até 3 parâmetros, 

sendo o 1° obrigatório contendo 'w http.ResponseWriter'

sendo o 2° obrigatório podendo ser error ou struct

sendo o 3° opcional podendo ser error ou struct (diferente do 2° parâmetro)

# Exemplo de utilização

<pre>
import "github.com/RodrigoRea/response" 


func respostaI(w http.ResponseWriter, r *http.Request) {
    data, err := busca()
    response.JSON(w, err, data)
}

func respostaII(w http.ResponseWriter, r *http.Request) {
    data, _ := busca()
    response.JSON(w, data)
}

func respostaIII(w http.ResponseWriter, r *http.Request) {    
    response.JSON(w, "Minha Mensagem")
}

func respostaIV(w http.ResponseWriter, r *http.Request) {  
    err := errors.New("Mensagem de erro")  
    response.JSON(w, err)
}

</pre>

# Exemplo de saída JSON
<pre>
    // respostaI
    {
        "Dados":{
            "Nome":"Rodrigo",
            "Cidade":"São Paulo"
        },
        "Erro": false,
        "Mensagem":""
    }


    // respostaII (Sem erro)
    {
        "Dados":{
            "Nome":"Rodrigo",
            "Cidade":"São Paulo"
        },
        "Erro": false,
        "Mensagem":""
    }

    // respostaII (Com erro)
    {
        "Dados":{},
        "Erro": true,
        "Mensagem":"Nenhum resultado encontrado"
    }


    // respostaIII
    {
        "Dados":{},
        "Erro": false,
        "Mensagem":"Minha Mensagem"
    }

    // respostaIV
    {
        "Dados":{},
        "Erro": true,
        "Mensagem":"Mensagem de erro"
    }
</pre>
