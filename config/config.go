package config

var BaseURL = "http://jw.glutnn.cn/academic/"
var Port = "3010"
var CaptchaPath = "getCaptcha.do?captchaCheckCode=0"
var CheckPath = "checkCaptcha.do?captchaCode="
var LoginPath = "j_acegi_security_check"
var List = "listLeft.do"

// 教室查询
var classroom_query = "teacher/teachresource/roomschedulequery.jsdo?groupId=&moduleId=1070"

// 教室查询数据
var classroom_query_week = "teacher/teachresource/roomschedule_week.jsdo"

// 本学期课表
var Current_semester_schedule = "student/currcourse/currcourse.jsdo?groupId=&moduleId=2000"

// 等级考试成绩
var Level_exam_scores = "student/queryscore/skilltestscore.jsdo"

// 专业教学计划维护
var major_teaching_plan_maintenance = "211"

// 校内学生申请
var on_campus_student_application = "83012"

// 缓考管理
var deferred_exam_management = ""

// 学生申请缓考
var student_apply_for_deferred_exam = "manager/deferredexam/studentApplyIndex.do?groupId=&moduleId=305"

// 综合审查结果
var comprehensive_review_results = "studentcheckscore/studentCheckresultList.do?groupId=&moduleId=430"

// 我的报表
var my_reports = "1600"

// 个人考勤信息
var personal_attendance_information = "7200"

// 学分互认
var credit_recognition = "manager/crossCrediting/schoolStudentApply.do?groupId=&moduleId=83012"

// 学生选课
var student_course_selection = "2050"

// 重修重考报名
var reexamination_registration = "student/remajor/signupsetting.jsdo?groupId=&moduleId=2010"

// post
// year=&term=&prop=&groupName=&para=0&sortColumn=&Submit=%E6%9F%A5%E8%AF%A2
// 个人成绩查询
var Personal_grades_inquiry = "manager/score/studentOwnScore.do?groupId=&moduleId=2021"

// 修改密码
var change_password = "sysmgr/user_password.jsdo?groupId=&moduleId=155"

// 学生考试安排
var Student_exam_arrangement = "student/exam/index.jsdo"

// 学生申请免修免考
var student_apply_for_exemption = ""

// 教学计划管理
var teaching_plan_management = "manager/studyschedule/majorschedule.jsdo?groupId=&moduleId=211"

// 个人教学计划
var personal_teaching_plan = ""

// 学籍信息
var student_status_information = "student/studentinfo/studentInfoModifyIndex.do?frombase=0&wantTag=0&groupId=&moduleId=2060"

// 免修免考申请管理
var exemption_examination_management = ""

// 等级考试报名
var level_exam_registration = ""

// 个人信息
var Person_info = "showPersonalInfo.do"
