package calculos

func Lambda(x, y float64) float64 {
    return x / y
}

func PesoProprio(gama, espessura float64) float64{	//Peso Próprio
	return gama * espessura
}

func CdTotalLaje (pplaje, ppcontrapiso, ppcemamica, sobrecarga float64) float64 {
	Pp := pplaje + ppcontrapiso + ppcemamica
	result := 1.4 * Pp + 1.5 * sobrecarga

	return result
}
