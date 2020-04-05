/*
 * Created on Mon Apr 06 2020
 *
 * Copyright (c) 2020 Konstantin Dobratulin
 */

#include <iostream>
#include <vector>

#include "./matrix_lib.h"

int main(int argc, char* argv[]) {
    std::vector<std::vector<float>> A;
    std::vector<std::vector<float>> B;

    // TODO(zsxoff): Remove hard-coded matrix.

    A.push_back(std::vector<float>{42, 8, 6, 4});
    A.push_back(std::vector<float>{10, 8, 6, 4});
    A.push_back(std::vector<float>{10, 8, 6, 4});

    B.push_back(std::vector<float>{10, 8});
    B.push_back(std::vector<float>{10, 8});
    B.push_back(std::vector<float>{10, 8});
    B.push_back(std::vector<float>{10, 8});

    // Get result vector.
    const std::vector<std::vector<float>> C = matmul(A, B);

    for (auto& row : C) {
        for (auto& element : row) {
            std::cout << element << " ";
        }
        std::cout << std::endl;
    }

    return 0;
}
