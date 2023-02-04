#ifndef __STUBS_H
#define __STUBS_H

#if defined(__MINGW32__) || defined(WIN32) || defined(__MINGW64__)

int pipe(int *fd);

int fcntl(int fd, int a, int b);

#define F_SETFD 0

#endif
#endif
