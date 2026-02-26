# 导入所需的库
from sklearn.datasets import fetch_20newsgroups
from sklearn.model_selection import train_test_split
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.naive_bayes import MultinomialNB
from sklearn.metrics import classification_report, accuracy_score # <<<<<<<<<< 导入 accuracy_score (可选，因为report也包含)
import numpy as np

# 1. 加载数据集
print("开始加载 20 Newsgroups 数据集...")
newsgroups_data = fetch_20newsgroups(subset='all', shuffle=True, random_state=42)
print(f"数据集加载完毕。共有 {len(newsgroups_data.data)} 篇文档。")
print(f"共有 {len(newsgroups_data.target_names)} 个类别: {newsgroups_data.target_names}")

X = newsgroups_data.data
y = newsgroups_data.target

# 2. 分割数据集
print("\n开始分割数据集...")
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.25, random_state=42)
print(f"训练集大小: {len(X_train)} 篇文档")
print(f"测试集大小: {len(X_test)} 篇文档")

# 3. 文本特征提取（向量化）
print("\n开始文本特征提取...")
vectorizer = CountVectorizer(lowercase=True,max_df=0.5, min_df=2)
X_train_counts = vectorizer.fit_transform(X_train)
X_test_counts = vectorizer.transform(X_test)
print(f"训练数据向量化后的维度: {X_train_counts.shape}")
print(f"测试数据向量化后的维度: {X_test_counts.shape}")

# 4. 训练朴素贝叶斯模型
print("\n开始训练朴素贝叶斯模型...")
model = MultinomialNB()
model.fit(X_train_counts, y_train)
print("模型训练完毕。")

# 5. 进行预测
print("\n开始在测试集上进行预测...")
y_pred = model.predict(X_test_counts)
print("预测完毕。")

# 6. 评估模型
print("\n模型评估结果 (完整报告):")
# 生成并打印完整的 classification_report 字符串
full_report_str = classification_report(y_test, y_pred, target_names=newsgroups_data.target_names)
print(full_report_str)

# output_dict=True 会让函数返回一个字典
report_dict = classification_report(y_test, y_pred, target_names=newsgroups_data.target_names, output_dict=True)

# 准确率 (Accuracy) - 这是针对所有样本的
overall_accuracy = report_dict['accuracy'] # 'accuracy'键直接给出总体准确率

# 对于多分类问题，通常关注 'macro avg' 或 'weighted avg' 的精确率、召回率、F1
# 'weighted avg' 会根据每个类的样本数量进行加权，通常更常用
weighted_avg_metrics = report_dict['weighted avg']
overall_precision = weighted_avg_metrics['precision']
overall_recall = weighted_avg_metrics['recall']
overall_f1_score = weighted_avg_metrics['f1-score']

print(f"  - 总体准确率 (Accuracy):   {overall_accuracy:.4f}")
print(f"  - 加权平均精确率 (Precision): {overall_precision:.4f}")
print(f"  - 加权平均召回率 (Recall):    {overall_recall:.4f}")
print(f"  - 加权平均F1分数 (F1-score):  {overall_f1_score:.4f}")
