import re  # 用于正则表达式处理文本
import random  # 用于随机选择
import numpy as np  # 用于数值计算
import os


def load_emails_from_folder(folder_path):
    """
    从指定文件夹加载所有 .txt 文件的内容。
    参数:
        folder_path (str): 包含 .txt 邮件文件的文件夹路径。
    返回:
        list: 一个包含所有邮件文本内容的列表。
    """
    emails = []
    if not os.path.isdir(folder_path):  # 检查路径是否存在且为文件夹
        print(f"错误：文件夹 '{folder_path}' 不存在或不是一个有效的目录。")
        return emails  # 返回空列表

    for filename in os.listdir(folder_path):  # 遍历文件夹中的所有文件和文件夹
        if filename.endswith(".txt"):  # 只处理 .txt 文件
            file_path = os.path.join(folder_path, filename)  # 构建完整的文件路径
            try:
                with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:  # 以只读模式打开文件
                    # errors='ignore' 会在遇到解码错误时跳过有问题的字符
                    # 你也可以尝试其他编码，如 'latin-1'，如果 'utf-8' 不行
                    emails.append(f.read())  # 读取文件全部内容并添加到列表
            except Exception as e:
                print(f"读取文件 '{file_path}' 时出错: {e}")
    return emails



ham_folder_path = "/email/ham"
spam_folder_path = "/email/spam"


print(f"正在从 '{ham_folder_path}' 加载正常邮件...")
ham_emails_full = load_emails_from_folder(ham_folder_path)
print(f"成功加载 {len(ham_emails_full)} 封正常邮件。")

print(f"正在从 '{spam_folder_path}' 加载垃圾邮件...")
spam_emails_full = load_emails_from_folder(spam_folder_path)
print(f"成功加载 {len(spam_emails_full)} 封垃圾邮件。")

# 检查是否成功加载了邮件
if not ham_emails_full or not spam_emails_full:
    print("\n错误：未能成功加载邮件数据。请检查文件夹路径和文件内容。程序将退出。")
    exit()  # 如果没有加载到邮件，则退出程序


all_emails = ham_emails_full + spam_emails_full
all_labels = [0] * len(ham_emails_full) + [1] * len(spam_emails_full)

email_data = list(zip(all_emails, all_labels))
random.seed(42)  # 设置随机种子以保证结果可复现
random.shuffle(email_data)  # 打乱数据顺序

# 分割训练集和测试集 (假设你希望保持40训练，剩余测试的比例)
# 你可能需要根据实际加载的邮件数量调整这里的分割逻辑
total_emails_loaded = len(all_emails)
num_train_samples = 40
num_test_samples = total_emails_loaded - num_train_samples

train_data = email_data[:num_train_samples]
test_data = email_data[num_train_samples:num_train_samples + num_test_samples]

train_texts = [doc for doc, label in train_data]
train_labels = [label for doc, label in train_data]
test_texts = [doc for doc, label in test_data]
test_labels = [label for doc, label in test_data]

print(f"\n总共加载邮件数量: {total_emails_loaded}")
print(f"训练集邮件数量: {len(train_texts)}")
print(f"测试集邮件数量: {len(test_texts)}")

def tokenize_email(email_text):
    """
    对邮件文本进行分词处理：
    1. 使用正则表达式 \W+ (非字母数字字符) 作为分隔符进行分词。
    2. 将所有词转换为小写。
    3. 返回词列表。
    """
    list_of_tokens = re.split(r'\W+', email_text)
    return [tok.lower() for tok in list_of_tokens if len(tok) > 0]  # 过滤空字符串


def create_vocab_list(dataset_texts):
    """
    创建词汇表：
    遍历所有训练文档，将所有出现过的词（去重）集合起来。
    返回一个包含所有不重复词的列表。
    """
    vocab_set = set([])  # 创建一个空集合
    for document_text in dataset_texts:
        vocab_set = vocab_set | set(tokenize_email(document_text))  # 取并集
    return sorted(list(vocab_set))  # 排序以保证一致性


print("\n正在创建词汇表...")
vocab_list = create_vocab_list(train_texts)
print(f"词汇表大小: {len(vocab_list)}")


def text_to_vector(vocab_list, input_text_tokens):
    """
    将输入的词列表转换为基于词汇表的特征向量（词集模型/二元模型）。
    如果词汇表中的词出现在输入词列表中，则向量对应位置为1，否则为0。
    """
    return_vec = [0] * len(vocab_list)  # 初始化一个长度为词汇表大小的全0向量
    for word in input_text_tokens:
        if word in vocab_list:
            return_vec[vocab_list.index(word)] = 1
    return return_vec


print("\n正在向量化训练数据...")
train_matrix = []  # 存储训练数据的向量
for text in train_texts:
    tokens = tokenize_email(text)
    train_matrix.append(text_to_vector(vocab_list, tokens))


def train_naive_bayes(train_matrix, train_category_labels):
    """
    训练朴素贝叶斯分类器参数。
    计算 P(类别) 和 P(词 | 类别)。
    使用拉普拉斯平滑处理。
    """
    num_train_docs = len(train_matrix)
    num_words = len(train_matrix[0])
    p_spam = sum(train_category_labels) / float(num_train_docs)
    p0_num = np.ones(num_words)
    p1_num = np.ones(num_words)
    count_docs_in_class0 = 0
    count_docs_in_class1 = 0

    for i in range(num_train_docs):
        if train_category_labels[i] == 1:
            p1_num += train_matrix[i]
            count_docs_in_class1 += 1
        else:
            p0_num += train_matrix[i]
            count_docs_in_class0 += 1

    p1_vect = np.log(p1_num / (count_docs_in_class1 + 2.0))
    p0_vect = np.log(p0_num / (count_docs_in_class0 + 2.0))
    return p0_vect, p1_vect, p_spam


print("\n正在训练朴素贝叶斯分类器...")
p0_vec, p1_vec, p_spam_overall = train_naive_bayes(np.array(train_matrix), np.array(train_labels))
print("分类器训练完毕。")


def classify_naive_bayes(test_vector, p0_log_vec, p1_log_vec, p_class1_log_prior):
    log_prob_spam = np.sum(test_vector * p1_log_vec) + np.log(p_class1_log_prior)
    log_prob_ham = np.sum(test_vector * p0_log_vec) + np.log(1.0 - p_class1_log_prior)
    if log_prob_spam > log_prob_ham:
        return 1
    else:
        return 0


print("\n开始在测试集上进行预测和评估...")
predictions = []
misclassified_emails = []
correct_predictions = 0

if not test_texts:  # 检查测试集是否为空
    print("警告：测试集为空，无法进行评估。")
else:
    for i in range(len(test_texts)):
        test_email_text = test_texts[i]
        actual_label = test_labels[i]
        tokenized_test_email = tokenize_email(test_email_text)
        test_email_vector = np.array(text_to_vector(vocab_list, tokenized_test_email))
        predicted_label = classify_naive_bayes(test_email_vector, p0_vec, p1_vec, p_spam_overall)
        predictions.append(predicted_label)

        if predicted_label == actual_label:
            correct_predictions += 1
        else:
            misclassified_emails.append({
                "text": test_email_text,
                "actual": "spam" if actual_label == 1 else "ham",
                "predicted": "spam" if predicted_label == 1 else "ham"
            })

    accuracy = correct_predictions / float(len(test_texts))
    print(f"\n测试集准确率: {accuracy:.4f} ({correct_predictions}/{len(test_texts)})")

    if misclassified_emails:
        print("\n预测错误的邮件:")
        for item in misclassified_emails:
            print(f"  - 内容: \"{item['text'][:100]}...\"")  # 只打印部分内容以防过长
            print(f"    实际类别: {item['actual']}, 预测类别: {item['predicted']}")
    else:
        print("\n所有测试邮件均预测正确！")