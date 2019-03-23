package controller

import (
	"cleanarchitecture/adapter/gateway"
	"cleanarchitecture/adapter/interfaces"
	"cleanarchitecture/domain"
	"cleanarchitecture/usecase"

	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"

)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(conn *gorm.DB, logger interfaces.Logger) *UserController {
	return &UserController {
		Interactor: usecase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				Conn: conn,
			},
			Logger: logger,
		},
	}
}

func (controller *UserController) Create(c interfaces.IContext) error{
	type (
		Request struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Age int `json:"age"`
		}
		Response struct {
			UserID int `json:"user_id"`
		}
	)
	req := Request{}
	c.Bind(&req)
	user := domain.User{Name: req.Name, Email: req.Email, Age: req.Age}

	id, err := controller.Interactor.Add(user)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot add user"))
		c.JSON(500, NewError(500, err.Error()))
		return err
	}
	res := Response{UserID: id}
	c.JSON(201, res)

	return nil
}