/*
 * Created on Mon Apr 06 2020
 *
 * Copyright (c) 2020 Konstantin Dobratulin
 */

#include "./matrix_lib.h"

const std::vector<std::vector<float>> matmul(
    const std::vector<std::vector<float>>& A,
    const std::vector<std::vector<float>>& B) {
    // Get matrix sizes.
    const size_t L = A.size();     // A Rows
    const size_t M = A[0].size();  // A Cols
    const size_t N = B[0].size();  // B Cols

    // Init new vector with sizes L rows, N cols.
    std::vector<std::vector<float>> C;
    C.assign(L, std::vector<float>(N));

    // Compute matrix multiplication.
    for (size_t i = 0; i < L; ++i) {
        for (size_t j = 0; j < N; ++j) {
            for (size_t k = 0; k < M; ++k) {
                C[i][j] += A[i][k] * B[k][j];
            }
        }
    }

    return C;
}
