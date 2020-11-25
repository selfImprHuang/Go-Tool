/*
 *  @Author : huangzj
 *  @Time : 2020/11/12 14:04
 *  @Description：
 */

package linqUse

import "time"

type Book struct {
	Name        string
	Author      string
	Money       float64
	WordsNum    int
	PublishTime time.Time
}

func MakeBook() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Effective Java",
		Author:      "Java",
		Money:       78,
		WordsNum:    9000,
		PublishTime: time.Date(2020, 2, 15, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Java语言",
		Author:      "Java",
		Money:       50,
		WordsNum:    3000,
		PublishTime: time.Date(2020, 2, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Lua语言",
		Author:      "Lua",
		Money:       75,
		WordsNum:    45000,
		PublishTime: time.Date(2020, 1, 10, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "React语言",
		Author:      "React",
		Money:       99,
		WordsNum:    14500,
		PublishTime: time.Date(2020, 7, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Red语言",
		Author:      "Red",
		Money:       28,
		WordsNum:    880,
		PublishTime: time.Date(2019, 4, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "JavaScript语言",
		Author:      "JavaScript",
		Money:       81,
		WordsNum:    3776,
		PublishTime: time.Date(2019, 5, 17, 10, 0, 0, 0, time.Local),
	})
	return bookList
}

func MakeBook1() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Effective Java",
		Author:      "Java",
		Money:       78,
		WordsNum:    9000,
		PublishTime: time.Date(2020, 2, 15, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Go语言(第二版)",
		Author:      "Go",
		Money:       50,
		WordsNum:    3000,
		PublishTime: time.Date(2020, 2, 1, 10, 0, 0, 0, time.Local),
	})
	return bookList
}
func MakeBook2() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	return bookList
}

func MakeBook3() []Book {
	bookList := make([]Book, 0)
	bookList = append(bookList, Book{
		Name:        "Go语言",
		Author:      "Go",
		Money:       100,
		WordsNum:    1000,
		PublishTime: time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
	})
	bookList = append(bookList, Book{
		Name:        "Go语言(第三版)",
		Author:      "Go",
		Money:       50,
		WordsNum:    3000,
		PublishTime: time.Date(2020, 2, 1, 10, 0, 0, 0, time.Local),
	})
	return bookList
}
