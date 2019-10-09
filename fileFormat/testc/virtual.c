// 原文：https://blog.csdn.net/china_jeffery/article/details/79626185
#include <windows.h>
#include <stdio.h>

int main() {
	SIZE_T size = 1 << 30; // 1G
	
	// 预定1G空间
	char *pVirtualAdd = (char *) VirtualAlloc(NULL, size, MEM_RESERVE, PAGE_READWRITE);
	if (pVirtualAdd == NULL)  {
		printf("Reserve 1G failed.\n");
		return 1;
	}
	
	// 验证分配粒度是不是64K
	int n = (long)pVirtualAdd % (64*1024);
	if (n == 0) {
		printf("64K\n");
	}
	
	printf("Reserve 1G\n");
	getchar();
	
	if (VirtualAlloc(pVirtualAdd, size, MEM_COMMIT, PAGE_READWRITE) == NULL) {
		printf("Commit 1G failed.\n");
		return 1;
	}
	
	printf("Alloc 1G\n");
	getchar();
	
	// 页面大小为4K，访问2560个页面
	for (int i = 0; i < 2560; i++) {
		char *p = pVirtualAdd + i * (4*1024);
		*p = 'A';
	}
	
	printf("Use 10M\n");
	getchar();
	
	return 0;
}