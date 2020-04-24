package booking_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/booking"
)

const booklanguageURL = api.BASE_URL + "booklanguage"

func CreateBookLanguage(language booking.BookingLanguage) (booking.BookingLanguage, error) {
	entity := booking.BookingLanguage{}
	resp, _ := api.Rest().SetBody(language).Post(booklanguageURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func updateBookLanguage(language booking.BookingLanguage) (booking.BookingLanguage, error) {
	entity := booking.BookingLanguage{}
	resp, _ := api.Rest().SetBody(language).Post(booklanguageURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookLanguage(id string) (booking.BookingLanguage, error) {
	entity := booking.BookingLanguage{}
	resp, _ := api.Rest().Get(booklanguageURL + "read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookLanguage(id string) (booking.BookingLanguage, error) {
	entity := booking.BookingLanguage{}
	resp, _ := api.Rest().Get(booklanguageURL + "delete?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookLanguages() ([]booking.BookingLanguage, error) {
	entity := []booking.BookingLanguage{}
	resp, _ := api.Rest().Get(booklanguageURL + "reads")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
