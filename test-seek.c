#include<stdio.h>
#include <unistd.h>
#include <sys/stat.h>
#include <fcntl.h>
#include<stdlib.h>
 #include <errno.h>
 #include <string.h>

int main() {

    FILE *fptr = fopen("path.txt", "r");
    char path[100];
    fgets(path, 100, fptr);

    int fd = open(path, O_RDONLY);
    off_t offset = lseek(fd, SEEK_SET, 22);

    if ((offset < 0)) {
        printf("seek failed: %s\n", strerror(errno));
        return offset;
    }

    char *data = (char*)malloc(20 * sizeof(char));
    ssize_t n = read(fd, data, 20);

    if (n > 0) {
        printf("data read: %s\n", data);
    } else {
        printf("could not read data\n");
        return -1;
    }

    close(fd);

    return 0;
}