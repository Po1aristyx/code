from PyQt5.QtWidgets import (QWidget, QVBoxLayout, QHBoxLayout, QPushButton, 
                           QTableWidget, QTableWidgetItem, QHeaderView, QMessageBox,
                           QLabel, QLineEdit, QFormLayout, QComboBox)
from PyQt5.QtCore import Qt
from sql_connection import connect_to_sql_server

class StudentInfoManager(QWidget):
    def __init__(self):
        super().__init__()
        self.init_ui()
        
    def init_ui(self):
        layout = QVBoxLayout()
        
        # 创建按钮布局
        button_layout = QHBoxLayout()
        
        # 创建功能按钮
        self.add_btn = QPushButton('添加学生')
        self.modify_btn = QPushButton('修改信息')
        self.delete_btn = QPushButton('删除学生')
        
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
            'F1': QLineEdit(),  # 学号
            'F2': QLineEdit(),  # 姓名
            'F3': QComboBox(),  # 性别
            'F4': QLineEdit(),  # 班级
            'F5': QLineEdit()   # 年龄
        }
        
        self.inputs['F3'].addItems(['男', '女'])
        
        # 添加输入框到表单
        self.form_layout.addRow('学号:', self.inputs['F1'])
        self.form_layout.addRow('姓名:', self.inputs['F2'])
        self.form_layout.addRow('性别:', self.inputs['F3'])
        self.form_layout.addRow('班级:', self.inputs['F4'])
        self.form_layout.addRow('年龄:', self.inputs['F5'])
        
        layout.addLayout(self.form_layout)
        
        # 创建表格
        self.table = QTableWidget()
        self.table.setColumnCount(5)
        self.table.setHorizontalHeaderLabels(['学号', '姓名', '性别', '班级', '年龄'])
        
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
        self.add_btn.clicked.connect(self.add_student)
        self.modify_btn.clicked.connect(self.modify_student)
        self.delete_btn.clicked.connect(self.delete_student)
        self.table.itemClicked.connect(self.load_student)
        
        # 加载学生数据
        self.load_students()
        
    def load_students(self):
        try:
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                cursor.execute("""
                    SELECT F1, F2, F3, F4, F5
                    FROM xsxx
                    ORDER BY F1
                """)
                
                students = cursor.fetchall()
                
                self.table.setRowCount(0)
                if students:
                    self.table.setRowCount(len(students))
                    for row, student in enumerate(students):
                        for col, value in enumerate(student):
                            item_text = ""
                            if value is not None:
                                if col in [3, 4]: # 班级和年龄
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
            QMessageBox.critical(self, '错误', f'加载学生信息失败: {str(e)}')
            
    def load_student(self, item):
        row = item.row()
        self.inputs['F1'].setText(self.table.item(row, 0).text())
        self.inputs['F2'].setText(self.table.item(row, 1).text())
        gender = self.table.item(row, 2).text()
        gender_index = self.inputs['F3'].findText(gender)
        if gender_index != -1:
            self.inputs['F3'].setCurrentIndex(gender_index)
        self.inputs['F4'].setText(self.table.item(row, 3).text())
        self.inputs['F5'].setText(self.table.item(row, 4).text())
            
    def add_student(self):
        try:
            # 获取输入数据
            data = {key: input_field.text().strip() if isinstance(input_field, QLineEdit) else input_field.currentText()
                   for key, input_field in self.inputs.items()}
            
            # 验证必填字段
            if not data['F1'] or not data['F2'] or not data['F4'] or not data['F5']:
                QMessageBox.warning(self, '警告', '请填写所有字段')
                return
                
            # 验证年龄
            try:
                age = int(data['F5'])
                if age < 0:
                    QMessageBox.warning(self, '警告', '年龄不能为负数')
                    return
            except ValueError:
                QMessageBox.warning(self, '警告', '请输入有效的年龄')
                return
                
            conn, cursor = connect_to_sql_server()
            if not conn or not cursor:
                return
                
            try:
                # 检查学号是否已存在
                cursor.execute("""
                    SELECT COUNT(*)
                    FROM xsxx
                    WHERE F1 = ?
                """, (data['F1'],))
                
                if cursor.fetchone()[0] > 0:
                    QMessageBox.warning(self, '警告', '该学号已存在')
                    return
                    
                # 插入数据
                cursor.execute("""
                    INSERT INTO xsxx (F1, F2, F3, F4, F5)
                    VALUES (?, ?, ?, ?, ?)
                """, (data['F1'], data['F2'], data['F3'], data['F4'], age))
                
                conn.commit()
                QMessageBox.information(self, '成功', '添加学生成功')
                
                # 清空输入框
                for input_field in self.inputs.values():
                    if isinstance(input_field, QLineEdit):
                        input_field.clear()
                    elif isinstance(input_field, QComboBox):
                        input_field.setCurrentIndex(0)
                    
                # 刷新表格
                self.load_students()
                
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'添加学生失败: {str(e)}')
            
    def modify_student(self):
        try:
            # 获取输入数据
            data = {key: input_field.text().strip() if isinstance(input_field, QLineEdit) else input_field.currentText()
                   for key, input_field in self.inputs.items()}
            
            # 验证必填字段
            if not data['F1'] or not data['F2'] or not data['F4'] or not data['F5']:
                QMessageBox.warning(self, '警告', '请填写所有字段')
                return
                
            # 验证年龄
            try:
                age = int(data['F5'])
                if age < 0:
                    QMessageBox.warning(self, '警告', '年龄不能为负数')
                    return
            except ValueError:
                QMessageBox.warning(self, '警告', '请输入有效的年龄')
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
                """, (data['F1'],))
                
                if cursor.fetchone()[0] == 0:
                    QMessageBox.warning(self, '警告', '该学号不存在')
                    return
                    
                # 更新数据
                cursor.execute("""
                    UPDATE xsxx
                    SET F2 = ?, F3 = ?, F4 = ?, F5 = ?
                    WHERE F1 = ?
                """, (data['F2'], data['F3'], data['F4'], age, data['F1']))
                
                conn.commit()
                QMessageBox.information(self, '成功', '修改学生信息成功')
                
                # 清空输入框
                for input_field in self.inputs.values():
                    if isinstance(input_field, QLineEdit):
                        input_field.clear()
                    elif isinstance(input_field, QComboBox):
                        input_field.setCurrentIndex(0)
                    
                # 刷新表格
                self.load_students()
                
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'修改学生信息失败: {str(e)}')
            
    def delete_student(self):
        try:
            student_id = self.inputs['F1'].text().strip()
            
            if not student_id:
                QMessageBox.warning(self, '警告', '请输入要删除的学号')
                return
                
            reply = QMessageBox.question(
                self, '确认',
                f'确定要删除学号为 {student_id} 的学生吗？',
                QMessageBox.Yes | QMessageBox.No,
                QMessageBox.No
            )
            
            if reply == QMessageBox.No:
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
                    
                # 删除学生
                cursor.execute("""
                    DELETE FROM xsxx
                    WHERE F1 = ?
                """, (student_id,))
                
                conn.commit()
                QMessageBox.information(self, '成功', '删除学生成功')
                
                # 清空输入框
                for input_field in self.inputs.values():
                    if isinstance(input_field, QLineEdit):
                        input_field.clear()
                    elif isinstance(input_field, QComboBox):
                        input_field.setCurrentIndex(0)
                    
                # 刷新表格
                self.load_students()
                
            finally:
                cursor.close()
                conn.close()
                
        except Exception as e:
            QMessageBox.critical(self, '错误', f'删除学生失败: {str(e)}') 