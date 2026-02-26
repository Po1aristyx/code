from sql_connection import connect_to_sql_server

def check_table_structure():
    conn, cursor = connect_to_sql_server()
    if not conn or not cursor:
        print("无法连接到数据库")
        return
        
    try:
        # 获取所有表名
        cursor.execute("""
            SELECT TABLE_NAME 
            FROM INFORMATION_SCHEMA.TABLES 
            WHERE TABLE_TYPE = 'BASE TABLE'
        """)
        tables = cursor.fetchall()
        
        for table in tables:
            table_name = table[0]
            print(f"\n{table_name} 表结构：")
            print("-" * 50)
            
            # 获取列信息
            cursor.execute(f"""
                SELECT COLUMN_NAME, DATA_TYPE, CHARACTER_MAXIMUM_LENGTH
                FROM INFORMATION_SCHEMA.COLUMNS
                WHERE TABLE_NAME = '{table_name}'
                ORDER BY ORDINAL_POSITION
            """)
            columns = cursor.fetchall()
            
            print("列名\t数据类型\t长度")
            print("-" * 50)
            for col in columns:
                print(f"{col[0]}\t{col[1]}\t{col[2]}")
            
            # 显示示例数据
            cursor.execute(f"SELECT TOP 5 * FROM {table_name}")
            rows = cursor.fetchall()
            
            if rows:
                print(f"\n{table_name} 示例数据：")
                print("-" * 50)
                for row in rows:
                    print(row)
            
            print("\n")
            
    except Exception as e:
        print(f"查询出错: {str(e)}")
        
    finally:
        cursor.close()
        conn.close()

def test_specific_query():
    conn, cursor = connect_to_sql_server()
    if not conn or not cursor:
        print("无法连接到数据库")
        return
        
    try:
        # 测试选课查询
        student_id = '2001140105'  # 使用示例学号
        print(f"\n测试选课查询 (学号: {student_id})：")
        print("-" * 50)
        query = """
            SELECT k.课程号, k.课程名, k.学分
            FROM xxxx x
            JOIN kc k ON x.F2 = k.课程号
            WHERE x.F1 = ?
        """
        cursor.execute(query, (student_id,))
        courses = cursor.fetchall()
        for course in courses:
            print(course)
            
        # 测试成绩查询
        print(f"\n测试成绩查询 (学号: {student_id})：")
        print("-" * 50)
        query = """
            SELECT k.课程号, k.课程名, k.学分, x.F3 as 成绩
            FROM xxxx x
            JOIN kc k ON x.F2 = k.课程号
            WHERE x.F1 = ?
        """
        cursor.execute(query, (student_id,))
        grades = cursor.fetchall()
        for grade in grades:
            print(grade)
            
    except Exception as e:
        print(f"查询出错: {str(e)}")
        
    finally:
        cursor.close()
        conn.close()

if __name__ == "__main__":
    check_table_structure()
    
    # 测试具体查询
    print("\n测试具体查询功能：")
    print("=" * 50)
    test_specific_query() 