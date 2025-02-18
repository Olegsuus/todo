package apperror

type AppError struct {
	BusinessError string
	UserError     string
	Status        int
}

func (a AppError) Error() string {
	return a.UserError
}

type ReqError struct {
	Status int    `json:"status,omitempty" bson:"status"`
	Text   string `json:"text" bson:"text"`
}

func (e ReqError) Error() string {
	return e.Text
}

var Decoding = ReqError{Status: 400, Text: "Ошибка при декодировании данных, возможны невалидные форматы данных. Просьба свериться с документацией swagger"}
var Database = ReqError{Status: 500, Text: "Ошибка при обращении в базу данных"}
var Encoding = ReqError{Status: 500, Text: "Ошибка при кодировании данных"}
