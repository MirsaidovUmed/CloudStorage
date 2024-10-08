package errors

import "errors"

var (
	ErrDataNotFound   = errors.New("не найдено")
	ErrAlreadyHasUser = errors.New("вы уже зарегестрированы")
	ErrWrongPassword  = errors.New("неправильный пароль")
	ErrAccessDenied   = errors.New("нет доступа")
	ErrUserNotFound   = errors.New("пользователь не найден")
	ErrChangeRole     = errors.New("невозможно обновить роль")
)
