package main
import (
	"fmt"
	"src/github.com/jlvmexicogit/asatidwh/variables"
)

func main() {
	estado,texto := variables.ConviertoaTexto(1223)
	fmt.Println("estado= ",estado)
	fmt.Println("estado= ",texto)
}