package main

import "gopkg.in/segmentio/analytics-go.v3"

/*
https://segment.com/docs/connections/sources/catalog/libraries/server/go/quickstart/

*/

var userId = "f6ca424242"

func main() {
	println("Hello segment")
	client := analytics.New("sQrOUWrWBnmWZEFUwNHhSpSJcGdDDDYF")
	defer client.Close()

	// call once to track = on sign-up
	client.Enqueue(analytics.Identify{
		UserId: userId,
		Traits: analytics.NewTraits().
			SetName("Michael Bolton").
			SetEmail("mbolton@example.com").
			Set("plan", "Enterprise").
			Set("friends", 42),
	})

	client.Enqueue(analytics.Track{
		Event:  "Signed Up",
		UserId: userId,
		Properties: analytics.NewProperties().
			Set("plan", "Enterprise"),
	})

	// You’ll want to track events that are indicators of success for your site, like Signed Up, Item Purchased or Article Bookmarked.

	client.Enqueue(analytics.Track{
		Event:  "Article Bookmarked",
		UserId: userId,
		Properties: analytics.NewProperties().
			Set("title", "Snow Fall").
			Set("subtitle", "The Avalanche at Tunnel Creek").
			Set("author", "John Branch"),
	})

}
