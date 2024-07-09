package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Note struct {
	Date    time.Time
	Content string
}

var notes_db []Note

func main() {
	fmt.Printf("\nBrainBucket CLI v1.0.0 : Made by @xoti$\n")
	fmt.Printf("\nBrainBucket --> Привет, меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое!\n")

	input_scanner := bufio.NewScanner(os.Stdin)

menu:
	for {
		fmt.Printf("\nМеню:")
		fmt.Printf("\n1. Открыть меню заметок")
		fmt.Printf("\n2. Закрыть меня")
		fmt.Printf("\n3. Credits")

		fmt.Printf("\n...\nЧто ты хочешь сделать: ")
		if !input_scanner.Scan() {
			fmt.Printf("\nBrainBucket --> Ошибка 100: не переживай, это я виноват... --> BrainBucket Status: :(\n")
			fmt.Printf("\nError description: %v\n", input_scanner.Err())
			return
		}
		menu_action := strings.TrimSpace(input_scanner.Text())

		switch menu_action {
		case "1":
			for {
				fmt.Printf("\nМеню заметок:")
				fmt.Printf("\n1. Добавить новую заметку")
				fmt.Printf("\n2. Попросить меня забыть заметку")
				fmt.Printf("\n3. Отредактировать заметку")
				fmt.Printf("\n4. Показать все заметки")
				fmt.Printf("\n5. Вернуться в главное меню")

				fmt.Printf("\n...\nЧто ты хочешь сделать: ")
				if !input_scanner.Scan() {
					fmt.Printf("\nBrainBucket --> Ошибка 200: не переживай, это я виноват... --> BrainBucket Status: :(\n")
					fmt.Printf("\nError description: %v\n", input_scanner.Err())
					return
				}
				submenu_action := strings.TrimSpace(input_scanner.Text())

				switch submenu_action {
				case "1":
					fmt.Printf("\nВведи текст новой заметки: ")
					if !input_scanner.Scan() {
						fmt.Printf("\nBrainBucket --> Ошибка 201: не переживай, это я виноват... --> BrainBucket Status: :(\n")
						fmt.Printf("\nError description: %v\n", input_scanner.Err())
						return
					}
					note_new_content := input_scanner.Text()

					new_channel := make(chan bool)
					go NewNote(&note_new_content, new_channel)
					<-new_channel

				case "2":
					if notes_db == nil {
						fmt.Printf("\nBrainBucket --> На данный момент у тебя нет заметок... Давай это исправим!\n")
					} else {
						fmt.Printf("\nFAQ: как удалить заметку?\n")
						fmt.Printf("\nBrainBucket --> Выбери, какая из заметок тебе больше нужна, запомни, какая она по счету. Потом введи ее номер здесь!\n")
						fmt.Printf("\nТвои заметки: \n%v\n", notes_db)

						fmt.Printf("\nВведи номер заметки, которую ты хочешь удалить: ")
						if !input_scanner.Scan() {
							fmt.Printf("\nBrainBucket --> Ошибка 202: не переживай, это я виноват... --> BrainBucket Status: :(\n")
							fmt.Printf("\nError description: %v\n", input_scanner.Err())
							return
						}
						note_delete_index := input_scanner.Text()

						delete_channel := make(chan bool)
						go RemoveNote(&note_delete_index, delete_channel)
						<-delete_channel
					}

				case "3":
					if notes_db == nil {
						fmt.Printf("\nBrainBucket --> На данный момент у тебя нет заметок... Давай это исправим!\n")
					} else {
						fmt.Printf("\nFAQ: как изменить заметку?\n")
						fmt.Printf("\nBrainBucket --> Выбери, какую из заметок ты хочешь изменить, запомни, какая она по счету. Потом введи ее номер здесь!\n")
						fmt.Printf("\nBrainBucket --> Твои заметки: \n%v\n", notes_db)

						fmt.Printf("\nВведи номер заметки, которую ты хочешь изменить: ")
						if !input_scanner.Scan() {
							fmt.Printf("\nBrainBucket --> Ошибка 2030: не переживай, это я виноват... --> BrainBucket Status: :(\n")
							fmt.Printf("\nError description: %v\n", input_scanner.Err())
							return
						}
						note_edit_index := input_scanner.Text()

						fmt.Printf("\nВведи новый текст для заметки: ")
						if !input_scanner.Scan() {
							fmt.Printf("\nBrainBucket --> Ошибка 2031: не переживай, это я виноват... --> BrainBucket Status: :(\n")
							fmt.Printf("\nError description: %v\n", input_scanner.Err())
							return
						}
						note_edit_content := input_scanner.Text()

						edit_channel := make(chan bool)
						go EditNote(&note_edit_index, &note_edit_content, edit_channel)
						<-edit_channel
					}

				case "4":
					if notes_db == nil {
						fmt.Printf("\nBrainBucket --> На данный момент у тебя нет заметок... Давай это исправим!\n")
					} else {
						fmt.Printf("\nBrainBucket --> Твои заметки: \n%v\n", notes_db)
					}

				case "5":
					continue menu

				default:
					fmt.Printf("\nBrainBucket --> Такого действия с заметками не существует... --> BrainBucket Status: :<\n")
				}
			}

		case "2":
			fmt.Printf("\nBrainBucket --> Было приятно повидаться!\n")
			return

		case "3":
			fmt.Printf("\nBrainBucket v1.0.0 : Made by @xoti$\n")
			fmt.Printf("\nBrainBucket --> Привет, меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое!\n")

		default:
			fmt.Printf("\nBrainBucket --> Такого пункта меню не существует... --> BrainBucket Status: :<\n")
		}
	}
}

func NewNote(note_new_content *string, new_channel chan bool) {
	notes_db = append(notes_db, Note{
		Date:    time.Now(),
		Content: *note_new_content,
	})

	fmt.Printf("\nBrainBucket --> Добавил новую заметку!\n")

	new_channel <- true
}

func RemoveNote(note_delete_index *string, delete_channel chan bool) {
	note_delete_index_int, err0r := strconv.Atoi(*note_delete_index)
	if err0r != nil {
		fmt.Printf("\nBrainBucket --> Ты ввел неверные данные, по ним не получится найти необходимую заметку... --> BrainBucket Status: :<\n")
		delete_channel <- false
	}

	if note_delete_index_int < 0 || note_delete_index_int > len(notes_db) {
		fmt.Printf("\nBrainBucket --> Ты выбрал несуществующую заметку... --> BrainBucket Status: :<\n")
		delete_channel <- false
	} else {
		fmt.Printf("\nBrainBucket --> Заметка %v удалена!\n", notes_db[note_delete_index_int])
		notes_db = append(notes_db[:note_delete_index_int], notes_db[note_delete_index_int+1:]...)
		delete_channel <- true
	}
}

func EditNote(note_edit_index *string, note_edit_content *string, edit_channel chan bool) {
	note_edit_index_int, err0r := strconv.Atoi(*note_edit_index)
	if err0r != nil {
		fmt.Printf("\nBrainBucket --> Ты ввел неверные данные, по ним не получится найти необходимую заметку... --> BrainBucket Status: :<\n")
		edit_channel <- false
	}

	if note_edit_index_int < 0 || note_edit_index_int > len(notes_db) {
		fmt.Printf("\nBrainBucket --> Ты выбрал несуществующую заметку... --> BrainBucket Status: :<\n")
		edit_channel <- false
	} else {
		notes_db[note_edit_index_int] = Note{
			Date:    time.Now(),
			Content: *note_edit_content,
		}
		fmt.Printf("\nBrainBucket --> Заметка %v изменена!\n", notes_db[note_edit_index_int])
		edit_channel <- true
	}
}
