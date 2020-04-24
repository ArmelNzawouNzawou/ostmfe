package booking

import "time"

type Booking struct {
	Id             string    `json:"id"`
	Date           time.Time `json:"date"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	NumberOfPeople int64     `json:"numberOfPeople"`
	Age            int64     `json:"age"`
	Grade          string    `json:"grade"`
	FaxNumber      string    `json:"faxNumber"`
}

type BookingAddress struct {
	AddressId   string `json:"addressId"`
	BookingId   string `json:"bookingId"`
	Description string `json:"description"`
}

type BookingTransport struct {
	TransportId string `json:"transportId"`
	BookingId   string `json:"bookingId"`
	Description string `json:"description"`
}

type BookingLanguage struct {
	LanguageId  string `json:"languageId"`
	BookId      string `json:"bookId"`
	Description string `json:"description"`
}

type BookType struct {
	Id          string `json:"id"`
	TypeName    string `json:"type_name"`
	Description string `json:"description"`
}
