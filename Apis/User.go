package Apis

mport (
 "net/http"
 "log"
 "fmt"
 "strconv"
 "gopkg.in/gin-gonic/gin.v1"
 . "newland/models"
)

IndexApi(c *gin.Context) {
 c.String(http.StatusOK, "It works")
}
