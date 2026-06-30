package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"github.com/raffy-io/bookings/internal/db"
	"github.com/raffy-io/bookings/internal/forms"
	"github.com/raffy-io/bookings/internal/models"
	"github.com/raffy-io/bookings/ui/layout"
	"github.com/raffy-io/bookings/ui/pages"
)

type Handlers struct {
	Session *scs.SessionManager
	Queries *db.Queries
}

func New(session *scs.SessionManager,queries *db.Queries) *Handlers {
	return &Handlers{
		Session: session,
		Queries:queries,
	}
}

func (h *Handlers) Homepage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	notif := &models.NotifType{
		ErrorMessage: h.Session.PopString(r.Context(), "error"),
	}
	component := pages.Homepage(notif)
	layout := layout.Base(path,"Welcome", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	component := pages.About()
	layout := layout.Base(path,"About Us", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) Generals(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	component := pages.Generals()
	layout := layout.Base(path,"General's Quarters", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) Majors(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	component := pages.Majors()
	layout := layout.Base(path,"Major's Suite", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) Availability(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	notif := &models.NotifType{}
	form := &forms.AvailableForm{}
	token := nosurf.Token(r)

	component := pages.Availability(notif, form, token)
	layout := layout.Base(path,"Check Availability", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) PostAvailability(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	notif := &models.NotifType{}
	form := &forms.AvailableForm{}
	token := nosurf.Token(r)

	// Assign and trim directly
	form.Arrival = strings.TrimSpace(r.Form.Get("start_date"))
	form.Departure = strings.TrimSpace(r.Form.Get("end_date"))

	// basic validation
	errors, isValid := forms.ValidateAvailable(form)

	if !isValid {
		path := r.URL.Path
		notif.ErrorsMap = errors

		component := pages.Availability(notif, form, token)
		layout := layout.Base(path,"Available Rooms", component)
		templ.Handler(layout).ServeHTTP(w, r)
		return
	}

	fmt.Println("Success! form submitted!")

}


func (h *Handlers) Booking(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	notif := &models.NotifType{
		ErrorMessage: h.Session.PopString(r.Context(), "error"),
	}
	form := &forms.BookingForm{}
	token := nosurf.Token(r)

	component := pages.Booking(notif, form, token)
	layout := layout.Base(path,"Make Booking", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) PostBooking(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	notif := &models.NotifType{}
	token := nosurf.Token(r)
	form := &forms.BookingForm{
		// Assign and trim directly
		FirstName: strings.TrimSpace(r.Form.Get("first_name")),
		LastName:  strings.TrimSpace(r.Form.Get("last_name")),
		Email:     strings.TrimSpace(r.Form.Get("email")),
		Phone:     strings.TrimSpace(r.Form.Get("phone")),
	}

	// basic validation
	errors, isValid := forms.ValidateReservation(form)

	if !isValid {
		path := r.URL.Path
		notif.ErrorsMap = errors
		component := pages.Booking(notif, form, token)
		layout := layout.Base(path,"Make Booking", component)
		w.WriteHeader(http.StatusUnprocessableEntity) // 422 status code
		templ.Handler(layout).ServeHTTP(w, r)
		return
	}

	data, err := h.Queries.CreateReservation(r.Context(),db.CreateReservationParams{
		FirstName: form.FirstName,
		LastName: form.LastName,
		Email: form.Email,
		Phone: form.Phone,
	})

	if err != nil {
		h.Session.Put(r.Context(), "error", "Failed to book reservation..please try again.")
		// The Professional Go Way (Structured Logging)
    	slog.Error("failed to create reservation in database", "error", err)
		http.Redirect(w,r,"/booking",http.StatusSeeOther)
		return 
	}

	// Notif data on redirect
	h.Session.Put(r.Context(), "flash", "Reservation completed successfully!")
	// Reservation summary on redirect
	h.Session.Put(r.Context(), "reservation", data)

	// Redirect
	http.Redirect(w, r, "/booking-summary", http.StatusSeeOther)
}

func (h *Handlers) BookingSummary(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	notif := &models.NotifType{
		SuccessMessage: h.Session.PopString(r.Context(), "flash"),
	}
	reservation, ok := h.Session.Get(r.Context(), "reservation").(*db.Reservation)
	if !ok {
		h.Session.Put(r.Context(), "error", "Can't get reservation from the session")
		fmt.Println("you will be redirected")
		http.Redirect(w,r,"/",http.StatusTemporaryRedirect)
		return
	}
	
	h.Session.Remove(r.Context(),"reservation")
	component := pages.BookingSummary(notif, reservation)
	layout := layout.Base(path,"Reservation Summary", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (h *Handlers) Contacts(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	component := pages.Contacts()
	layout := layout.Base(path,"Contact Us", component)
	templ.Handler(layout).ServeHTTP(w, r)
}