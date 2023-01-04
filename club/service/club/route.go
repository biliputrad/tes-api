package club

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"test-api/club/service/club/repository"
	"test-api/club/service/club/usecase"
)

func RouteClub(db *gorm.DB, routerGroup *gin.RouterGroup) {
	clubRepository := repository.NewClubRepository(db)
	clubUseCase := usecase.NewClubUseCase(clubRepository)
	handlerClub := NewClubHandler(clubUseCase)

	routerGroup.POST("/create", handlerClub.CreateClubHandler)
	routerGroup.PUT("/update", handlerClub.UpdateRecordMatchHandler)
	routerGroup.GET("/get", handlerClub.GetAllClubHandler)
	routerGroup.POST("/contain-letter", handlerClub.ContainLetterHandler)
}
