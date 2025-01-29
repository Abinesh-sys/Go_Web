package routes
import (

	"go_web/controllers"
	"net/http"
)

func Registerroutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.ContactIndex)
	mux.HandleFunc("/contacts", controllers.ContactCreate)
	return mux
}