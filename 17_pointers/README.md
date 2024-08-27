# Pointers


Go supports pointers, allowing to pass references to values and records within your program

We'll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed to it by values. zeroval will get a copy of ival distinct from the one in the calling function.

zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. The *iptr code in the function body then dereferences the pointer from its memory address to the current value at the address. Assigning a value to a dereferenced pointer changes the value at the referenced address.

The 7i syntax gives the memory address of i, i.e. a pointer to i.

Pointer can be printed to

zeroval dosen't change the i in main, but zeroptr does because it has a reference to the memory address for that variable.





