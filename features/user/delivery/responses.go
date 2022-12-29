package delivery

import "cornjobmailer/features/user/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type LoginResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Token  string `json:"token"`
}

type RegistResponse struct {
	Name   string `json:"name" `
	Email  string `json:"email"`
	Status string `json:"status"`
}

type UserResponse struct {
	ID     uint   `json:"id_user"`
	Name   string `json:"name" `
	Email  string `json:"email"`
	Status string `json:"status"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "login":
		cnv := core.(domain.UserCore)
		res = LoginResponse{Name: cnv.Name, Email: cnv.Email, Status: cnv.Status, Token: cnv.Token}
	case "reg":
		cnv := core.(domain.UserCore)
		res = RegistResponse{Name: cnv.Name, Email: cnv.Email, Status: cnv.Status}
	case "user":
		cnv := core.(domain.UserCore)
		res = UserResponse{
			ID:     cnv.ID,
			Name:   cnv.Name,
			Email:  cnv.Email,
			Status: cnv.Status,
		}
	case "all":
		var arr []UserResponse
		cnv := core.([]domain.UserCore)
		for _, val := range cnv {
			arr = append(arr, UserResponse{
				ID:     val.ID,
				Name:   val.Name,
				Email:  val.Email,
				Status: val.Status,
			})
		}
		res = arr
	}
	return res
}
