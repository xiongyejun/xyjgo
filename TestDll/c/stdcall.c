#include "go.h"
#include "stdlib.h"

GoInt __stdcall sum(GoInt a, GoInt b) {
	return Sum(a, b);
}

GoString __stdcall getStr() {
	return GetStr();
}

GoSlice __stdcall reurnSlice() {
	return ReurnSlice();
}

void __stdcall cfree(void *p) {
	Free(p);
}