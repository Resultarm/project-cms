package http

import (
	"backend/db"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("secret-key")

func InitServer(port string, db *db.Db) {
	dbCollections = db

	r := gin.Default()

	// Access cors - All
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3008", "http://localhost:3006"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "jwt-token"},
		AllowCredentials: true,
	}))

	r.GET("/read-logo", getImage)

	// 5 Endpoints
	// Link tree
	r.GET("/read-linktree", read)

	// Resultarmind
	r.GET("/read-menu", getMenu)
	r.GET("/read-access-link", getAccessLink)
	r.GET("/read-banner", getBanner)
	r.GET("/read-section", getSection)
	r.GET("/read-cards", getCards)
	r.GET("/read-multi-sections", getMultiSection)
	r.GET("/read-multi-sections-item", getMultiSectionItem)
	//r.GET("/read-team", getTeam)
	r.GET("/read-members", getMembers)
	r.GET("/read-team", getTeam)
	r.GET("/read-questions", getQuestions)
	r.GET("/read-question", getQuestion)
	r.GET("/read-footer", getFooter)

	r.Static("/media", "media")

	r.POST("/login", login)

	routes := r.Group("/")
	routes.Use(Auth)
	{
		routes.GET("/verify", verify)

		routes.POST("/create-logo", createImage)
		routes.POST("/create-linktree", createLinksData)
		routes.PUT("/update-linktree/:id", update)
		routes.DELETE("/delete-linktree/:id", delete)
		routes.DELETE("/delete-logo", deleteImage)

		// Resultarmind
		routes.POST("/create-menu", createUpdateMenu)
		routes.POST("/create-access-link", createUpdateAccessLink)
		routes.POST("/create-banner", createUpdateBanner)
		routes.POST("/create-section", createUpdateSection)
		routes.POST("/create-cards", createUpdateCards)
		routes.POST("/create-multi-sections", createUpdateMultiSection)
		routes.POST("/create-multi-sections-item", createMultiSectionItem)
		routes.POST("/create-member", createMember)
		routes.POST("/create-team", createUpdateTeam)
		routes.POST("/create-questions", createUpdateQuestions)
		routes.POST("/create-question", createQuestion)
		routes.POST("/create-footer", createUpdateFooter)

		routes.DELETE("/delete-member/:id", deleteMember)
		routes.DELETE("/delete-multi-section-item/:id", deleteMultiSectionItem)
		routes.DELETE("/delete-question/:id", deleteQuestion)
	}

	r.Run(fmt.Sprintf(":%s", port))
}
