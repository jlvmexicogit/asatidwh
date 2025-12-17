package main
import (
	"fmt"
	"runtime"
	"src/github.com/jlvmexicogit/asatidwh/variables"
)

func main() {
	estado,texto := variables.ConviertoaTexto(1223)
	fmt.Println("estado= ",estado)
	fmt.Println("estado= ",texto)
	
	if os := runtime.GOOS; os=="Linux." || os=="OS X." {
		fmt.Println("Esto no es windows",os)
	} else {
		fmt.Println("Esto si es windows",os)
	}
}