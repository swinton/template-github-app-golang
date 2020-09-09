package main

import (
	"log"

	"github.com/google/go-github/github"
	"github.com/swinton/go-probot/probot"
)

func main() {
	probot.HandleEvent("issues", func(ctx *probot.Context) error {
		// Because we're listening for "issues" we know the payload is a *github.IssuesEvent
		event := ctx.Payload.(*github.IssuesEvent)
		log.Printf("🌈 Got issues %+v\n", event)

		return nil
	})
	probot.Start()
}
