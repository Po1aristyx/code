#include <stdio.h>
#include <stdlib.h> // For qsort
#include <string.h> // For memcpy (optional, can do manual struct copy)
#include <unistd.h>

#define MAX_PARTITIONS 20 // 假设系统中最大分区表项数（空闲+已分配）
#define MEMORY_SIZE 256   // 总内存大小，单位KB
#define OS_SIZE 10        // 操作系统占用大小

// 作业ID常量 (用于初始作业)
#define OS_JOB_ID 0
#define JOB1_ID 1
#define JOB2_ID 2
#define JOB4_ID 4

// 空闲区结构
typedef struct {
    int start_address;
    int size;
} FreeBlock;

// 已分配区结构
typedef struct {
    int job_id;
    int start_address;
    int size;
} AllocatedBlock;

FreeBlock free_list[MAX_PARTITIONS];
int free_block_count = 0;

AllocatedBlock allocated_list[MAX_PARTITIONS];
int allocated_block_count = 0;

int next_assignable_job_id = 5; // 用于分配新的作业ID，接续初始作业ID

// 比较函数，用于qsort对free_list按大小升序排序（大小相同则按地址升序）
int compare_free_blocks(const void *a, const void *b) {
    FreeBlock *blockA = (FreeBlock *)a;
    FreeBlock *blockB = (FreeBlock *)b;

    if (blockA->size < blockB->size) return -1;
    if (blockA->size > blockB->size) return 1;
    // 大小相同，按起始地址排序
    if (blockA->start_address < blockB->start_address) return -1;
    if (blockA->start_address > blockB->start_address) return 1;
    return 0;
}

// 对空闲区列表进行排序
void sort_free_list() {
    if (free_block_count > 1) {
        qsort(free_list, free_block_count, sizeof(FreeBlock), compare_free_blocks);
    }
}

// 打印空闲区列表
void print_free_list() {
    printf("--- Free List (Best-Fit Sorted by Size, then Address) ---\n");
    if (free_block_count == 0) {
        printf("  No free blocks.\n");
    } else {
        for (int i = 0; i < free_block_count; i++) {
            printf("  Start: %3d KB, Size: %3d KB\n", free_list[i].start_address, free_list[i].size);
        }
    }
    printf("-------------------------------------------------------\n");
}

// 打印已分配区列表
void print_allocated_list() {
    printf("--- Allocated List ---\n");
    if (allocated_block_count == 0) {
        printf("  No allocated blocks (except OS if initialized separately).\n");
    } else {
        // Sort by start_address for consistent printing (optional)
        AllocatedBlock temp_list[MAX_PARTITIONS];
        memcpy(temp_list, allocated_list, allocated_block_count * sizeof(AllocatedBlock));
        // Simple bubble sort for allocated list display
        for(int i=0; i < allocated_block_count -1; ++i) {
            for(int j=0; j < allocated_block_count -i -1; ++j) {
                if(temp_list[j].start_address > temp_list[j+1].start_address) {
                    AllocatedBlock t = temp_list[j];
                    temp_list[j] = temp_list[j+1];
                    temp_list[j+1] = t;
                }
            }
        }

        for (int i = 0; i < allocated_block_count; i++) {
            printf("  Job ID: %2d, Start: %3d KB, Size: %3d KB\n", temp_list[i].job_id, temp_list[i].start_address, temp_list[i].size);
        }
    }
    printf("------------------------\n");
}

// 初始化内存状态 (根据图1)
void initialize_memory() {
    // 清空列表
    free_block_count = 0;
    allocated_block_count = 0;

    // 1. 添加已分配区 (OS 和初始作业)
    // 操作系统
    allocated_list[allocated_block_count].job_id = OS_JOB_ID;
    allocated_list[allocated_block_count].start_address = 0;
    allocated_list[allocated_block_count].size = OS_SIZE;
    allocated_block_count++;

    // 作业1
    allocated_list[allocated_block_count].job_id = JOB1_ID;
    allocated_list[allocated_block_count].start_address = 10;
    allocated_list[allocated_block_count].size = 10;
    allocated_block_count++;

    // 作业4
    allocated_list[allocated_block_count].job_id = JOB4_ID;
    allocated_list[allocated_block_count].start_address = 20;
    allocated_list[allocated_block_count].size = 25;
    allocated_block_count++;

    // 作业2
    allocated_list[allocated_block_count].job_id = JOB2_ID;
    allocated_list[allocated_block_count].start_address = 65;
    allocated_list[allocated_block_count].size = 45;
    allocated_block_count++;

    // 2. 添加空闲区
    // 空闲区1
    free_list[free_block_count].start_address = 45;
    free_list[free_block_count].size = 20;
    free_block_count++;

    // 空闲区2
    free_list[free_block_count].start_address = 110;
    free_list[free_block_count].size = MEMORY_SIZE - 110; // 146KB
    free_block_count++;

    // 3. 排序空闲区列表
    sort_free_list();
}

// 最佳适应算法分配内存
int allocate_memory_best_fit(int job_id, int requested_size) {
    if (requested_size <= 0) {
        printf("Error: Requested size must be positive.\n");
        return 0; // 分配失败
    }
    if (allocated_block_count >= MAX_PARTITIONS) {
        printf("Error: Allocated list is full. Cannot allocate Job %d.\n", job_id);
        return 0;
    }


    int best_fit_idx = -1;
    // free_list 已经按大小排序，第一个满足条件的就是最佳的
    for (int i = 0; i < free_block_count; i++) {
        if (free_list[i].size >= requested_size) {
            best_fit_idx = i;
            break; // 因为已排序，第一个找到的就是最小的足够大的块
        }
    }

    if (best_fit_idx == -1) {
        printf("Allocation Failed: No sufficient free block for Job %d (Size: %dKB).\n", job_id, requested_size);
        return 0; // 分配失败
    }

    FreeBlock chosen_block = free_list[best_fit_idx];

    // 添加到已分配列表
    allocated_list[allocated_block_count].job_id = job_id;
    allocated_list[allocated_block_count].start_address = chosen_block.start_address;
    allocated_list[allocated_block_count].size = requested_size;
    allocated_block_count++;

    // 处理空闲区列表
    if (chosen_block.size == requested_size) {
        // 正好匹配，从空闲列表移除该块
        for (int i = best_fit_idx; i < free_block_count - 1; i++) {
            free_list[i] = free_list[i + 1];
        }
        free_block_count--;
    } else {
        // 分割空闲块
        free_list[best_fit_idx].start_address = chosen_block.start_address + requested_size;
        free_list[best_fit_idx].size = chosen_block.size - requested_size;
        // 因为大小改变，需要重新排序
        sort_free_list();
    }

    printf("Allocation Successful: Job %d (Size: %dKB) allocated at Address %dKB.\n", job_id, requested_size, chosen_block.start_address);
    return 1; // 分配成功
}

// 回收内存
void deallocate_memory(int job_id_to_release) {
    int allocated_idx = -1;
    for (int i = 0; i < allocated_block_count; i++) {
        if (allocated_list[i].job_id == job_id_to_release) {
            allocated_idx = i;
            break;
        }
    }

    if (allocated_idx == -1) {
        printf("Deallocation Failed: Job %d not found in allocated list.\n", job_id_to_release);
        return;
    }

    if (free_block_count >= MAX_PARTITIONS) {
        printf("Error: Free list is full. Cannot add released block from Job %d without merging.\n", job_id_to_release);
        // This is a simplification; in a real system, merging should still proceed if it reduces free_block_count
    }


    AllocatedBlock released_block_info = allocated_list[allocated_idx];
    printf("Deallocating Job %d (Start: %dKB, Size: %dKB).\n",
           released_block_info.job_id, released_block_info.start_address, released_block_info.size);

    // 从已分配列表移除
    for (int i = allocated_idx; i < allocated_block_count - 1; i++) {
        allocated_list[i] = allocated_list[i + 1];
    }
    allocated_block_count--;

    // --- 合并逻辑 ---
    int current_addr = released_block_info.start_address;
    int current_size = released_block_info.size;
    int merged_something_in_iteration;

    do {
        merged_something_in_iteration = 0;
        for (int i = 0; i < free_block_count; ) { // Manual increment for i
            // 检查下邻 (free_list[i] is below current_addr)
            if (free_list[i].start_address + free_list[i].size == current_addr) {
                current_addr = free_list[i].start_address;
                current_size += free_list[i].size;
                // 移除 free_list[i]
                for (int j = i; j < free_block_count - 1; j++) {
                    free_list[j] = free_list[j + 1];
                }
                free_block_count--;
                merged_something_in_iteration = 1;
                continue; // 重新扫描，因为列表改变了
            }
            // 检查上邻 (free_list[i] is above current_addr)
            else if (current_addr + current_size == free_list[i].start_address) {
                current_size += free_list[i].size;
                // 移除 free_list[i]
                for (int j = i; j < free_block_count - 1; j++) {
                    free_list[j] = free_list[j + 1];
                }
                free_block_count--;
                merged_something_in_iteration = 1;
                continue; // 重新扫描
            }
            i++;
        }
    } while (merged_something_in_iteration && free_block_count < MAX_PARTITIONS); // Continue merging if something changed

    // 将最终（可能合并后的）块添加到空闲列表
    if (free_block_count < MAX_PARTITIONS) {
        free_list[free_block_count].start_address = current_addr;
        free_list[free_block_count].size = current_size;
        free_block_count++;
        sort_free_list();
    } else {
         printf("Critical Error: Free list full even after attempting merge for released Job %d. Block lost.\n", job_id_to_release);
    }
    printf("Deallocation Complete for Job %d.\n", job_id_to_release);
}


int main() {
    initialize_memory();
    printf("Initial Memory State:\n");
    print_free_list();
    print_allocated_list();
    printf("=======================================================\n\n");

    // 1. 作业申请队列
    printf("--- Processing Job Allocation Requests ---\n");
    int job_A_id = next_assignable_job_id++;
    printf("Requesting Job %d (Size 15KB):\n", job_A_id);
    allocate_memory_best_fit(job_A_id, 15);
    print_free_list();
    print_allocated_list();
    printf("--------------------------------------------\n");

    int job_B_id = next_assignable_job_id++;
    printf("Requesting Job %d (Size 50KB):\n", job_B_id);
    allocate_memory_best_fit(job_B_id, 50);
    print_free_list();
    print_allocated_list();
    printf("--------------------------------------------\n");

    int job_C_id = next_assignable_job_id++;
    printf("Requesting Job %d (Size 100KB):\n", job_C_id);
    allocate_memory_best_fit(job_C_id, 100);
    print_free_list();
    print_allocated_list();
    printf("--------------------------------------------\n");

    int job_D_id = next_assignable_job_id++;
    printf("Requesting Job %d (Size 30KB):\n", job_D_id); // This might fail or take the last large chunk
    allocate_memory_best_fit(job_D_id, 30);
    print_free_list();
    print_allocated_list();
    printf("=======================================================\n\n");


    // 2. 作业释放序列
    printf("--- Processing Job Release Requests ---\n");
    printf("Releasing Job %d (Original Job1):\n", JOB1_ID);
    deallocate_memory(JOB1_ID);
    print_free_list();
    print_allocated_list();
    printf("--------------------------------------------\n");

    printf("Releasing Job %d (Newly allocated Job A):\n", job_A_id);
    deallocate_memory(job_A_id);
    print_free_list();
    print_allocated_list();
    printf("--------------------------------------------\n");

    printf("Releasing Job %d (Original Job4):\n", JOB4_ID);
    deallocate_memory(JOB4_ID);
    print_free_list();
    print_allocated_list();
    printf("--------------------------------------------\n");
    
    printf("Releasing Job %d (Newly allocated Job C if it was allocated):\n", job_C_id);
    deallocate_memory(job_C_id); // Attempt to release Job C
    print_free_list();
    print_allocated_list();
    printf("=======================================================\n\n");

    printf("Final Memory State:\n");
    print_free_list();
    print_allocated_list();

    return 0;
}