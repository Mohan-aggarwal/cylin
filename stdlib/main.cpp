//
// Created by prabhakardev on 2022-02-14.
//
#include <iostream>
using namespace std;

typedef union {int a; float b;} C;

int main(){
    C c = 2.0f;
    cout<<c.b<<endl;
    return 0;
}
