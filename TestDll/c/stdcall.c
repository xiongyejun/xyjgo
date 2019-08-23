#include "go.h"

GoInt __stdcall sum(GoInt a, GoInt b) {
	return Sum(a, b);
}

GoString __stdcall getStr() {
	return GetStr();
}