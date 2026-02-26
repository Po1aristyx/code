-- 删除已存在的表
IF OBJECT_ID('成绩表', 'U') IS NOT NULL
    DROP TABLE 成绩表;
IF OBJECT_ID('选课表', 'U') IS NOT NULL
    DROP TABLE 选课表;
IF OBJECT_ID('学生信息', 'U') IS NOT NULL
    DROP TABLE 学生信息;
IF OBJECT_ID('kc', 'U') IS NOT NULL
    DROP TABLE kc;
IF OBJECT_ID('用户表', 'U') IS NOT NULL
    DROP TABLE 用户表;

-- 创建课程表(kc)
CREATE TABLE kc (
    课程号 VARCHAR(20) PRIMARY KEY,
    课程名 VARCHAR(50) NOT NULL,
    学分 FLOAT NOT NULL,
    学期 VARCHAR(20) NOT NULL
);

-- 创建学生信息表
CREATE TABLE 学生信息 (
    学号 VARCHAR(20) PRIMARY KEY,
    姓名 VARCHAR(20) NOT NULL,
    班级 VARCHAR(20) NOT NULL
);

-- 创建用户表
CREATE TABLE 用户表 (
    用户名 VARCHAR(20) PRIMARY KEY,
    密码 VARCHAR(100) NOT NULL,
    角色 VARCHAR(10) NOT NULL CHECK (角色 IN ('学生', '管理员')),
    学号 VARCHAR(20),
    FOREIGN KEY (学号) REFERENCES 学生信息(学号)
);

-- 创建选课表
CREATE TABLE 选课表 (
    学号 VARCHAR(20),
    课程号 VARCHAR(20),
    PRIMARY KEY (学号, 课程号),
    FOREIGN KEY (课程号) REFERENCES kc(课程号),
    FOREIGN KEY (学号) REFERENCES 学生信息(学号)
);

-- 创建成绩表
CREATE TABLE 成绩表 (
    学号 VARCHAR(20),
    课程号 VARCHAR(20),
    成绩 FLOAT NOT NULL,
    PRIMARY KEY (学号, 课程号),
    FOREIGN KEY (学号, 课程号) REFERENCES 选课表(学号, 课程号)
);

-- 插入测试数据
-- 课程数据
INSERT INTO kc (课程号, 课程名, 学分, 学期) VALUES
('CS001', '计算机导论', 3.0, '2024春季'),
('CS002', '数据库原理', 4.0, '2024春季'),
('CS003', '操作系统', 4.0, '2024春季'),
('CS004', 'Java程序设计', 3.5, '2024春季');

-- 学生信息数据
INSERT INTO 学生信息 (学号, 姓名, 班级) VALUES
('2024001', '张三', '计算机2401'),
('2024002', '李四', '计算机2401'),
('2024003', '王五', '计算机2402');

-- 用户数据
INSERT INTO 用户表 (用户名, 密码, 角色, 学号) VALUES
('admin', 'admin123', '管理员', NULL),
('2024001', '123456', '学生', '2024001'),
('2024002', '123456', '学生', '2024002'),
('2024003', '123456', '学生', '2024003');

-- 选课数据
INSERT INTO 选课表 (学号, 课程号) VALUES
('2024001', 'CS001'),
('2024001', 'CS002'),
('2024002', 'CS002'),
('2024002', 'CS003'),
('2024003', 'CS001'),
('2024003', 'CS002'),
('2024003', 'CS003');

-- 成绩数据
INSERT INTO 成绩表 (学号, 课程号, 成绩) VALUES
('2024001', 'CS001', 85),
('2024001', 'CS002', 92),
('2024002', 'CS002', 88),
('2024002', 'CS003', 90),
('2024003', 'CS001', 95),
('2024003', 'CS002', 87),
('2024003', 'CS003', 91); 