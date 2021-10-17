package store

import "errors"

var prestamo_a int32 = 300
var prestamo_b int32 = 500
var prestamo_c int32 = 700

func (repo Repository) CreditAssignment(investment int32) (int32, int32, int32, error) {
	var a, b, c int32 = 0, 0, 0
	if investment < prestamo_a {
		return 0, 0, 0, errors.New("la inversion no puede ser menor al monto de prestamo.")
	}
	return calcular(investment, a, b, c)
}

func calcular(monto int32, a, b, c int32) (int32, int32, int32, error) {
	if monto == prestamo_a {
		return a + 1, b, c, nil
	} else if monto == prestamo_b {
		return a, b + 1, c, nil
	} else if monto == prestamo_c {
		return a, b, c + 1, nil
	} else if monto == (prestamo_c + prestamo_a + prestamo_b) {
		return a + 1, b + 1, c + 1, nil
	} else if int(monto%prestamo_a) == 0 && int((monto-prestamo_a)%prestamo_a) == 0 {
		return calcular(monto-prestamo_a, a+1, b, c)
	} else if int(monto%prestamo_b) == 0 && int((monto-prestamo_b)%prestamo_b) == 0 {
		return calcular(monto-prestamo_b, a, b+1, c)
	} else if int(monto%prestamo_c) == 0 && int((monto-prestamo_c)%prestamo_c) == 0 {
		return calcular(monto-prestamo_c, a, b, c+1)
	} else if int((monto-prestamo_c)%prestamo_b) == 0 {
		return calcular(monto-prestamo_c, a, b, c+1)
	} else if int((monto-prestamo_a)%prestamo_b) == 0 {
		return calcular(monto-prestamo_a, a+1, b, c)
	} else if int((monto-prestamo_a*2)%prestamo_b) == 0 {
		return calcular(monto-prestamo_a*2, a+2, b, c)
	} else if int((monto-prestamo_a*3)%prestamo_b) == 0 {
		return calcular(monto-prestamo_a*3, a+3, b, c)
	}
	return 0, 0, 0, errors.New("no cudara los datos")
}
