BINARY_NAME=athleteiqbox
MAIN_GO_PATH=./main.go
DOCKER_IMAGE=athleteiqbox-builder

all: raspberry upload run

raspberry:
	@echo "Compiling for Raspberry Pi using Docker with CGO enabled..."
	@docker run --rm \
	   -v $(PWD):/app \
	   -w /app \
	   -e CGO_ENABLED=1 \
	   ghcr.io/kilianp07/pi-go-build-images:bullseye-arm64-v1.23.0 \
	   go build -o ./build/raspberry/$(BINARY_NAME) $(MAIN_GO_PATH)
	@echo "Done!"

clean:
	@echo "Cleaning up..."
	@rm -rf ./build/raspberry
	@echo "Done!"

upload:
	@echo "Uploading to Raspberry Pi..."
	@scp ./build/raspberry/$(BINARY_NAME) $(SSH_TARGET):AthleteIQBox/
	@scp ./box.json $(SSH_TARGET):AthleteIQBox/
	@echo "Done!"

run:
	@echo "Running on Raspberry Pi..."
	@ssh $(SSH_TARGET) "cd AthleteIQBox && sudo ./$(BINARY_NAME) --conf box.json"
	@echo "Done!"

test:
	@echo "Testing on local..."
	@GOOS=$(HOSTOS) go test -v ./...
	@echo "Done!"

kill:
	@echo "Killing on Raspberry Pi..."
	@ssh $(SSH_TARGET) "sudo pkill $(BINARY_NAME)"
	@echo "Done!"
