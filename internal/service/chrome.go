package service

import (
	"campushelphub/internal/config"
	"campushelphub/internal/errors"
	"campushelphub/model"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type ChromeService struct {
	config  *config.Config
	errs    *errors.Error
	service *selenium.Service
	caps    selenium.Capabilities
}

// NewChromeService 创建Chrome服务，部分配置项从配置文件中获取
func NewChromeService(cfg *config.Config, errs *errors.Error) *ChromeService {
	logFile, err := os.OpenFile(cfg.ChromeVerify.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件失败: %v", err))
	}
	defer logFile.Close()
	driverOpts := []selenium.ServiceOption{
		selenium.Output(logFile), // 驱动进程的stdout/stderr写入日志
	}
	service, err := selenium.NewChromeDriverService(cfg.ChromeVerify.ChromeDriverPath, cfg.ChromeVerify.Port, driverOpts...)
	if err != nil {
		panic(fmt.Sprintf("启动ChromeDriver服务失败: %v", err))
	}
	// 注意：这里不使用defer service.Stop()，避免程序退出时停止服务

	// 配置Chrome浏览器（保持运行）
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Path: cfg.ChromeVerify.ChromeBinaryPath,
		Args: []string{
			"--no-sandbox",            // 解决权限问题
			"--disable-dev-shm-usage", // 禁用临时目录限制
			"--disable-blink-features=AutomationControlled", // 反检测
			"--disable-popup-blocking",                      // 禁用弹窗拦截
			"--persistent-logging=1",                        // 开启Chrome持久日志
			"--log-level=0",                                 // Chrome日志级别（0=详细，3=只报错）
			"--headless",                                    // 无头模式
		},
		// 禁用Chrome的自动退出机制
		ExcludeSwitches: []string{"enable-automation", "disable-background-networking"},
	}
	caps.AddChrome(chromeCaps)
	return &ChromeService{
		config:  cfg,
		errs:    errs,
		service: service,
		caps:    caps,
	}
}

func (s *ChromeService) VerifyStudent(verify *model.ChromeStudentVerify) *errors.Error {
	studentID := verify.StudentID
	password := verify.Password
	// 连接Selenium并打开浏览器会话
	wd, err := selenium.NewRemote(s.caps, s.config.ChromeVerify.DriverURL)
	if err != nil {
		return s.errs.NewError(errors.ErrChromeOpen, http.StatusInternalServerError, err)
	}
	defer wd.Quit() // 程序结束时关闭关于会话

	// 打开认证页面
	if err = wd.Get(s.config.ChromeVerify.URL); err != nil {
		return s.errs.NewError(errors.ErrChromeOpen, http.StatusInternalServerError, err)
	}

	// 显式等待和输入学号
	wd.SetImplicitWaitTimeout(10 * time.Second) // 全局隐式等待（兜底）
	usernameInput, err := wd.FindElement(selenium.ByID, "j_username")
	if err != nil {
		return s.errs.NewError(errors.ErrChromeOpenURL, http.StatusInternalServerError, err)
	} else {
		if err := usernameInput.SendKeys(studentID); err != nil {
			return s.errs.NewError(errors.ErrChromeInteraction, http.StatusInternalServerError, err)
		}
	}

	// 输入密码
	passwordInput, err := wd.FindElement(selenium.ByID, "j_password")
	if err != nil {
		return s.errs.NewError(errors.ErrChromeOpenURL, http.StatusInternalServerError, err)
	} else {
		if err := passwordInput.SendKeys(password); err != nil {
			return s.errs.NewError(errors.ErrChromeInteraction, http.StatusInternalServerError, err)
		}
	}

	// 点击登录按钮
	loginButton, err := wd.FindElement(selenium.ByID, "loginButton")
	if err != nil {
		return s.errs.NewError(errors.ErrChromeOpenURL, http.StatusInternalServerError, err)
	} else {
		if err := loginButton.Click(); err != nil {
			return s.errs.NewError(errors.ErrChromeInteraction, http.StatusInternalServerError, err)
		}
	}

	// 等待跳转并验证结果
	time.Sleep(3 * time.Second)
	currentURL, err := wd.CurrentURL()
	if err != nil {
		return s.errs.NewError(errors.ErrChromeInteraction, http.StatusInternalServerError, err)
	} else {
		if strings.Contains(currentURL, "/"+s.config.ChromeVerify.VerifySign) {
			return nil
		} else {
			return s.errs.NewError(errors.ErrChromeVerifyFailed, http.StatusInternalServerError, nil)
		}
	}
}
