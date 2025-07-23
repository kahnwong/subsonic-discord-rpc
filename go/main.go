package main

import (
	"fmt"
	"github.com/kahnwong/rich-go/client"
	"os"
	"time"
)

func main() {
	err := client.Login(os.Getenv("DISCORD_APP_ID"))
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(client.Activity{
		ActivityType: client.ActivityTypes.Listening,
		State:        "Heyy!!!",
		Details:      "I'm running on rich-go :)",
		SmallImage:   "https://github.com/kahnwong/dashboard-icons/blob/master/rpc/intellij.png?raw=true",
	})

	if err != nil {
		panic(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 10)
}
