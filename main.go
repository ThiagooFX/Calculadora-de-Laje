package main

import (
	"calcularLaje/db"
	"fmt"
)

func main() {
    tipoLaje := "2B"
    lambda := 1.25
    
    params, err := db.GetLajeParams(tipoLaje, lambda)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    fmt.Printf("Para laje tipo %s com λ=%.2f:\n", tipoLaje, lambda)
    fmt.Printf("μx: %.2f\n", params.MuX)
    fmt.Printf("μ'x: %.2f\n", params.MuXLinha)
    fmt.Printf("μy: %.2f\n", params.MuY)
    fmt.Printf("μ'y: %.2f\n", params.MuYLinha)
}