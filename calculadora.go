package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

type calc struct {

}

//receiver function nospermite hacer calc.operate, hace que calc pueda ser usada por el struct
func (calc) operate(entrada string, operador string) int {
  valores := strings.Split(entrada, operador)
  op_1 := parsear(valores[0])
  op_2 := parsear(valores[1])
  switch operador {
    case "+":
    return op_1 + op_2
  case "-":
    return op_1 - op_2
  case "*":
    return op_1 * op_2
  case "/":
    return op_1 / op_2
  default:
    return 0

  }
}

func parsear(entrada string) int {
  op_1, _ := strconv.Atoi(entrada)
  return op_1
  
}

func leerEntrada() string {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

func main_2() {
  entrada := leerEntrada()
  operador := leerEntrada()

  c := calc{}
  fmt.Println(c.operate(entrada, operador))
}
