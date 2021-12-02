package main

import "fmt"

type taskList struct {
  tasks []*task // slice of pointers to  the task 
}

func (tl *taskList) addToList(t *task){
  tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) removeFromList(index int) {
  tl.tasks = append(tl.tasks[:index], tl.tasks[index + 1:]...)
}

type task struct {
  name string
  description string
  complete bool 
}

func(t *task) markComplete(){
  t.complete = true
}

func (t *task) updateDescrition(descr string) {
  t.description = descr
}

func (t *task) updateName(name string) {
  t.name = name
}



func main() {
  t1 := &task {
    name: "Completar todo list",
    description: "to do list in go",
  }

  t2 := &task {
    name: "Completar todo list",
    description: "to do list in go",  
  }

  t3 := &task {
    name: "complete portfolio",
    description: "complete portfolio",
  }

  list := &taskList {
    tasks: []*task{
      t1, t2,
    },
  }


  //fmt.Println(list.tasks[0])

  // fmt.Printf("%+v\n", t1)
  // t1.markComplete()
  // t1.updateName("holo")
  // fmt.Printf("%+v\n", t1)

  list.addToList(t3)

  fmt.Println(len(list.tasks))

  list.removeFromList(1)

  fmt.Println(len(list.tasks))

}
