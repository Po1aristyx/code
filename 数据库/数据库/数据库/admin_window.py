from PyQt5.QtWidgets import QMainWindow, QWidget
from PyQt5 import uic
from student_info_manager import StudentInfoManager
from grade_manager import GradeManager
from admin_query import AdminQuery

class AdminWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        
        # 从 .ui 文件加载 UI
        uic.loadUi('ui/admin_window.ui', self)
        
        self.init_ui()

    def init_ui(self):
        # 初始化页面
        self.student_info_manager = StudentInfoManager()
        self.grade_manager = GradeManager()
        self.admin_query = AdminQuery()

        # 将页面添加到 stackedWidget
        self.stackedWidget.addWidget(self.student_info_manager)
        self.stackedWidget.addWidget(self.grade_manager)
        self.stackedWidget.addWidget(self.admin_query)

        # 连接菜单操作到槽函数
        self.action_add_student.triggered.connect(lambda: self.show_page(self.student_info_manager, "add"))
        self.action_delete_student.triggered.connect(lambda: self.show_page(self.student_info_manager, "delete"))
        self.action_edit_student.triggered.connect(lambda: self.show_page(self.student_info_manager, "edit"))
        
        # 将查询学生信息连接到 AdminQuery 页面
        self.action_query_student.triggered.connect(self.show_query_page)
        
        self.action_add_grade.triggered.connect(lambda: self.show_page(self.grade_manager, "add"))
        self.action_delete_grade.triggered.connect(lambda: self.show_page(self.grade_manager, "delete"))
        self.action_edit_grade.triggered.connect(lambda: self.show_page(self.grade_manager, "edit"))

        # 将查询学生成绩连接到 AdminQuery 页面
        self.action_query_grade.triggered.connect(self.show_query_page)

        # 设置默认显示的页面
        self.show_query_page()

    def show_page(self, page_widget, mode):
        self.stackedWidget.setCurrentWidget(page_widget)
        if hasattr(page_widget, 'set_mode'):
            page_widget.set_mode(mode)

    def show_query_page(self):
        self.stackedWidget.setCurrentWidget(self.admin_query)
