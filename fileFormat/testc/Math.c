double Add(double a, double b) {
    return a+b;
}

double Sub(double a, double b) {
    return a-b;
}

double Mul(double a, double b) {
    return a*b;
}

// 生成dll
// cl Math.c /LDd /DEF Math.def
// /LDd    表示生成Debug版的DLL，不加任何参数生成exe可执行文件
// /LD    生成Release版的DLL