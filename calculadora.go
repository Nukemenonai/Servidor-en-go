package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func calculadora() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()

  operacion:= scanner.Text()
  fmt.Println(operacion)

  valores := strings.Split(operacion, "+")
  fmt.Println(valores)
  fmt.Println(valores[0] + valores[1])

  //aplicar casting
  op_1, err := strconv.Atoi(valores[0])
  if err != nil {
    fmt.Println(err)
  }
  op_2, err2 := strconv.Atoi(valores[1])
  if err2 != nil {
    fmt.Println(err2)
  }


  fmt.Println(op_1 + op_2)
}
