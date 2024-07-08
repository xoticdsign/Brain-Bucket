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

var notes []Note

func main() {

	input_scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nBrainBucket v1.0.0 : Made by @xoti$\n")
	fmt.Printf("\nОписание:\nПривет! Меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое :)\n")

menu:

	for {

		fmt.Printf("\nМеню:")
		fmt.Printf("\n1. Доступные команды")
		fmt.Printf("\n2. Описание программы")
		fmt.Printf("\n3. Закрыть программу")
		fmt.Printf("\n...\nВыбери действие: ")

		if !input_scanner.Scan() {
			fmt.Printf("\nОшибка ввода в меню :( - %v\n", input_scanner.Err())
			return
		}

		menu_action := strings.TrimSpace(input_scanner.Text())

		switch menu_action {

		case "1":

			for {

				fmt.Printf("\n1. Добавить заметку")
				fmt.Printf("\n2. Удалить заметку")
				fmt.Printf("\n3. Редактировать заметку")
				fmt.Printf("\n4. Показать все заметки")
				fmt.Printf("\n5. Вернуться в меню")
				fmt.Printf("\n...\nВыбери действие: ")

				if !input_scanner.Scan() {
					fmt.Printf("\nОшибка ввода в субменю :( - %v\n", input_scanner.Err())
					return
				}

				submenu_action := strings.TrimSpace(input_scanner.Text())

				switch submenu_action {

				case "1":

					fmt.Printf("\nВведи текст заметки: ")

					if !input_scanner.Scan() {
						fmt.Printf("\nОшибка ввода заметки :( - %v\n", input_scanner.Err())
						return
					}

					note_content := input_scanner.Text()

					writing_note := make(chan bool)

					go NewNote(&note_content, writing_note)

					<-writing_note

				case "2":

					fmt.Printf("\nКак удалить заметку:\nПопроси меня показать все заметки и выбери, какая из них тебе больше нужна, запомни, какая она по счету!\nПотом введи ее номер здесь :)\n")
					fmt.Printf("\nВведи номер заметки, которую ты хочешь удалить: ")

					if !input_scanner.Scan() {
						fmt.Printf("\nОшибка выбора заметки :O - %v\n", input_scanner.Err())
						return
					}

					note_number := input_scanner.Text()

					deletion_channel := make(chan bool)

					go RemoveNote(&note_number, deletion_channel)

					<-deletion_channel

				case "3":

				case "4":

					fmt.Printf("\nТвои заметки:\n%v\n", notes)

				case "5":

					continue menu

				default:

					fmt.Printf("\nТакого действия с заметками не существует...\n")

				}

			}

		case "2":

			fmt.Printf("\nBrainBucket v1.0.0 : Made by @xoti$\n")
			fmt.Printf("\nОписание:\nПривет! Меня зовут BrainBucket, я помогу тебе запомнить важное и вспомнить необходимое :)\n")

		case "3":

			fmt.Printf("\nБыло приятно повидаться!\n")

			return

		default:

			fmt.Printf("\nТы выбрал несуществующее действие, попробуй еще раз :O\n")

		}

	}

}

func NewNote(note_content *string, writing_note chan bool) {

	notes = append(notes, Note{
		Date:    time.Now(),
		Content: *note_content,
	})

	fmt.Printf("Добавил твою заметку!\n")

	writing_note <- true

}

func RemoveNote(note_number *string, deletion_channel chan bool) {

	note_number_int, err0r := strconv.Atoi(*note_number)
	if err0r != nil {
		fmt.Printf("Недопустимое значение :< - %v\n", *note_number)
		deletion_channel <- false
	} else {
		notes = append(notes[:note_number_int], notes[note_number_int+1:]...)
		fmt.Printf("Заметка %v удалена!\n", notes[note_number_int])
		deletion_channel <- true
	}

}
