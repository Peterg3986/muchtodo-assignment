// @title           MuchToDo API
// @version         1.0
// @description     This is an API for MuchToDo application with user authentication.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support - Innocent
// @contact.url    https://github.com/Innocent9712
// @contact.email  innocent@altschoolafrica.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description "Type 'Bearer' followed by a space and a JWT token."
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/auth"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/cache"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/config"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/database"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/handlers"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/middleware"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/routes"
	"github.com/Innocent9712/much-to-do/Server/MuchToDo/internal/logger"

	// Swagger imports
	_ "github.com/Innocent9712/much-to-do/Server/MuchToDo/docs" // This is required for swag to find your docs
)

const usernameCacheSentinelKey = "username_cache_initialized"
const usernameCacheTTL = 24 * time.Hour

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

package main

import (
    "context"
    "log"
    "muchtodo/internal/config"
)

func main() {
    // Connect to MongoDB
    client := config.ConnectMongo()
    defer func() {
        if err := client.Disconnect(context.Background()); err != nil {
            log.Fatalf("Error disconnecting from MongoDB: %v", err)
        }
    }()

    // Start your server or routes here
    log.Println("🚀 MuchToDo API running on :8080")
    // Example: start Gin server or whatever you use
    // router := setupRouter()
    // router.Run(":8080")
}
