from PyQt5.QtWidgets import QWidget, QHeaderView, QTableWidgetItem
from PyQt5 import uic
from sql_connection import connect_to_sql_server

class GradeQuery(QWidget):
    def __init__(self, student_id):
        super().__init__()
        self.student_id = student_id
        
        # 从 .ui 文件加载 UI
        uic.loadUi('ui/grade_query.ui', self)
        
        self.init_ui()

    def init_ui(self):
        # 初始化UI组件
        self.setup_table()
        self.load_grades()
        self.calculate_and_display_stats()
        self.stylize_ui()

    def setup_table(self):
        # 设置表格列
        self.grade_table.setColumnCount(4)
        self.grade_table.setHorizontalHeaderLabels(['课程号', '课程名', '学分', '成绩'])
        # 设置表格样式
        self.grade_table.horizontalHeader().setSectionResizeMode(QHeaderView.Stretch)
        self.grade_table.setAlternatingRowColors(True)
        self.grade_table.setStyleSheet("""
            QTableWidget {
                background-color: #f0f0f0;
                border: 1px solid #dcdcdc;
                gridline-color: #dcdcdc;
            }
            QHeaderView::section {
                background-color: #e0e0e0;
                padding: 4px;
                border: 1px solid #dcdcdc;
            }
        """)

    def load_grades(self):
        # 从数据库加载成绩
        conn, cursor = connect_to_sql_server()
        if not conn:
            return

        try:
            query = """
                SELECT a.F2, a.F3, a.F5, b.F3
                FROM kcb a
                JOIN cjb b ON a.F1 = b.F2
                WHERE b.F1 = ?
            """
            cursor.execute(query, self.student_id)
            records = cursor.fetchall()

            self.grade_table.setRowCount(len(records))
            for row_index, row_data in enumerate(records):
                for col_index, col_data in enumerate(row_data):
                    self.grade_table.setItem(row_index, col_index, QTableWidgetItem(str(col_data)))
        except Exception as e:
            print(f"Error loading grades: {e}")
        finally:
            cursor.close()
            conn.close()

    def calculate_and_display_stats(self):
        # 计算并显示统计数据
        total_credits = 0
        total_weighted_score = 0
        grades = []

        for row in range(self.grade_table.rowCount()):
            credit_item = self.grade_table.item(row, 2)
            grade_item = self.grade_table.item(row, 3)

            if credit_item and grade_item:
                try:
                    credit = float(credit_item.text())
                    grade = float(grade_item.text())
                    grades.append(grade)
                    total_credits += credit
                    total_weighted_score += credit * grade
                except ValueError:
                    # 跳过无法转换的行
                    continue

        avg_grade = total_weighted_score / total_credits if total_credits > 0 else 0
        max_grade = max(grades) if grades else 0
        min_grade = min(grades) if grades else 0

        stats_text = f"""
        <b>统计信息：</b><br>
        <hr>
        <b>总学分:</b> {total_credits:.2f}<br>
        <b>平均分:</b> {avg_grade:.2f}<br>
        <b>最高分:</b> {max_grade:.2f}<br>
        <b>最低分:</b> {min_grade:.2f}<br>
        """
        self.stats_display.setHtml(stats_text)

    def stylize_ui(self):
        # 统一设置整个页面的样式
        self.setStyleSheet("""
            QWidget {
                font-family: 'Microsoft YaHei';
                font-size: 14px;
            }
            QLabel#title_label {
                font-size: 20px;
                font-weight: bold;
                color: #333;
                margin: 10px;
            }
            QTextEdit {
                background-color: #f9f9f9;
                border: 1px solid #e0e0e0;
                border-radius: 5px;
                padding: 10px;
            }
        """)