package Routes
import (
	Controllers2 "freshers-bootcamp/day3/exercise2/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers2.GetStudents)
		grp1.POST("user", Controllers2.CreateStudent)
		grp1.GET("user/:id", Controllers2.GetStudentByID)
		grp1.PUT("user/:id", Controllers2.UpdateStudent)
		grp1.DELETE("user/:id", Controllers2.DeleteStudent)
	}
	return r
}