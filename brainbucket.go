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
	Date    string
	Content Content
}

type Content struct {
	Title   string
	Content string
}

var notes_db []Note

func main() {
	fmt.Printf("\nBrainBucket CLI v1.0.0 : Made by xoti$\n")
	fmt.Printf("\nBrainBucket --> Привет, меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое!\n")

	input_scanner := bufio.NewScanner(os.Stdin)

menu:
	for {
		fmt.Printf("\nМеню:")
		fmt.Printf("\n1. Открыть меню заметок")
		fmt.Printf("\n2. Закрыть меня")
		fmt.Printf("\n3. Credits and License")

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
					fmt.Printf("\nВведи заголовок новой заметки (нажми Enter, если хочешь пропустить этот шаг): ")
					if !input_scanner.Scan() {
						fmt.Printf("\nBrainBucket --> Ошибка 2010: не переживай, это я виноват... --> BrainBucket Status: :(\n")
						fmt.Printf("\nError description: %v\n", input_scanner.Err())
						return
					}
					note_new_title := input_scanner.Text()

					fmt.Printf("Введи текст новой заметки: ")
					if !input_scanner.Scan() {
						fmt.Printf("\nBrainBucket --> Ошибка 2011: не переживай, это я виноват... --> BrainBucket Status: :(\n")
						fmt.Printf("\nError description: %v\n", input_scanner.Err())
						return
					}
					note_new_content := input_scanner.Text()

					new_channel := make(chan bool)
					go NewNote(&note_new_title, &note_new_content, new_channel)
					<-new_channel

				case "2":
					if notes_db == nil {
						fmt.Printf("\nBrainBucket --> На данный момент у тебя нет заметок... Давай это исправим!\n")
					} else {
						fmt.Printf("\nBrainBucket --> Выбери, какая из заметок тебе больше нужна, запомни, какая она по счету. Потом введи ее номер здесь!\n")
						fmt.Print("\nТвои заметки:\n")
						ShowNotes()
						fmt.Printf("\nВведи номер заметки, которую ты хочешь удалить: ")
						if !input_scanner.Scan() {
							fmt.Printf("\nBrainBucket --> Ошибка 202: не переживай, это я виноват... --> BrainBucket Status: :(\n")
							fmt.Printf("\nError description: %v\n", input_scanner.Err())
							return
						}
						note_delete_index := input_scanner.Text()
						note_delete_index_int, err0r := strconv.Atoi(note_delete_index)
						if err0r != nil {
							fmt.Printf("\nBrainBucket --> Ты ввел неверные данные, по ним не получится найти необходимую заметку... --> BrainBucket Status: :<\n")
							continue menu
						}
						if note_delete_index_int < 0 || note_delete_index_int > len(notes_db) {
							fmt.Printf("\nBrainBucket --> Ты выбрал несуществующую заметку... --> BrainBucket Status: :<\n")
						} else {
							delete_channel := make(chan bool)
							go RemoveNote(&note_delete_index_int, delete_channel)
							<-delete_channel
						}
					}

				case "3":
					if notes_db == nil {
						fmt.Printf("\nBrainBucket --> На данный момент у тебя нет заметок... Давай это исправим!\n")
					} else {
						fmt.Printf("\nBrainBucket --> Выбери, какую из заметок ты хочешь изменить, запомни, какая она по счету. Потом введи ее номер здесь!\n")
						fmt.Print("\nТвои заметки:\n")
						ShowNotes()

						fmt.Printf("\nВведи номер заметки, которую ты хочешь изменить: ")
						if !input_scanner.Scan() {
							fmt.Printf("\nBrainBucket --> Ошибка 2030: не переживай, это я виноват... --> BrainBucket Status: :(\n")
							fmt.Printf("\nError description: %v\n", input_scanner.Err())
							return
						}
						note_edit_index := input_scanner.Text()

						note_edit_index_int, err0r := strconv.Atoi(note_edit_index)
						if err0r != nil {
							fmt.Printf("\nBrainBucket --> Ты ввел неверные данные, по ним не получится найти необходимую заметку... --> BrainBucket Status: :<\n")
							continue menu
						}

						if note_edit_index_int < 0 || note_edit_index_int > len(notes_db) {
							fmt.Printf("\nBrainBucket --> Ты выбрал несуществующую заметку... --> BrainBucket Status: :<\n")
						} else {
							fmt.Printf("\nВведи новый заголовок для заметки (нажми Enter, если не хочешь его изменять): ")
							if !input_scanner.Scan() {
								fmt.Printf("\nBrainBucket --> Ошибка 2031: не переживай, это я виноват... --> BrainBucket Status: :(\n")
								fmt.Printf("\nError description: %v\n", input_scanner.Err())
								return
							}
							note_edit_title := input_scanner.Text()
							fmt.Printf("Введи новый текст для заметки (нажми Enter, если не хочешь его изменять): ")
							if !input_scanner.Scan() {
								fmt.Printf("\nBrainBucket --> Ошибка 2032: не переживай, это я виноват... --> BrainBucket Status: :(\n")
								fmt.Printf("\nError description: %v\n", input_scanner.Err())
								return
							}
							note_edit_content := input_scanner.Text()
							edit_channel := make(chan bool)
							go EditNote(&note_edit_index_int, &note_edit_title, &note_edit_content, edit_channel)
							<-edit_channel
						}
					}

				case "4":
					if notes_db == nil {
						fmt.Printf("\nBrainBucket --> На данный момент у тебя нет заметок... Давай это исправим!\n")
					} else {
						fmt.Print("\nТвои заметки:\n")
						ShowNotes()
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
			PrintLicense("LICENSE")

		default:
			fmt.Printf("\nBrainBucket --> Такого пункта меню не существует... --> BrainBucket Status: :<\n")
		}
	}
}

func NewNote(note_new_title *string, note_new_content *string, new_channel chan bool) {
	defer close(new_channel)

	if len(*note_new_content) == 0 {
		fmt.Printf("\nBrainBucket --> В твоей заметке ничего нет... --> BrainBucket Status: :<\n")
		new_channel <- false
	} else {
		notes_db = append(notes_db, Note{
			Date: time.Now().Format("02 Jan 2006, 15:04"),
			Content: Content{
				Title:   *note_new_title,
				Content: *note_new_content,
			},
		})
		fmt.Printf("\nBrainBucket --> Добавил новую заметку!\n")
		new_channel <- true
	}
}

func RemoveNote(note_delete_index_int *int, delete_channel chan bool) {
	defer close(delete_channel)

	fmt.Printf("\nBrainBucket --> Заметка удалена!\n")
	notes_db = append(notes_db[:*note_delete_index_int], notes_db[*note_delete_index_int+1:]...)
	delete_channel <- true
}

func EditNote(note_edit_index_int *int, note_edit_title *string, note_edit_content *string, edit_channel chan bool) {
	defer close(edit_channel)

	if len(*note_edit_title) == 0 && len(*note_edit_content) == 0 {
		fmt.Printf("\nBrainBucket --> Нет изменений... --> BrainBucket Status: :<\n")
		edit_channel <- true
	} else {
		switch {
		case len(*note_edit_title) == 0:
			notes_db[*note_edit_index_int] = Note{
				Date: time.Now().Format("02 Jan 2006, 15:04"),
				Content: Content{
					Title:   notes_db[*note_edit_index_int].Content.Title,
					Content: *note_edit_content,
				},
			}
			fmt.Printf("\nBrainBucket --> Заметка изменена!\n")
			edit_channel <- true

		case len(*note_edit_content) == 0:
			notes_db[*note_edit_index_int] = Note{
				Date: time.Now().Format("02 Jan 2006, 15:04"),
				Content: Content{
					Title:   *note_edit_title,
					Content: notes_db[*note_edit_index_int].Content.Content,
				},
			}
			fmt.Printf("\nBrainBucket --> Заметка изменена!\n")
			edit_channel <- true

		default:
			notes_db[*note_edit_index_int] = Note{
				Date: time.Now().Format("02 Jan 2006, 15:04"),
				Content: Content{
					Title:   *note_edit_title,
					Content: *note_edit_content,
				},
			}
			fmt.Printf("\nBrainBucket --> Заметка изменена!\n")
			edit_channel <- true
		}
	}
}

func ShowNotes() {
	for _, note := range notes_db {
		fmt.Printf("\n%v *. %v\n-----------------\n%v\n", note.Date, note.Content.Title, note.Content.Content)
	}
}

func PrintLicense(file_path string) {
	file, err0r := os.Open(file_path)
	if err0r != nil {
		fmt.Printf("\nBrainBucket --> Ошибка чтения файла... --> BrainBucket Status: :(\n")
		return
	}
	defer file.Close()

	file_scanner := bufio.NewScanner(file)
	for file_scanner.Scan() {
		fmt.Printf("\n%v", file_scanner.Text())
	}
	if err0r := file_scanner.Err(); err0r != nil {
		fmt.Printf("\nBrainBucket --> Ошибка чтения файла... --> BrainBucket Status: :(")
		return
	}
	fmt.Printf("\n")
}
