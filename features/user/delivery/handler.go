package delivery

import (
	"cornjobmailer/features/user/domain"
	"cornjobmailer/utils/common"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

var validate = validator.New()

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}

	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
	e.GET("/user", handler.ShowAll())
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		er := validate.Struct(input)
		if er != nil {
			if strings.Contains(er.Error(), "min") {
				return c.JSON(http.StatusBadRequest, FailResponse("min. 4 character"))
			} else if strings.Contains(er.Error(), "max") {
				return c.JSON(http.StatusBadRequest, FailResponse("max. 30 character"))
			} else if strings.Contains(er.Error(), "email") {
				return c.JSON(http.StatusBadRequest, FailResponse("must input valid email"))
			}
			return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
		}

		cnv := ToDomain(input)
		res, err := uh.srv.Register(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success register user", ToResponse(res, "reg")))
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := uh.srv.Login(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("password not match."))
			} else if strings.Contains(err.Error(), "aktivasi") {
				return c.JSON(http.StatusBadRequest, FailResponse("belum diaktivasi"))
			} else {
				return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
			}
		} else if strings.Contains(res.Status, "pending") {
			return c.JSON(http.StatusBadRequest, FailResponse("belum diaktivasi"))
		} else if res.ID != 0 {
			res.Token = common.GenerateToken(uint(res.ID))
			return c.JSON(http.StatusAccepted, SuccessResponse("Success to login", ToResponse(res, "login")))
		}

		return nil

	}
}

func (uh *userHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uh.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all data", ToResponse(res, "all")))
	}
}
