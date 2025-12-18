package ejercicios
import (
	"strconv"
)
func Ejercicio01(cadena string) (int,string) {
	var texto string
	//entero =strconv.Atoi(cadena)
	var entero int
	entero,err:=strconv.Atoi(cadena)
	if err != nil {
		return 0, "Hubo un error "+err.Error()
	}	
	if entero>100 {
		texto="Es mayor a 100"
	}  else {
		texto="Es menor a 100"
	} 
	return entero,texto
}