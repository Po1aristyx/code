#include <stdio.h>
#include <unistd.h> // for pipe, fork, read, write, close
#include <string.h> // for strcpy
#include <stdlib.h> // for exit


int main()
{   int x,fd[2];
    char buf[30],s[30];
    pipe(fd);
    while ((x=fork())==-1); // 1. 潜在的繁忙等待和多次fork尝试
    if (x==0)
    {
        close(fd[0]);
        printf("Child Process!\n");
        strcpy(buf,"This is an example\n");
        write(fd[1],buf,30); // 2. 写入固定30字节
        exit(0); // 3. 子进程写完后没有关闭fd[1]就退出了
    }
    else{
        close(fd[1]);
        printf("Parent Process!\n");
        read(fd[0],s,30); // 4. 父进程读取
        printf("%s\n",s); // 5. 打印读取的内容
        // 6. 父进程没有关闭fd[0]
        // 7. 父进程没有wait子进程
    }
}