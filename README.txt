
# run the tests
GO111MODULE=off go run -tags tests ./test/

# run the main
GO111MODULE=off go run -tags main ./main/
go run main/*.go

# build the main
GO111MODULE=off go build -o main.elf ./main/
