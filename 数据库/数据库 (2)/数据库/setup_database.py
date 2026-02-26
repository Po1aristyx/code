from sql_connection import connect_to_sql_server

def execute_sql_file(file_path):
    conn, cursor = connect_to_sql_server()
    if not conn or not cursor:
        print("无法连接到数据库")
        return
    
    try:
        # 读取SQL文件
        with open(file_path, 'r', encoding='utf-8') as file:
            sql_script = file.read()
        
        # 分割并执行SQL命令
        commands = sql_script.split(';')
        for command in commands:
            if command.strip():
                cursor.execute(command)
        
        # 提交事务
        conn.commit()
        print("数据库表创建成功！")
        
    except Exception as e:
        print(f"执行SQL脚本时出错: {str(e)}")
        conn.rollback()
    
    finally:
        cursor.close()
        conn.close()
        print("数据库连接已关闭")

if __name__ == "__main__":
    execute_sql_file("d:\\shujuku\\create_tables.sql") 