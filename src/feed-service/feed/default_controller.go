package feed
import(
	"net/http"
	"fmt"
	"strconv"

)
type DefaultController struct {
}
func NewDefaultController(
) *DefaultController {
	return &DefaultController{}
}

func (ctrl *DefaultController) GetFeed(w http.ResponseWriter, r *http.Request) {	
		// TODO Correct Implementation
		resp, err := http.Get("http://localhost:3000/posts")
		fmt.Println(resp)
		
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (ctrl *DefaultController) GetFeedWithAmount(w http.ResponseWriter, r *http.Request) {
	amount := r.Context().Value("amount").(string)

	amo, err := strconv.ParseInt(amount, 10, 64)
		fmt.Println(amo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
}
	
