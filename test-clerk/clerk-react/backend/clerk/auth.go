package clerk

import (
	"context"
	"fmt"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

// func GetUserId(c *gin.Context, jwt string) {
// 	clerk.SetKey(os.Getenv("CLERK_SEC_KEY"))
// 	sessionToken, err := clerk.VerifyToken(jwt)

// 	if err != nil {
// 		c.JSON(403, gin.H{"error": "invalid auth token provided"})
// 		return
// 	}
// 	fmt.Printf("sdölfajs %s", sessionToken.Subject)
// }

func PrintUserData() {
	ctx := context.Background()
	clerk.SetKey(os.Getenv("CLERK_SEC_KEY"))

	resource, err := user.Get(ctx, os.Getenv("USER_ID"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*resource.FirstName)
	}
}
