from PyQt5.QtWidgets import (QWidget, QVBoxLayout, QHBoxLayout, QPushButton, 
                           QTableWidget, QTableWidgetItem, QHeaderView, QStackedWidget,
                           QLabel, QLineEdit, QMessageBox)
from PyQt5.QtCore import Qt
from sql_connection import connect_to_sql_server

class AdminQuery(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        # 创建主布局
        layout = QVBoxLayout()
        
        # 创建按钮布局
        button_layout = QHBoxLayout()
        
        # 创建三个查询按钮
        self.course_btn = QPushButton('开课查询')
        self.grade_btn = QPushButton('学生成绩表')
        self.personal_btn = QPushButton('个人总成绩表')
        
        # 设置按钮样式
        button_style = """
            QPushButton {
                background-color: #4CAF50;
                color: white;
                padding: 8px 16px;
                border: none;
                border-radius: 4px;
                font-size: 14px;
                min-width: 100px;
            }
            QPushButton:hover {
                background-color: #45a049;
            }
            QPushButton:pressed {
                background-color: #3d8b40;
            }
        """
        self.course_btn.setStyleSheet(button_style)
        self.grade_btn.setStyleSheet(button_style)
        self.personal_btn.setStyleSheet(button_style)
        
        # 添加按钮到布局
        button_layout.addWidget(self.course_btn)
        button_layout.addWidget(self.grade_btn)
        button_layout.addWidget(self.personal_btn)
        button_layout.addStretch()
        
        # 添加按钮布局到主布局
        layout.addLayout(button_layout)
        
        # 创建堆叠窗口部件
        self.stacked_widget = QStackedWidget()
        
        # 创建三个查询页面
        self.course_page = CourseQueryPage()
        self.grade_page = GradeQueryPage()
        self.personal_page = PersonalGradeQueryPage()
        
        # 添加页面到堆叠窗口
        self.stacked_widget.addWidget(self.course_page)
        self.stacked_widget.addWidget(self.grade_page)
        self.stacked_widget.addWidget(self.personal_page)
        
        # 添加堆叠窗口到主布局
        layout.addWidget(self.stacked_widget)
        
        # 设置布局
        self.setLayout(layout)
        
        # 连接按钮信号
        self.course_btn.clicked.connect(self.show_course_query)
        self.grade_btn.clicked.connect(self.show_grade_query)
        self.personal_btn.clicked.connect(self.show_personal_query)
        
        # 默认显示开课查询
        self.show_course_query()
        
    def show_course_query(self):
        self.stacked_widget.setCurrentWidget(self.course_page)
        
    def show_grade_query(self):
        self.stacked_widget.setCurrentWidget(self.grade_page)
        
    def show_personal_query(self):
        self.stacked_widget.setCurrentWidget(self.personal_page)

class CourseQueryPage(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        layout = QVBoxLayout()
        
        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(3)
        self.table.setHorizontalHeaderLabels(['课程号', '课程名', '学分'])
        
        # 设置表格样式
        self.table.setStyleSheet("""
            QTableWidget {
                border: 1px solid #ddd;
                border-radius: 4px;
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
        
        # 查询并显示所有课程
        self.query_courses()
        
    def query_courses(self):
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 查询所有课程信息
                cursor.execute("""
                    SELECT F1, F2, F3
                    FROM kc
                    ORDER BY F1
                """)
                
                courses = cursor.fetchall()
                
                # 清空表格
                self.table.setRowCount(0)
                
                # 填充数据
                if courses:
                    self.table.setRowCount(len(courses))
                    for row, course in enumerate(courses):
                        for col, value in enumerate(course):
                            item = QTableWidgetItem(str(value))
                            item.setTextAlignment(Qt.AlignCenter)
                            self.table.setItem(row, col, item)
                            
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'查询失败: {str(e)}')

class GradeQueryPage(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        layout = QVBoxLayout()
        
        # 创建搜索区域
        search_layout = QHBoxLayout()
        
        # 添加班级输入框
        self.class_label = QLabel('班级:')
        self.class_input = QLineEdit()
        self.class_input.setPlaceholderText('输入班级查询')
        search_layout.addWidget(self.class_label)
        search_layout.addWidget(self.class_input)
        
        # 添加课程号输入框
        self.course_label = QLabel('课程号:')
        self.course_input = QLineEdit()
        self.course_input.setPlaceholderText('输入课程号查询')
        search_layout.addWidget(self.course_label)
        search_layout.addWidget(self.course_input)
        
        # 添加查询按钮
        self.query_btn = QPushButton('查询')
        self.query_btn.setStyleSheet("""
            QPushButton {
                background-color: #4CAF50;
                color: white;
                padding: 8px 16px;
                border: none;
                border-radius: 4px;
            }
            QPushButton:hover {
                background-color: #45a049;
            }
        """)
        self.query_btn.clicked.connect(self.query_grades)
        search_layout.addWidget(self.query_btn)
        
        layout.addLayout(search_layout)
        
        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(5)
        self.table.setHorizontalHeaderLabels(['学号', '姓名', '班级', '课程名', '成绩'])
        
        # 设置表格样式
        self.table.setStyleSheet("""
            QTableWidget {
                border: 1px solid #ddd;
                border-radius: 4px;
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
        for i in range(5):
            header.setSectionResizeMode(i, QHeaderView.Stretch)
            
        layout.addWidget(self.table)
        self.setLayout(layout)
        
    def query_grades(self):
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 获取查询条件
                class_name = self.class_input.text().strip()
                course_id = self.course_input.text().strip()
                
                # 构建查询条件
                query = """
                    SELECT x.F1, s.F2, s.F4, k.F2, x.F3
                    FROM xxxx x
                    JOIN xsxx s ON x.F1 = s.F1
                    JOIN kc k ON x.F2 = k.F1
                    WHERE 1=1
                """
                params = []
                
                if class_name:
                    query += " AND s.F4 = ?"
                    params.append(class_name)
                    
                if course_id:
                    query += " AND k.F1 = ?"
                    params.append(course_id)
                    
                query += " ORDER BY s.F4, x.F1"
                
                # 执行查询
                cursor.execute(query, params)
                grades = cursor.fetchall()
                
                # 清空表格
                self.table.setRowCount(0)
                
                # 填充数据
                if grades:
                    self.table.setRowCount(len(grades))
                    for row, grade in enumerate(grades):
                        for col, value in enumerate(grade):
                            item_text = ""
                            if value is not None:
                                if col == 2:  # 班级列
                                    try:
                                        item_text = str(int(float(str(value))))
                                    except (ValueError, TypeError):
                                        item_text = str(value)
                                else:
                                    item_text = str(value)
                            item = QTableWidgetItem(item_text)
                            item.setTextAlignment(Qt.AlignCenter)
                            self.table.setItem(row, col, item)
                            
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'查询失败: {str(e)}')

class PersonalGradeQueryPage(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        layout = QVBoxLayout()
        
        # 创建搜索区域
        search_layout = QHBoxLayout()
        
        # 添加学号输入框
        self.id_label = QLabel('学号:')
        self.id_input = QLineEdit()
        self.id_input.setPlaceholderText('输入学号查询')
        search_layout.addWidget(self.id_label)
        search_layout.addWidget(self.id_input)
        
        # 添加查询按钮
        self.query_btn = QPushButton('查询')
        self.query_btn.setStyleSheet("""
            QPushButton {
                background-color: #4CAF50;
                color: white;
                padding: 8px 16px;
                border: none;
                border-radius: 4px;
            }
            QPushButton:hover {
                background-color: #45a049;
            }
        """)
        self.query_btn.clicked.connect(self.query_personal_grades)
        search_layout.addWidget(self.query_btn)
        
        layout.addLayout(search_layout)
        
        # 创建学生信息标签
        self.info_label = QLabel()
        self.info_label.setStyleSheet("font-size: 14px; margin: 10px 0;")
        layout.addWidget(self.info_label)
        
        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(3)
        self.table.setHorizontalHeaderLabels(['课程名', '学分', '成绩'])
        
        # 设置表格样式
        self.table.setStyleSheet("""
            QTableWidget {
                border: 1px solid #ddd;
                border-radius: 4px;
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
        
        # 创建统计信息标签
        self.summary_label = QLabel()
        self.summary_label.setStyleSheet("font-size: 14px; margin: 10px 0;")
        layout.addWidget(self.summary_label)
        
        self.setLayout(layout)
        
    def query_personal_grades(self):
        try:
            student_id = self.id_input.text().strip()
            if not student_id:
                QMessageBox.warning(self, '警告', '请输入学号')
                return
                
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 查询学生基本信息
                cursor.execute("""
                    SELECT F2, F4
                    FROM xsxx
                    WHERE F1 = ?
                """, (student_id,))
                
                student = cursor.fetchone()
                if not student:
                    QMessageBox.warning(self, '警告', '学号不存在')
                    return
                    
                # 显示学生信息
                self.info_label.setText(f"姓名：{student[0]}    班级：{student[1]}")
                
                # 查询所有课程成绩
                cursor.execute("""
                    SELECT k.F2, k.F3, x.F3
                    FROM xxxx x
                    JOIN kc k ON x.F2 = k.F1
                    WHERE x.F1 = ?
                    ORDER BY k.F1
                """, (student_id,))
                
                grades = cursor.fetchall()
                
                # 清空表格
                self.table.setRowCount(0)
                
                if grades:
                    # 计算统计信息
                    total_credit = 0
                    weighted_score = 0
                    
                    # 填充数据
                    self.table.setRowCount(len(grades))
                    for row, grade in enumerate(grades):
                        for col, value in enumerate(grade):
                            item = QTableWidgetItem(str(value))
                            item.setTextAlignment(Qt.AlignCenter)
                            self.table.setItem(row, col, item)
                            
                        # 计算统计信息
                        credit = float(grade[1])
                        score = float(grade[2])
                        total_credit += credit
                        weighted_score += credit * score
                        
                    # 计算加权平均分
                    weighted_average = weighted_score / total_credit if total_credit > 0 else 0
                    
                    # 显示统计信息
                    self.summary_label.setText(
                        f"总学分：{total_credit:.1f}    "
                        f"加权平均分：{weighted_average:.2f}"
                    )
                else:
                    self.summary_label.setText("暂无成绩记录")
                    
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'查询失败: {str(e)}') 