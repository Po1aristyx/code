from sql_connection import connect_to_sql_server

def create_cjpm_table():
    conn, cursor = connect_to_sql_server()
    if not conn or not cursor:
        print("无法连接到数据库")
        return
        
    try:
        # 创建cjpm表
        cursor.execute("""
            IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'cjpm')
            CREATE TABLE cjpm (
                student_id NVARCHAR(50),
                course_id NVARCHAR(50),
                student_name NVARCHAR(50),
                course_name NVARCHAR(50),
                gender NVARCHAR(10),
                score FLOAT,
                age INT,
                credit FLOAT,
                class_name NVARCHAR(50),
                PRIMARY KEY (student_id, course_id)
            )
        """)
        
        conn.commit()
        print("成绩维护表(cjpm)创建成功")
        
    except Exception as e:
        print(f"创建成绩维护表时出错: {str(e)}")
        conn.rollback()
        
    finally:
        cursor.close()
        conn.close()

if __name__ == "__main__":
    create_cjpm_table() 