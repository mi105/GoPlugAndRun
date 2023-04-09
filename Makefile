BIN_DIR := bin

all: demo

demo: demo.go cat_plugin.so dog_plugin.so
	go build -o $(BIN_DIR)/demo demo.go

cat_plugin.so: cat_plugin.go
	go build -buildmode=plugin -o $(BIN_DIR)/cat_plugin.so cat_plugin.go

dog_plugin.so: dog_plugin.go
	go build -buildmode=plugin -o $(BIN_DIR)/dog_plugin.so dog_plugin.go

clean:
	rm -r $(BIN_DIR)
