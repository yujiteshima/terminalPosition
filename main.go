package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	options := []string{"選択肢1", "選択肢2", "選択肢3"}

	selectedIndex := 0

	for {
		clearScreen()

		rows, err := getTerminalRows()
		if err != nil {
			fmt.Println("エラー：ターミナルのサイズを取得できませんでした。")
			return
		}

		startIndex := selectedIndex - (rows / 2)
		if startIndex < 0 {
			startIndex = 0
		} else if startIndex > len(options)-rows {
			startIndex = len(options) - rows
		}
		endIndex := startIndex + rows
		if endIndex > len(options) {
			endIndex = len(options)
		}

		for i := startIndex; i < endIndex; i++ {
			prefix := " "
			if i == selectedIndex {
				prefix = ">"
			}
			fmt.Printf("%s %s\n", prefix, options[i])
		}

		var input string
		fmt.Print("選択してください: ")
		fmt.Scanln(&input)

		switch input {
		case "j":
			if selectedIndex < len(options)-1 {
				selectedIndex++
			}
		case "k":
			if selectedIndex > 0 {
				selectedIndex--
			}
		case "":
			fmt.Printf("%sが選択されました\n", options[selectedIndex])
			return
		default:
			fmt.Println("無効な入力です")
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getTerminalRows() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	sizeStr := strings.TrimSpace(string(out))
	sizeArr := strings.Split(sizeStr, " ")
	rows, err := strconv.Atoi(sizeArr[0])
	if err != nil {
		return 0, err
	}

	return rows, nil
}
