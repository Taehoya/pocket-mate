package route

import (
	"database/sql"

	controller "github.com/Taehoya/pocket-mate/pkg/controllers"
	"github.com/Taehoya/pocket-mate/pkg/repository"
	usecase "github.com/Taehoya/pocket-mate/pkg/usecases"
	"github.com/gin-gonic/gin"
)

func NewTripRouter(db *sql.DB, group *gin.RouterGroup) {
	tripReository := repository.NewTripRepository(db)
	tripUseCase := usecase.NewTripUseCase(tripReository)
	tripController := controller.NewTripController(tripUseCase)

	group.GET("/trip", tripController.GetTripAll)
}
