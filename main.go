package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tarea struct {
	ID          int
	Descripcion string
	Estado   bool
}

func main() {
	tareas, err := cargarTareas()
	if err != nil {
		fmt.Println("Error al cargar tareas: ", err)
		return
	}

	lector := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n-----ToDo List-----\n1. Lista de tareas.\n2. Agregar tarea.\n3. Completar tarea.\n4. Salir.")
		fmt.Print("Eliga una opcion: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			mostrarTareas(tareas)
		case 2:
			fmt.Print("Ingrese la nueva tarea: ")
			descripcion, _ := lector.ReadString('\n')
			descripcion = strings.TrimSpace(descripcion)
			tareas = agregarTarea(tareas, descripcion)
		case 3:
			fmt.Print("Seleccione el ID de la tarea a completar: ")
			var id int
			fmt.Scanln(&id)
			tareas = completarTarea(tareas, id)
		case 4:
			fmt.Println("Adios!")
			return
		default:
			fmt.Println("Opcion invalida,seleccione otra.")
		}
	}
}


const nombreArchivoToDo = "todolist.csv"

func cargarTareas() ([]Tarea, error) {
	var tareas []Tarea

	archivo, err := os.Open(nombreArchivoToDo)
	if err != nil {
		if os.IsNotExist(err) {
			return tareas, nil
		}
		return nil, err
	}
	defer archivo.Close()

	lector := csv.NewReader(archivo)
	datos, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, dato := range datos {
		id, _ := strconv.Atoi(dato[0])
		estado, _ := strconv.ParseBool(dato[2])
		tareas = append(tareas, Tarea{ID: id, Descripcion: dato[1], Estado: estado})
	}

	return tareas, nil
}

func guardarTareas(tareas []Tarea) error {
	archivo, err := os.Create(nombreArchivoToDo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	escritor := csv.NewWriter(archivo)
	defer escritor.Flush()

	for _, tarea := range tareas {
		escritor.Write([]string{
			strconv.Itoa(tarea.ID),
			tarea.Descripcion,
			strconv.FormatBool(tarea.Estado),
		})
	}

	return nil
}

func agregarTarea(tareas []Tarea, descripcion string) []Tarea {
	nuevaTarea := Tarea{
		ID:          len(tareas) + 1,
		Descripcion: descripcion,
		Estado:   false,
	}
	tareas = append(tareas, nuevaTarea)
	guardarTareas(tareas) 
	return tareas
}

func mostrarTareas(tareas []Tarea) {
	fmt.Println("\nLista de tareas:")
	for _, tarea := range tareas {
		estado := "❌"
		if tarea.Estado {
			estado = "✅"
		}
		fmt.Printf("[%d] %s - %s\n", tarea.ID, tarea.Descripcion, estado)
	}
}

func completarTarea(tareas []Tarea, id int) []Tarea {
	for i, tarea := range tareas {
		if tarea.ID == id {
			tareas[i].Estado = true
			break
		}
	}
	guardarTareas(tareas)
	return tareas
}
