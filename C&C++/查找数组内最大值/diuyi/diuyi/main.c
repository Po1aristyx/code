#include <stdio.h>
#include <stdlib.h>
typedef struct Data
{
    int min;
    int max;
}Data;
int MIN;      /* 全局变量 */
int fun1(int a[], int n)
/*通过函数return返回最大值，通过全局变量MIN带回最小值*/
{
    int i, max;
    max = MIN = a[0]; //给最大值最小值赋初值
    for (i = 0; i < n; i++)
    {
        if (a[i] > max) max = a[i];
        if (a[i] < MIN) MIN = a[i];
    }
    return(max);
}
int* fun2(int a[], int n)
/*将最大、最小值放到数组b中，通过return返回*/
{
    static int b[2];

    int i;
    b[0] = b[1] = a[0];  //给最大值最小值赋初值
    for (i = 1; i < n; i++)
    {
        if (a[i] > b[0])	b[0] = a[i];
        if (a[i] < b[1]) 	b[1] = a[i];
    }
    return(b);
}
Data* fun3(int a[], int n)
/*将最大、最小值放到结构体指针p中，通过return返回*/
{
    Data* p;
    int i;
    p = (Data*)malloc(sizeof(Data*));
    p->max = p->min = a[0];
    for (i = 1; i < n; i++)
    {
        if (a[i] > p->max) 	p->max = a[i];
        if (a[i] < p->min)	p->min = a[i];
    }
    return(p);
}
Data fun4(int a[], int n)
/*将最大、最小值放到结构体p中，通过结构体p带回返回值*/
{
    Data p;
    int i;
    p.max = p.min = a[0];      //给最大值最小值赋初值
    for (i = 1; i < n; i++)
    {
        if (a[i] > p.max) 		p.max = a[i];
        if (a[i] < p.min)		p.min = a[i];
    }
    return(p);
}
void fun5(int a[], int n, int* p, int* q)
/*用指针带回返回值，指针p指向最大值，指针q指向最小值*/
{
    int i;
    *p = *q = a[0];
    for (i = 1; i < n; i++)
    {
        if (*p < a[i]) 	*p = a[i];
        if (*q > a[i])	*q = a[i];
    }
}

int main()
{
    int mData[5] = { 2,3,5,4,1 };
    //     在此处定义变量

    //以下为fun1示例，仿照该示例，编写调用fun2-5函数，并用printf输出结果。
    int rmax;
    rmax = fun1(mData, 5);
    printf("fun1:min=%d,max=%d.\n", MIN, rmax);
    //fun2
    int* r2 = fun2(mData, 5);
    printf("fun2: min = %d, max = %d\n", r2[1], r2[0]);
    //fun3
    Data* r3 = fun3(mData, 5);
    printf("fun3: min = %d, max = %d\n", r3->min, r3->max);
    free(r3);
    //fun4
    Data r4 = fun4(mData, 5);
    printf("fun4: min = %d, max = %d\n", r4.min, r4.max);
    //fun5
    int r5min, r5max;
    fun5(mData, 5, &r5min, &r5max);
    printf("fun5: min = %d, max = %d\n", r5min, r5max);


    system("pause");
    exit(0);
}