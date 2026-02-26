import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split, GridSearchCV, cross_val_score
from sklearn.tree import DecisionTreeClassifier, export_graphviz
# from sklearn.preprocessing import LabelEncoder, StandardScaler, OneHotEncoder # LabelEncoder might not be needed if Sex is mapped
# from sklearn.impute import SimpleImputer # Not used directly, manual imputation
# from sklearn.compose import ColumnTransformer # Not used in this manual flow
# from sklearn.pipeline import Pipeline # Not used in this manual flow
from sklearn.metrics import accuracy_score
import graphviz
import matplotlib.pyplot as plt
import seaborn as sns

# (1) 导入库 (已完成)

# (3) 使用 Pandas 库读取数据文件
try:
    train_df = pd.read_csv('train.csv')
    test_df = pd.read_csv('test.csv')
    print("train.csv 和 test.csv 加载成功。")
except FileNotFoundError:
    print("错误：'train.csv' 或 'test.csv' 文件未找到。请确保文件存在于正确的路径。")
    # 为了能继续运行，您需要确保文件存在
    exit() # 如果文件不存在，则退出

print("训练数据 (train.csv) 前5行:")
print(train_df.head())
print("\n测试数据 (test.csv) 前5行:")
print(test_df.head())

# 保存测试集的 PassengerId 用于后续提交 (如果需要)
test_passenger_ids = test_df['PassengerId']

# --- 数据预处理与特征工程 ---
print("\n--- 数据预处理与特征工程 ---")

# 合并训练集和测试集以便统一处理特征 (注意：Survived列只在训练集中)
train_labels = train_df['Survived']
train_features = train_df.drop('Survived', axis=1)
all_features_df = pd.concat([train_features, test_df], ignore_index=True)

print(f"合并后的数据集形状: {all_features_df.shape}")

# 特征工程和缺失值处理

# 1. Pclass: 无需处理，已经是数值型。

# 2. Name: 不再提取 Title。Name 列将在后续被删除。

# 3. Sex: 转换为数值
all_features_df['Sex'] = all_features_df['Sex'].map({'male': 0, 'female': 1}).astype(int)

# 4. Age: 缺失值填充 (改为按 Pclass 的中位数填充)
print("\n--- 填充 Age 缺失值 (按 Pclass 中位数) ---")
age_median_by_pclass = all_features_df.groupby('Pclass')['Age'].median()
print("各 Pclass 对应的 Age 中位数:")
print(age_median_by_pclass)

for p_class in all_features_df['Pclass'].unique():
    if p_class in age_median_by_pclass:
        all_features_df.loc[(all_features_df['Age'].isnull()) & (all_features_df['Pclass'] == p_class), 'Age'] = age_median_by_pclass[p_class]
    else:
        # 理论上 Pclass 1,2,3 应该都有中位数，这是一个备用方案
        median_age_overall = all_features_df['Age'].median()
        print(f"警告: Pclass {p_class} 没有计算出 Age 中位数，将使用全局中位数 {median_age_overall} 进行填充。")
        all_features_df.loc[(all_features_df['Age'].isnull()) & (all_features_df['Pclass'] == p_class), 'Age'] = median_age_overall

# 确保所有 Age 缺失值都被填充 (如果某个Pclass的Age全为NaN，则上面的循环可能未填充)
if all_features_df['Age'].isnull().any():
    print(f"Age 仍有 {all_features_df['Age'].isnull().sum()} 个缺失值，用全局中位数填充剩余部分。")
    all_features_df['Age'].fillna(all_features_df['Age'].median(), inplace=True)
else:
    print("Age 缺失值已全部填充。")


# 5. SibSp 和 Parch: 可以创建 FamilySize
all_features_df['FamilySize'] = all_features_df['SibSp'] + all_features_df['Parch'] + 1
all_features_df['IsAlone'] = (all_features_df['FamilySize'] == 1).astype(int)

# 6. Ticket: 结构复杂，这里暂时不进行深度处理，直接删除。

# 7. Fare: 测试集中有一个缺失值，用中位数填充
all_features_df['Fare'].fillna(all_features_df['Fare'].median(), inplace=True)

# 8. Cabin: 缺失值很多。创建一个新特征 HasCabin
all_features_df['HasCabin'] = all_features_df['Cabin'].notna().astype(int)

# 9. Embarked: 训练集中有两个缺失值，用众数填充
embarked_mode = all_features_df['Embarked'].mode()[0]
all_features_df['Embarked'].fillna(embarked_mode, inplace=True)
print(f"\nEmbarked 缺失值用众数 '{embarked_mode}' 填充。")

# 选择要用于模型的特征
# 删除原始的 Name, Ticket, Cabin, PassengerId (从特征集中)
# SibSp, Parch 已被 FamilySize 和 IsAlone 替代或包含
features_to_drop = ['Name', 'Ticket', 'Cabin', 'PassengerId', 'SibSp', 'Parch']
all_features_processed = all_features_df.drop(columns=features_to_drop)
print("\n处理后的特征列:")
print(all_features_processed.columns)

# 将分类特征进行独热编码 (现在只有 Embarked 是 object 类型需要编码)
categorical_cols = all_features_processed.select_dtypes(include='object').columns.tolist()
print(f"\n需要独热编码的分类特征: {categorical_cols}") # 应该主要是 ['Embarked']

all_features_encoded = pd.get_dummies(all_features_processed, columns=categorical_cols, dummy_na=False)
print("\n独热编码后的数据前5行:")
print(all_features_encoded.head())
print(f"编码后的特征数量: {all_features_encoded.shape[1]}")


# 分离回训练集和测试集
num_train_samples = len(train_labels)
X_train_processed = all_features_encoded.iloc[:num_train_samples]
X_test_processed = all_features_encoded.iloc[num_train_samples:]
y_train = train_labels # Survived 列

feature_names_final = X_train_processed.columns.tolist()
class_names_titanic = ['Not Survived', 'Survived'] # 0 和 1

print(f"\n最终训练特征形状: {X_train_processed.shape}")
print(f"最终测试特征形状: {X_test_processed.shape}")


# (2) 导入 Scikit-learn 提供的决策树模块 DecisionTreeClassifier() (已导入)
# (4) 构造决策树模型
print("\n--- 构建和评估决策树模型 ---")
# 基础模型
dt_clf = DecisionTreeClassifier(random_state=42)
dt_clf.fit(X_train_processed, y_train)

# (5) 模型的评估和预测
# 训练集准确率
train_accuracy_titanic = dt_clf.score(X_train_processed, y_train)
print(f"基础决策树在训练集上的准确率: {train_accuracy_titanic:.4f} (可能过拟合)")

# 使用交叉验证评估模型
cv_scores = cross_val_score(DecisionTreeClassifier(random_state=42), X_train_processed, y_train, cv=5)
print(f"交叉验证平均准确率 (基础模型): {np.mean(cv_scores):.4f} +/- {np.std(cv_scores):.4f}")


# (6) 模型参数的选择、调整（选择最优参数）- 使用 GridSearchCV
print("\n--- 使用 GridSearchCV 进行参数调优 ---")
param_grid = {
    'criterion': ['gini', 'entropy'],
    'max_depth': [3, 5, 7, 10, None], # 决策树最大深度
    'min_samples_split': [2, 5, 10, 20], # 内部节点再划分所需最小样本数
    'min_samples_leaf': [1, 2, 5, 10],   # 叶节点最少样本数
    'max_features': ['sqrt', 'log2', None] # 划分时考虑的最大特征数
}

grid_search = GridSearchCV(DecisionTreeClassifier(random_state=42), param_grid, cv=5, scoring='accuracy', n_jobs=-1)
grid_search.fit(X_train_processed, y_train)

best_params = grid_search.best_params_
best_score = grid_search.best_score_
print(f"最佳参数: {best_params}")
print(f"最佳交叉验证准确率: {best_score:.4f}")

# 使用最佳参数训练最终模型
best_dt_clf = grid_search.best_estimator_

# 评估调优后的模型
train_accuracy_best = best_dt_clf.score(X_train_processed, y_train)
print(f"调优后决策树在训练集上的准确率: {train_accuracy_best:.4f}")

# 选择对分类结果有关键作用的特征 (特征重要性)
print("\n--- 特征重要性 ---")
importances = best_dt_clf.feature_importances_
feature_importance_df = pd.DataFrame({
    'feature': feature_names_final,
    'importance': importances
}).sort_values(by='importance', ascending=False)

print("最重要的特征:")
print(feature_importance_df.head(10))

plt.figure(figsize=(10, max(6, len(feature_names_final) // 2))) # 调整图像大小以适应特征数量
sns.barplot(x='importance', y='feature', data=feature_importance_df.head(min(15, len(feature_names_final))))
plt.title('Top Feature Importances from Best Decision Tree (Max 15)')
# plt.show() # 在Jupyter中可取消注释
plt.savefig('titanic_feature_importances_no_title.png', bbox_inches='tight')
print("特征重要性图已保存为 titanic_feature_importances_no_title.png")


# (7) 决策树的可视化、用 graphviz 将决策树呈现出来。
print("\n--- 可视化决策树 (调优后，可能限制深度以便查看) ---")
try:
    # 可视化一个限制深度的版本，以便观察
    dot_data_best_viz = export_graphviz(best_dt_clf, out_file=None,
                                    feature_names=feature_names_final,
                                    class_names=class_names_titanic, # 使用 ['Not Survived', 'Survived']
                                    filled=True, rounded=True,
                                    special_characters=True,
                                    max_depth=5, # 限制可视化深度为3
                                    fontname="Microsoft YaHei") # 尝试指定中文字体
    graph_best_viz = graphviz.Source(dot_data_best_viz, encoding='utf-8')
    graph_best_viz.render("titanic_decision_tree_best_depth3_no_title", view=False, format='png')
    print(f"调优后决策树 (深度限制为3的可视化) 已保存为 titanic_decision_tree_best_depth3_no_title.png")
    # if graph_best_viz: display(graph_best_viz) # 在Jupyter中取消注释以显示

    # 如果想可视化完整的最佳树 (如果其 max_depth 不是 None 且不太大)
    actual_max_depth = best_dt_clf.get_depth()
    if best_params.get('max_depth') is not None and best_params.get('max_depth') <= 5:
        dot_data_full_best = export_graphviz(best_dt_clf, out_file=None,
                                        feature_names=feature_names_final,
                                        class_names=class_names_titanic,
                                        filled=True, rounded=True,
                                        special_characters=True,
                                        fontname="Microsoft YaHei")
        graph_full_best = graphviz.Source(dot_data_full_best, encoding='utf-8')
        graph_full_best.render("titanic_decision_tree_best_full_no_title", view=False, format='png')
        print(f"调优后完整决策树 (实际深度: {actual_max_depth}) 已保存为 titanic_decision_tree_best_full_no_title.png")
        # if graph_full_best: display(graph_full_best)
    else:
        print(f"调优后的完整树实际深度为 {actual_max_depth}，可能较深，未生成完整可视化。上面已生成深度为3的版本。")

except Exception as e:
    print(f"决策树可视化失败: {e}.")
    print("请确保Graphviz已正确安装并配置在系统PATH中，并且支持所选字体（如 Microsoft YaHei）。")


# 预测测试集 (如果需要提交结果)
print("\n--- 预测测试集 ---")
test_predictions = best_dt_clf.predict(X_test_processed)

# 创建提交文件
submission_df = pd.DataFrame({
    'PassengerId': test_passenger_ids,
    'Survived': test_predictions
})
submission_df.to_csv('titanic_submission_dt_no_title.csv', index=False)
print("预测结果已保存到 'titanic_submission_dt_no_title.csv'")
