package calculos

func MomentoLajeBalanco(cd, lx float64) float64 {
	momento := cd * ((lx * lx) / 2)

	return momento
}

func MomentoLajeLambdaMenorQueDois(u, cd, lx float64) ( float64) {
	momentoU := (u * cd * (lx * lx)) / 100

	return momentoU
}
