go build -tags=name

go test -tags=integration
go test -v -tags integration
go test -v -tags integration,db
