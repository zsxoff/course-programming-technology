#!/usr/bin/bash

matrix_A="/matrices/matrix_3x3_1.txt"
matrix_B="/matrices/matrix_3x3_2.txt"

docker build -t "laboratory-1" .

docker run \
--rm \
-it \
-v ${PWD}/matrices:/matrices \
laboratory-1 ${matrix_A} ${matrix_B}