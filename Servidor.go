package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strings"
)

func servidor() {
	/* Protocolo y puerto para escuchar la info*/
	s, err := net.Listen("tcp", ":9999")

	/*Validacion del error, si es diferente de nulo,
	muestra el error que hay en el sistema*/
	if err != nil {
		fmt.Println(err)
		return
	}
	/*Ciclo para escuchar a los clientes*/
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		/*Funcion que leera la informacion del cliente*/
		go Cliente(c)
	}
}

/* estructura que toma la informacion de la persona que se suscribe a un canal especifico*/
type InfoPersona struct {
	Nombre string
	Email  string
	Canal  string
}

func Cliente(c net.Conn) {
	var persona InfoPersona
	err := gob.NewDecoder(c).Decode(&persona)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Datos: \n", persona)
		if strings.Contains(persona.Canal, "1") == true {
			fmt.Println("Hola")
		} else {
			fmt.Println("Chao")
		}

	}

}

func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)
}
