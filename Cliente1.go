package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type InfoPersona struct {
	Nombre string
	Email  string
	Canal  string
}

func cliente(persona InfoPersona) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	/*Solicitar datos de usuario*/
	leer := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu Nombre: ")
	name, _ := leer.ReadString('\n')
	fmt.Println("Ingresa tu e-mail: ")
	mail, _ := leer.ReadString('\n')
	fmt.Println("Ingresa el canal 1 o 2: ")
	chanel, _ := leer.ReadString('\n')

	persona := InfoPersona{
		Nombre: "Nombre: " + name,
		Email:  "e-mail: " + mail,
		Canal:  "Canal: " + chanel,
	}
	go cliente(persona)
	var input string
	fmt.Scanln(&input)
}
