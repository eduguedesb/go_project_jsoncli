package main

import (
    "flag"
    "fmt"
    "project_jsoncli/jsonparser"
    "log"
    "os"
)

func main() {
    // Definindo os comandos da CLI
    readCmd := flag.NewFlagSet("read", flag.ExitOnError)
    filterCmd := flag.NewFlagSet("filter", flag.ExitOnError)
    formatCmd := flag.NewFlagSet("format", flag.ExitOnError)
    writeCmd := flag.NewFlagSet("write", flag.ExitOnError)

    // Flags para cada comando
    filterKey := filterCmd.String("key", "", "Chave para filtrar")
    filterValue := filterCmd.String("value", "", "Valor para filtrar")
    outputFile := writeCmd.String("out", "output.json", "Arquivo de saída")

    if len(os.Args) < 2 {
        fmt.Println("Subcomando é necessário: read, filter, format ou write")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "read":
        readCmd.Parse(os.Args[2:])
        jsonparser.ReadJSON()
    case "filter":
        filterCmd.Parse(os.Args[2:])
        if *filterKey == "" || *filterValue == "" {
            fmt.Println("A chave e o valor para filtrar são obrigatórios")
            os.Exit(1)
        }
        jsonparser.FilterJSON(*filterKey, *filterValue)
    case "format":
        formatCmd.Parse(os.Args[2:])
        jsonparser.FormatJSON()
    case "write":
        writeCmd.Parse(os.Args[2:])
        jsonparser.WriteJSON(*outputFile)
    default:
        fmt.Println("Comando desconhecido")
        os.Exit(1)
    }
}
