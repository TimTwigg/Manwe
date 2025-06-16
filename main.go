package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	routes "github.com/TimTwigg/EncounterManagerBackend/server"
	io "github.com/TimTwigg/EncounterManagerBackend/utils/io"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	dashboard "github.com/supertokens/supertokens-golang/recipe/dashboard"
	dashboardmodels "github.com/supertokens/supertokens-golang/recipe/dashboard/dashboardmodels"
	emailpassword "github.com/supertokens/supertokens-golang/recipe/emailpassword"
	session "github.com/supertokens/supertokens-golang/recipe/session"
	thirdparty "github.com/supertokens/supertokens-golang/recipe/thirdparty"
	tpmodels "github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	usermetadata "github.com/supertokens/supertokens-golang/recipe/usermetadata"
	supertokens "github.com/supertokens/supertokens-golang/supertokens"
)

func cleanup() {
	asset_utils.CloseDB(asset_utils.DB)
	logger.Init("Database closed.")
	logger.Init("Server stopped.")
}

func main() {
	logger.Init("Reading environment variables...")
	api_key, err := io.GetEnvVar("SUPERTOKENS_API_KEY")
	if err != nil {
		logger.Error("Error reading environment variables: " + err.Error())
		os.Exit(1)
		return
	}
	google_client_id, err := io.GetEnvVar("GOOGLE_CLIENT_ID")
	if err != nil {
		logger.Error("Error reading environment variables: " + err.Error())
		os.Exit(1)

		return
	}
	google_client_secret, err := io.GetEnvVar("GOOGLE_CLIENT_SECRET")
	if err != nil {
		logger.Error("Error reading environment variables: " + err.Error())
		os.Exit(1)
		return
	}

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
			APIKey:        api_key,
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
			thirdparty.Init(&tpmodels.TypeInput{
				SignInAndUpFeature: tpmodels.TypeInputSignInAndUp{
					Providers: []tpmodels.ProviderInput{
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "google",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     google_client_id,
										ClientSecret: google_client_secret,
									},
								},
							},
						},
					},
				},
			}),
			dashboard.Init(&dashboardmodels.TypeInput{
				Admins: &[]string{
					"tim@twiggusa.com",
				},
			}),
			usermetadata.Init(nil),
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

	if err := http.ListenAndServe("localhost:8080", routes.CORSMiddleware(session.VerifySession(nil, http.HandlerFunc(routes.HandleRoute)))); err != nil {
		asset_utils.CloseDB(asset_utils.DB)
		log.Fatal(err)
	}
}
