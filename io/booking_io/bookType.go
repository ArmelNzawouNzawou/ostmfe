package booking_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/booking"
)

const booktypeURL = api.BASE_URL + "booktype"

func CreateBookType(bookType booking.BookType) (booking.BookType, error) {
	entity := booking.BookType{}
	resp, _ := api.Rest().SetBody(bookType).Post(booktypeURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func updateBookType(bookType booking.BookType) (booking.BookType, error) {
	entity := booking.BookType{}
	resp, _ := api.Rest().SetBody(bookType).Post(booktypeURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookType(id string) (booking.BookType, error) {
	entity := booking.BookType{}
	resp, _ := api.Rest().Get(booktypeURL + "read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookType(id string) (booking.BookType, error) {
	entity := booking.BookType{}
	resp, _ := api.Rest().Get(booktypeURL + "delete?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookTypes() ([]booking.BookType, error) {
	entity := []booking.BookType{}
	resp, _ := api.Rest().Get(booktypeURL + "reads")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
