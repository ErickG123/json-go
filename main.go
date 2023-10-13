package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	// Utilizando Tags (anotações)
	// O Unmarshal olha para as Tags
	// Se eu colocar um "-" na tag ele vai omitir a informação
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  1000,
	}

	// Transformando uma Struct em JSON
	// O .Marshal transforma a Struc em JSON
	// Uso o Marshal para salvar o valor
	result, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	// Eu converto para String, pois ele é retornado
	// em bytes
	println(string(result))

	// Criando um Encoder
	// os.Stdout, sai pela saida padrão (terminal)
	// Uso o Encoder quando eu quero entregar o valor
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	// Pegando o JSON e transformando em uma Struct
	jsonPuro := []byte(`{"n": 2, "s": 200}`)
	var contaX Conta
	// Como a Struct está em branco, eu tenho que acessar
	// o seu endereço na memória
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}

	println(contaX.Saldo)
}
