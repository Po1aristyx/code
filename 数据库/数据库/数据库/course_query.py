from PyQt5.QtWidgets import QWidget, QHeaderView, QTableWidgetItem
from PyQt5 import uic
from sql_connection import connect_to_sql_server

class CourseQuery(QWidget):
    def __init__(self, student_id):
        super().__init__()
        self.student_id = student_id
        
        # 从 .ui 文件加载 UI
        uic.loadUi('ui/course_query.ui', self)
        
        self.init_ui()

    def init_ui(self):
        self.setup_table()
        self.load_courses()
        self.stylize_ui()

    def setup_table(self):
        # 设置表格列
        self.course_table.setColumnCount(6)
        self.course_table.setHorizontalHeaderLabels(['课程号', '课程名', '学分', '学时', '考核方式', '任课教师'])
        # 设置表格样式
        self.course_table.horizontalHeader().setSectionResizeMode(QHeaderView.Stretch)
        self.course_table.setAlternatingRowColors(True)
        self.course_table.setStyleSheet("""
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

    def load_courses(self):
        # 从数据库加载课程
        conn, cursor = connect_to_sql_server()
        if not conn:
            return

        try:
            # SQL 查询：连接 cjb 和 kcb 表，根据学生ID筛选该学生选择的课程
            query = """
                SELECT a.F1, a.F3, a.F5, a.F4, a.F6, a.F7
                FROM kcb a
                JOIN cjb b ON a.F1 = b.F2
                WHERE b.F1 = ?
            """
            cursor.execute(query, self.student_id)
            records = cursor.fetchall()

            self.course_table.setRowCount(len(records))
            for row_index, row_data in enumerate(records):
                for col_index, col_data in enumerate(row_data):
                    self.course_table.setItem(row_index, col_index, QTableWidgetItem(str(col_data)))
        except Exception as e:
            print(f"Error loading courses: {e}")
        finally:
            cursor.close()
            conn.close()
    
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
        """)
