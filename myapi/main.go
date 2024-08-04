package main
import (
    "log"
    "net/http"
    _"github.com/yourname/reponame/controllers"
    _"github.com/gorilla/mux"
    _"github.com/yourname/reponame/services"
    "github.com/yourname/reponame/api"
    _"github.com/go-sql-driver/mysql"
    "database/sql"
	"fmt"
	"os"
)

var (
    dbUser = os.Getenv("DB_USER")
    dbPassword = os.Getenv("DB_PASSWORD")
    dbDatabase = os.Getenv("DB_NAME")
    dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
)

func main() {
    db, err := sql.Open("mysql", dbConn)
    if err != nil {
        log.Println("fail to connect DB")
        return
    }

    r := api.NewRouter(db)

    log.Println("server start at port 8080")
    // ListenAndServe 関数にて、サーバーを起動
    log.Fatal(http.ListenAndServe(":8080", r))
}