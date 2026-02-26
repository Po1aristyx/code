from PyQt5.QtWidgets import (QWidget, QVBoxLayout, QHBoxLayout, QLabel, 
                           QLineEdit, QPushButton, QMessageBox, QTableWidget,
                           QTableWidgetItem, QHeaderView)
from PyQt5.QtCore import Qt
from sql_connection import connect_to_sql_server

class GradeQuery(QWidget):
    def __init__(self, student_id=None):
        super().__init__()
        self.student_id = student_id  # 如果是学生登录，会传入学号
        self.init_ui()
        if self.student_id:
            self.query_grades()
        
    def init_ui(self):
        # 创建主布局
        layout = QVBoxLayout()
        
        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(5)
        self.table.setHorizontalHeaderLabels(['学号', '姓名', '班级', '课程名', '成绩'])
        self.table.horizontalHeader().setSectionResizeMode(QHeaderView.Stretch)
        self.table.setStyleSheet("""
            QTableWidget {
                border: 1px solid #CCCCCC;
                border-radius: 6px;
                font-size: 14px;
            }
            QHeaderView::section {
                background-color: #f0f0f0;
                padding: 8px;
                border: none;
                border-bottom: 1px solid #CCCCCC;
                font-weight: bold;
            }
        """)
        layout.addWidget(self.table)
        
        # 添加总成绩和平均分显示
        self.summary_label = QLabel()
        self.summary_label.setStyleSheet("""
            QLabel {
                font-size: 14px;
                color: #333333;
                padding: 10px;
                margin-top: 10px;
            }
        """)
        layout.addWidget(self.summary_label)
        
        self.setLayout(layout)

    def query_grades(self):
        # 使用已登录的学生ID
        student_id = self.student_id

        if not student_id:
            return

        try:
            # 连接数据库
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                QMessageBox.critical(self, '错误', '无法连接到数据库')
                return

            # 查询学生基本信息
            cursor.execute("""
                SELECT F2 as name
                FROM xsxx
                WHERE F1 = ?
            """, (student_id,))

            student = cursor.fetchone()
            if not student:
                QMessageBox.warning(self, '警告', '学号不存在！')
                return

            student_name = student[0]

            # 查询学生的所有课程成绩和学分
            cursor.execute("""
                SELECT 
                    x.F1 as student_id,
                    s.F2 as student_name,
                    CAST(s.F4 AS INT) as class,
                    k.F2 as course_name,
                    k.F3 as credit,  
                    x.F3 as score
                FROM xxxx x
                JOIN xsxx s ON x.F1 = s.F1
                JOIN kc k ON x.F2 = k.F1
                WHERE x.F1 = ?
                ORDER BY k.F1
            """, (student_id,))

            grades = cursor.fetchall()

            # 清空表格
            self.table.setRowCount(0)

            if grades:
                # 计算总分和平均分
                total_score = 0
                total_credit = 0
                weighted_score = 0

                # 设置表格列数为6（包含新增的学分列）
                self.table.setColumnCount(6)
                self.table.setHorizontalHeaderLabels(['学号', '姓名', '班级', '课程名称', '学分', '成绩'])

                # 显示每门课程的成绩
                self.table.setRowCount(len(grades))
                for row, grade in enumerate(grades):
                    for col, value in enumerate(grade):
                        # 对于班级列（第3列，索引为2），确保显示为整数
                        if col == 2:  # 班级列
                            value = int(float(value))
                        item = QTableWidgetItem(str(value) if value is not None else '')
                        item.setTextAlignment(Qt.AlignCenter)
                        self.table.setItem(row, col, item)

                    # 计算总分和加权分数
                    score = float(grade[5])  # 成绩在第6列
                    credit = float(grade[4])  # 学分在第5列
                    total_score += score
                    total_credit += credit
                    weighted_score += score * credit

                # 计算平均分
                average_score = total_score / len(grades)
                weighted_average = weighted_score / total_credit if total_credit > 0 else 0

                # 显示统计信息
                summary = f"总学分: {total_credit:.1f}   "
                summary += f"平均分: {average_score:.2f}   "
                summary += f"加权平均分: {weighted_average:.2f}"
                self.summary_label.setText(summary)
            else:
                self.summary_label.setText("暂无成绩记录")

        except Exception as e:
            QMessageBox.critical(self, '错误', f'查询失败: {str(e)}')
        finally:
            # 确保数据库连接关闭
            if 'cursor' in locals():
                cursor.close()
            if 'conn' in locals():
                conn.close()

def main():
    # 测试代码
    grade_query = GradeQuery()
    
    # 测试查询学生成绩
    test_student_id = '1'
    print(f"\n测试查询学号 {test_student_id} 的成绩：")
    grade_query.query_grades()
    
    # 关闭连接
    grade_query.close()

if __name__ == "__main__":
    main() 