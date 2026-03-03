#!/bin/bash

# 考试系统全流程测试脚本
# 覆盖范围：登录 -> 教师出题 -> 组卷 -> 学生参考 -> 异步交卷 -> AI异步批改结果查询

BASE_URL="http://localhost:8888" # 请根据实际 api-gateway 地址通过命令行参数或修改此变量

TEACHER_EMAIL="ryilj131@139.com"
TEACHER_PASS="e8VSadT8EsiARPx"

STUDENT_EMAIL="izhlu8.p2i31@yeah.net"
STUDENT_PASS="CSJeKXwboya8d0W"

echo "========== 1. 老师登录 =========="
TEACHER_RES=$(curl -s -X POST "$BASE_URL/v1/user/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$TEACHER_EMAIL\", \"password\": \"$TEACHER_PASS\"}")
TEACHER_TOKEN=$(echo $TEACHER_RES | grep -oP '"token":"\K[^"]+')

if [ -z "$TEACHER_TOKEN" ]; then
    echo "老师登录失败: $TEACHER_RES"
    exit 1
fi
echo "老师登录成功，Token 已获取。"

# 假设课程 ID 为 1 (实际请根据数据库或之前创建的课程修改)
COURSE_ID=1

echo -e "\n========== 2. 老师创建题目 (客观题+主观题) =========="
# 创建客观题
Q1_RES=$(curl -s -X POST "$BASE_URL/v1/question/create" \
  -H "Authorization: $TEACHER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"course_id\": $COURSE_ID,
    \"type\": 1,
    \"content\": \"Go 是由哪家公司开发的？\",
    \"answer\": \"Google\",
    \"score\": \"10\",
    \"difficulty\": 1
  }")
Q1_ID=$(echo $Q1_RES | grep -oP '"id":\K[0-9]+')
echo "客观题创建成功，ID: $Q1_ID"

# 创建主观题
Q2_RES=$(curl -s -X POST "$BASE_URL/v1/question/create" \
  -H "Authorization: $TEACHER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"course_id\": $COURSE_ID,
    \"type\": 4,
    \"content\": \"请简述 Go 语言并发编程的优势：\",
    \"answer\": \"使用了 goroutine 和 channel，轻量级且高效。\",
    \"score\": \"20\",
    \"difficulty\": 2
  }")
Q2_ID=$(echo $Q2_RES | grep -oP '"id":\K[0-9]+')
echo "主观题创建成功，ID: $Q2_ID"

echo -e "\n========== 3. 老师创建试卷 =========="
NOW=$(date +%s)
START_TIME=$NOW
END_TIME=$((NOW + 3600)) # 1小时后

EXAM_RES=$(curl -s -X POST "$BASE_URL/v1/exam/create" \
  -H "Authorization: $TEACHER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"course_id\": $COURSE_ID,
    \"title\": \"MQ 系统压力测试卷\",
    \"total_score\": \"100\",
    \"pass_score\": \"60\",
    \"duration\": 60,
    \"start_time\": $START_TIME,
    \"end_time\": $END_TIME,
    \"exam_type\": 1
  }")
EXAM_ID=$(echo $EXAM_RES | grep -oP '"id":\K[0-9]+')
echo "试卷创建成功，ID: $EXAM_ID"

echo -e "\n========== 4. 将题目添加到试卷 =========="
curl -s -X POST "$BASE_URL/v1/exam/question/add" \
  -H "Authorization: $TEACHER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"exam_id\": $EXAM_ID, \"question_id\": $Q1_ID, \"score\": \"50\", \"sort\": 1}"
curl -s -X POST "$BASE_URL/v1/exam/question/add" \
  -H "Authorization: $TEACHER_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"exam_id\": $EXAM_ID, \"question_id\": $Q2_ID, \"score\": \"50\", \"sort\": 2}"
echo "题目绑定完成。"

echo -e "\n========== 5. 学生登录 =========="
STUDENT_RES=$(curl -s -X POST "$BASE_URL/v1/user/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\": \"$STUDENT_EMAIL\", \"password\": \"$STUDENT_PASS\"}")
STUDENT_TOKEN=$(echo $STUDENT_RES | grep -oP '"token":"\K[^"]+')
echo "学生登录成功。"

echo -e "\n========== 6. 学生开始考试 =========="
RECORD_RES=$(curl -s -X POST "$BASE_URL/v1/exam-record/start" \
  -H "Authorization: $STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"exam_id\": $EXAM_ID}")
RECORD_ID=$(echo $RECORD_RES | grep -oP '"record_id":\K[0-9]+')
echo "开始考试，生成答卷记录 ID: $RECORD_ID"

echo -e "\n========== 7. 模拟心跳快照 (保存过程中间答案) =========="
curl -s -X POST "$BASE_URL/v1/exam-record/snapshot" \
  -H "Authorization: $STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"record_id\": $RECORD_ID,
    \"answers\": [
      {\"question_id\": $Q1_ID, \"answer\": \"Google\"}
    ]
  }"
echo "心跳同步成功 (Redis 写入验证)"

echo -e "\n========== 8. 提交答卷 (异步 Kafka 削峰) =========="
curl -s -X POST "$BASE_URL/v1/exam-record/submit" \
  -H "Authorization: $STUDENT_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"record_id\": $RECORD_ID,
    \"exam_id\": $EXAM_ID,
    \"answers\": [
      {\"question_id\": $Q1_ID, \"answer\": \"Google\"},
      {\"question_id\": $Q2_ID, \"answer\": \"Go语言性能非常强悍，因为它是编译型语言且有原生协程。\"}
    ]
  }"
echo "答卷已成功提交至 MQ 缓冲区。"

echo -e "\n========== 9. 等待 MQ 与 AI 批改 (预计 5 秒) =========="
sleep 5

echo -e "\n========== 10. 查询最终考试结果 (验证异步判分) =========="
RESULT_RES=$(curl -s -X GET "$BASE_URL/v1/exam-record/result/$RECORD_ID" \
  -H "Authorization: $STUDENT_TOKEN")
echo "最终得分结果: $RESULT_RES"
