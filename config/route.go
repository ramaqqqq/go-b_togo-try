package config

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	userHandler "bookingoto-try/features/user/handler"
	userRepo "bookingoto-try/features/user/repo"
	userUseCase "bookingoto-try/features/user/usecase"
)

type Presenter struct {
	Msg string
}

func Route(r *mux.Router, db *gorm.DB) Presenter {

	userRepo := userRepo.NewUserRepo(db)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler.NewUserHandler(r, userUseCase)

	return Presenter{
		Msg: "<<present>>",
	}
}
