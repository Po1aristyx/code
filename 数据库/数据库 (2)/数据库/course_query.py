import pyodbc
from PyQt5.QtWidgets import QWidget, QVBoxLayout, QTableWidget, QTableWidgetItem, QLabel, QHeaderView
from sql_connection import connect_to_sql_server

class CourseQuery(QWidget):
    def __init__(self, student_id):
        super().__init__()
        self.student_id = student_id
        self.conn, self.cursor = connect_to_sql_server()
        self.init_ui()
        self.query_student_courses()

    def init_ui(self):
        # 设置窗口标题和大小
        self.setWindowTitle('课程查询')
        self.setMinimumSize(600, 400)

        # 创建主布局
        layout = QVBoxLayout()
        
        # 添加标题标签
        title = QLabel('我的课程信息')
        title.setStyleSheet("""
            QLabel {
                font-size: 18px;
                font-weight: bold;
                color: #333333;
                margin: 10px 0;
            }
        """)
        layout.addWidget(title)

        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(3)
        self.table.setHorizontalHeaderLabels(['课程号', '课程名', '学分'])
        
        # 设置表格样式
        self.table.setStyleSheet("""
            QTableWidget {
                border: 1px solid #ddd;
                border-radius: 5px;
                background-color: white;
            }
            QHeaderView::section {
                background-color: #f5f5f5;
                padding: 5px;
                border: 1px solid #ddd;
                font-weight: bold;
            }
        """)
        
        # 设置列宽自动调整
        header = self.table.horizontalHeader()
        for i in range(3):
            header.setSectionResizeMode(i, QHeaderView.Stretch)
        
        layout.addWidget(self.table)
        self.setLayout(layout)

    def query_student_courses(self):
        try:
            # 查询学生所报课程
            query = """
                SELECT k.F1, k.F2, k.F3
                FROM kc k
                INNER JOIN xxxx x ON k.F1 = x.F2
                WHERE x.F1 = ?
            """
            self.cursor.execute(query, (self.student_id,))
            courses = self.cursor.fetchall()
            
            # 清空表格
            self.table.setRowCount(0)
            
            if not courses:
                return
            
            # 填充表格
            for row, course in enumerate(courses):
                self.table.insertRow(row)
                for col, value in enumerate(course):
                    item = QTableWidgetItem(str(value))
                    self.table.setItem(row, col, item)

        except Exception as e:
            print(f"查询出错: {str(e)}")

    def closeEvent(self, event):
        if self.cursor:
            self.cursor.close()
        if self.conn:
            self.conn.close()
        event.accept()

def main():
    # 测试代码
    course_query = CourseQuery("TEST001")
    
    # 查看表结构
    print("\n查看数据库表结构：")
    course_query.show_table_structure()
    
    # 测试查询所有课程
    print("\n测试查询所有开课情况：")
    course_query.query_courses()
    
    # 测试查询特定学生的课程
    print("\n测试查询学生选课情况：")
    course_query.query_student_courses()
    
    # 关闭连接
    course_query.close()

if __name__ == "__main__":
    main() 