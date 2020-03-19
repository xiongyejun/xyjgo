#include "go.h"

struct Sprintf_return __stdcall gosprintf(GoInt p0, GoInt p1, GoInt p2) {
	return Sprintf(p0, p1, p2);
}

GoInt __stdcall sum(GoInt a, GoInt b) {
	return Sum(a, b);
}

void __stdcall cfree(void *p) {
	Free(p);
}