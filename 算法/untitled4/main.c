#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define MAX_TYPES 200
#define MAX_NAME_LEN 25
#define MAX_ROOM_SOCKETS_N 100
#define MAX_DEVICES_M 100
#define MAX_ADAPTERS_K 100

typedef struct {
    char name[MAX_NAME_LEN]; // For devices, this is device name; for room sockets, it's the socket type itself
    int type_id;          // For devices, this is plug_type_id; for room sockets, it's socket_type_id
} NamedItem;

// Adapters are defined by what socket they provide and what plug they need
// typedef struct {
// int provides_socket_type_id;
// int needs_plug_type_id;
// } Adapter; // Not strictly needed as a struct if directly used in can_transform

NamedItem room_socket_instances[MAX_ROOM_SOCKETS_N];
int num_room_sockets_n;

NamedItem devices[MAX_DEVICES_M];
int num_devices_m;

// Adapter info will be directly used to populate can_transform
int num_adapters_k;

char type_id_to_name_map[MAX_TYPES][MAX_NAME_LEN];
int unique_type_count = 0;

int can_transform[MAX_TYPES][MAX_TYPES];

int match_for_room_socket[MAX_ROOM_SOCKETS_N];
int visited_this_dfs_pass[MAX_ROOM_SOCKETS_N];

int get_or_create_type_id(const char* type_name) {
    for (int i = 0; i < unique_type_count; ++i) {
        if (strcmp(type_id_to_name_map[i], type_name) == 0) {
            return i;
        }
    }
    if (unique_type_count < MAX_TYPES) {
        strncpy(type_id_to_name_map[unique_type_count], type_name, MAX_NAME_LEN -1);
        type_id_to_name_map[unique_type_count][MAX_NAME_LEN-1] = '\0';
        return unique_type_count++;
    }
    fprintf(stderr, "Error: Exceeded MAX_TYPES limit.\n");
    exit(1);
}

int find_match_for_device_dfs(int device_idx) {
    int device_plug_type_id = devices[device_idx].type_id;

    for (int rs_idx = 0; rs_idx < num_room_sockets_n; ++rs_idx) {
        if (visited_this_dfs_pass[rs_idx]) {
            continue;
        }

        int room_socket_actual_type_id = room_socket_instances[rs_idx].type_id;

        if (can_transform[device_plug_type_id][room_socket_actual_type_id]) {
            visited_this_dfs_pass[rs_idx] = 1;

            if (match_for_room_socket[rs_idx] == -1 ||
                find_match_for_device_dfs(match_for_room_socket[rs_idx])) {
                match_for_room_socket[rs_idx] = device_idx;
                return 1;
            }
        }
    }
    return 0;
}
void solve() {
    unique_type_count = 0;
    memset(can_transform, 0, sizeof(can_transform));
    // Initialize diagonal for can_transform later, after unique_type_count is somewhat stable or use MAX_TYPES
    for (int i = 0; i < MAX_TYPES; ++i) { // Initialize for all possible types
        can_transform[i][i] = 1;
    }


    char line_buffer[256]; // Buffer to read each line

    // --- 0. 读取 n, m, k ---
    if (fgets(line_buffer, sizeof(line_buffer), stdin) == NULL) {
        fprintf(stderr, "错误：无法读取 n, m, k 行。\n");
        return;
    }
    if (sscanf(line_buffer, "%d %d %d", &num_room_sockets_n, &num_devices_m, &num_adapters_k) != 3) {
        fprintf(stderr, "错误：n, m, k 输入格式错误。应为三个整数，例如 '4 5 3'。\n");
        return;
    }

    // Validate n, m, k against defined maximums (important!)
    if (num_room_sockets_n > MAX_ROOM_SOCKETS_N || num_room_sockets_n <= 0 ||
        num_devices_m > MAX_DEVICES_M || num_devices_m <= 0 || // Assuming m can be 0 if no devices
        num_adapters_k > MAX_ADAPTERS_K || num_adapters_k < 0) { // k can be 0
        fprintf(stderr, "错误：n, m, 或 k 的值超出范围或无效。\n");
        printf("n (插座数): %d (最大 %d)\n", num_room_sockets_n, MAX_ROOM_SOCKETS_N);
        printf("m (设备数): %d (最大 %d)\n", num_devices_m, MAX_DEVICES_M);
        printf("k (适配器数): %d (最大 %d)\n", num_adapters_k, MAX_ADAPTERS_K);
        return;
    }


    char buffer1[MAX_NAME_LEN];
    char buffer2[MAX_NAME_LEN];

    // --- 1. 输入发布室插座 ---
    // printf("调试：准备读取 %d 个发布室插座\n", num_room_sockets_n);
    for (int i = 0; i < num_room_sockets_n; ++i) {
        if (fgets(line_buffer, sizeof(line_buffer), stdin) == NULL) {
            fprintf(stderr, "错误：读取第 %d 个发布室插座时提前结束输入。\n", i + 1);
            return;
        }
        if (sscanf(line_buffer, "%s", buffer1) != 1) {
            fprintf(stderr, "错误：第 %d 个发布室插座格式错误。\n", i + 1);
            return;
        }
        // Remove trailing newline from buffer1 if necessary (though sscanf %s handles whitespace)
        buffer1[strcspn(buffer1, "\r\n")] = 0;
        if (strlen(buffer1) == 0) {
            fprintf(stderr, "错误：第 %d 个发布室插座类型为空。\n", i + 1);
            return;
        }
        strcpy(room_socket_instances[i].name, buffer1);
        room_socket_instances[i].type_id = get_or_create_type_id(buffer1);
    }

    // --- 2. 输入设备 ---
    // printf("调试：准备读取 %d 个设备\n", num_devices_m);
    for (int i = 0; i < num_devices_m; ++i) {
        if (fgets(line_buffer, sizeof(line_buffer), stdin) == NULL) {
            fprintf(stderr, "错误：读取第 %d 个设备时提前结束输入。\n", i + 1);
            return;
        }
        if (sscanf(line_buffer, "%s %s", buffer1, buffer2) != 2) {
            fprintf(stderr, "错误：第 %d 个设备格式错误。应为 '设备名 插头类型'。\n", i + 1);
            return;
        }
        buffer1[strcspn(buffer1, "\r\n")] = 0; // Device name
        buffer2[strcspn(buffer2, "\r\n")] = 0; // Plug type
        if (strlen(buffer1) == 0 || strlen(buffer2) == 0) {
            fprintf(stderr, "错误：第 %d 个设备的名称或插头类型为空。\n", i + 1);
            return;
        }
        strcpy(devices[i].name, buffer1);
        devices[i].type_id = get_or_create_type_id(buffer2);
    }

    // --- 3. 输入适配器 ---
    // printf("调试：准备读取 %d 个适配器\n", num_adapters_k);
    for (int i = 0; i < num_adapters_k; ++i) {
        if (fgets(line_buffer, sizeof(line_buffer), stdin) == NULL) {
            fprintf(stderr, "错误：读取第 %d 个适配器时提前结束输入。\n", i + 1);
            return;
        }
        if (sscanf(line_buffer, "%s %s", buffer1, buffer2) != 2) {
            fprintf(stderr, "错误：第 %d 个适配器格式错误。应为 '提供插座类型 需要插头类型'。\n", i + 1);
            return;
        }
        buffer1[strcspn(buffer1, "\r\n")] = 0; // Provides socket
        buffer2[strcspn(buffer2, "\r\n")] = 0; // Needs plug
        if (strlen(buffer1) == 0 || strlen(buffer2) == 0) {
            fprintf(stderr, "错误：第 %d 个适配器的提供类型或需求类型为空。\n", i + 1);
            return;
        }
        int provides_id = get_or_create_type_id(buffer1);
        int needs_id = get_or_create_type_id(buffer2);
        can_transform[needs_id][provides_id] = 1;
    }
    // printf("调试：所有输入读取完毕。\n");

    // --- 4. 计算传递闭包 (Floyd-Warshall for reachability) ---
    for (int k_intermediate = 0; k_intermediate < unique_type_count; ++k_intermediate) {
        for (int i_start_plug = 0; i_start_plug < unique_type_count; ++i_start_plug) {
            for (int j_end_socket = 0; j_end_socket < unique_type_count; ++j_end_socket) {
                if (can_transform[i_start_plug][k_intermediate] && can_transform[k_intermediate][j_end_socket]) {
                    can_transform[i_start_plug][j_end_socket] = 1;
                }
            }
        }
    }

    // --- 5. 执行匈牙利算法进行二分图匹配 ---
    int matched_devices_count = 0;
    for (int i = 0; i < num_room_sockets_n; ++i) {
        match_for_room_socket[i] = -1;
    }

    for (int dev_idx = 0; dev_idx < num_devices_m; ++dev_idx) {
        for (int rs_idx = 0; rs_idx < num_room_sockets_n; ++rs_idx) {
            visited_this_dfs_pass[rs_idx] = 0;
        }
        if (find_match_for_device_dfs(dev_idx)) {
            matched_devices_count++;
        }
    }

    // --- 6. 输出结果 ---
    printf("%d\n", num_devices_m - matched_devices_count);
}

int main() {
    int num_test_cases = 1;
    char first_line_buffer[256];

    if (fgets(first_line_buffer, sizeof(first_line_buffer), stdin) == NULL) {
        fprintf(stderr, "错误：无法读取测试用例数量行。\n");
        return 1;
    }
    if (sscanf(first_line_buffer, "%d", &num_test_cases) != 1) {
        fprintf(stderr, "错误：测试用例数量行格式错误。\n");
        return 1;
    }
    if (num_test_cases <= 0) { // 通常测试用例数至少为1
        fprintf(stderr, "错误：无效的测试用例数量 %d。\n", num_test_cases);
        return 1;
    }


    while (num_test_cases--) {
        solve();
    }

    return 0;
}