# Go programs link with C/C++ programs

- C and Go together in the same environment. While it is possible,
it requires special interface software and can give up the memory safety and stack management properties that Go provides. 
It is important to be cautious when using C libraries in Go code as it introduces an element of risk.
Three Go compiler implementations are supported by the Go team - gc, gccgo, and gollvm. 
The cgo program provides a "foreign function interface" to allow safe calling of C libraries from Go code.
SWIG extends this capability to C++ libraries. However, linking code from these compilers directly with GCC/LLVM-compiled
C or C++ programs requires an understanding of the calling conventions for all languages involved and caution for stack limits when calling C or C++ from Go.

# Goroutines intead of threads

- Goroutines are a way to make concurrency easy to use by multiplexing independently executing functions onto a set of threads.
  When a coroutine blocks, the runtime automatically moves other coroutines on the same operating system thread to a different,
  runnable thread so they won't be blocked. Goroutines are very cheap and have little overhead beyond the memory for the stack,
  which is just a few kilobytes. The CPU overhead averages about three cheap instructions per function call. Resizable,
  bounded stacks are used to make the stacks small, and the runtime grows and shrinks the memory for storing the stack automatically.
  It is practical to create hundreds of thousands of goroutines in the same address space.
