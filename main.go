package main

import "fmt"

func main() {
	var num1, num2 float32 
	var op int

	fmt.Printf("Ingrese el primer numero \n")
	fmt.Scanf("%f", &num1)
	
	fmt.Printf("Ingrese su segundo numero \n")
	fmt.Scanf("%f", &num2)
	
	fmt.Printf("Que operacion desea realizar. 1-suma, 2-resta, 3-multiplica, 4-divi \n")
	fmt.Scanf("%d", &op)
	
	result, typeOp := validate(op, num1, num2)

	fmt.Printf("El resultado de la %v es: %.2f \n", typeOp, result)
}

func validate(op int, num1 float32, num2 float32) (float32, string) {

	switch op {
	case 1:
		return sum(num1, num2)

	case 2:
		return rest(num1,num2)

	case 3:
		return multi(num1, num2)

	case 4:
		return divi(num1,num2)

	default : return sum(num1, num2) // tiene que retornar algo si no se cumple alguna opcion
	}
}

func sum(num1 float32, num2 float32) (float32, string) {
	return (num1 + num2 ), "suma"
}

func rest(num1 float32, num2 float32) (float32, string) {
	return (num1 - num2), ("resta")
}

func multi(num1 float32, num2 float32) (float32, string) {
	return (num1 * num2), "multiplicacion"
}

func divi(num1 float32, num2 float32) (float32, string) {
	return (num1 / num2), "division"
}