package booking_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/booking"
)

const bookingURL = api.BASE_URL + "booking/"

func CreateBooking(mybooking booking.Booking) (booking.Booking, error) {
	entity := booking.Booking{}
	resp, _ := api.Rest().SetBody(mybooking).Post(bookingURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBooking(mybooking booking.Booking) (booking.Booking, error) {
	entity := booking.Booking{}
	resp, _ := api.Rest().SetBody(mybooking).Post(bookingURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBooking(id string) (booking.Booking, error) {
	entity := booking.Booking{}
	resp, _ := api.Rest().Get(bookingURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBooking(id string) (booking.Booking, error) {
	entity := booking.Booking{}
	resp, _ := api.Rest().Get(bookingURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookings() ([]booking.Booking, error) {
	entity := []booking.Booking{}
	resp, _ := api.Rest().Get(bookingURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
