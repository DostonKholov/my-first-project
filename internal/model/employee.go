package model

type Employee struct {
	Id          int    `json:"id"`          // Уникальный идентификатор сотрудника
	LastName    string `json:"lastname"`    // Фамилия сотрудника
	FirstName   string `json:"firstname"`   // Имя сотрудника
	MiddleName  string `json:"middlename"`  // Отчество сотрудника
	Position    string `json:"position"`    // Должность
	Department  string `json:"department"`  // Название отдела
	Email       string `json:"email"`       // Электронная почта
	PhoneNumber string `json:"phonenumber"` // Номер телефона
	HireDate    string `json:"hiredate"`    // Дата приёма на работу
	Status      string `json:"status"`      // Дата приёма на работу
	PhotoUrl    string `json:"photourl"`    // Ссылка на фотографию
	Notes       string `json:"notes"`       // Дополнительные заметки
}
