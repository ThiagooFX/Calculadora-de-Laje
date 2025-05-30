package calculos

// calcular espessura da laje
func Espessura(lx, ly, cobrimento, engaste float64) float64 {
	var l float64

	if ly*0.7 < lx {
		l = ly * 0.7
		return l
	} else {
		l = lx
		return l
	}

	d := 2.5 - (0.1 * engaste) * l

	H := d + 0.5 + cobrimento

	return H
}

func Lambda(lx, ly float64) float64 {
	return lx / ly
}

func PesoProprio(gama, espessura float64) float64 { //Peso Próprio
	return gama * espessura
}

func CdTotalLaje(pplaje, ppcontrapiso, ppcemamica, sobrecarga float64) float64 {
	Pp := pplaje + ppcontrapiso + ppcemamica
	result := 1.4*Pp + 1.5*sobrecarga

	return result
}
