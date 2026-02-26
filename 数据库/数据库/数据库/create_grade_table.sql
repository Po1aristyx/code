-- 创建成绩查询表
CREATE TABLE cjcx (
    student_id VARCHAR(20),                -- 学号
    total_score DECIMAL(10, 2)            -- 总成绩
);

-- 插入测试数据
INSERT INTO cjcx (student_id, total_score)
VALUES ('2024001', 85.5); 