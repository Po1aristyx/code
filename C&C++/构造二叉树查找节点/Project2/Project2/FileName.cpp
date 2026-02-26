#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct TreeNode {
    char val;
    struct TreeNode* left;
    struct TreeNode* right;
} TreeNode;

int index = 0;
TreeNode* buildTree(char* preorder) {
    if (preorder[index] == '#' || preorder[index] == '\0') {
        index++;  // 空结点，直接返回 NULL
        return NULL;
    }
    TreeNode* node = (TreeNode*)malloc(sizeof(TreeNode));
    node->val = preorder[index++];
    node->left = buildTree(preorder);
    node->right = buildTree(preorder);
    return node;
}
bool findPath(TreeNode* root, char x, char* path, int* path_len) {
    if (root == NULL) {
        return false;
    }
    path[(*path_len)++] = root->val;

    if (root->val == x) {
        return true;
    }
    if (findPath(root->left, x, path, path_len) || findPath(root->right, x, path, path_len)) {
        return true;
    }
    (*path_len)--;
    return false;
}

int main() {
    char preorder[100];
    char x;
    scanf("%s", preorder);
    scanf(" %c", &x);

    TreeNode* root = buildTree(preorder);
    char path[100];
    int path_len = 0;
    if (findPath(root, x, path, &path_len)) {
        // 输出路径
        for (int i = 0; i < path_len; i++) {
            printf("%c ", path[i]);
        }
        printf("\n");
    }
    else {
        printf("Node not found\n");
    }

    return 0;
}