package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) agregarALista(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) eliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) imprimirLista() {
	for _, tarea := range tl.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}

func (tl *taskList) imprimirListaCompletadas() {
	for _, tarea := range tl.tasks {
		if tarea.completado == true {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompletada() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go en esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python en esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de Node",
		descripcion: "Completar mi curso de Node en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarALista(t3)
	lista.eliminarDeLista(1)
	lista.agregarALista(t2)

	lista.imprimirLista()

	t1.marcarCompletada()
	fmt.Println("------>> COMPLETADAS ------------------->>")
	lista.imprimirListaCompletadas()

}
