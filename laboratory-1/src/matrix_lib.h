/*
 * Created on Mon Apr 06 2020
 *
 * Copyright (c) 2020 Konstantin Dobratulin
 */

#ifndef LABORATORY_1_SRC_MATRIX_LIB_H_
#define LABORATORY_1_SRC_MATRIX_LIB_H_

#include <vector>

extern "C" const std::vector<std::vector<float>> matmul(
    const std::vector<std::vector<float>>& A,
    const std::vector<std::vector<float>>& B);

#endif  // LABORATORY_1_SRC_MATRIX_LIB_H_
