package main

// + Инициализировать проект, создаёт файл go.mod: go mod init homework-01
// + git checkout branchname
// Выкачать зависимости для проекта: go mod tidy
//
// "Hello 🗺️!"
// // Далее Commit
// Push to server

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() { // Имя функции с Большой буквы, чтобы была видна в других пакетах
	fmt.Println(emoji.Sprint("Hello :world_map:!"))
}
