package main

import (
	"bufio"
	"fmt"
	"os"
)

type Nombre struct {
	Fname string // Nombre
	Lname string // Apellido
}

func main() {
	fmt.Print("Introduce el nombre del archivo de texto: ")
	var filename string
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()


	var nombres []Nombre
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		linea := scanner.Text()
		var nombre, apellido string
		_, err := fmt.Sscanf(linea, "%s %s", &nombre, &apellido)
		if err != nil {
			fmt.Println("Error al leer la línea:", err)
			continue
		}
		nombres = append(nombres, Nombre{Fname: nombre, Lname: apellido})
	}


	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	for _, n := range nombres {
		fmt.Printf("Nombre: %s, Apellido: %s\n", n.Fname, n.Lname)
	}
}
