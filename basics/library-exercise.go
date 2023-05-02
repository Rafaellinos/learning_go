package main

import (
	"fmt"
	"time"
)

type Member struct {
	name string
}

type Title string

type BookInfo struct {
	id           int
	available    bool
	checkOutDate time.Time
	checkInDate  time.Time
}

type Library struct {
	name    string
	books   map[Title]BookInfo
	members []Member
}

func appendMember(members *[]Member, name string) {
	*members = append(*members, Member{name})
}

func rentBook(books *map[Title]BookInfo, bookTitle Title) {
	bookInfo, found := (*books)[bookTitle]
	if !found {
		panic(fmt.Sprintf("Book %s not found", bookTitle))
	}
	if !bookInfo.available {
		panic(fmt.Sprintf("Book %s not available", bookTitle))
	}
	(*books)[bookTitle] = BookInfo{
		id:           bookInfo.id,
		available:    false,
		checkInDate:  time.Now(),
		checkOutDate: bookInfo.checkOutDate,
	}
}

func recallBook(books *map[Title]BookInfo, bookTitle Title) {
	bookInfo, found := (*books)[bookTitle]
	if !found {
		panic(fmt.Sprintf("Book %s not found", bookTitle))
	}
	if bookInfo.available {
		panic(fmt.Sprintf("Book %s already available", bookTitle))
	}
	(*books)[bookTitle] = BookInfo{
		id:           bookInfo.id,
		available:    true,
		checkInDate:  bookInfo.checkInDate,
		checkOutDate: time.Now(),
	}
}

func libraryExercise() {
	rafael := Member{
		"Rafael",
	}
	var members []Member
	members = append(members, rafael)
	appendMember(&members, "yasmin")

	books := map[Title]BookInfo{
		"Clean Code":         {1, true, time.Time{}, time.Time{}},
		"Fluent Python":      {2, true, time.Time{}, time.Time{}},
		"Project Phoenix":    {3, true, time.Time{}, time.Time{}},
		"Clean Architecture": {4, true, time.Time{}, time.Time{}},
	}

	stateLibrary := Library{
		"State Library",
		books,
		members,
	}

	fmt.Println(stateLibrary)
	rentBook(&stateLibrary.books, "Clean Code")
	fmt.Println(stateLibrary)
	recallBook(&stateLibrary.books, "Clean Code")
	fmt.Println(stateLibrary)
}
