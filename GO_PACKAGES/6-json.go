package main

// json no contexto do go é um pacote que consegue serializar dados e desserializar dados em json

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonString := `{"name":"Diego", "age":22}`
	var person Person
	// O método json.Unmarshal desserializa (converte) uma string JSON em uma estrutura Go.
	// Ele pega os dados JSON e preenche os campos da estrutura correspondente.
	json.Unmarshal([]byte(jsonString), &person)
	fmt.Printf("Name: %s, Idade: %d\n", person.Name, person.Age)

	person.Name = "Mayara"
	person.Age = 23

	// O método json.Marshal serializa (converte) uma estrutura Go em uma string JSON.
	// Ele pega os dados da estrutura e os transforma em um formato JSON.
	jsonData, _ := json.Marshal(person)

	fmt.Println(string(jsonData))
}
