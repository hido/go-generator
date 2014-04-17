package estimate

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.PrintLn("Gaussian")
	}
	filename := os.Args[0]
	switch method := os.Args[1]; method {
	case "Gaussian":
		fmt.PrintLn("Gaussian")
	case "Cumurative":
		fmt.PrintLn("Cumurative")
	case "AR":
		fmt.PrintLn("AR")
	default:
		fmt.PrintLn("NON SUPPORTED METHOD")
	}
}
