package api

type Message struct {
	Status int
	Msg    string
}

func SuccessMessage(msg string) Message {
	return Message{
		Status: 200,
		Msg:    msg,
	}
}

func GenericError() Message {
	return Message{
		Status: 500,
		Msg:    "something went wrong",
	}
}
func BadRequestError(msg string) Message {
	return Message{
		Status: 400,
		Msg:    msg,
	}
}
func NoContentError() Message {
	return Message{
		Status: 206,
		Msg:    "no resource present",
	}
}
func NotFoundError(msg string) Message {
	return Message{
		Status: 404,
		Msg:    msg,
	}
}
