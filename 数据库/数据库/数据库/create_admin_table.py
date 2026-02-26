from sql_connection import connect_to_sql_server

def create_admin_table():
    conn, cursor = connect_to_sql_server()
    if not conn or not cursor:
        print("无法连接到数据库")
        return
        
    try:
        # 创建admin表
        cursor.execute("""
            IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'admin')
            CREATE TABLE admin (
                username NVARCHAR(50) PRIMARY KEY,
                password NVARCHAR(50) NOT NULL
            )
        """)
        
        # 检查是否已存在管理员账户
        cursor.execute("SELECT COUNT(*) FROM admin WHERE username = 'admin'")
        count = cursor.fetchone()[0]
        
        if count == 0:
            # 插入管理员账户
            cursor.execute("""
                INSERT INTO admin (username, password)
                VALUES (?, ?)
            """, ('admin', '123'))
            
        conn.commit()
        print("管理员表创建成功，初始账户已设置")
        
    except Exception as e:
        print(f"创建管理员表时出错: {str(e)}")
        conn.rollback()
        
    finally:
        cursor.close()
        conn.close()

if __name__ == "__main__":
    create_admin_table() 