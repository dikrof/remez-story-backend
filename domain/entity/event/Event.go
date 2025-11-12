package event

type Event struct {
	ID          EventID
	Code        EventCode
	Title       string
	Description string
	Deprecated  bool
}
