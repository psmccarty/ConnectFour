test:
	go test .

cover:
	go test -coverprofile=coverFile.out
	go tool cover -html=coverFile.out

run:
	go run connectfour.go

clean:
	del *~
	del coverFile.out
