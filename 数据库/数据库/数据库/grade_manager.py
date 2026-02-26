from PyQt5.QtWidgets import (QWidget, QVBoxLayout, QHBoxLayout, QPushButton, 
                           QTableWidget, QTableWidgetItem, QHeaderView, QMessageBox,
                           QLabel, QLineEdit, QFormLayout)
from PyQt5.QtCore import Qt
from sql_connection import connect_to_sql_server

class GradeManager(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        layout = QVBoxLayout()
        
        # 创建按钮布局
        button_layout = QHBoxLayout()
        
        # 创建功能按钮
        self.add_btn = QPushButton('添加成绩')
        self.modify_btn = QPushButton('修改成绩')
        self.delete_btn = QPushButton('删除成绩')
        
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
        """
        self.add_btn.setStyleSheet(button_style)
        self.modify_btn.setStyleSheet(button_style)
        self.delete_btn.setStyleSheet(button_style)
        
        # 添加按钮到布局
        button_layout.addWidget(self.add_btn)
        button_layout.addWidget(self.modify_btn)
        button_layout.addWidget(self.delete_btn)
        button_layout.addStretch()
        
        layout.addLayout(button_layout)
        
        # 创建表单区域
        self.form_layout = QFormLayout()
        
        # 创建输入框
        self.inputs = {
            'student_id': QLineEdit(),    # 学号
            'course_id': QLineEdit(),     # 课程号
            'score': QLineEdit()          # 成绩
        }
        
        # 添加输入框到表单
        self.form_layout.addRow('学号:', self.inputs['student_id'])
        self.form_layout.addRow('课程号:', self.inputs['course_id'])
        self.form_layout.addRow('成绩:', self.inputs['score'])
        
        layout.addLayout(self.form_layout)
        
        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(5)
        self.table.setHorizontalHeaderLabels(['学号', '姓名', '课程号', '课程名', '成绩'])
        
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
        
        # 连接信号
        self.add_btn.clicked.connect(self.add_grade)
        self.modify_btn.clicked.connect(self.modify_grade)
        self.delete_btn.clicked.connect(self.delete_grade)
        self.table.itemClicked.connect(self.load_grade)
        
        # 加载成绩数据
        self.load_grades()
        
    def load_grades(self):
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                cursor.execute("""
                    SELECT x.F1, s.F2, x.F2, k.F2, x.F3
                    FROM xxxx x
                    JOIN xsxx s ON x.F1 = s.F1
                    JOIN kc k ON x.F2 = k.F1
                    ORDER BY x.F1, x.F2
                """)
                
                grades = cursor.fetchall()
                
                self.table.setRowCount(0)
                if grades:
                    self.table.setRowCount(len(grades))
                    for row, grade in enumerate(grades):
                        for col, value in enumerate(grade):
                            item = QTableWidgetItem(str(value))
                            item.setTextAlignment(Qt.AlignCenter)
                            self.table.setItem(row, col, item)
                            
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'加载成绩信息失败: {str(e)}')
            
    def load_grade(self, item):
        row = item.row()
        self.inputs['student_id'].setText(self.table.item(row, 0).text())
        self.inputs['course_id'].setText(self.table.item(row, 2).text())
        self.inputs['score'].setText(self.table.item(row, 4).text())
            
    def add_grade(self):
        try:
            # 获取输入数据
            student_id = self.inputs['student_id'].text().strip()
            course_id = self.inputs['course_id'].text().strip()
            score = self.inputs['score'].text().strip()
            
            # 验证必填字段
            if not all([student_id, course_id, score]):
                QMessageBox.warning(self, '警告', '请填写所有字段')
                return
                
            # 验证成绩
            try:
                score_value = float(score)
                if not (0 <= score_value <= 100):
                    QMessageBox.warning(self, '警告', '成绩必须在0-100之间')
                    return
            except ValueError:
                QMessageBox.warning(self, '警告', '请输入有效的成绩')
                return
                
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 检查学号是否存在
                cursor.execute("""
                    SELECT COUNT(*)
                    FROM xsxx
                    WHERE F1 = ?
                """, (student_id,))
                
                if cursor.fetchone()[0] == 0:
                    QMessageBox.warning(self, '警告', '该学号不存在')
                    return
                    
                # 检查课程号是否存在
                cursor.execute("""
                    SELECT COUNT(*)
                    FROM kc
                    WHERE F1 = ?
                """, (course_id,))
                
                if cursor.fetchone()[0] == 0:
                    QMessageBox.warning(self, '警告', '该课程号不存在')
                    return
                    
                # 检查是否已存在该成绩记录
                cursor.execute("""
                    SELECT COUNT(*)
                    FROM xxxx
                    WHERE F1 = ? AND F2 = ?
                """, (student_id, course_id))
                
                if cursor.fetchone()[0] > 0:
                    QMessageBox.warning(self, '警告', '该成绩记录已存在')
                    return
                    
                # 插入成绩记录
                cursor.execute("""
                    INSERT INTO xxxx (F1, F2, F3)
                    VALUES (?, ?, ?)
                """, (student_id, course_id, score_value))
                
                conn.commit()
                QMessageBox.information(self, '成功', '添加成绩成功')
                
                # 清空输入框
                for input_field in self.inputs.values():
                    input_field.clear()
                    
                # 刷新表格
                self.load_grades()
                
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'添加成绩失败: {str(e)}')
            
    def modify_grade(self):
        try:
            # 获取输入数据
            student_id = self.inputs['student_id'].text().strip()
            course_id = self.inputs['course_id'].text().strip()
            score = self.inputs['score'].text().strip()
            
            # 验证必填字段
            if not all([student_id, course_id, score]):
                QMessageBox.warning(self, '警告', '请填写所有字段')
                return
                
            # 验证成绩
            try:
                score_value = float(score)
                if not (0 <= score_value <= 100):
                    QMessageBox.warning(self, '警告', '成绩必须在0-100之间')
                    return
            except ValueError:
                QMessageBox.warning(self, '警告', '请输入有效的成绩')
                return
                
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 检查成绩记录是否存在
                cursor.execute("""
                    SELECT COUNT(*)
                    FROM xxxx
                    WHERE F1 = ? AND F2 = ?
                """, (student_id, course_id))
                
                if cursor.fetchone()[0] == 0:
                    QMessageBox.warning(self, '警告', '该成绩记录不存在')
                    return
                    
                # 更新成绩
                cursor.execute("""
                    UPDATE xxxx
                    SET F3 = ?
                    WHERE F1 = ? AND F2 = ?
                """, (score_value, student_id, course_id))
                
                conn.commit()
                QMessageBox.information(self, '成功', '修改成绩成功')
                
                # 清空输入框
                for input_field in self.inputs.values():
                    input_field.clear()
                    
                # 刷新表格
                self.load_grades()
                
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'修改成绩失败: {str(e)}')
            
    def delete_grade(self):
        try:
            # 获取输入数据
            student_id = self.inputs['student_id'].text().strip()
            course_id = self.inputs['course_id'].text().strip()
            
            # 验证必填字段
            if not all([student_id, course_id]):
                QMessageBox.warning(self, '警告', '请填写学号和课程号')
                return
                
            reply = QMessageBox.question(
                self, '确认',
                f'确定要删除该成绩记录吗？',
                QMessageBox.Yes | QMessageBox.No,
                QMessageBox.No
            )
            
            if reply == QMessageBox.No:
                return
                
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 检查成绩记录是否存在
                cursor.execute("""
                    SELECT COUNT(*)
                    FROM xxxx
                    WHERE F1 = ? AND F2 = ?
                """, (student_id, course_id))
                
                if cursor.fetchone()[0] == 0:
                    QMessageBox.warning(self, '警告', '该成绩记录不存在')
                    return
                    
                # 删除成绩记录
                cursor.execute("""
                    DELETE FROM xxxx
                    WHERE F1 = ? AND F2 = ?
                """, (student_id, course_id))
                
                conn.commit()
                QMessageBox.information(self, '成功', '删除成绩成功')
                
                # 清空输入框
                for input_field in self.inputs.values():
                    input_field.clear()
                    
                # 刷新表格
                self.load_grades()
                
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'删除成绩失败: {str(e)}') 