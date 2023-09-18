package route

import (
	"database/sql"

	controller "github.com/Taehoya/pocket-mate/pkg/controllers/trip"
	repository "github.com/Taehoya/pocket-mate/pkg/repositories/trip"
	usecase "github.com/Taehoya/pocket-mate/pkg/usecases/trip"
	"github.com/gin-gonic/gin"
)

func NewTripRouter(db *sql.DB, group *gin.RouterGroup) {
	tripReository := repository.NewTripRepository(db)
	tripUseCase := usecase.NewTripUseCase(tripReository)
	tripController := controller.NewTripController(tripUseCase)

	group.GET("/trip", tripController.GetTripAll)
}
