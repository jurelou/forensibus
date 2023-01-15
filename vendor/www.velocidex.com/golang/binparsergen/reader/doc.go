package reader

// This reader makes using the generated code more efficient since it
// caches entire pages from the file into memory locally. Thus
// avoiding the system call overhead for parsing fields in the same
// struct.
