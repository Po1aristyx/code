#include <stdio.h>
#include <stdbool.h>

bool isInsertionSort(int original[], int middle[], int n) {
    int i = 0;
    while (i < n - 1 && middle[i] <= middle[i + 1]) {
        i++;
    }
    for (int j = i + 1; j < n; j++) {
        if (original[j] != middle[j]) {
            return false;
        }
    }
    return true;
}

void nextInsertionSortStep(int middle[], int n) {
    int i = 0;
    while (i < n - 1 && middle[i] <= middle[i + 1]) {
        i++;
    }
    int key = middle[i + 1];
    int j = i;
    while (j >= 0 && middle[j] > key) {
        middle[j + 1] = middle[j];
        j--;
    }
    middle[j + 1] = key;
}

void merge(int arr[], int left, int mid, int right) {
    int n1 = mid - left + 1;
    int n2 = right - mid;
    int L[n1], R[n2];
    for (int i = 0; i < n1; i++) {
        L[i] = arr[left + i];
    }
    for (int j = 0; j < n2; j++) {
        R[j] = arr[mid + 1 + j];
    }
    int i = 0, j = 0, k = left;
    while (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
            arr[k++] = L[i++];
        }
        else {
            arr[k++] = R[j++];
        }
    }
    while (i < n1) {
        arr[k++] = L[i++];
    }
    while (j < n2) {
        arr[k++] = R[j++];
    }
}

void nextMergeSortStep(int middle[], int n, int currentSize) {
    for (int left = 0; left < n - currentSize; left += 2 * currentSize) {
        int mid = left + currentSize - 1;
        int right = (left + 2 * currentSize - 1 < n - 1) ? (left + 2 * currentSize - 1) : (n - 1);
        merge(middle, left, mid, right);
    }
}

int main() {
    int n;
    scanf("%d", &n);
    int original[n], middle[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &original[i]);
    }
    for (int i = 0; i < n; i++) {
        scanf("%d", &middle[i]);
    }

    if (isInsertionSort(original, middle, n)) {
        printf("Insertion Sort\n");
        nextInsertionSortStep(middle, n);
    }
    else {
        printf("Merge Sort\n");
        int currentSize = 1;
        while (currentSize < n) {
            bool sorted = true;
            for (int i = 0; i < n; i += 2 * currentSize) {
                int left = i;
                int mid = (i + currentSize - 1 < n) ? i + currentSize - 1 : n - 1;
                int right = (i + 2 * currentSize - 1 < n) ? i + 2 * currentSize - 1 : n - 1;
                if (right >= n) break;
                for (int j = left; j < right; j++) {
                    if (middle[j] > middle[j + 1]) {
                        sorted = false;
                        break;
                    }
                }
                if (!sorted) break;
            }
            if (!sorted) break;
            currentSize *= 2;
        }
        nextMergeSortStep(middle, n, currentSize);
    }

    for (int i = 0; i < n; i++) {
        if (i > 0) printf(" ");
        printf("%d", middle[i]);
    }
    printf("\n");

    return 0;
}