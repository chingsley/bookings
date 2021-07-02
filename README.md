## Description

_Bookings_ is a Booking and Reservations app for a Bed & Breakfast business. This app is built using the Go programming language.

### Key Functionality

- Showcase the available property, usually available rooms
- Customers can book a room for one or more nights
- Customers can check the availability of a room
- Notify guests, and notify property owner
- An admin dashboard system for property owners to:

  - review existing bookings
  - manage the booking calendar
  - change or cancel a booking

### Key System Components:

- Authentication system
- Database
- Notification system

### Technologies Used

- go version go1.16.5 darwin/amd64

## Dependencies

- [Chi Router](https://github.com/go-chi/chi/v5) v2.4.0
- [Alex Edwards SCS](https://github.com/alexedwards/scs) session manager v5.0.3
- [justinas nosurf](https://github.com/justinas/nosurf) v1.1.1

### To run the app locally:

- Clone the repo, then run the command below:
- `go run cmd/web/*.go`
