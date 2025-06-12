package clerk

import (
	"context"
	"fmt"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/tokens"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context, jwt string) {

	clerk.SetKey(os.Getenv("CLERK_SEC_KEY"))
	sessionToken, err := tokens.Verify(jwt)

	if err != nil {
		c.JSON(403, gin.h{"invalid auth token provided"})
		return
	}
	fmt.Printf("sdölfajs %s", sessionToken.Subject)
}

func PrintUserData() {

	// Each operation requires a context.Context as the first argument.
	ctx := context.Background()

	// Set the API key
	clerk.SetKey(os.Getenv("CLERK_SEC_KEY"))

	// Get
	resource, err := user.Get(ctx, os.Getenv("USER_ID"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*resource.FirstName)
	}
}
