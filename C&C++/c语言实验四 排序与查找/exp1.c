#include<stdio.h>
int main() {
    int n = 0;
    scanf("%d", &n);
    int arr[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    int check = 0;
    scanf("%d", &check);
    int left = 0;
    int count = 0;
    int right = n - 1;
    while (left <= right) {
        count++;
        int mid = (left + right) / 2;
        if (arr[mid] > check) {
            right = mid - 1;
        }
        if (arr[mid] < check) {
            left = mid + 1;
        }
        if (arr[mid] == check) {
            printf("%d\n", mid);
            printf("%d\n", count);
            break;
        }
    }
    if (left > right) {
        printf("%d\n", -1);
        printf("%d\n", count);
    }
    return 0;
}