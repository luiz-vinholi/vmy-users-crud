package usecases

import (
	"vmytest/src/domain/entities"
	"vmytest/src/infra/repositories"
)

type GetUsersResult struct {
	Total int64           `json:"total"`
	Users []entities.User `json:"users"`
}

type Pagination struct {
	Limit  *int
	Offset *int
}

func GetUsers(pagination Pagination) (*GetUsersResult, error) {
	pag := handlePagination(pagination)
	result, err := usersRepo.GetUsers(pag)
	if err != nil {
		return nil, err
	}
	var eusers []entities.User
	for _, user := range result.Users {
		euser := entities.User{
			Id:        user.Id.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			BirthDate: user.BirthDate,
		}
		if user.Address != nil {
			euser.Address = &entities.Address{
				Street:  user.Address.Street,
				City:    user.Address.City,
				State:   user.Address.State,
				Country: user.Address.Country,
			}
		}
		euser.SetAge()
		eusers = append(eusers, euser)
	}
	response := &GetUsersResult{
		Total: int64(result.Total),
		Users: eusers,
	}
	return response, nil
}

func handlePagination(pagination Pagination) repositories.Pagination {
	var limit int
	var offset int
	if pagination.Limit == nil {
		limit = 10
	} else {
		limit = *pagination.Limit
	}
	if pagination.Offset == nil {
		offset = 0
	} else {
		offset = *pagination.Offset
	}
	return repositories.Pagination{
		Limit:  limit,
		Offset: offset,
	}
}
