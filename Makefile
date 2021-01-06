APPS     :=	edge operator
TEST_PATH = ./...

all:	$(APPS)

$(APPS):
	go build -o bin/$@ cmd/$@/*.go

clean:
	rm -rf bin/*

test:
	go test $(TEST_PATH)

coverage:
	go test $(TEST_PATH) -coverprofile=c.out || true
	go tool cover -html=c.out
	go tool cover -func=c.out | tail -1
	rm c.out
