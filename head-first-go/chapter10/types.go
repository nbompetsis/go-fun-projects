package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1900 {
		return errors.New("invalid year")
	}

	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month > 12 || month < 0 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day > 31 || day < 1 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

func main() {
	date := Date{}
	err := date.SetYear(1990)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)

	fmt.Println(date.Day())
	fmt.Println(date.Month())
	fmt.Println(date.Year())

}
