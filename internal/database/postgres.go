package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.mod/internal/model"
	"log"
	"os"
)

func NewConnectPostgres() *sql.DB {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := os.Getenv("DB_PASSWORD")
	dbname := viper.GetString("db.dbname")

	dbParams := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dbParams)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	log.Println("Connected to Postgres successfully")
	return db
}

// Сотрудник

// Получить одного сотрудника по ID
func (d *Database) GetEmployeeByID(id int64) (model.Employee, error) {
	query := `SELECT id, lastname, firstname, middlename, position, department, email, phonenumber, hiredate, status, photourl, notes FROM employees WHERE id=$1`
	row := d.Connection.QueryRow(query, id)

	var employee model.Employee
	err := row.Scan(
		&employee.Id,
		&employee.LastName,
		&employee.FirstName,
		&employee.MiddleName,
		&employee.Position,
		&employee.Department,
		&employee.Email,
		&employee.PhoneNumber,
		&employee.HireDate,
		&employee.Status,
		&employee.PhotoUrl,
		&employee.Notes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, fmt.Errorf("сотрудник с id %d не найден", id)
		}
		return employee, fmt.Errorf("ошибка получения сотрудника: %v", err)
	}

	return employee, nil
}

// Получить всех сотрудников
func (d *Database) GetAllEmployees() ([]model.Employee, error) {
	query := `SELECT id, lastname, firstname, middlename, position, department, email, phonenumber, hiredate, status, photourl, notes FROM employees`
	rows, err := d.Connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка сотрудников: %v", err)
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(
			&emp.Id,
			&emp.LastName,
			&emp.FirstName,
			&emp.MiddleName,
			&emp.Position,
			&emp.Department,
			&emp.Email,
			&emp.PhoneNumber,
			&emp.HireDate,
			&emp.Status,
			&emp.PhotoUrl,
			&emp.Notes,
		); err != nil {
			return nil, fmt.Errorf("ошибка чтения сотрудников: %v", err)
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

// Создать нового сотрудника
func (d *Database) CreateEmployee(employee model.Employee) error {
	query := `INSERT INTO employees (lastname, firstname, middlename, position, department, email, phonenumber, hiredate, status, photourl, notes)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := d.Connection.Exec(query,
		employee.LastName,
		employee.FirstName,
		employee.MiddleName,
		employee.Position,
		employee.Department,
		employee.Email,
		employee.PhoneNumber,
		employee.HireDate,
		employee.Status,
		employee.PhotoUrl,
		employee.Notes,
	)
	if err != nil {
		return fmt.Errorf("ошибка создания сотрудника: %v", err)
	}
	return nil
}

// Удалить сотрудника по ID
func (d *Database) DeleteEmployee(id int64) error {
	query := `DELETE FROM employees WHERE id=$1`
	_, err := d.Connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления сотрудника: %v", err)
	}
	return nil
}

// Обновить данные сотрудника
func (d *Database) UpdateEmployee(id int64, employee model.Employee) error {
	query := `UPDATE employees 
              SET lastname=$1, firstname=$2, middlename=$3, position=$4, department=$5, email=$6, phonenumber=$7, hiredate=$8, status=$9, photourl=$10, notes=$11
              WHERE id=$12`

	_, err := d.Connection.Exec(query,
		employee.LastName,
		employee.FirstName,
		employee.MiddleName,
		employee.Position,
		employee.Department,
		employee.Email,
		employee.PhoneNumber,
		employee.HireDate,
		employee.Status,
		employee.PhotoUrl,
		employee.Notes,
		id,
	)
	if err != nil {
		return fmt.Errorf("ошибка обновления сотрудника: %v", err)
	}
	return nil
}

// Пользователи

// Получить одного пользователя по ID
func (d *Database) GetUserByID(id int64) (model.User, error) {
	query := `SELECT id, username, password, role FROM users WHERE id=$1`

	var user model.User
	row := d.Connection.QueryRow(query, id)
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("пользователь с id %d не найден", id)
		}
		return user, fmt.Errorf("ошибка при получении пользователя: %v", err)
	}

	return user, nil
}

// Получить всех пользователей
func (d *Database) GetAllUsers() ([]model.User, error) {
	query := `SELECT id, username, password, role FROM users`
	rows, err := d.Connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
		); err != nil {
			return nil, fmt.Errorf("ошибка сканирования пользователя: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// Создать нового пользователя
func (d *Database) CreateUser(user model.User) error {
	query := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`
	_, err := d.Connection.Exec(query, user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("ошибка добавления пользователя: %v", err)
	}
	return nil
}

// Удалить пользователя по ID
func (d *Database) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := d.Connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %v", err)
	}
	return nil
}

// Обновить пользователя
func (d *Database) UpdateUser(id int64, user model.User) error {
	query := `UPDATE users SET username=$1, password=$2, role=$3 WHERE id=$4`
	_, err := d.Connection.Exec(query, user.Username, user.Password, user.Role, id)
	if err != nil {
		return fmt.Errorf("ошибка обновления пользователя: %v", err)
	}
	return nil
}
