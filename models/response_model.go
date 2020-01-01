package models

type Response struct {
	Status  string      `json:"status" bson:"status"`
	Data    interface{} `json:"data" bson:"data"`
	Message string      `json:"message" bson:"message"`
}

type ResponseStatusEnum struct {
	ERROR string
	OK    string
}

var ResponseStatus = ResponseStatusEnum{
	ERROR: "ERROR",
	OK:    "OK",
}
