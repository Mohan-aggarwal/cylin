//
// Created by prabhakardev on 2022-02-14.
//
#include <iostream>

extern "C" void println(int x){
    std::cout<<x<<std::endl;
}

extern "C" void println(float x){
    std::cout<<x<<std::endl;
}
