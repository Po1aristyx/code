#include <stdio.h>
#include <string.h>
#include <limits.h> // Potentially useful for other algorithms, not strictly needed for this greedy one
#include <stdlib.h>
#include <windows.h> // For SetConsoleOutputCP and SetConsoleCP

#define MAX_DENOMINATIONS 6
#define MAX_NAME_LEN 10

typedef struct {
    char name[MAX_NAME_LEN];
    int value_in_jiao;
} Denomination;

// Helper to parse "X元Y角", "X元", "Y角" string to jiao
// This function is from your provided code.
int parse_amount_to_jiao(const char* amount_str_const) {
    int yuan = 0;
    int jiao = 0;
    char temp_str[100];
    // Make a mutable copy to use with strtok-like modifications (e.g., inserting '\0')
    strncpy(temp_str, amount_str_const, sizeof(temp_str) - 1);
    temp_str[sizeof(temp_str) - 1] = '\0'; // Ensure null termination

    char* yuan_ptr = strstr(temp_str, "元");
    char* jiao_ptr = strstr(temp_str, "角");

    if (yuan_ptr != NULL) {
        *yuan_ptr = '\0'; // Null-terminate at "元" to get the yuan part
        yuan = atoi(temp_str);
        // Check for jiao part after "元"
        if (jiao_ptr != NULL && jiao_ptr > yuan_ptr) {
            // Move past "元" (and its null terminator if we were to consider that,
            // but yuan_ptr still points to the original string's "元" location)
            // The part for jiao starts after "元".
            jiao = atoi(yuan_ptr + strlen("元")); // Get part after "元"
        }
    } else if (jiao_ptr != NULL) { // No "元", but "角" is present
        *jiao_ptr = '\0'; // Null-terminate at "角"
        jiao = atoi(temp_str);
    } else { // No "元" and no "角", assume the whole string is yuan value
        // atoi will return 0 for non-numeric strings or empty strings
        yuan = atoi(temp_str);
    }
    return yuan * 10 + jiao;
}

// Function to calculate and display the change
void calculate_and_display_change(const char* input_amount_str, const Denomination denominations[], int num_denoms) {
    int target_jiao = parse_amount_to_jiao(input_amount_str);

    if (target_jiao <= 0) {
        // Check if the original input string was empty or represented zero,
        // to give a more specific "no change needed" or "invalid input" message.
        char trimmed_input[100];
        strncpy(trimmed_input, input_amount_str, sizeof(trimmed_input)-1);
        trimmed_input[sizeof(trimmed_input)-1] = '\0';
        // A simple check for common zero inputs. A more robust trim might be needed for spaces.
        if (strlen(trimmed_input) == 0 || strcmp(trimmed_input, "0") == 0 ||
            strcmp(trimmed_input, "0元") == 0 || strcmp(trimmed_input, "0角") == 0) {
            printf("无需找零或输入为零。\n");
        } else {
            printf("解析金额 \"%s\" 失败或金额无效。请输入有效金额 (如 '5元2角', '10元', '3角') 且金额需为正。\n", input_amount_str);
        }
        printf("------------------------------------\n");
        return;
    }

    printf("\n为金额 \"%s\" (即 %d角) 计算找零：\n", input_amount_str, target_jiao);

    int change_given_counts[MAX_DENOMINATIONS] = {0}; // Assumes MAX_DENOMINATIONS matches num_denoms
    int total_coins_count = 0;
    int remaining_amount = target_jiao;

    for (int i = 0; i < num_denoms; ++i) {
        if (remaining_amount == 0) {
            break;
        }
        // Denominations should be sorted descending for greedy algorithm to work optimally for standard currency
        if (denominations[i].value_in_jiao <= remaining_amount) {
            int count = remaining_amount / denominations[i].value_in_jiao;
            if (count > 0) {
                change_given_counts[i] = count;
                total_coins_count += count;
                remaining_amount %= denominations[i].value_in_jiao;
            }
        }
    }

    if (remaining_amount == 0) {
        printf("最优找零方案 (贪心算法)：\n");
        for (int i = 0; i < num_denoms; ++i) {
            if (change_given_counts[i] > 0) {
                printf("- %s: %d 张/枚\n", denominations[i].name, change_given_counts[i]);
            }
        }
        printf("总计数量: %d 枚/张\n", total_coins_count);
    } else {
        // This case should ideally not happen with standard denominations like RMB for any positive amount
        printf("无法使用当前面额精确找零 \"%s\"，剩余 %d角。\n", input_amount_str, remaining_amount);
    }
    printf("------------------------------------\n");
}

int main() {
    // 设置控制台输出和输入编码为 UTF-8 (代码页 65001) 以支持中文
    SetConsoleOutputCP(CP_UTF8);
    SetConsoleCP(CP_UTF8);

    Denomination denominations[MAX_DENOMINATIONS] = {
            {"5元", 50},
            {"2元", 20},
            {"1元", 10},
            {"5角", 5},
            {"2角", 2},
            {"1角", 1}
    };
    // Ensure denominations are sorted by value descending for the greedy algorithm.
    // The above list is already sorted.
    int num_denoms = MAX_DENOMINATIONS;

    printf("--- 超市现金收支系统找零 ---\n");
    printf("可用面额有: ");
    for(int i=0; i < num_denoms; ++i) {
        printf("%s%s", denominations[i].name, (i == num_denoms - 1) ? "" : ", ");
    }
    printf("\n");

    char input_buffer[100];
    while (1) {
        printf("\n请输入需要找零的金额 : ");
        if (fgets(input_buffer, sizeof(input_buffer), stdin) == NULL) {
            // EOF (Ctrl+Z on Windows, Ctrl+D on Linux/macOS) or error
            printf("\n读取输入失败或已达文件末尾，正在退出...\n");
            break;
        }

        // Remove trailing newline character if present
        input_buffer[strcspn(input_buffer, "\n")] = 0;

        // Check for exit command
        if (strcmp(input_buffer, "exit") == 0) {
            printf("正在退出系统...\n");
            break;
        }

        // Handle empty input after removing newline (user just pressed Enter)
        if (strlen(input_buffer) == 0) {
            printf("请输入有效金额或 'exit'。\n");
            printf("------------------------------------\n");
            continue;
        }

        calculate_and_display_change(input_buffer, denominations, num_denoms);
    }

    return 0;
}