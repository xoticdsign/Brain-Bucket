package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Date    time.Time
	Content string
}

func main() {
	// var notes []Note
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nBrainBucket v1.0.0 : Made by @xoti$\n")
	fmt.Printf("\nОписание:\nПривет! Меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое :)\n")

	for {
		fmt.Printf("\nМеню:")
		fmt.Printf("\n1. Доступные команды")
		fmt.Printf("\n2. Описание программы")
		fmt.Printf("\n3. Закрыть программу")

		fmt.Printf("\n...\nВыбери действие: ")
		if !scanner.Scan() {
			fmt.Printf("\nОшибка ввода в меню :( - %v\n", scanner.Err())
			return
		}
		menu_action := strings.TrimSpace(scanner.Text())

		switch menu_action {

		// Доступные команды

		case "1":
			for {
				fmt.Printf("\n1. Добавить заметку")
				fmt.Printf("\n2. Удалить заметку")
				fmt.Printf("\n3. Редактировать заметку")
				fmt.Printf("\n4. Показать все заметки")
				fmt.Printf("\n5. Вернуться в меню\n")

				fmt.Printf("\n...\nВыбери действие: ")
				if !scanner.Scan() {
					fmt.Printf("\nОшибка ввода в субменю :( - %v\n", scanner.Err())
					return
				}
				submenu_action := strings.TrimSpace(scanner.Text())

				switch submenu_action {

				// Добавить заметку

				case "1":
					fmt.Printf("\n...\nВведи текст заметки: ")
					if !scanner.Scan() {
						fmt.Printf("\nОшибка ввода заметки :( - %v\n", scanner.Err())
						return
					}
					// note_content := scanner.Text()

				// Удалить заметку

				case "2":

				// Редактировать заметку

				case "3":

				// Показать все заметки

				case "4":

				// Вернуться в меню

				case "5":
					break // Need to be changed

				// Неправильный ввод пользователя в субменю

				default:
					fmt.Printf("\nТакого действия с заметками не существует...\n")
				}
			}

		// Описание программы

		case "2":
			fmt.Printf("\nBrainBucket v1.0.0 : Made by @xoti$\n")
			fmt.Printf("\nОписание:\nПривет! Меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое :)\n")

		// Закрыть программу

		case "3":
			fmt.Printf("\nБыло приятно повидаться!\n")
			return

		// Неправильный ввод пользователя в меню

		default:
			fmt.Printf("\nТы выбрал несуществующее действие, попробуй еще раз :O\n")
		}
	}
}

func NewNote(notes []Note, note_content *string, writing_note chan []Note) {}
