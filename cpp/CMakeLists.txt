cmake_minimum_required(VERSION 3.16)
set(CMAKE_VERBOSE_MAKEFILE ON)
set(CMAKE_CXX_STANDARD 20)
project(TirePressure VERSION 1.0
        DESCRIPTION "see top level readme."
        LANGUAGES CXX)

# uncomment these lines to enable coverage measurements using gcov
#set(CMAKE_CXX_FLAGS "--coverage")

enable_testing()
add_subdirectory(src)
add_subdirectory(test-catch2)
add_subdirectory(test-doctest)
add_subdirectory(test-gtest)

