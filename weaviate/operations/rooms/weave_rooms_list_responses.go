package rooms


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/weaviate/models"
)

/*WeaveRoomsListOK Successful response

swagger:response weaveRoomsListOK
*/
type WeaveRoomsListOK struct {

	// In: body
	Payload *models.RoomsListResponse `json:"body,omitempty"`
}

// NewWeaveRoomsListOK creates WeaveRoomsListOK with default headers values
func NewWeaveRoomsListOK() *WeaveRoomsListOK {
	return &WeaveRoomsListOK{}
}

// WithPayload adds the payload to the weave rooms list o k response
func (o *WeaveRoomsListOK) WithPayload(payload *models.RoomsListResponse) *WeaveRoomsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave rooms list o k response
func (o *WeaveRoomsListOK) SetPayload(payload *models.RoomsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveRoomsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}