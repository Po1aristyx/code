from PyQt5.QtWidgets import (QMainWindow, QWidget, QVBoxLayout, QHBoxLayout,
                             QLabel, QLineEdit, QPushButton, QMessageBox,
                             QStackedWidget, QFormLayout, QComboBox, QTableWidget,
                             QTableWidgetItem, QHeaderView)
from PyQt5.QtCore import Qt
# 从 main 导入 LoginWindow，需要处理循环引用的问题
# from main import LoginWindow
from sql_connection import connect_to_sql_server
from student_info_manager import StudentInfoManager
from grade_manager import GradeManager
from admin_query import AdminQuery
import main

# A global to hold the LoginWindow class
LoginWindow = None

def set_login_window(login_window_class):
    global LoginWindow
    LoginWindow = login_window_class

class StudentManagementWidget(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        layout = QVBoxLayout()
        
        # 创建输入表单
        form_layout = QVBoxLayout()
        
        # 学号输入
        id_layout = QHBoxLayout()
        self.id_label = QLabel('学号:')
        self.id_input = QLineEdit()
        id_layout.addWidget(self.id_label)
        id_layout.addWidget(self.id_input)
        form_layout.addLayout(id_layout)
        
        # 姓名输入
        name_layout = QHBoxLayout()
        self.name_label = QLabel('姓名:')
        self.name_input = QLineEdit()
        name_layout.addWidget(self.name_label)
        name_layout.addWidget(self.name_input)
        form_layout.addLayout(name_layout)
        
        # 年龄输入
        age_layout = QHBoxLayout()
        self.age_label = QLabel('年龄:')
        self.age_input = QLineEdit()
        age_layout.addWidget(self.age_label)
        age_layout.addWidget(self.age_input)
        form_layout.addLayout(age_layout)
        
        # 班级输入
        class_layout = QHBoxLayout()
        self.class_label = QLabel('班级:')
        self.class_input = QLineEdit()
        class_layout.addWidget(self.class_label)
        class_layout.addWidget(self.class_input)
        form_layout.addLayout(class_layout)
        
        # 性别选择
        gender_layout = QHBoxLayout()
        self.gender_label = QLabel('性别:')
        self.gender_combo = QComboBox()
        self.gender_combo.addItems(['男', '女'])
        gender_layout.addWidget(self.gender_label)
        gender_layout.addWidget(self.gender_combo)
        form_layout.addLayout(gender_layout)
        
        layout.addLayout(form_layout)
        
        # 按钮布局
        button_layout = QHBoxLayout()
        
        self.add_button = QPushButton('添加')
        self.add_button.clicked.connect(self.add_student)
        button_layout.addWidget(self.add_button)
        
        self.modify_button = QPushButton('修改')
        self.modify_button.clicked.connect(self.modify_student)
        button_layout.addWidget(self.modify_button)
        
        self.delete_button = QPushButton('删除')
        self.delete_button.clicked.connect(self.delete_student)
        button_layout.addWidget(self.delete_button)
        
        self.search_button = QPushButton('查询')
        self.search_button.clicked.connect(self.search_student)
        button_layout.addWidget(self.search_button)
        
        self.grade_manager_btn = QPushButton('学生成绩管理')
        self.grade_manager_btn.clicked.connect(self.show_grade_manager)
        button_layout.addWidget(self.grade_manager_btn)
        
        self.back_btn = QPushButton('返回')
        self.back_btn.clicked.connect(self.go_back)
        
        button_layout.addWidget(self.back_btn)
        
        button_layout.addStretch()
        
        layout.addLayout(button_layout)
        
        # 添加表格显示学生信息
        self.table = QTableWidget()
        self.table.setColumnCount(5)
        self.table.setHorizontalHeaderLabels(['学号', '姓名', '性别', '班级', '年龄'])
        self.table.horizontalHeader().setSectionResizeMode(QHeaderView.Stretch)
        layout.addWidget(self.table)
        
        self.setLayout(layout)
        
        # 初始化显示所有学生
        self.load_students()
        
    def load_students(self):
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                QMessageBox.critical(self, '错误', '无法连接到数据库')
                return
                
            # 修改查询顺序：F1(学号), F2(姓名), F3(性别), F4(班级), F5(年龄)
            cursor.execute("SELECT F1, F2, F3, F4, F5 FROM XSXX")
            students = cursor.fetchall()
            
            self.table.setRowCount(len(students))
            for i, student in enumerate(students):
                for j, value in enumerate(student):
                    if j == 4:  # 年龄字段是第5列
                        self.table.setItem(i, j, QTableWidgetItem(str(int(float(value))) if value else ''))
                    elif j == 3:  # 班级字段是第4列
                        class_value = str(value).strip() if value else ''
                        if '.' in class_value:  # 如果是浮点数形式，去掉小数部分
                            class_value = class_value.split('.')[0]
                        self.table.setItem(i, j, QTableWidgetItem(class_value))
                    else:
                        self.table.setItem(i, j, QTableWidgetItem(str(value).strip() if value else ''))
                    
        except Exception as e:
            QMessageBox.critical(self, '错误', f'加载学生信息失败: {str(e)}')
        finally:
            if cursor:
                cursor.close()
            if conn:
                conn.close()
                
    def validate_input(self, student_id, name, age, class_name):
        if not student_id:
            QMessageBox.warning(self, '警告', '学号不能为空')
            return False
            
        if not student_id.isdigit():
            QMessageBox.warning(self, '警告', '请输入数字！')
            return False
            
        if not name:
            QMessageBox.warning(self, '警告', '姓名不能为空')
            return False
            
        if not age:
            QMessageBox.warning(self, '警告', '年龄不能为空')
            return False
            
        try:
            age_num = int(age)
            if age_num <= 0 or age_num > 150:
                QMessageBox.warning(self, '警告', '请输入有效的年龄')
                return False
        except ValueError:
            QMessageBox.warning(self, '警告', '年龄必须是整数')
            return False
            
        if not class_name:
            QMessageBox.warning(self, '警告', '班级不能为空')
            return False
            
        return True
        
    def add_student(self):
        student_id = self.id_input.text().strip()
        name = self.name_input.text().strip()
        age = self.age_input.text().strip()
        class_name = self.class_input.text().strip()
        gender = self.gender_combo.currentText()
        
        if not self.validate_input(student_id, name, age, class_name):
            return
            
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                QMessageBox.critical(self, '错误', '无法连接到数据库')
                return
                
            # 检查学号是否已存在
            cursor.execute("SELECT COUNT(*) FROM XSXX WHERE F1 = ?", (student_id,))
            if cursor.fetchone()[0] > 0:
                QMessageBox.warning(self, '警告', '该学号已存在！请重新输入！')
                return
                
            # 插入新学生记录：F1(学号), F2(姓名), F3(性别), F4(班级), F5(年龄)
            cursor.execute("""
                INSERT INTO XSXX (F1, F2, F3, F4, F5)
                VALUES (?, ?, ?, ?, ?)
            """, (student_id, name, gender, class_name, int(age)))
            
            conn.commit()
            QMessageBox.information(self, '成功', '添加学生记录成功！')
            
            # 刷新表格
            self.load_students()
            
            # 清空输入框
            self.id_input.clear()
            self.name_input.clear()
            self.age_input.clear()
            self.class_input.clear()
            
        except Exception as e:
            QMessageBox.critical(self, '错误', f'添加学生记录失败: {str(e)}')
        finally:
            if cursor:
                cursor.close()
            if conn:
                conn.close()
                
    def search_student(self):
        student_id = self.id_input.text().strip()
        
        if not student_id:
            QMessageBox.warning(self, '警告', '请输入要查询的学号')
            return
            
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                QMessageBox.critical(self, '错误', '无法连接到数据库')
                return
                
            cursor.execute("""
                SELECT F1, F2, F3, F4, F5
                FROM XSXX
                WHERE F1 = ?
            """, (student_id,))
            
            student = cursor.fetchone()
            
            if student:
                self.name_input.setText(str(student[1]).strip() if student[1] else '')
                self.gender_combo.setCurrentText(str(student[2]).strip() if student[2] else '')
                # 处理班级字段
                class_value = str(student[3]).strip() if student[3] else ''
                if '.' in class_value:  # 如果是浮点数形式，去掉小数部分
                    class_value = class_value.split('.')[0]
                self.class_input.setText(class_value)
                self.age_input.setText(str(int(float(student[4]))) if student[4] else '')
            else:
                QMessageBox.warning(self, '警告', '未找到该学号的学生信息')
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'查询学生信息失败: {str(e)}')
        finally:
            if cursor:
                cursor.close()
            if conn:
                conn.close()
                
    def modify_student(self):
        student_id = self.id_input.text().strip()
        name = self.name_input.text().strip()
        age = self.age_input.text().strip()
        class_name = self.class_input.text().strip()
        gender = self.gender_combo.currentText()
        
        if not student_id:
            QMessageBox.warning(self, '警告', '请输入要修改的学号')
            return
            
        if not self.validate_input(student_id, name, age, class_name):
            return
            
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                QMessageBox.critical(self, '错误', '无法连接到数据库')
                return
                
            # 检查学号是否存在
            cursor.execute("SELECT COUNT(*) FROM XSXX WHERE F1 = ?", (student_id,))
            if cursor.fetchone()[0] == 0:
                QMessageBox.warning(self, '警告', '该学号不存在！')
                return
                
            # 更新学生记录：F2(姓名), F3(性别), F4(班级), F5(年龄)
            cursor.execute("""
                UPDATE XSXX
                SET F2 = ?, F3 = ?, F4 = ?, F5 = ?
                WHERE F1 = ?
            """, (name, gender, class_name, int(age), student_id))
            
            conn.commit()
            QMessageBox.information(self, '成功', '修改学生记录成功！')
            
            # 刷新表格
            self.load_students()
            
        except Exception as e:
            QMessageBox.critical(self, '错误', f'修改学生记录失败: {str(e)}')
        finally:
            if cursor:
                cursor.close()
            if conn:
                conn.close()
                
    def delete_student(self):
        student_id = self.id_input.text().strip()
        
        if not student_id:
            QMessageBox.warning(self, '警告', '请输入要删除的学号')
            return
            
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                QMessageBox.critical(self, '错误', '无法连接到数据库')
                return
                
            # 检查学号是否存在
            cursor.execute("SELECT COUNT(*) FROM XSXX WHERE F1 = ?", (student_id,))
            if cursor.fetchone()[0] == 0:
                QMessageBox.warning(self, '警告', '该学号不存在！')
                return
                
            # 删除确认
            reply = QMessageBox.question(self, '确认', '确定要删除该学生记录吗？',
                                       QMessageBox.Yes | QMessageBox.No, QMessageBox.No)
            
            if reply == QMessageBox.Yes:
                # 删除学生记录
                cursor.execute("DELETE FROM XSXX WHERE F1 = ?", (student_id,))
                conn.commit()
                QMessageBox.information(self, '成功', '删除学生记录成功！')
                
                # 清空输入框
                self.id_input.clear()
                self.name_input.clear()
                self.age_input.clear()
                self.class_input.clear()
                
                # 刷新表格
                self.load_students()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'删除学生记录失败: {str(e)}')
        finally:
            if cursor:
                cursor.close()
            if conn:
                conn.close()

    def show_grade_manager(self):
        self.stacked_widget.setCurrentWidget(self.grade_manager_page)
        
    def go_back(self):
        self.login_window = main.LoginWindow()
        self.login_window.show()
        self.close()

class AdminWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.initUI()
        
    def initUI(self):
        # 创建中心部件和布局
        central_widget = QWidget()
        self.setCentralWidget(central_widget)
        layout = QVBoxLayout(central_widget)
        
        # 创建按钮布局
        button_layout = QHBoxLayout()
        
        # 创建功能按钮
        query_btn = QPushButton('查询')
        student_info_btn = QPushButton('学生信息管理')
        grade_btn = QPushButton('学生成绩管理')
        
        # 设置按钮样式
        button_style = """
            QPushButton {
                background-color: #4CAF50;
                color: white;
                padding: 10px 20px;
                border: none;
                border-radius: 6px;
                font-size: 14px;
                min-width: 120px;
            }
            QPushButton:hover {
                background-color: #45a049;
            }
            QPushButton:pressed {
                background-color: #3d8b40;
            }
        """
        query_btn.setStyleSheet(button_style)
        student_info_btn.setStyleSheet(button_style)
        grade_btn.setStyleSheet(button_style)
        
        # 添加按钮到布局
        button_layout.addWidget(query_btn)
        button_layout.addWidget(student_info_btn)
        button_layout.addWidget(grade_btn)
        button_layout.addStretch()
        
        # 添加按钮布局到主布局
        layout.addLayout(button_layout)
        
        # 创建堆叠窗口部件
        self.stacked_widget = QStackedWidget()
        
        # 创建各个功能页面
        self.query_page = AdminQuery()
        self.student_info_page = StudentInfoManager()
        self.grade_page = GradeManager()
        
        # 添加页面到堆叠窗口
        self.stacked_widget.addWidget(self.query_page)
        self.stacked_widget.addWidget(self.student_info_page)
        self.stacked_widget.addWidget(self.grade_page)
        
        # 添加堆叠窗口到主布局
        layout.addWidget(self.stacked_widget)
        
        # 连接按钮信号
        query_btn.clicked.connect(self.show_query)
        student_info_btn.clicked.connect(self.show_student_info)
        grade_btn.clicked.connect(self.show_grade_manager)
        
        # 设置窗口属性
        self.setWindowTitle('学生成绩管理系统 - 管理员')
        self.setGeometry(100, 100, 1000, 600)
        
        # 默认显示查询页面
        self.show_query()
        
        self.back_btn = QPushButton('返回')
        self.back_btn.setStyleSheet(button_style)
        self.back_btn.clicked.connect(self.go_back)
        button_layout.addWidget(self.back_btn)
        
    def show_query(self):
        self.stacked_widget.setCurrentWidget(self.query_page)
        
    def show_student_info(self):
        self.stacked_widget.setCurrentWidget(self.student_info_page)
        
    def show_grade_manager(self):
        self.stacked_widget.setCurrentWidget(self.grade_page) 

    def go_back(self):
        self.login_window = main.LoginWindow()
        self.login_window.show()
        self.close() 