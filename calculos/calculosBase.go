package calculos

import "math"

// calcular espessura da laje
func Espessura(lx, ly, cobrimento, engaste float64) float64 {
	l := lx

	if ly*0.7 < lx {
		l = ly * 0.7
	}

	d := 2.5 - (0.1*engaste)*l

	h := d + 0.5 + cobrimento

	return h
}

func Lambda(lx, ly float64) float64 {
	lambda := ly / lx
	passo := 0.05
	result := math.Ceil(lambda/passo) * passo
	return result
}

func PesoProprio(carga, espessura float64) float64 { //Peso Próprio
	return carga * espessura
}

func CdTotalLaje(pplaje, ppcontrapiso, ppcemamica, sobrecarga float64) float64 {
	Pp := pplaje + ppcontrapiso + ppcemamica
	result := 1.4*Pp + 1.5*sobrecarga

	return result
}
