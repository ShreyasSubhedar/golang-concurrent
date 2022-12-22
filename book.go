package main

import (
	"fmt"
)

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%v\n"+
			"Published:\t\t%v\n", b.Title, b.Author, b.YearPublished)
}

// books := make([]Book, 0)

var books = []Book{
	Book{
		ID:            1,
		Title:         "Vayuputras",
		Author:        "Amish",
		YearPublished: 2007,
	},
	Book{
		ID:            2,
		Title:         "Nagas",
		Author:        "Amish",
		YearPublished: 2005,
	},
	Book{
		ID:            3,
		Title:         "Meluha",
		Author:        "Amish",
		YearPublished: 2003,
	},
	Book{
		ID:            4,
		Title:         "Book4",
		Author:        "ADSa",
		YearPublished: 2001,
	},
	Book{
		ID:            5,
		Title:         "Book5",
		Author:        "ADSa",
		YearPublished: 2001,
	},
	Book{
		ID:            6,
		Title:         "Book6",
		Author:        "ADSa",
		YearPublished: 2001,
	},
	Book{
		ID:            7,
		Title:         "Book7",
		Author:        "ADSa",
		YearPublished: 2001,
	},
	Book{
		ID:            8,
		Title:         "Book8",
		Author:        "ADSa",
		YearPublished: 2001,
	},
	Book{
		ID:            9,
		Title:         "Book9",
		Author:        "ADSa",
		YearPublished: 2001,
	},
	Book{
		ID:            10,
		Title:         "Book10",
		Author:        "ADSa",
		YearPublished: 2001,
	},
}
