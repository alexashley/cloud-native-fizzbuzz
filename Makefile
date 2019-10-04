MAKEFLAGS += --silent

default:
	echo "No default"

build:
	cd counting && go build -o ../bin/counting 
