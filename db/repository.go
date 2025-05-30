package db

import (
    "encoding/json"
    "os"
    "fmt"
)
type Dados struct {
    Lambda   float64 `json:"lambda"`
    MuX      float64 `json:"mu_x"`
    MuXLinha float64 `json:"mu'_x"`
    MuY      float64 `json:"mu_y"`
    MuYLinha float64 `json:"mu'_y"`
}

type Tabela struct {
    Descricao string              `json:"descricao"`
    Dados     map[string][]Dados  `json:"dados"`
}

type Database struct {
    Tabelas map[string]Tabela `json:"tabelas"`
}

var dbInstance *Database

func LoadDatabase() error {
    file, err := os.Open("db/db.json")
    if err != nil {
        return fmt.Errorf("erro ao abrir arquivo: %v", err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    dbInstance = &Database{}
    if err := decoder.Decode(dbInstance); err != nil {
        return fmt.Errorf("erro ao decodificar JSON: %v", err)
    }

    return nil
}

func GetLajeParams(tipoLaje string, lambda float64) (Dados, error) {
    var db Database
    file, err := os.Open("db/db.json")
    if err != nil {
        return Dados{}, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&db); err != nil {
        return Dados{}, err
    }

    for _, tabela := range db.Tabelas {
        for tipo, dados := range tabela.Dados {
            if tipo == tipoLaje {
                for _, dado := range dados {
                    if dado.Lambda == lambda {
                        return dado, nil
                    }
                }
            }
        }
    }

    return Dados{}, fmt.Errorf("parâmetros não encontrados para tipoLaje %s e lambda %.2f", tipoLaje, lambda)
}