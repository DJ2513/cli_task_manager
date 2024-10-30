package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		return
	}
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	err = storage.Save(todos)
	if err != nil {
		return
	}
}
