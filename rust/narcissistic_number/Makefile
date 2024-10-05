# Variables
CARGO = cargo

# Default target: just build the project
all: build

# Build the project
build:
	$(CARGO) build

# Run the project (useful for when you have a main.rs)
run:
	$(CARGO) run

# Run the tests
test:
	$(CARGO) test

# Run benchmarks
bench:
	$(CARGO) bench

# Clean the project (remove target directory)
clean:
	$(CARGO) clean

# Rebuild the project from scratch
rebuild: clean build
