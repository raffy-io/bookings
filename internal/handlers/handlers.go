package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"

	"github.com/justinas/nosurf"
	"github.com/raffy-io/bookings/internal/forms"
	"github.com/raffy-io/bookings/internal/models"
	"github.com/raffy-io/bookings/ui/layout"
	"github.com/raffy-io/bookings/ui/pages"
)

type Handlers struct {
	Session *scs.SessionManager
}


func New(session *scs.SessionManager) *Handlers {
	return &Handlers{
		Session: session,
	}
}

// Home renders home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	component := pages.Home()
	layout := layout.Base("Welcome", component)
	templ.Handler(layout).ServeHTTP(w, r)

}

// About renders about page
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	component := pages.AboutUs()
	layout := layout.Base("About Us", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

// Generals renders generals quarter page
func (h *Handlers) Generals(w http.ResponseWriter, r *http.Request) {
	component := pages.GeneralsQuarters()
	layout := layout.Base("General's Quarters", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

// Majors renders majors suite page
func (h *Handlers) Majors(w http.ResponseWriter, r *http.Request) {
	component := pages.MajorsSuite()
	layout := layout.Base("Major's Suite", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

// Availability renders availability page
func (h *Handlers) Availability(w http.ResponseWriter, r *http.Request) {
	notif := &models.NotifType{}
	availabilityForm := &forms.AvailableForm{}
	token := nosurf.Token(r)

	component := pages.AvailableRooms(notif,availabilityForm, token)
	layout := layout.Base("Available Rooms", component)

	templ.Handler(layout).ServeHTTP(w, r)
}

// PostAvailability handles post requests
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
		notif.ErrorsMap = errors

		if notif.ErrorsMap["arrival"] != ""{
			notif.ErrorMessage = notif.ErrorsMap["arrival"]
		} else if notif.ErrorsMap["departure"] != "" {
			notif.ErrorMessage = notif.ErrorsMap["departure"]
		}

		component := pages.AvailableRooms(notif, form, token)
		layout := layout.Base("Available Rooms", component)
		templ.Handler(layout).ServeHTTP(w, r)
		return
	}

	fmt.Println("Success! form submitted!")
	
}

// Reservation renders make reservation page
func (h *Handlers) Reservation(w http.ResponseWriter, r *http.Request) {
	notif := &models.NotifType{}
	form := &forms.ReserveForm{}
	token := nosurf.Token(r)

	component := pages.MakeReservation(notif, form,token)
	layout := layout.Base("Make Reservation", component)
	templ.Handler(layout).ServeHTTP(w, r)
}

// PostReservation handles post requests
func (h *Handlers) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

	notif := &models.NotifType{}
	form := &forms.ReserveForm{}
	token := nosurf.Token(r)

	// Assign and trim directly
    form.FirstName = strings.TrimSpace(r.Form.Get("first_name"))
    form.LastName = strings.TrimSpace(r.Form.Get("last_name"))
    form.Email = strings.TrimSpace(r.Form.Get("email"))
    form.Phone = strings.TrimSpace(r.Form.Get("phone"))

	// basic validation 
	errors, isValid := forms.ValidateReservation(form)

	if !isValid {
		notif.ErrorsMap = errors
		component := pages.MakeReservation(notif, form,token)
		layout := layout.Base("Make Reservation", component)
		templ.Handler(layout).ServeHTTP(w, r)

		return
	}

	fmt.Println("Success! form is submitted!")
	
}

// Contacts renders contacts page
func (h *Handlers) Contacts(w http.ResponseWriter, r *http.Request) {
	component := pages.Contacts()
	layout := layout.Base("Contact Us", component)
	templ.Handler(layout).ServeHTTP(w, r)
}
