package services

import "CloudStorage/pkg/errors"

func (s *Service) DeleteUser(userId int) (err error) {
	_, err = s.Repo.GetUserByID(userId)
	if err != nil {
		if err == errors.ErrDataNotFound {
			return errors.ErrUserNotFound
		}
		return err
	}
	err = s.Repo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return
}
