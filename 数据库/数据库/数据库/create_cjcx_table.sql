-- 如果表存在则删除
IF OBJECT_ID('cjcx', 'U') IS NOT NULL
    DROP TABLE cjcx;

-- 创建新的cjcx表
CREATE TABLE cjcx (
    student_id NVARCHAR(255),
    student_name NVARCHAR(255),
    gender NVARCHAR(255),
    class_no FLOAT,
    course_id NVARCHAR(255),
    course_name NVARCHAR(255),
    score FLOAT
);

-- 插入数据
INSERT INTO cjcx (student_id, student_name, gender, class_no, course_id, course_name, score)
SELECT 
    s.F1 as student_id,
    s.F2 as student_name,
    s.F3 as gender,
    s.F4 as class_no,
    k.F1 as course_id,
    k.F2 as course_name,
    x.F3 as score
FROM xxxx x
JOIN xsxx s ON x.F1 = s.F1
JOIN kc k ON x.F2 = k.F1; 