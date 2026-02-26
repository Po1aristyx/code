import numpy as np
import matplotlib.pyplot as plt
from sklearn.cluster import KMeans

# (1) K-means 聚类随机数据
# 随机生成 500 个二维数据点
X_random = np.random.rand(500, 2) * 100  # 乘以 100 使数据分布在 0-100 之间

# K 的取值
k_values = [2, 3, 4, 5]

# 创建子图，2行2列
fig1, axes1 = plt.subplots(2, 2, figsize=(12, 10))
axes1 = axes1.flatten()  # 将 2x2 的 axes 数组展平为一维，方便索引

print("WCSS (Inertia) for different K values on random data:")

for i, k in enumerate(k_values):
    # 初始化 KMeans 模型
    # n_init='auto' (or a specific number like 10) is recommended for scikit-learn >= 1.4
    # For older versions, n_init=10 is a good default.
    # random_state 保证结果可复现
    kmeans_random = KMeans(n_clusters=k, random_state=42,
                           n_init='auto' if hasattr(KMeans(), 'n_init') and KMeans().n_init == 'auto' else 10)

    # 拟合模型并预测簇标签
    labels_random = kmeans_random.fit_predict(X_random)

    # 获取簇中心
    centroids_random = kmeans_random.cluster_centers_

    # 获取簇内平方和 (WCSS) / Inertia
    wcss_random = kmeans_random.inertia_
    print(f"  K={k}, WCSS = {wcss_random:.2f}")

    # 可视化
    ax = axes1[i]
    scatter = ax.scatter(X_random[:, 0], X_random[:, 1], c=labels_random, cmap='viridis', alpha=0.7, edgecolors='k',
                         s=50)
    ax.scatter(centroids_random[:, 0], centroids_random[:, 1], marker='X', s=200, color='red', edgecolor='black',
               label='Centroids')
    # 在标题中加入 WCSS
    ax.set_title(f'K={k}\nWCSS = {wcss_random:.2f}')
    ax.set_xlabel('Feature 1')
    ax.set_ylabel('Feature 2')
    ax.legend(*scatter.legend_elements(), title='Clusters')
    ax.grid(True)

plt.tight_layout()
plt.suptitle('K-means Clustering on Randomly Generated 2D Data', fontsize=16, y=1.03)  # Adjusted y for suptitle
plt.show()