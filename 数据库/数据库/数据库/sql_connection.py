import pyodbc

def connect_to_sql_server():
    try:
        # 连接参数
        server = r'LAPTOP-HSJUDHJM\POLARIS'  # 服务器名称
        database = 'keshe'  # 数据库名称
        
        # 使用 Windows 身份认证的连接字符串
        conn_str = (
            f'DRIVER={{SQL Server}};'
            f'SERVER={server};'
            f'DATABASE={database};'
            'Trusted_Connection=yes;'  # 使用 Windows 身份认证
        )
        
        # 建立连接
        conn = pyodbc.connect(conn_str)
        print("成功连接到 SQL Server!")
        
        # 创建游标
        cursor = conn.cursor()
        
        return conn, cursor
        
    except Exception as e:
        print(f"连接错误: {str(e)}")
        return None, None

def execute_sql_script(file_path):
    """执行SQL脚本文件"""
    conn, cursor = connect_to_sql_server()
    if conn:
        try:
            # 读取SQL文件
            with open(file_path, 'r', encoding='utf-8') as file:
                sql_script = file.read()
            
            # 分割并执行多个SQL语句
            statements = sql_script.split(';')
            for statement in statements:
                if statement.strip():
                    cursor.execute(statement)
            
            # 提交事务
            conn.commit()
            print(f"成功执行SQL脚本: {file_path}")
            
        except Exception as e:
            print(f"执行SQL脚本错误: {str(e)}")
            conn.rollback()
        finally:
            cursor.close()
            conn.close()
            print("数据库连接已关闭")

def main():
    conn, cursor = connect_to_sql_server()
    if conn:
        try:
            # 测试连接是否成功
            cursor.execute("SELECT @@VERSION")
            row = cursor.fetchone()
            print(f"SQL Server 版本: {row[0]}")
            
        finally:
            # 关闭连接
            cursor.close()
            conn.close()
            print("数据库连接已关闭")

if __name__ == "__main__":
    main() 