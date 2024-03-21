package controllers

import (
	"log"
	"net/http"
	"strconv"

	m "latihan_echo/models"

	"github.com/labstack/echo/v4"
)

func GetAllUsersEcho(c echo.Context) error {
	result, err := ModelGetAllUsersEcho()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func InsertUserEcho(c echo.Context) error {
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")

	result, err := ModelInsertUser(name, age, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUserEcho(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")

	result, err := ModelUpdateUser(id, name, age, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUserEcho(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))

	result, err := ModelDeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ModelGetAllUsersEcho() (m.UsersResponse, error) {
	db := connect()
	defer db.Close()
	var res m.UsersResponse

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return res, err
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.Usertype); err != nil {
			log.Println(err)
			return res, err
		} else {
			users = append(users, user)
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = users

	return res, nil
}

func ModelInsertUser(name string, age int, address string) (m.UserResponse, error) {
	db := connect()
	defer db.Close()
	var res m.UserResponse

	stmt, errQuery := db.Exec("INSERT INTO users(name, age, address) values (?,?,?)",
		name,
		age,
		address,
	)

	if errQuery == nil {
		res.Status = http.StatusOK
		res.Message = "Success"
	} else {
		res.Status = http.StatusInternalServerError
		res.Message = "Insert Failed !"
	}

	var user m.User
	lastInsertId, err := stmt.LastInsertId()
	if err != nil {
		return res, err
	}
	user.ID = int(lastInsertId)
	user.Name = name
	user.Age = age
	user.Address = address
	res.Data = user

	return res, nil
}

func ModelUpdateUser(id int, name string, age int, address string) (m.UserResponse, error) {
	db := connect()
	defer db.Close()
	var res m.UserResponse

	query := "UPDATE users SET name = ?, age = ?, address = ? WHERE id = ? "

	stmt, err := db.Prepare(query)
	if err != nil {
		return res, err
	}

	result, errQuery := stmt.Exec(name, age, address, id)
	if errQuery != nil {
		return res, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return res, nil
	}

	var user m.User
	user.Name = name
	user.Age = age
	user.Address = address

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = user

	return res, nil
}

func ModelDeleteUser(id int) (m.UserResponse, error) {
	var res m.UserResponse

	db := connect()
	defer db.Close()

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		id,
	)
	if errQuery != nil {
		return res, errQuery
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	return res, nil
}
