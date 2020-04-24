package booking_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/booking"
)

const bookaddressURL = api.BASE_URL + "bookaddress/"

func CreateBookingAddress(address booking.BookingAddress) (booking.BookingAddress, error) {
	entity := booking.BookingAddress{}
	resp, _ := api.Rest().SetBody(address).Post(bookaddressURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBookingAddress(address booking.BookingAddress) (booking.BookingAddress, error) {
	entity := booking.BookingAddress{}
	resp, _ := api.Rest().SetBody(address).Post(bookaddressURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookingAddress(id string) (booking.BookingAddress, error) {
	entity := booking.BookingAddress{}
	resp, _ := api.Rest().Get(bookaddressURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookingAddress(id string) (booking.BookingAddress, error) {
	entity := booking.BookingAddress{}
	resp, _ := api.Rest().Get(bookaddressURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookingAddresses() ([]booking.BookingAddress, error) {
	entity := []booking.BookingAddress{}
	resp, _ := api.Rest().Get(bookaddressURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
