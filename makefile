BINARY_NAME=athleteiqbox
MAIN_GO_PATH=./main.go

all: raspberry upload run

raspberry:
	@echo "Compiling for Raspberry Pi..."
	@GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -o ./build/raspberry/$(BINARY_NAME) $(MAIN_GO_PATH)
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
