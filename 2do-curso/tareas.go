package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *task) marcarCompleto() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de go",
		descripcion: "Completar mi curso de go de platzi en esta semana",
	}
	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de platzi en esta semana",
	}
	t3 := &task{
		nombre:      "Completar mi curso de NodeJS",
		descripcion: "Completar mi curso de NodeJS de platzi en esta semana",
	}
	lista := taskList{
		tasks: []*task{t1, t2},
	}
	fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)
	fmt.Println(len(lista.tasks))
	lista.eliminarDeLista(1)
	fmt.Println(len(lista.tasks))
}
