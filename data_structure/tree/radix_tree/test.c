#include <stdio.h>
//nginx radix tree
int main(void)
{
    int preallocate = 7;
    int mask = 0;
    int inc = 0x80000000;
    int key;
    while (preallocate--) {

        key = 0;
        mask >>= 1;//决定此次初始化的最大深度，例如mask=11100000,则最多初始化前三位
        mask |= 0x80000000;

        printf("mask:%x", mask);
        printf("inc:%x\n", inc);
        do {//根据inc的值添加节点
            key += inc;//随着不断的累加，位数不断像高位移动，直到32位的全部为0，这里key是int，共32位是关键，累加后移动到高位的直接舍弃了
            printf("  k:%x  ", key);

        } while (key);

        inc >>= 1;
        printf("\n");
    }
}