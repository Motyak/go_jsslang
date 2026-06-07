
# run the tests
GO111MODULE=off go run -tags tests ./test/

# run the main
GO111MODULE=off go run ./main/

# build the main
GO111MODULE=off go build -o main.elf ./main/

--- for convenience

source env.sh

gorun test/tests.go
gorun test/tommystr

cd test
gorun tommystr
