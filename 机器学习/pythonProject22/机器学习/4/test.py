import shutil
print(f"Path to dot.exe according to shutil: {shutil.which('dot.exe')}")

import graphviz
print(f"Python graphviz library version: {graphviz.__version__}")

# 尝试用一个最简单的 DOT 字符串来生成图像
try:
    dot_simple = graphviz.Source('digraph { a -> b }')
    dot_simple.render('test_graphviz_output', view=False, format='png')
    print("Successfully rendered 'test_graphviz_output.png' (and .dot file)")
except Exception as e:
    print(f"Error during simple graphviz render: {e}")