#define _GNU_SOURCE

#include <limits.h>
#include <errno.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>
#include <fcntl.h>
#include <signal.h>
#include <sys/stat.h>
#include <sys/socket.h>
#include <sys/sendfile.h>
#include <netinet/in.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>

int main(const int argc, const char *const argv[])
{
    const pid_t orig_parent = getppid();
    const struct sockaddr_in self_addr = {
        .sin_family = AF_INET,
        .sin_port = htons(9000),
        .sin_addr = {
            htonl(INADDR_LOOPBACK)
        }
    };
    const int listen_sk = socket(AF_INET, SOCK_STREAM, 0);
    const int public_dir = open(argv[1], O_PATH);
    struct sockaddr client_addr;
    socklen_t addr_len;

    if (argc < 2) {
        dprintf(STDERR_FILENO,
            "usage: %s <dir to serve files from>\n",
            argv[0]);
        return 1;
    }

    if (bind(listen_sk, (struct sockaddr *)&self_addr, sizeof(self_addr))) {
        perror("bind");
        return 1;
    }

    if (listen(listen_sk, 8)) {
        perror("listen");
        return 1;
    }

    printf("[+] Listening; press Ctrl-C to exit...\n");

    while (orig_parent == getppid()) {
        const int sk = accept(listen_sk, &client_addr, &addr_len);

        if (sk < 0) {
            perror("[-] accept");
            break;
        }

        printf("[+] Accepted Connection\n");

        serve_file(sk, public_dir);
        close(sk);
    }

    return 0;
}
