import sys
from PyQt5.QtWidgets import (QApplication, QMainWindow, QWidget,QLineEdit,QMessageBox, QButtonGroup)
from PyQt5.QtCore import Qt
from PyQt5.QtGui import QPalette, QBrush, QPixmap
from PyQt5 import uic
from sql_connection import connect_to_sql_server
from grade_query import GradeQuery
from course_query import CourseQuery
import admin_window

class LoginWindow(QWidget):
    def __init__(self):
        super().__init__()
        uic.loadUi('ui/login.ui', self)
        
        # 设置背景图片
        palette = self.palette()
        background = QPixmap("maodie.png")
        if not background.isNull():
            background = background.scaled(self.size(), Qt.IgnoreAspectRatio, Qt.SmoothTransformation)
            palette.setBrush(QPalette.Window, QBrush(background))
            self.setPalette(palette)
            self.setAutoFillBackground(True)
        else:
            print("Warning: Background image 'maodie.png' not found.")

        self.login_type_group = QButtonGroup(self)
        self.login_type_group.addButton(self.student_radio)
        self.login_type_group.addButton(self.admin_radio)
        
        self.login_button.clicked.connect(self.login)
        self.student_radio.toggled.connect(self.update_input_labels)
        
        self.update_input_labels()
        
    def update_input_labels(self):
        if self.student_radio.isChecked():
            self.id_label.setText('学号:')
            self.name_label.setText('姓名:')
            self.name_input.setEchoMode(QLineEdit.Normal)
        else:
            self.id_label.setText('用户名:')
            self.name_label.setText('密码:')
            self.name_input.setEchoMode(QLineEdit.Password)
            
    def login(self):
        conn, cursor = connect_to_sql_server()
        if not conn or not cursor:
            QMessageBox.critical(self, '错误', '无法连接到数据库')
            return
            
        try:
            if self.student_radio.isChecked():
                student_id = self.id_input.text().strip()
                student_name = self.name_input.text().strip()
                
                if not student_id or not student_name:
                    QMessageBox.warning(self, '警告', '请输入学号和姓名')
                    return
                    
                cursor.execute("SELECT COUNT(*) FROM xsxx WHERE F1 = ? AND F2 = ?", (student_id, student_name))
                count = cursor.fetchone()[0]
                
                if count > 0:
                    self.main_window = MainWindow(student_id, student_name)
                    self.main_window.show()
                    self.close()
                else:
                    QMessageBox.warning(self, '登录失败', '学号或姓名错误')
            else:
                username = self.id_input.text().strip()
                password = self.name_input.text().strip()
                
                if not username or not password:
                    QMessageBox.warning(self, '警告', '请输入用户名和密码')
                    return
                    
                cursor.execute("SELECT COUNT(*) FROM admin WHERE username = ? AND password = ?", (username, password))
                count = cursor.fetchone()[0]
                
                if count > 0:
                    self.admin_window = admin_window.AdminWindow()
                    self.admin_window.show()
                    self.close()
                else:
                    QMessageBox.warning(self, '登录失败', '用户名或密码错误')
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'登录失败: {str(e)}')
            
        finally:
            cursor.close()
            conn.close()

class MainWindow(QMainWindow):
    def __init__(self, student_id, student_name):
        super().__init__()
        self.student_id = student_id
        self.student_name = student_name
        
        uic.loadUi('ui/main_window.ui', self)
        
        self.initUI()
        
    def initUI(self):
        # 创建并添加页面
        self.grade_query_page = GradeQuery(self.student_id)
        self.stacked_widget.addWidget(self.grade_query_page)
        
        self.course_query_page = CourseQuery(self.student_id)
        self.stacked_widget.addWidget(self.course_query_page)
        
        # 连接按钮信号
        self.grade_query_btn.clicked.connect(self.show_grade_query)
        self.course_query_btn.clicked.connect(self.show_course_query)
        
        # 设置窗口标题
        self.setWindowTitle(f'学生成绩管理系统 - {self.student_name}')
        
        # 默认显示成绩查询页面
        self.show_grade_query()
        
    def show_grade_query(self):
        self.stacked_widget.setCurrentWidget(self.grade_query_page)
        
    def show_course_query(self):
        self.stacked_widget.setCurrentWidget(self.course_query_page)

if __name__ == '__main__':
    app = QApplication(sys.argv)
    login_win = LoginWindow()
    login_win.show()
    sys.exit(app.exec_())