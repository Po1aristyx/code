import numpy as np
import matplotlib.pyplot as plt
from sklearn.datasets import load_iris
from sklearn.cluster import KMeans, AgglomerativeClustering, DBSCAN
from sklearn.preprocessing import StandardScaler  # DBSCAN 对尺度敏感，可以考虑使用
from sklearn.metrics import silhouette_score  # 导入轮廓系数

# 加载 Iris 数据集
iris = load_iris()
X_iris = iris.data
y_iris = iris.target  # 真实标签

# 为了可视化，我们选取花瓣长度和花瓣宽度 (特征索引 2 和 3)
# 但聚类本身会在所有4个特征上进行
X_iris_visualization = X_iris[:, [2, 3]]  # Petal length and Petal width

# 特征名称（用于绘图标签）
feature_names_visualization = [iris.feature_names[2], iris.feature_names[3]]


# (2.1) K-means 聚类 (K=3，因为已知有3个类别)
# Adjust n_init for compatibility with scikit-learn versions
kmeans_iris = KMeans(n_clusters=3, random_state=45,
                     n_init='auto' if hasattr(KMeans(), 'n_init') and KMeans(n_clusters=3,
                                                                             random_state=42).n_init == 'auto' else 10)
labels_kmeans = kmeans_iris.fit_predict(X_iris)
silhouette_kmeans = silhouette_score(X_iris, labels_kmeans)
print(f"K-means Silhouette Score: {silhouette_kmeans:.3f}")

# (2.2) AGNES 层次聚类 (n_clusters=3)
agnes = AgglomerativeClustering(n_clusters=3)
labels_agnes = agnes.fit_predict(X_iris)
silhouette_agnes = silhouette_score(X_iris, labels_agnes)
print(f"AGNES Silhouette Score: {silhouette_agnes:.3f}")

# (2.3) DBSCAN 密度聚类
# DBSCAN 的参数 eps 和 min_samples 可能需要根据数据调整
# 用户提供的参数
dbscan_eps = 0.5
dbscan_min_samples = 5
dbscan = DBSCAN(eps=dbscan_eps, min_samples=dbscan_min_samples)
labels_dbscan = dbscan.fit_predict(X_iris)
n_clusters_dbscan = len(set(labels_dbscan)) - (1 if -1 in labels_dbscan else 0)
n_noise_dbscan = list(labels_dbscan).count(-1)

silhouette_dbscan = float('nan')  # Default to NaN
dbscan_message = f"DBSCAN (eps={dbscan_eps}, min_samples={dbscan_min_samples}) found {n_clusters_dbscan} clusters and {n_noise_dbscan} noise points."

# 轮廓系数要求至少有2个簇，并且每个簇的样本数不能太少 (n_labels in [2, n_samples - 1])
if n_clusters_dbscan >= 2 and n_clusters_dbscan < len(X_iris):
    try:
        silhouette_dbscan = silhouette_score(X_iris, labels_dbscan)
        dbscan_message += f"\nDBSCAN Silhouette Score: {silhouette_dbscan:.3f}"
    except ValueError as e:
        dbscan_message += f"\nDBSCAN: Could not compute silhouette score. Error: {e}"
else:
    dbscan_message += f"\nDBSCAN: Silhouette score not computed (requires >=2 and <n_samples clusters)."
print(dbscan_message)
print("-" * 45)

# 可视化聚类结果
fig2, axes2 = plt.subplots(2, 2, figsize=(15, 13))  # Adjusted figsize
axes2 = axes2.flatten()


# 绘图函数，避免代码重复
def plot_clusters(ax, data, labels, title, feature_names, true_labels=False, silhouette_val=None):
    unique_labels = np.unique(labels)
    # Ensure enough colors for unique labels
    n_plot_labels = len(unique_labels)
    colors = plt.cm.viridis(np.linspace(0, 1, n_plot_labels))

    legend_handles = []

    for i, k in enumerate(unique_labels):
        current_color = colors[i]
        marker = 'o'
        point_label_text = ''

        if k == -1 and not true_labels:  # DBSCAN中的噪声点
            current_color = [0, 0, 0, 1]  # 黑色
            marker = 'x'
            point_label_text = 'Noise'
        else:
            if true_labels:
                # Ensure k is a valid index for iris.target_names
                if 0 <= k < len(iris.target_names):
                    point_label_text = f'{iris.target_names[k]}'
                else:
                    point_label_text = f'Unknown Label {k}'  # Should not happen with y_iris
            else:
                point_label_text = f'Cluster {int(k)}'

        class_member_mask = (labels == k)
        xy = data[class_member_mask]
        # Check if xy is not empty before plotting
        if xy.shape[0] > 0:
            handle = ax.scatter(xy[:, 0], xy[:, 1], s=50, c=[current_color], edgecolor='k', marker=marker, alpha=0.7,
                                label=point_label_text)
            if point_label_text:
                legend_handles.append(handle)

    plot_title_full = title
    if silhouette_val is not None:
        if not np.isnan(silhouette_val):
            plot_title_full += f'\nSilhouette: {silhouette_val:.3f}'
        else:
            plot_title_full += f'\nSilhouette: N/A'

    ax.set_title(plot_title_full)
    ax.set_xlabel(feature_names[0])
    ax.set_ylabel(feature_names[1])
    if legend_handles:
        ax.legend(handles=legend_handles, loc='best')
    ax.grid(True)


# 图1: 真实标签
plot_clusters(axes2[0], X_iris_visualization, y_iris, 'Iris Dataset', feature_names_visualization,
              true_labels=True)

# 图2: K-means 结果
plot_clusters(axes2[1], X_iris_visualization, labels_kmeans, f'K-means(K=3)', feature_names_visualization,
              silhouette_val=silhouette_kmeans)

# 图3: AGNES 结果
plot_clusters(axes2[2], X_iris_visualization, labels_agnes, f'AGNES(n_clusters=3)',
              feature_names_visualization, silhouette_val=silhouette_agnes)

# 图4: DBSCAN 结果
dbscan_plot_title = f'DBSCAN (eps={dbscan_eps}, min_samp={dbscan_min_samples})'
if n_clusters_dbscan > 0:
    dbscan_plot_title += f'\n({n_clusters_dbscan} clusters, {n_noise_dbscan} noise)'
else:
    dbscan_plot_title += f'\n(0 clusters, {n_noise_dbscan} noise)'

plot_clusters(axes2[3], X_iris_visualization, labels_dbscan, dbscan_plot_title,
              feature_names_visualization, silhouette_val=silhouette_dbscan)

plt.tight_layout(rect=[0, 0, 1, 0.96])
plt.show()
