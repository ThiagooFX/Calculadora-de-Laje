package main

import (
	"calcularLaje/calculos"
	"calcularLaje/db"
	"fmt"
)

func main() {

    var lx, ly, qtdEngaste, cargaCeramica, cargaContrapiso, cargaLaje, cargaSobrecarga, cobrimento float64
    var tipoLaje string

    fmt.Println("Selecione a opção que melhor representa o emprego da laje: ")
    fmt.Println("1 - COBERTURA NÃO EM BALANÇO")
    fmt.Println("2 - LAJES DE PISO NÃO EM BALANÇO")
    fmt.Println("3 - LAJE EM BALANÇO")
    fmt.Println("4 - LAJES QUE SUPORTAM VEÍCULOS DE PESO TOTAL MENOR OU IGUAL A 30KN")
    fmt.Println("5 - LAJES QUE SUPORTAM VEÍCULOS DE PESO TOTAL MAIOR QUE 30KN")

    fmt.Println("Insira o valor de LX: ")
    fmt.Scan(&lx)
    fmt.Println("Insira o valor de LY: ")
    fmt.Scan(&ly)
    fmt.Println("Insira o tipo de Laje: ")
    fmt.Scan(&tipoLaje)
    fmt.Println("Insira a quantidade de apoios engastados: ")
    fmt.Scan(&qtdEngaste)
    fmt.Println("Insira a carga da cerâmica: ")
    fmt.Scan(&cargaCeramica)
    fmt.Println("Insira a carga do contrapiso: ")
    fmt.Scan(&cargaContrapiso)
    fmt.Println("Insira a carga da laje: ")
    fmt.Scan(&cargaLaje)
    fmt.Println("Insira a carga de sobrecarga: ")
    fmt.Scan(&cargaSobrecarga)
    fmt.Println("Insira o cobrimento ")
    fmt.Scan(&cobrimento)


    // Achar a lambida
    lambda := calculos.Lambda(lx, ly)

    // Achar os valore de "u"
    params, err := db.GetLajeParams(tipoLaje, lambda)
    if err != nil {
        fmt.Printf("Erro: %v\n", err)
        return
    }

    // Calcular espessura da laje
    espessuraLaje := calculos.Espessura(lx, ly, cobrimento, qtdEngaste)

    // Calcular peso próprio de cada item que compõe a laje
    ppCeramica := calculos.PesoProprio(cargaCeramica, espessuraLaje)
    ppContrapiso := calculos.PesoProprio(cargaContrapiso, espessuraLaje)
    ppLaje := calculos.PesoProprio(cargaLaje, espessuraLaje)

    // Calcular Cd total da laje
    cdTotalLaje := calculos.CdTotalLaje(ppLaje, ppContrapiso, ppCeramica, cargaSobrecarga)

    // Calcular momento em cada direção:
    momentoUx := calculos.MomentoLajeLambdaMenorQueDois(params.MuX, cdTotalLaje, lx)
    momentoUxLinha := calculos.MomentoLajeLambdaMenorQueDois(params.MuXLinha, cdTotalLaje, lx)
    momentoUy := calculos.MomentoLajeLambdaMenorQueDois(params.MuY, cdTotalLaje, lx)
    momentoUyLinha := calculos.MomentoLajeLambdaMenorQueDois(params.MuYLinha, cdTotalLaje, lx)

    fmt.Println(lambda)
    fmt.Printf("Momento em UX: %f", momentoUx)
    fmt.Printf("Momento em U'X: %f", momentoUxLinha)
    fmt.Printf("Momento em UY: %f", momentoUy)
    fmt.Printf("Momento em U'Y: %f", momentoUyLinha)
    
    
}