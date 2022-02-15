package main

// + Инициализировать проект, создаёт файл go.mod: go mod init homework-01
// + git checkout branchname
// Выкачать зависимости для проекта: go mod tidy
//
// "Hello 🗺️!"
// // Далее Commit
// Push to server

// git config --global user.name "Ta4hfc5"
// git config --global user.email johndoe@example.com

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() { // Имя функции с Большой буквы, чтобы была видна в других пакетах
	emoji.Sprint("Hello :world_map:!") // ответ функции
	return
}

// fmt.Println(emoji.Sprint("Hello :world_map: !")) // Вывод в строку терминала

// git config --list
// Comment: Commit HW-00 Push-08
