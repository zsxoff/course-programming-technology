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
    const unsigned int A_ROWS = A.size();
    const unsigned int A_COLS = A[0].size();

    const unsigned int B_ROWS = B.size();
    const unsigned int B_COLS = B[0].size();

    // TODO(zsxoff): Possibly check sizes.

    // Synonims to sizes.
    const size_t L = A_ROWS;
    const size_t M = A_COLS;
    const size_t N = B_COLS;

    // Init new vector with sizes M rows, N cols.
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
