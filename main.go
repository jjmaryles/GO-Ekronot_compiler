package main

import (
	"GO-Ekronot_compiler/jack-compiler/comp-engine"
	"fmt"
	"strings"
	"sync"
)

func main() {
	// process each file as separate goroutine
	wg := sync.WaitGroup{}
	defer wg.Wait()
	files := Targil3Files()
	for _, f := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			comp_engine.Compile(f, strings.Replace(f, ".jack", ".vm", 1))
		}(f)
	}

	fmt.Println("done")
}

func Targil5Files() []string {
	var files []string
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Average\\Main.jack")

	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\ComplexArrays\\Main.jack")

	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\ConvertToBin\\Main.jack")

	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Pong\\Ball.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Pong\\Bat.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Pong\\Main.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Pong\\PongGame.jack")

	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Seven\\Main.jack")

	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Square\\Main.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Square\\Square.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\nand2tetris\\projects\\11\\Square\\SquareGame.jack")

	return files
}

func Targil3Files() []string {
	var files []string

	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\Snake\\Snake.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\Snake\\SnakeGame.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\Snake\\SnakePart.jack")
	files = append(files, "C:\\Users\\jjmar\\OneDrive\\Desktop\\Snake\\Main.jack")

	return files
}
