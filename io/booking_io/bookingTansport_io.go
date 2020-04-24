package booking_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/booking"
)

const bookingtransport = api.BASE_URL + "bookingtransport"

func CreateBookingTransport(transport booking.BookingTransport) (booking.BookingTransport, error) {
	entity := booking.BookingTransport{}
	resp, _ := api.Rest().SetBody(transport).Post(bookingtransport + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBookingTransport(transport booking.BookingTransport) (booking.BookingTransport, error) {
	entity := booking.BookingTransport{}
	resp, _ := api.Rest().SetBody(transport).Post(bookingtransport + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookingTransport(id string) (booking.BookingTransport, error) {
	entity := booking.BookingTransport{}
	resp, _ := api.Rest().Get(bookingtransport + "read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookingTransport(id string) (booking.BookingTransport, error) {
	entity := booking.BookingTransport{}
	resp, _ := api.Rest().Get(bookingtransport + "delete?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookingTransports() (booking.BookingTransport, error) {
	entity := booking.BookingTransport{}
	resp, _ := api.Rest().Get(bookingtransport + "reads")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
