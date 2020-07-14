package models

type booking struct {
	Name     string         `json:"name"`
	Bookings []bookingarray `json:"bookings"`
}
type bookingarray struct {
	Name       string   `json:"name"`
	From       string   `json:"from"`
	To         string   `json:"to"`
	Date       string   `json:"date"`
	Rupees     string   `json:"rupees"`
	Duration   string   `json:"duration"`
	Seatno     []string `json:"seatno"`
	Passengers []string `json:"passengers"`
}
type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type travel struct {
	Name     string `json:"name"`
	From     string `json:"from"`
	To       string `json:"to"`
	Date     string `json:"date"`
	Rupees   string `json:"rupees"`
	Duration string `json:"duration"`
}
type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Houseno  string `json:"houseno"`
	Street   string `json:"street"`
	City     string `json:"city"`
	Pincode  string `json:"pincode"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
}
