package main

/*
#include <stdio.h>
#include <stdlib.h>

int* localAddr()
{
	int a = 0, b = 0;
	return &b;
}

int* mallocAddr()
{
	int* a = (int*)malloc(sizeof(int));
	*a = 100;
	return a;
}

void addOne(int *a)
{
	*a += 1;
}

void try()
{
	int *a = localAddr();
	*a += 1;
	// addOne(a);
}
*/
import "C"

// https://github.com/golang/go/issues/9733
func main() {
	C.try()
	// fmt.Println(C.localAddr())
	// fmt.Println(C.localAddr())
	// fmt.Println(C.localAddr())
	// fmt.Println(C.localAddr())
	// *C.localAddr() = 0
	// fmt.Println(*C.localAddr())
	// fmt.Println(C.mallocAddr())
	// fmt.Println(*C.mallocAddr())
}
