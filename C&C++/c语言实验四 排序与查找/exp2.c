#include <stdio.h>
#include <stdlib.h>

// 合并两个子数组
void merge(long long int arr[], int left, int mid, int right) {
    int n1 = mid - left + 1;
    int n2 = right - mid;

    // 创建临时数组
    long long int* L = (long long int*)malloc(n1 * sizeof(long long int));
    long long int* R = (long long int*)malloc(n2 * sizeof(long long int));

    // 拷贝数据到临时数组 L[] 和 R[]
    for (int i = 0; i < n1; i++)
        L[i] = arr[left + i];
    for (int j = 0; j < n2; j++)
        R[j] = arr[mid + 1 + j];

    // 合并临时数组到 arr[]
    int i = 0, j = 0, k = left;
    while (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
            arr[k] = L[i];
            i++;
        }
        else {
            arr[k] = R[j];
            j++;
        }
        k++;
    }

    // 拷贝剩余元素
    while (i < n1) {
        arr[k] = L[i];
        i++;
        k++;
    }
    while (j < n2) {
        arr[k] = R[j];
        j++;
        k++;
    }

    // 释放内存
    free(L);
    free(R);
}

// 递归实现归并排序
void mergeSort(long long int arr[], int left, int right) {
    if (left < right) {
        int mid = left + (right - left) / 2;
        mergeSort(arr, left, mid);
        mergeSort(arr, mid + 1, right);
        merge(arr, left, mid, right);
    }
}

int main() {
    int n;

    // 读取整数 n
    scanf("%d", &n);

    // 分配内存以保存 n 个长整型整数
    long long int* numbers = (long long int*)malloc(n * sizeof(long long int));
    if (numbers == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return 1;
    }

    // 读取 n 个整数
    for (int i = 0; i < n; i++) {
        scanf("%lld", &numbers[i]);
    }

    // 使用归并排序
    mergeSort(numbers, 0, n - 1);

    // 输出排序后的结果
    for (int i = 0; i < n; i++) {
        if (i > 0) printf(" ");
        printf("%lld", numbers[i]);
    }
    printf("\n");

    // 释放内存
    free(numbers);

    return 0;
}