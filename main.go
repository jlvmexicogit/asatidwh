package main
import (
	"fmt"
	"runtime"
	"src/github.com/jlvmexicogit/asatidwh/variables"
	"src/github.com/jlvmexicogit/asatidwh/ejercicios"
)

func main() {
	estado,texto := variables.ConviertoaTexto(1223)
	fmt.Println("estado= ",estado)
	fmt.Println("estado= ",texto)
	
	if os := runtime.GOOS; os=="linux" || os=="OS X." {
		fmt.Println("Esto no es windows",os)
	} else {
		fmt.Println("Esto si es windows",os)
	}
	
	switch os:=runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux",os)
	case "debian":
		fmt.Println("Esto es linux",os)
	default:
		fmt.Printf("%s \n",os)
	}

	entero,texto := ejercicios.Ejercicio01("191")
	fmt.Println("entero= ",entero)
	fmt.Println("texto= ",texto)
}