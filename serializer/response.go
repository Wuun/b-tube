package serializer

//Response is use to format  server's response.
type Response struct {
	StatusCode int
	Msg        string
	Error      string
	Data       interface{}
}
