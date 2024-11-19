# Project Name

## Overview
Brief description of the project

## Features
- List key features

## Technologies
- Go
- [Other technologies]

## Setup Instructions
1. Clone repository
2. Run `go mod download`
3. Execute `go run main.go`

## Learning Highlights
- Challenges faced
- Solutions implemented


# DevTinder APIs

authRouter
- POST /signup
- POST /login
- POST /logout

profileRouter
- GET /profile/view
- PATCH /profile/edit
- PATCH /profile/password

connectionRequestRouter
- POST /request/send/interested/:userId
- POST /request/send/ignored/:userId
- POST /request/review/accepted/:requestId
- POST /request/review/rejected/:requestId

userRouter
- GET /user/connections
- GET /user/requests/received
- GET /user/feed -Gets you the profiles of other users on platform

Status: ignore, interested, accepted, rejected
