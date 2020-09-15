# :rainbow: `template-github-app-golang` ![](https://github.com/swinton/template-github-app-golang/workflows/Build/badge.svg)
> Template for GitHub Apps in Golang

## About

This is a template repo for quickly getting started building GitHub Apps in Golang.

## Getting started

1. [Generate](../../generate) a repo from this template
1. Configure your local environment, providing values for the environment variables listed in [`.envrc.example`](.envrc.example)
1. Run app, e.g. using [`reflex`](https://github.com/cespare/reflex): `reflex -s go run main.go`
1. Expose the app to the interwebs, e.g. using [`ngrok`](https://ngrok.com): `ngrok http 8000`
    - Or, if you have an [ngrok configuration file](https://ngrok.com/docs#config) for a tunnel named `main`, then: `ngrok start main`
1. Add logic to [the handler](main.go) as required
