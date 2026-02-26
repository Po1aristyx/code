#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_STRING_LEN 25
#define MAX_N_ITEMS 105 // Max for physical sockets or devices
#define MAX_ADAPTER_RULES 105
#define MAX_POTENTIAL_STRINGS (MAX_N_ITEMS + MAX_N_ITEMS + MAX_ADAPTER_RULES * 2)
#define MAX_UNIQUE_TYPES MAX_POTENTIAL_STRINGS

// --- Type Management ---
char unique_type_names[MAX_UNIQUE_TYPES][MAX_STRING_LEN];
int num_unique_types = 0;
char all_strings_temp_storage[MAX_POTENTIAL_STRINGS][MAX_STRING_LEN];
int temp_string_count = 0;

// --- Adapter Graph (for BFS reachability) ---
struct AdjListNode {
    int dest_type_id;
    struct AdjListNode* next;
};
struct AdjListNode* adapter_adj_list[MAX_UNIQUE_TYPES];

// --- Input Storage ---
char physical_socket_initial_types_str[MAX_N_ITEMS][MAX_STRING_LEN];
int physical_socket_initial_type_ids[MAX_N_ITEMS];
int n_physical_sockets;

char device_needs_type_str[MAX_N_ITEMS][MAX_STRING_LEN];
int device_needs_type_ids[MAX_N_ITEMS];
int n_devices;

struct AdapterRuleRaw {
    char provides_socket_str[MAX_STRING_LEN];
    char needs_plug_str[MAX_STRING_LEN];
} adapter_rules_raw[MAX_ADAPTER_RULES];
int n_adapter_rules;

// --- Reachability ---
bool can_ps_provide_type[MAX_N_ITEMS][MAX_UNIQUE_TYPES]; // can_ps_provide_type[ps_idx][type_id]

// --- Bipartite Matching ---
struct AdjListNodeBipartite {
    int physical_socket_idx; // Right side node
    struct AdjListNodeBipartite* next;
};
struct AdjListNodeBipartite* bipartite_adj[MAX_N_ITEMS]; // For devices (left side)
int match_to_socket[MAX_N_ITEMS]; // match_to_socket[ps_idx] = device_idx, or -1
bool visited_dfs_round[MAX_N_ITEMS]; // For physical sockets in one DFS round for matching

// --- Helper Functions for Type Management ---
int compare_string_arrays_qsort(const void* a, const void* b) {
    return strcmp((const char*)a, (const char*)b);
}

void add_string_to_temp(const char* s) {
    if (temp_string_count < MAX_POTENTIAL_STRINGS) {
        strcpy(all_strings_temp_storage[temp_string_count++], s);
    }
}

int get_type_id_by_name(const char* name) {
    for (int i = 0; i < num_unique_types; ++i) {
        if (strcmp(name, unique_type_names[i]) == 0) return i;
    }
    return -1; // Should not happen if all types are processed
}

void build_unique_types() {
    if (temp_string_count == 0) return;
    qsort(all_strings_temp_storage, temp_string_count, MAX_STRING_LEN, compare_string_arrays_qsort);
    num_unique_types = 0;
    if (num_unique_types < MAX_UNIQUE_TYPES) { // Check before first strcpy
        strcpy(unique_type_names[num_unique_types++], all_strings_temp_storage[0]);
    }
    for (int i = 1; i < temp_string_count; ++i) {
        if (strcmp(all_strings_temp_storage[i], unique_type_names[num_unique_types - 1]) != 0) {
            if (num_unique_types < MAX_UNIQUE_TYPES) {
                strcpy(unique_type_names[num_unique_types++], all_strings_temp_storage[i]);
            } else break;
        }
    }
}

// --- Helper Functions for Graphs ---
void add_adapter_edge(int u_type_id, int v_type_id) { // N -> P
    struct AdjListNode* newNode = (struct AdjListNode*)malloc(sizeof(struct AdjListNode));
    newNode->dest_type_id = v_type_id;
    newNode->next = adapter_adj_list[u_type_id];
    adapter_adj_list[u_type_id] = newNode;
}

void add_bipartite_edge(int device_idx, int ps_idx) {
    struct AdjListNodeBipartite* newNode = (struct AdjListNodeBipartite*)malloc(sizeof(struct AdjListNodeBipartite));
    newNode->physical_socket_idx = ps_idx;
    newNode->next = bipartite_adj[device_idx];
    bipartite_adj[device_idx] = newNode;
}

void bfs_for_reachability(int start_type_id, bool* result_reachable_types) {
    for (int i = 0; i < num_unique_types; ++i) result_reachable_types[i] = false;
    if (start_type_id < 0 || start_type_id >= num_unique_types) return;

    int queue[MAX_UNIQUE_TYPES];
    int head = 0, tail = 0;

    result_reachable_types[start_type_id] = true;
    queue[tail++] = start_type_id;

    while (head < tail) {
        int u_type_id = queue[head++];
        struct AdjListNode* temp = adapter_adj_list[u_type_id];
        while (temp != NULL) {
            int v_type_id = temp->dest_type_id;
            if (!result_reachable_types[v_type_id]) {
                result_reachable_types[v_type_id] = true;
                if (tail < MAX_UNIQUE_TYPES) queue[tail++] = v_type_id;
            }
            temp = temp->next;
        }
    }
}

// --- Bipartite Matching DFS ---
bool find_augmenting_path_dfs(int dev_idx) {
    struct AdjListNodeBipartite* temp = bipartite_adj[dev_idx];
    while (temp != NULL) {
        int ps_idx = temp->physical_socket_idx;
        if (!visited_dfs_round[ps_idx]) {
            visited_dfs_round[ps_idx] = true;
            if (match_to_socket[ps_idx] < 0 || find_augmenting_path_dfs(match_to_socket[ps_idx])) {
                match_to_socket[ps_idx] = dev_idx;
                return true;
            }
        }
        temp = temp->next;
    }
    return false;
}


int main() {
    // Initialization
    for(int i=0; i<MAX_UNIQUE_TYPES; ++i) adapter_adj_list[i] = NULL;
    for(int i=0; i<MAX_N_ITEMS; ++i) bipartite_adj[i] = NULL;
    for(int i=0; i<MAX_N_ITEMS; ++i) {
        for(int j=0; j<MAX_UNIQUE_TYPES; ++j) {
            can_ps_provide_type[i][j] = false;
        }
    }


    // 1. Read physical socket info
    scanf("%d", &n_physical_sockets);
    for (int i = 0; i < n_physical_sockets; ++i) {
        scanf("%s", physical_socket_initial_types_str[i]);
        add_string_to_temp(physical_socket_initial_types_str[i]);
    }

    // 2. Read device info
    scanf("%d", &n_devices);
    char temp_dev_name[MAX_STRING_LEN];
    for (int i = 0; i < n_devices; ++i) {
        scanf("%s %s", temp_dev_name, device_needs_type_str[i]);
        add_string_to_temp(device_needs_type_str[i]);
    }

    // 3. Read adapter info
    scanf("%d", &n_adapter_rules);
    for (int i = 0; i < n_adapter_rules; ++i) {
        scanf("%s %s", adapter_rules_raw[i].provides_socket_str, adapter_rules_raw[i].needs_plug_str);
        add_string_to_temp(adapter_rules_raw[i].provides_socket_str);
        add_string_to_temp(adapter_rules_raw[i].needs_plug_str);
    }

    // 4. Build unique type list and get IDs
    build_unique_types();
    for (int i = 0; i < n_physical_sockets; ++i) {
        physical_socket_initial_type_ids[i] = get_type_id_by_name(physical_socket_initial_types_str[i]);
    }
    for (int i = 0; i < n_devices; ++i) {
        device_needs_type_ids[i] = get_type_id_by_name(device_needs_type_str[i]);
    }

    // 5. Build adapter graph (N -> P)
    for (int i = 0; i < n_adapter_rules; ++i) {
        int id_n_needs = get_type_id_by_name(adapter_rules_raw[i].needs_plug_str);
        int id_p_provides = get_type_id_by_name(adapter_rules_raw[i].provides_socket_str);
        if (id_n_needs != -1 && id_p_provides != -1) {
            add_adapter_edge(id_n_needs, id_p_provides);
        }
    }

    // 6. Compute reachability for each physical socket
    for (int i = 0; i < n_physical_sockets; ++i) {
        bfs_for_reachability(physical_socket_initial_type_ids[i], can_ps_provide_type[i]);
    }

    // 7. Build bipartite graph (Device -> PhysicalSocket)
    for (int i = 0; i < n_devices; ++i) { // For each device i
        int needed_type_id = device_needs_type_ids[i];
        if (needed_type_id == -1) continue;
        for (int j = 0; j < n_physical_sockets; ++j) { // For each physical socket j
            if (can_ps_provide_type[j][needed_type_id]) {
                add_bipartite_edge(i, j); // Device i can be powered by physical socket j
            }
        }
    }

    // 8. Max Bipartite Matching
    int matched_devices_count = 0;
    for(int i=0; i<MAX_N_ITEMS; ++i) match_to_socket[i] = -1; // Initialize match_to_socket for physical sockets

    for (int i = 0; i < n_devices; ++i) { // For each device
        for(int j=0; j<n_physical_sockets; ++j) visited_dfs_round[j] = false; // Reset visited for this DFS round
        if (find_augmenting_path_dfs(i)) {
            matched_devices_count++;
        }
    }

    // 9. Output
    printf("%d\n", n_devices - matched_devices_count);

    // --- Cleanup (Free dynamically allocated memory) ---
    for(int i=0; i<num_unique_types; ++i) {
        struct AdjListNode* curr = adapter_adj_list[i];
        while(curr) { struct AdjListNode* temp = curr; curr = curr->next; free(temp); }
    }
    for(int i=0; i<n_devices; ++i) {
        struct AdjListNodeBipartite* curr = bipartite_adj[i];
        while(curr) { struct AdjListNodeBipartite* temp = curr; curr = curr->next; free(temp); }
    }

    getchar(); // To pause console on some systems
    system("pause");
    return 0;
}