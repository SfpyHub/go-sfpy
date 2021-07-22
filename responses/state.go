package responses

type State string

const STARTED State = "STARTED"
const PAID State = "PAID"
const REFUNDED State = "REFUNDED"
const CANCELLED State = "CANCELLED"
const ERRORED State = "ERRORED"

var StringToState = map[string]State{
	"STARTED":   STARTED,
	"PAID":      PAID,
	"REFUNDED":  REFUNDED,
	"CANCELLED": CANCELLED,
	"ERRORED":   ERRORED,
}
