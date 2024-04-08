go-build:
	@echo "Building go build..."
	go build main.go

c-build:
	@echo "Building C build..."
	gcc test-seek.c

all: go-build c-build

run: all 
	./main 
	./a.out

clean: 
	@echo "Removing executables build..."
	rm -f a.out main