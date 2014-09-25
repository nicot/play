#include<stdio.h>

int p2(int n) {
	return !((n+ ~0) & n) & !!n;
}

main() {
	int i;
	for (i = -1000000; i < 100000; i++) {
		if (p2(i)) {
			printf("%d\n", i);
		}
	}
}
