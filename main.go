package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"teste-golang/pacotes"
	"teste-golang/utils"
)

var meuArray [5]int = [5]int{1, 2, 3, 4, 5}

type Pessoa struct {
	Id        int        `json:"id,omitempty"`
	Nome      string     `json:"nome"`
	Telefones []Telefone `json:"telefones"`
	Documento `json:"documento"`
}

type Telefone struct {
	Ddd    int `json:"ddd"`
	Numero int `json:"numero"`
}

type Documento struct {
	Cpf string `json:"cpf"`
}

type Gente interface {
	imprimirDados()
}

func (pes *Pessoa) atualizarDados(id int, nome string) {
	pes.Id = id
	pes.Nome = nome
}

func (pes Pessoa) imprimirDados() {
	fmt.Println(pes.Id, pes.Nome)
}

func funcaoQueImplementaInterface(g Gente) {
	g.imprimirDados()
}

func main() {
	valor1 := 10
	valor2 := 43
	resultado := somador(valor1, valor2)
	fmt.Printf("%v, %T\n", resultado, resultado)
	infoMeuPc()

	soma := 0
	for _, numero := range meuArray {
		soma += numero
	}

	fmt.Println(soma)
	exemploSlices()

	//instancia de objeto
	pessoa1 := Pessoa{
		Id:   1,
		Nome: "antigo nome",
		Telefones: []Telefone{
			{Ddd: 11, Numero: 9999999},
			{Ddd: 22, Numero: 123456789}},
		Documento: Documento{Cpf: "12345678901"},
	}

	funcaoQueImplementaInterface(pessoa1)
	pessoa1.atualizarDados(99, "novo nome")
	funcaoQueImplementaInterface(pessoa1)

	//funções anônimas
	func(pes Pessoa) {
		fmt.Println("Id vezes 10 é igual a: ", pes.Id*10)
	}(pessoa1)

	//convertendo pra json
	serializarParaJsonComMarshal(pessoa1)
	serializarParaJsonComEncoder(pessoa1)
	deserializarJsonComUnmarshal()
	deserializarJsonComDecoder()

	pacotes.PrintHello()
	utils.PrintBye()
}

func somador(valor1, valor2 int) int {
	return valor1 + valor2
}

func infoMeuPc() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

func exemploSlices() {
	valores := []int{1, 4, 2, 3, 5}
	corte := valores[0:] //se não delimitar um index final, o corte pega o array todo
	fmt.Println(corte)

	sort.Ints(valores)
	fmt.Println(valores)
}

func serializarParaJsonComMarshal(pessoa Pessoa) {
	objJson, _ := json.Marshal(pessoa)
	fmt.Println(string(objJson))
}

func serializarParaJsonComEncoder(pessoa Pessoa) {
	codificado := json.NewEncoder(os.Stdout)
	codificado.Encode(pessoa)
}

func deserializarJsonComUnmarshal() {
	var dadosJson = []byte(`{"id":100,"nome":"Bruno Noguchi","telefones":[{"ddd":99,"numero":123456789}],"documento":{"cpf":"12365478955"}}`)
	var pessoa Pessoa
	json.Unmarshal(dadosJson, &pessoa)

	fmt.Printf("%+v\n", pessoa)
}

func deserializarJsonComDecoder() {
	var dadosJson = `{"id":100,"nome":"Bruno Noguchi","telefones":[{"ddd":99,"numero":123456789}],"documento":{"cpf":"12365478955"}}`
	var pessoa Pessoa
	json.NewDecoder(strings.NewReader(dadosJson)).Decode(&pessoa)

	fmt.Printf("%+v\n", pessoa)
}
