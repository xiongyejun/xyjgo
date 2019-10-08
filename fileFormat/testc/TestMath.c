#include <stdio.h>

__declspec(dllimport) double Sub(double a, double b);


int main() {
    double result = Sub(3.0, 2.0);
    printf("Result = %f\n", result);
    return 0;
}

// cl /c TestMath.c
// link TestMath.obj Math.lib