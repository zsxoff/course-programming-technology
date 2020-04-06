/*
 * Created on Mon Apr 06 2020
 *
 * Copyright (c) 2020 Konstantin Dobratulin
 */

#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

#include "./matrix_lib.h"

std::vector<std::vector<float>> fromfile(const std::string& filename,
                                         const char delim) {
    // Init result vector.
    std::vector<std::vector<float>> result;

    // Parse file.
    std::ifstream filestream(filename);
    std::string buffer_line;

    while (std::getline(filestream, buffer_line)) {
        // Parse line to buffer vector.
        std::stringstream string_stream(buffer_line);
        std::string buffer_item;
        std::vector<float> buffer_floats;

        while (std::getline(string_stream, buffer_item, delim)) {
            buffer_floats.push_back(std::stof(buffer_item));
        }

        // Push back buffer to result vector.
        result.push_back(buffer_floats);
    }

    return result;
}

void print_matrix(const std::vector<std::vector<float>>& arr) {
    for (auto& row : arr) {
        for (auto& element : row) {
            std::cout << element << "\t";
        }
        std::cout << std::endl;
    }
}

int main(int argc, char* argv[]) {
    // ------------------------------------------------------------------------
    // Get files paths.

    if (argc < 3) {
        std::cout << "! Arguments error !" << std::endl;
        std::cout << std::endl;
        std::cout << "Please specify files with matrices as args:" << std::endl;
        std::cout << "./app matrix_A.txt matrix_B.txt" << std::endl;
        std::cout << std::endl;

        return -1;
    }

    const std::string p1 = argv[1];
    const std::string p2 = argv[2];

    // ------------------------------------------------------------------------
    // Get matrices from file.

    const char delimiter = ' ';

    const std::vector<std::vector<float>> A = fromfile(p1, delimiter);

    const std::vector<std::vector<float>> B = fromfile(p2, delimiter);

    // ------------------------------------------------------------------------
    // Print matrix.

    const size_t a_rows = A.size();
    const size_t a_cols = A[0].size();

    const size_t b_rows = B.size();
    const size_t b_cols = B[0].size();

    std::cout << "Matrix A: " << a_rows << "x" << a_cols << std::endl;
    std::cout << std::endl;
    print_matrix(A);
    std::cout << std::endl;

    std::cout << "Matrix B: " << b_rows << "x" << b_cols << std::endl;
    std::cout << std::endl;
    print_matrix(B);
    std::cout << std::endl;

    // ------------------------------------------------------------------------
    // Get result vector.

    const std::vector<std::vector<float>> C = matmul(A, B);

    // ------------------------------------------------------------------------
    // Print result.

    const size_t c_rows = C.size();
    const size_t c_cols = C[0].size();

    std::cout << "Result matrix: " << c_rows << "x" << c_cols << std::endl;
    std::cout << std::endl;
    print_matrix(C);
    std::cout << std::endl;

    return 0;
}
