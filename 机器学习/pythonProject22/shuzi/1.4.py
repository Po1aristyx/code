import numpy as np
import matplotlib.pyplot as plt
plt.rcParams['font.sans-serif'] = ['SimHei']  # 指定默认字体为黑体
plt.rcParams['axes.unicode_minus'] = False  # 解决保存图像是负号'-'显示为方块的问题


def dconv(seq1_vals, seq2_vals, k_axis_seq1, k_axis_seq2, title1="序列1", title2="序列2", title_conv="卷积结果"):
    conv_result_vals = np.convolve(seq1_vals, seq2_vals)

    k_conv_start = k_axis_seq1[0] + k_axis_seq2[0]
    k_conv_axis = np.arange(k_conv_start, k_conv_start + len(conv_result_vals))

    plt.figure(figsize=(10, 8))

    plt.subplot(3, 1, 1)
    plt.stem(k_axis_seq1, seq1_vals, basefmt=" ", use_line_collection=True)
    plt.title(title1 + f' (k from {k_axis_seq1[0]} to {k_axis_seq1[-1]})')
    plt.xlabel('k')
    plt.ylabel('值')
    plt.xticks(np.arange(min(k_axis_seq1) - 1, max(k_axis_seq1) + 2, 1))
    plt.yticks(np.arange(0, max(seq1_vals) + 1, 1))  # 调整y轴刻度
    plt.grid(True)

    plt.subplot(3, 1, 2)
    plt.stem(k_axis_seq2, seq2_vals, basefmt=" ", use_line_collection=True)
    plt.title(title2 + f' (k from {k_axis_seq2[0]} to {k_axis_seq2[-1]})')
    plt.xlabel('k')
    plt.ylabel('值')
    plt.xticks(np.arange(min(k_axis_seq2) - 1, max(k_axis_seq2) + 2, 1))
    plt.yticks(np.arange(0, max(seq2_vals) + 1, 1))  # 调整y轴刻度
    plt.grid(True)

    plt.subplot(3, 1, 3)
    plt.stem(k_conv_axis, conv_result_vals, basefmt=" ", use_line_collection=True)
    plt.title(title_conv + f' (k from {k_conv_axis[0]} to {k_conv_axis[-1]})')
    plt.xlabel('k')
    plt.ylabel('值')
    plt.xticks(np.arange(min(k_conv_axis) - 1, max(k_conv_axis) + 2, 1))
    plt.yticks(np.arange(0, max(conv_result_vals) + 1, 1))  # 调整y轴刻度
    plt.grid(True)

    plt.tight_layout()
    plt.show()

    return conv_result_vals, k_conv_axis

k_h_axis = np.arange(0, 4)
h_values = np.array([1, 1, 1, 1])
k_f_axis = np.arange(0, 3)
f_values = np.array([1, 1, 1])
yzs_values, k_yzs_axis = dconv(f_values, h_values, k_f_axis, k_h_axis,
                               title1="激励 f(k)",
                               title2="单位响应 h(k)",
                               title_conv="零状态响应 y_zs(k) = f(k) * h(k)")

print("激励 f(k):")
print("k_f:", k_f_axis)
print("f_values:", f_values)
print("-" * 30)
print("单位响应 h(k):")
print("k_h:", k_h_axis)
print("h_values:", h_values)
print("-" * 30)
print("零状态响应 y_zs(k):")
print("k_yzs:", k_yzs_axis)
print("yzs_values:", yzs_values)