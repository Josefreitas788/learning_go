package main

import (
    "fmt"
    "linha_de_comando/app"
    "log"
    "os"
)

func main() {
    fmt.Println("Iniciando")
    aplicacao := app.Gerar()
    if err := aplicacao.Run(os.Args); err != nil {
            log.Fatal(err)

    }

}
