#if defined(__MINGW32__) || defined(WIN32) || defined(__MINGW64__)

int pipe(int *fd) {
    return -1;
}

int fcntl(int fd, int a, int b) {
    return -1;
}

#define F_SETFD 0

#endif
