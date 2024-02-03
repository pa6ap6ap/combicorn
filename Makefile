coverage.html: combicorn.test
	go tool cover -html=cover.out -o coverage.html

combicorn.test: combicorn
	go test -coverprofile cover.out .

combicorn: *.go
	go build .
