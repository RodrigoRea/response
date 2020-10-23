# response

install:

go get github.com/RodrigoRea/response

# Exmplo de utilização

<pre>
import "github.com/RodrigoRea/response" 

.
.
.
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

<pre>
    {
        "Dados":{},
        "Erro":"",
        "Mensagem":""
    }
</pre>
