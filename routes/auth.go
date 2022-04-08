package routes

import (
	"github.com/danilopolani/gocialite/structs"
	"learn-gin/config"
	"learn-gin/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Redirect to correct oAuth URL
func RedirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GH"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GH"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/github/callback",
		},
		"google": {
			"clientID":     os.Getenv("CLIENT_ID_G"),
			"clientSecret": os.Getenv("CLIENT_SECRET_G"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{},
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func CallbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	//jika user belum terdaftar, tambahkan ke dalam database
	var newUser = getOrRegisterUser(provider, user)

	c.JSON(200, gin.H{
		"data":    newUser,
		"token":   token,
		"message": "berhasil login",
	})

	// Print in terminal user information
	//fmt.Printf("%#v", token)
	//fmt.Printf("%#v", user)
	//fmt.Printf("%#v", provider)

	// If no errors, show provider name
	//c.Writer.Write([]byte("Hi, " + user.FullName))
}

//set fungsigetOrRegisterUser() -> mengembalikan objek data user
//fungsi cek apakah user sudah terdaftar atau belum
func getOrRegisterUser(provider string, user *structs.User) models.User {
	//set objek cetakan user
	var userData models.User
	config.DB.Where("provider = ? AND social_id = ?", provider, user.ID).First(&userData)

	//jika data tidak ada, create user baru
	if userData.ID == 0 {
		newUser := models.User{
			FullName: user.FullName,
			Email:    user.Email,
			SocialId: user.ID,
			Provider: provider,
			Avatar:   user.Avatar,
		}
		config.DB.Create(&newUser)
		return newUser
	} else {
		return userData
	}
}
