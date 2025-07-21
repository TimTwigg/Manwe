package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	routes "github.com/TimTwigg/Manwe/server"
	io "github.com/TimTwigg/Manwe/utils/io"
	logger "github.com/TimTwigg/Manwe/utils/log"
	dashboard "github.com/supertokens/supertokens-golang/recipe/dashboard"
	dashboardmodels "github.com/supertokens/supertokens-golang/recipe/dashboard/dashboardmodels"
	emailpassword "github.com/supertokens/supertokens-golang/recipe/emailpassword"
	epmodels "github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	session "github.com/supertokens/supertokens-golang/recipe/session"
	sessmodels "github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	thirdparty "github.com/supertokens/supertokens-golang/recipe/thirdparty"
	tpmodels "github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	usermetadata "github.com/supertokens/supertokens-golang/recipe/usermetadata"
	supertokens "github.com/supertokens/supertokens-golang/supertokens"
)

func cleanup() {
	asset_utils.DBPool.Close()
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
	pool, err := asset_utils.GetDB()
	if err != nil {
		logger.Error("Error loading database: " + err.Error())
		return
	}
	asset_utils.DBPool = pool

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
			emailpassword.Init(&epmodels.TypeInput{
				// Override the SignUp function to add ID to database
				Override: &epmodels.OverrideStruct{
					Functions: func(originalImplementation epmodels.RecipeInterface) epmodels.RecipeInterface {
						originalSignUp := *originalImplementation.SignUp
						(*originalImplementation.SignUp) = func(email, password, tenantId string, userContext supertokens.UserContext) (epmodels.SignUpResponse, error) {
							response, err := originalSignUp(email, password, tenantId, userContext)
							if err != nil {
								return epmodels.SignUpResponse{}, err
							}
							if response.OK != nil {
								user := response.OK.User
								_ = asset_utils.UpsertUser(user.ID)
							}
							return response, nil
						}
						return originalImplementation
					},
				},
			}),
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
				// Override the SignUp function to add ID to database
				Override: &tpmodels.OverrideStruct{
					Functions: func(originalImplementation tpmodels.RecipeInterface) tpmodels.RecipeInterface {
						originalSignInUp := *originalImplementation.SignInUp
						(*originalImplementation.SignInUp) = func(thirdPartyID string, thirdPartyUserID string, email string, oAuthTokens map[string]interface{}, rawUserInfoFromProvider tpmodels.TypeRawUserInfoFromProvider, tenantId string, userContext *map[string]interface{}) (tpmodels.SignInUpResponse, error) {
							response, err := originalSignInUp(thirdPartyID, thirdPartyUserID, email, oAuthTokens, rawUserInfoFromProvider, tenantId, userContext)
							if err != nil {
								return tpmodels.SignInUpResponse{}, err
							}
							if response.OK != nil {
								user := response.OK.User
								_ = asset_utils.UpsertUser(user.ID)
							}
							return response, nil
						}
						return originalImplementation
					},
				},
			}),
			dashboard.Init(&dashboardmodels.TypeInput{
				Admins: &[]string{
					"tim@twiggusa.com",
				},
			}),
			usermetadata.Init(nil),
			session.Init(&sessmodels.TypeInput{
				Override: &sessmodels.OverrideStruct{
					Functions: func(originalImplementation sessmodels.RecipeInterface) sessmodels.RecipeInterface {
						originalCreateNewSession := *originalImplementation.CreateNewSession
						(*originalImplementation.CreateNewSession) = func(userID string, accessTokenPayload, sessionDataInDatabase map[string]interface{}, disableAntiCsrf *bool, tenantId string, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
							_ = asset_utils.UpsertUser(userID)
							return originalCreateNewSession(userID, accessTokenPayload, sessionDataInDatabase, disableAntiCsrf, tenantId, userContext)
						}

						return originalImplementation
					},
				},
			}),
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
		asset_utils.DBPool.Close()
		log.Fatal(err)
	}
}
