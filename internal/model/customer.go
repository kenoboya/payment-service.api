package model

type CustomerData struct {
	CustomerID string
	Name       string
	Phone      string
	// Нужно будет с Auth брать полную информацию про пользователя(зарег), но данные могут быть null
	// Это значит что, на goods.api нужно будет проверять зарегестрирован ли пользователь по параметру user_id
	// Если это так, то нужно будет делать запрос на auth.service за данными
	// Логику goods.api, auth.api нужно дописать и подравить
}