package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	routes "github.com/TimTwigg/EncounterManagerBackend/server"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	emailpassword "github.com/supertokens/supertokens-golang/recipe/emailpassword"
	session "github.com/supertokens/supertokens-golang/recipe/session"
	supertokens "github.com/supertokens/supertokens-golang/supertokens"
)

func cleanup() {
	asset_utils.CloseDB(asset_utils.DB)
	logger.Init("Database closed.")
	logger.Init("Server stopped.")
}

func main() {

	logger.Init("Loading database...")
	database, err := asset_utils.GetDB()
	if err != nil {
		logger.Error("Error loading database: " + err.Error())
		return
	}
	asset_utils.DB = database

	logger.Init("Initializing Authentication...")
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err = supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "https://st-dev-d95ae1a0-41c5-11f0-b541-0565ea7e6a4b.aws.supertokens.io",
			APIKey:        "s6Lz5lmzArecujwftOaEdVBRGC",
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "Olorin",
			APIDomain:       "http://localhost:8080",
			WebsiteDomain:   "http://localhost:5173",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	logger.Init("Server started on port 8080")
	if err := http.ListenAndServe("localhost:8080", routes.Middleware(http.HandlerFunc(routes.HandleRoute))); err != nil {
		asset_utils.CloseDB(asset_utils.DB)
		log.Fatal(err)
	}
}
