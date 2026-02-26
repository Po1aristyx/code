import numpy as np
from scipy import signal
import matplotlib.pyplot as plt

# 设置中文字体，确保图形可以正常显示中文
plt.rcParams['font.sans-serif'] = ['SimHei']  # 'SimHei' 是黑体
plt.rcParams['axes.unicode_minus'] = False  # 解决负号'-'显示为方块的问题

# --- 题一：设计巴特沃斯低通IIR数字滤波器 ---

# 数字滤波器指标
wp_d = 0.2 * np.pi  # 数字通带截止频率 (rad/sample)
ws_d = 0.3 * np.pi  # 数字阻带截止频率 (rad/sample)
Rp = 1              # 通带最大衰减 (dB)
Rs = 15             # 阻带最小衰减 (dB)


fs_for_bilinear = 1.0 # 假设数字系统的采样频率为1Hz，或理解为T=1

Wp_analog = 2 * fs_for_bilinear * np.tan(wp_d / (2 * fs_for_bilinear))
Ws_analog = 2 * fs_for_bilinear * np.tan(ws_d / (2 * fs_for_bilinear))


print(f"数字通带截止频率 wp_d: {wp_d/np.pi:.2f} pi rad/sample")
print(f"数字阻带截止频率 ws_d: {ws_d/np.pi:.2f} pi rad/sample")
print(f"预畸变后的模拟通带截止频率 Wp_analog: {Wp_analog:.4f} rad/s")
print(f"预畸变后的模拟阻带截止频率 Ws_analog: {Ws_analog:.4f} rad/s")

# 计算巴特沃斯滤波器的阶数N和截止频率Wn (模拟)
N, Wn_analog = signal.buttord(Wp_analog, Ws_analog, Rp, Rs, analog=True)
print(f"模拟巴特沃斯滤波器阶数 N = {N}")
print(f"模拟巴特沃斯滤波器截止频率 Wn_analog = {Wn_analog:.4f} rad/s")

# 设计模拟巴特沃斯滤波器原型
b_analog, a_analog = signal.butter(N, Wn_analog, btype='low', analog=True, output='ba')
print("模拟滤波器系数:")
print(f"b_analog = {b_analog}")
print(f"a_analog = {a_analog}")

# 用双线性变换法将模拟滤波器转换为数字滤波器
# fs 参数是数字系统的采样频率，与预畸变时使用的 fs_for_bilinear 对应
b_digital, a_digital = signal.bilinear(b_analog, a_analog, fs=fs_for_bilinear)
print("\n数字滤波器系数:")
print(f"b_digital = {b_digital}")
print(f"a_digital = {a_digital}")

# 计算数字滤波器的频率响应
w_rad, h_freq_resp = signal.freqz(b_digital, a_digital, worN=8000) # w_rad 是 rad/sample

# 绘制幅频响应和相频响应
plt.figure(figsize=(12, 8))

plt.subplot(2, 1, 1)
plt.plot(w_rad / np.pi, 20 * np.log10(abs(h_freq_resp)))
plt.title('数字滤波器幅频响应')
plt.xlabel('归一化频率 (×π rad/sample)')
plt.ylabel('幅度 (dB)')
plt.grid(True)
plt.axvline(wp_d / np.pi, color='green', linestyle='--', label=f'通带截止 {wp_d/np.pi:.2f}π')
plt.axvline(ws_d / np.pi, color='red', linestyle='--', label=f'阻带起始 {ws_d/np.pi:.2f}π')
plt.axhline(-Rp, color='green', linestyle=':', label=f'-{Rp} dB')
plt.axhline(-Rs, color='red', linestyle=':', label=f'-{Rs} dB')
plt.legend()
plt.ylim([-60, 5]) # 根据实际情况调整Y轴范围

plt.subplot(2, 1, 2)
angles = np.unwrap(np.angle(h_freq_resp))
plt.plot(w_rad / np.pi, angles)
plt.title('数字滤波器相频响应')
plt.xlabel('归一化频率 (×π rad/sample)')
plt.ylabel('相位 (radians)')
plt.grid(True)
plt.axvline(wp_d / np.pi, color='green', linestyle='--')
plt.axvline(ws_d / np.pi, color='red', linestyle='--')

plt.tight_layout()
plt.show()




# --- 题二：用所设计的滤波器对实际心电图信号采样序列进行滤波处理 ---
x_n_ecg = np.array([-4, -2, 0, -4, -6, -4, -2, -4, -6, -6, -4, -4, -6, -6, -2,
                    6, 12, 8, 0, -16, -38, -60, -84, -90, -66, -32, -4, -2,
                    -4, 8, 12, 12, 10, 6, 6, 6, 4, 0, 0, 0, 0, 0, -2, -4,
                    0, 0, 0, -2, -2, 0, 0, -2, -2, -2, -2, 0], dtype=float)

# 对心电信号进行滤波
y_n_ecg_filtered = signal.lfilter(b_digital, a_digital, x_n_ecg)

# 绘制滤波前后的心电图信号
n_samples = len(x_n_ecg)
n_axis = np.arange(n_samples)

plt.figure(figsize=(12, 8))

plt.subplot(2, 1, 1)
plt.plot(n_axis, x_n_ecg, label='原始心电信号 x(n)')
plt.title('滤波前的心电信号')
plt.xlabel('采样点 n')
plt.ylabel('幅度')
plt.grid(True)
plt.legend()

plt.subplot(2, 1, 2)
plt.plot(n_axis, y_n_ecg_filtered, label='滤波后的心电信号 y(n)', color='red')
plt.title('滤波后的心电信号 (低通滤波)')
plt.xlabel('采样点 n')
plt.ylabel('幅度')
plt.grid(True)
plt.legend()

plt.tight_layout()
plt.show()