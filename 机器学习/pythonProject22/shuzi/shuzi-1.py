import numpy as np

f1 = np.array([1, 1, 1])
k_f1 = np.arange(0, 3)


f2 = np.array([1, 2, 3])
k_f2 = np.arange(1, 4)
conv_result = np.convolve(f1, f2)
print("题一：f1(k) * f2(k) =", conv_result)
