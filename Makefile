BIN_DIR="bin"
BIN_NAME="family-tree"

clean:
	@go clean
	@rm -rf $(BIN_DIR)

build: clean
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BIN_NAME) .
