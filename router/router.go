package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"nngllgjw/config"
	"nngllgjw/controller"
	"strings"
)

func splitFunc(s, sep string) []string {
	return strings.Split(s, sep)
}
func Router() {

	router := gin.Default()
	store := cookie.NewStore([]byte("secret")) // 设置存储会话的密钥
	router.Use(sessions.Sessions("session", store))
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		jsessionid, img := controller.Captcha()
		c.HTML(200, "login.html", gin.H{
			"jsessionid": jsessionid,
			"image":      img,
			"message":    "请输入账号密码登录",
		})
	})

	router.POST("/", func(c *gin.Context) {
		// 获取表单提交的数据
		captcha := c.PostForm("captcha")
		jsessionid := c.PostForm("jsessionid")
		username := c.PostForm("username")
		password := c.PostForm("password")
		jsessionid = strings.Replace(jsessionid, " ", "", -1)
		result := controller.CheckCaptcha(jsessionid, captcha)
		if result {
			islogin := controller.Login(username, password, jsessionid, captcha)
			if islogin == "" {
				jsessionid, img := controller.Captcha()
				c.HTML(http.StatusOK, "login.html", gin.H{
					"jsessionid": jsessionid,
					"image":      img,
					"message":    "账号或密码错误！请重新输入！",
				})
			} else {
				session := sessions.Default(c)
				session.Set("loggedIn", true)
				session.Set("username", username)
				session.Set("jsessionid", islogin)
				session.Save()
				c.Redirect(http.StatusFound, "/system") // 重定向到 "/system" 页面
			}
		} else {
			jsessionid, img := controller.Captcha()
			c.HTML(http.StatusOK, "login.html", gin.H{
				"jsessionid": jsessionid,
				"image":      img,
				"message":    "验证码错误！请重新输入！",
			})
		}
	})
	router.GET("/system", func(c *gin.Context) {

		session := sessions.Default(c)
		loggedIn := session.Get("loggedIn")
		jsessionid := session.Get("jsessionid")
		if jsessionid == nil {
			c.Redirect(http.StatusFound, "/")

		}
		stringValue := jsessionid.(string)
		cookie_list := &http.Cookie{
			Name:  "JSESSIONID",
			Value: stringValue,
		}
		if loggedIn == true {

			PersonalInfo, _ := controller.GetPersonalInfo(cookie_list)
			sessionS := sessions.Default(c)
			sessionS.Save()
			c.HTML(http.StatusOK, "system.html", gin.H{
				"message":      "欢迎登录系统",
				"PersonalInfo": PersonalInfo,
			})
		} else {
			c.Redirect(http.StatusFound, "/login")
		}
	})

	//router.GET("/coursearrange", func(c *gin.Context) {
	//	session := sessions.Default(c)
	//	loggedIn := session.Get("loggedIn")
	//	jsessionid := session.Get("jsessionid")
	//	stringValue := jsessionid.(string)
	//	cookie_list := &http.Cookie{
	//		Name:  "JSESSIONID",
	//		Value: stringValue,
	//	}
	//	if loggedIn == true {
	//		type CourseSchedule struct {
	//			Days      [][]string
	//			Timeslots [][]string
	//		}
	//		CurrcourseAsExcel, err := controller.GetCurrcourseAsExcel(cookie_list)
	//		Day1 := CurrcourseAsExcel[1][1:]
	//		Day2 := CurrcourseAsExcel[2][1:]
	//		Day3 := CurrcourseAsExcel[3][1:]
	//		Day4 := CurrcourseAsExcel[4][1:]
	//		Day5 := CurrcourseAsExcel[5][1:]
	//		Day6 := CurrcourseAsExcel[6][1:]
	//		Day7 := CurrcourseAsExcel[7][1:]
	//		Time := [][]string{
	//			{"08:40-09:20"},
	//			{"09:25-10:05"},
	//			{"10:25-11:05"},
	//			{"11:10-11:50"},
	//			{"14:30-15:10"},
	//			{"15:15-15:55"},
	//			{"16:05-16:45"},
	//			{"16:50-17:30"},
	//			{"19:30-20:10"},
	//			{"20:15-20:55"},
	//			{"21:00-21:40"},
	//		}
	//		schedule := CourseSchedule{
	//			Days: [][]string{
	//				Day1, Day2, Day3, Day4, Day5, Day6, Day7,
	//			},
	//			Timeslots: Time,
	//		}
	//		fmt.Println(schedule.Days[0])
	//		if err != nil {
	//			return
	//		}
	//		c.HTML(http.StatusOK, "course.html", gin.H{
	//			"schedule": schedule,
	//			"message":  "课表查询",
	//		})
	//	} else {
	//		c.Redirect(http.StatusFound, "/")
	//	}
	//})
	router.GET("/GPA", func(c *gin.Context) {
		session := sessions.Default(c)
		loggedIn := session.Get("loggedIn")
		jsessionid := session.Get("jsessionid")
		if jsessionid == nil {
			c.Redirect(http.StatusFound, "/")
		}
		if loggedIn == true {
			c.HTML(http.StatusOK, "gpa.html", gin.H{
				"message": "请选择年份和学期",
			})
		} else {
			c.Redirect(http.StatusFound, "/")
		}
	})
	router.POST("/GPA", func(c *gin.Context) {
		session := sessions.Default(c)
		loggedIn := session.Get("loggedIn")
		jsessionid := session.Get("jsessionid")
		if jsessionid == nil {
			c.Redirect(http.StatusFound, "/")

		}
		if loggedIn == true {
			stringValue := jsessionid.(string)
			cookie_list := &http.Cookie{
				Name:  "JSESSIONID",
				Value: stringValue,
			}
			sessionS := sessions.Default(c)
			year := c.PostForm("year")
			term := c.PostForm("term")
			GPAS, _, ScoreExcel, Fail := controller.GetAllStudentOwnScores(cookie_list, year, term)
			sessionS.Save()
			c.HTML(http.StatusOK, "gpa.html", gin.H{
				"message":    "GPA计算",
				"GPAS":       GPAS,
				"ScoreExcel": ScoreExcel,
				"Fail":       Fail,
			})
		} else {
			c.Redirect(http.StatusFound, "/")
		}
	})

	router.GET("/exam", func(c *gin.Context) {
		session := sessions.Default(c)
		loggedIn := session.Get("loggedIn")
		jsessionid := session.Get("jsessionid")
		if jsessionid == nil {
			c.Redirect(http.StatusFound, "/")

		}
		if loggedIn == true {
			stringValue := jsessionid.(string)
			cookie_list := &http.Cookie{
				Name:  "JSESSIONID",
				Value: stringValue,
			}
			sessionS := sessions.Default(c)
			AllEaxm, _ := controller.GetStudentExam(cookie_list)
			sessionS.Save()
			c.HTML(http.StatusOK, "exam.html", gin.H{
				"message": "考试安排",
				"AllEaxm": AllEaxm,
			})
		} else {
			c.Redirect(http.StatusFound, "/")
		}
	})
	router.GET("/lexam", func(c *gin.Context) {
		session := sessions.Default(c)
		loggedIn := session.Get("loggedIn")
		jsessionid := session.Get("jsessionid")
		if jsessionid == nil {
			c.Redirect(http.StatusFound, "/")

		}
		if loggedIn == true {
			stringValue := jsessionid.(string)
			cookie_list := &http.Cookie{
				Name:  "JSESSIONID",
				Value: stringValue,
			}
			sessionS := sessions.Default(c)
			AllLExam, _ := controller.GetStudentLExam(cookie_list)
			sessionS.Save()
			c.HTML(http.StatusOK, "lexam.html", gin.H{
				"message":  "成绩考试查询",
				"AllLExam": AllLExam,
			})
		} else {
			c.Redirect(http.StatusFound, "/")
		}
	})
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/")

	})
	router.GET("/getCaptcha", func(c *gin.Context) {
		jsessionid, captchaImage := controller.Captcha()
		c.JSON(200, gin.H{
			"jsessionid":   jsessionid,
			"captchaImage": captchaImage,
		})
	})
	router.Run(":" + config.Port)
}
