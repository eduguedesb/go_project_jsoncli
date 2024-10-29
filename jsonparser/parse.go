package jsonparser

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

// Função para ler JSON
func ReadJSON() {
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatalf("Erro ao ler o arquivo JSON: %s", err)
    }
    fmt.Println(string(data))
}

// Função para filtrar JSON por chave e valor
func FilterJSON(key, value string) {
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatalf("Erro ao ler o arquivo JSON: %s", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        log.Fatalf("Erro ao parsear JSON: %s", err)
    }

    var filteredData []map[string]interface{}
    for _, item := range jsonData {
        if item[key] == value {
            filteredData = append(filteredData, item)
        }
    }

    result, err := json.MarshalIndent(filteredData, "", "  ")
    if err != nil {
        log.Fatalf("Erro ao formatar JSON: %s", err)
    }
    fmt.Println(string(result))
}

// Função para formatar JSON com indentação
func FormatJSON() {
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatalf("Erro ao ler o arquivo JSON: %s", err)
    }

    var jsonData interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        log.Fatalf("Erro ao parsear JSON: %s", err)
    }

    result, err := json.MarshalIndent(jsonData, "", "  ")
    if err != nil {
        log.Fatalf("Erro ao formatar JSON: %s", err)
    }
    fmt.Println(string(result))
}

// Função para escrever JSON em um novo arquivo
func WriteJSON(filename string) {
    data, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatalf("Erro ao ler o arquivo JSON: %s", err)
    }

    if err := ioutil.WriteFile(filename, data, 0644); err != nil {
        log.Fatalf("Erro ao escrever JSON: %s", err)
    }

    fmt.Printf("Dados salvos no arquivo %s\n", filename)
}
