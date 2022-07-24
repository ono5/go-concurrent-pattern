package main

import (
	"context"
	"fmt"
	"withvalue/session"
)

type key int

var userKey key = 0

func main() {
	ctx := context.Background()
	ctx = session.WithUser(ctx, 1)

	// Becase returing interface type, we need to change it into int type
	userID := session.GetUserID(ctx)
	if userID == nil {
		fmt.Println("Not Logged in!")
		return
	}

	fmt.Println(*userID)

	uId := ctx.Value(userKey)
	fmt.Println(uId)
}
