CXX=g++
CFLAGS=-Wall -O2

all: lib app

lib: src/matrix_lib.cpp
	${CXX} ${CFLAGS} -shared -fPIC $< -o ./libs/libmatrix.so

app: src/main.cpp
	${CXX} ${CFLAGS} -L ./libs -lmatrix ./src/main.cpp -o app -Wl,-rpath=./libs

clean:
	rm -f ./libs/* ./app

.PHONY: all lib app clean