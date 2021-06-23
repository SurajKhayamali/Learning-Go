package utils

type HTTPRequest struct {
	Method string
}

func SwitchDemo() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get request")
		// With Go, we have implicit break
		// if you do want to fall from one case to next, we can do that by providing a fallthrough statement which we provide by using 'fallthrough' keyword
		// fallthrough
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}
}
