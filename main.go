package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// type Tarea struct{
	// 	Id int;
	// 	Tarea string;
	// 	Estado string
	// }

	toDoListFile := "Todo.csv"

	verificadorArchivo(toDoListFile)

	menu := "Todo List \n-------------- \n1.Mostar tareas.\n2.Agregar tarea.\n3.Borrar tarea.\n4.Cambiar estado de tarea.(Completa/Incompleta)\n5.Cambiar tarea."

	fmt.Println(menu)

	var seleccion string

	fmt.Println("Ingresar opcion seleccionada: ")
	fmt.Scanln(&seleccion)

	switch seleccion {
	case "1":
		leerYMostarCSV(toDoListFile)
	case "2":
		agregarTarea(toDoListFile)
	case "3":
		var indexLineaABorrar string
		leerYMostarCSV(toDoListFile)
		fmt.Println("Ingresar el numero de la linea a borrar:")
		fmt.Scanln(&indexLineaABorrar)
		borrarTarea(toDoListFile, indexLineaABorrar, 0)
		leerYMostarCSV(toDoListFile)
	case "4":
		var indexEstadoAActualizar string
		leerYMostarCSV(toDoListFile)
		fmt.Println("Ingresar el numero de linea a cambiar de estado:")
		fmt.Scanln(&indexEstadoAActualizar)
		actualizarEstado(toDoListFile, indexEstadoAActualizar, 0)
		leerYMostarCSV(toDoListFile)


	}
}
func actualizarEstado(nombreArchivo, valor string, indiceTareas int) {
	toDo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo.", err)
	}
	defer toDo.Close()
	lector := csv.NewReader(toDo)
	lineas, err := lector.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo.", err)
	}
	var lineasFiltradas [][]string
	// completa := "Completa"
	// incompleta := "Incompleta"
	for _, linea := range lineas {
		if indiceTareas >= len(linea) {
			fmt.Println("Indice fuera de rango.")
		}
		if linea[indiceTareas] != valor {
			lineasFiltradas = append(lineasFiltradas, linea)
		}
		if linea[indiceTareas] == valor {
			if lineas[indiceTareas][2] == "Completa" {
				linea[2] = "Incompleta"
				fmt.Println(linea)
				lineasFiltradas = append(lineasFiltradas, linea)
			} else {
				linea[2] = "Completa"
				lineasFiltradas = append(lineasFiltradas, linea)

			}
		}
	}
	toDo, err = os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error creando el archivo: ", err)
	}
	defer toDo.Close()
	escritor := csv.NewWriter(toDo)
	defer escritor.Flush()

	for _, linea := range lineasFiltradas {
		if err := escritor.Write(linea); err != nil {
			fmt.Println("Error escribiendo una fila: ", err)
		}
	}
}

func borrarTarea(nombreArchivo, valor string, indiceTareas int) {
	toDo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
	}
	defer toDo.Close()

	lector := csv.NewReader(toDo)
	lineas, err := lector.ReadAll()
	if err != nil {
		fmt.Println("Error leyendo el archivo: ", err)
	}
	var lineasFiltradas [][]string
	for i, linea := range lineas {
		if indiceTareas >= len(linea) {
			fmt.Println("Indice fuera de rango: ", i)
		}
		if linea[indiceTareas] != valor {
			lineasFiltradas = append(lineasFiltradas, linea)
		}
	}
	toDo, err = os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error creando el archivo: ", err)
	}
	defer toDo.Close()
	escritor := csv.NewWriter(toDo)
	defer escritor.Flush()

	for _, linea := range lineasFiltradas {
		if err := escritor.Write(linea); err != nil {
			fmt.Println("Error escribiendo una fila: ", err)
		}
	}

}

func agregarTarea(nombreArchivo string) {
	toDo, err := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Error al abrir archivo CSV.", err)
		return
	}
	defer toDo.Close()

	escritor := csv.NewWriter(toDo)
	defer escritor.Flush()

	nuevaTarea := []string{"X", "", ""}

	var nombreTareaNueva string
	var estadoTareaNueva string
	fmt.Println("Ingrese nombre de tarea nueva: ")
	fmt.Scanln(&nombreTareaNueva)
	nuevaTarea[1] = nombreTareaNueva
	fmt.Println("Ingrese estado de tarea nueva: ")
	fmt.Scanln(&estadoTareaNueva)
	nuevaTarea[2] = estadoTareaNueva

	if err := escritor.Write(nuevaTarea); err != nil {
		fmt.Println("Error al escribir el archivo.", err)
		return
	}

}

func verificadorArchivo(nombreArchivo string) {
	if _, err := os.Stat(nombreArchivo); err == nil {
		fmt.Println("El archivo ya existe, no se creara uno nuevo.")
	} else if os.IsNotExist(err) {
		fmt.Println("El archivo no existe , se creara uno nuevo.")
		crearCSV(nombreArchivo)
	}
}

func crearCSV(nombreArchivo string) {
	toDo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear archivo CSV.", err)
		return
	}
	defer toDo.Close()

	escritor := csv.NewWriter(toDo)
	defer escritor.Flush()

	tareas := [][]string{
		{"1", "Tarea 1", "Completa"},
		{"2", "Tarea 2", "Incompleta"},
		{"3", "Tarea 3", "Completa"},
	}
	for _, linea := range tareas {
		if err := escritor.Write(linea); err != nil {
			fmt.Println("Error al escribir el archivo.", err)
			return
		}
	}

	fmt.Println("Archivo CSV creado:", nombreArchivo)
}

func leerYMostarCSV(nombreArchivo string) {
	toDo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println("Error al abrir archivo CSV.", err)
		return
	}
	defer toDo.Close()

	lector := csv.NewReader(toDo)

	lineas, err := lector.ReadAll()
	if err != nil {
		fmt.Println("Error leyendo archivo CSV.", err)
		return
	}
	fmt.Println("Contenido archivo CSV:")
	for _, linea := range lineas {
		fmt.Println(linea)
	}
}
