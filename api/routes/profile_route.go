package route 

import(
	"time"

    "github.com/PARMESHWARPANWAR/dev-tinder/api/controller"
	"github.com/PARMESHWARPANWAR/dev-tinder/bootstrap"
	"github.com/PARMESHWARPANWAR/dev-tinder/domain"
	"github.com/PARMESHWARPANWAR/dev-tinder/mongo"
	"github.com/PARMESHWARPANWAR/dev-tinder/repository"
	"github.com/PARMESHWARPANWAR/dev-tinder/usecase"
	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
} 