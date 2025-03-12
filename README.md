# Goofy Jokes

A collaborative Go project to share dad jokes, knock-knock jokes, and one-liners!

## How to Contribute
1. Fork this repo.
2. Create a branch: `git checkout -b my-joke`.
3. Add your joke to `internal/jokes/dad.go`, `knock.go`, or `oneliners.go`.
4. Test it: `go run cmd/server/main.go` and visit `localhost:8080`.
5. Commit and push: `git commit -m "Add my awesome joke" && git push origin my-joke`.
6. Open a Pull Request!

## Running Locally
- Install Go: [golang.org](https://golang.org/)
- Run: `go run cmd/server/main.go`
- Open `localhost:8080` in your browser.

## Debugging Tips
- Use `fmt.Println` to log values.