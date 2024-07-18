package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponseRegister struct {
	Valid bool `json:"valid"`
}

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ApiResponseRegister{Valid: true})
	}
}

// Register recieves a RequestPayloadRegister, then checks if the username is valid,
// inserts into user_info, generates UserDataSession, then sends a ResponseRegisterLogin
//func Register(db *dbHandler.DBHandler) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var registerPayload types.RequestPayloadRegister
//		registerResponse := types.ResponseRegisterLogin{
//			Valid: false,
//		}
//
//		// Bind request body
//		if err := ctx.BindJSON(&registerPayload); err != nil {
//			fmt.Printf("Error binding json: %+v\n", err)
//			ctx.JSON(http.StatusInternalServerError, registerResponse)
//			return
//		}
//		fmt.Printf("%+v\n", registerPayload)
//
//		// Check if username currently exists
//		_, err := db.GetUserIDByUsername(registerPayload.Username)
//		if err != nil {
//			if !errors.Is(err, pgx.ErrNoRows) {
//				fmt.Printf("Error checking username validity: %+v\n", err)
//				ctx.JSON(http.StatusInternalServerError, registerResponse)
//				return
//			}
//		} else {
//			fmt.Printf("Username '%s' already exists", registerPayload.Username)
//			ctx.JSON(http.StatusUnauthorized, registerResponse)
//			return
//		}
//
//		// Insert User
//		userID, err := db.InsertUser()
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, registerResponse)
//			fmt.Printf("Error inserting user: %+v\n", err)
//			return
//		}
//
//		UserData := types.UserData{
//			UserID:    userID,
//			Username:  registerPayload.Username,
//			Password:  registerPayload.Password,
//			FirstName: registerPayload.FirstName,
//			LastName:  registerPayload.LastName,
//		}
//
//		// Insert UserData
//		if err := db.InsertUserData(UserData); err != nil {
//			ctx.JSON(http.StatusInternalServerError, registerResponse)
//			fmt.Printf("Error inserting user: %+v\n", err)
//			return
//		}
//
//		// create and insert game session
//		// TODO: randomly select spellID, currently hardcoded
//		gameSession := types.GameSession{
//			GameSessionID: types.GameSessionID(fmt.Sprintf("%d", userID)),
//			UserID:        userID,
//			SpellID:       1,
//			Rounds:        0,
//		}
//		if err := db.InsertGameSession(gameSession); err != nil {
//			ctx.JSON(http.StatusInternalServerError, registerResponse)
//			fmt.Printf("error inserting game session: %+v", err)
//			return
//		}
//
//		accessToken, err := auth.CreateJWTFromUserID(userID)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, registerResponse)
//			return
//		}
//
//		registerResponse = types.ResponseRegisterLogin{
//			Valid: true,
//			Tokens: types.AllTokens{
//				AccessToken:  types.AccessToken{AccessToken: accessToken},
//				RefreshToken: types.RefreshToken{RefreshToken: accessToken},
//			},
//			UserData: types.ResponseUserData{
//				UserID:    UserData.UserID,
//				Username:  UserData.Username,
//				FirstName: UserData.FirstName,
//				LastName:  UserData.LastName,
//			},
//		}
//
//		ctx.JSON(http.StatusCreated, registerResponse)
//
//		// Backup
//		//cmd := exec.Command("./backup.sh")
//		//if err := cmd.Run(); err != nil {
//		//	log.Printf("cError backing up Postgres: %+v\n", err)
//		//}
//	}
//}
