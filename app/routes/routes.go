// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Main(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Main", args).Url
}


type tCategory struct {}
var Category tCategory


func (_ tCategory) Index(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Index", args).Url
}

func (_ tCategory) Add(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Add", args).Url
}

func (_ tCategory) Edit(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Edit", args).Url
}

func (_ tCategory) Delete(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Delete", args).Url
}


type tContent struct {}
var Content tContent


func (_ tContent) Index(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Index", args).Url
}

func (_ tContent) Left(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Left", args).Url
}

func (_ tContent) List(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.List", args).Url
}

func (_ tContent) Keywords(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Keywords", args).Url
}

func (_ tContent) Delete(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Delete", args).Url
}

func (_ tContent) CateNameSearch(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Content.CateNameSearch", args).Url
}

func (_ tContent) Push(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Push", args).Url
}

func (_ tContent) Remove(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Remove", args).Url
}

func (_ tContent) Comment(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Comment", args).Url
}

func (_ tContent) Relationlist(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Relationlist", args).Url
}

func (_ tContent) AddContent(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.AddContent", args).Url
}

func (_ tContent) Add(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Add", args).Url
}

func (_ tContent) Edit(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Content.Edit", args).Url
}


type tFocus struct {}
var Focus tFocus


func (_ tFocus) Index(
		focus interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focus", focus)
	return revel.MainRouter.Reverse("Focus.Index", args).Url
}

func (_ tFocus) Add(
		focus interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focus", focus)
	return revel.MainRouter.Reverse("Focus.Add", args).Url
}

func (_ tFocus) Edit(
		focus interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focus", focus)
	return revel.MainRouter.Reverse("Focus.Edit", args).Url
}


type tFocusCate struct {}
var FocusCate tFocusCate


func (_ tFocusCate) Index(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Index", args).Url
}

func (_ tFocusCate) Add(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Add", args).Url
}

func (_ tFocusCate) Edit(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Edit", args).Url
}

func (_ tFocusCate) Delete(
		focusCate interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "focusCate", focusCate)
	return revel.MainRouter.Reverse("FocusCate.Delete", args).Url
}


type tCopyfrom struct {}
var Copyfrom tCopyfrom


func (_ tCopyfrom) Index(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Index", args).Url
}

func (_ tCopyfrom) Add(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Add", args).Url
}

func (_ tCopyfrom) Edit(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Edit", args).Url
}

func (_ tCopyfrom) Delete(
		copyfrom interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "copyfrom", copyfrom)
	return revel.MainRouter.Reverse("Copyfrom.Delete", args).Url
}


type tExtend struct {}
var Extend tExtend


func (_ tExtend) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Extend.Index", args).Url
}


type tAnnounce struct {}
var Announce tAnnounce


func (_ tAnnounce) Index(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Index", args).Url
}

func (_ tAnnounce) Add(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Add", args).Url
}

func (_ tAnnounce) Edit(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Edit", args).Url
}

func (_ tAnnounce) Delete(
		announce interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "announce", announce)
	return revel.MainRouter.Reverse("Announce.Delete", args).Url
}


type tComplaints struct {}
var Complaints tComplaints


func (_ tComplaints) Index(
		complaints interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "complaints", complaints)
	return revel.MainRouter.Reverse("Complaints.Index", args).Url
}

func (_ tComplaints) Delete(
		complaints interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "complaints", complaints)
	return revel.MainRouter.Reverse("Complaints.Delete", args).Url
}


type tModule struct {}
var Module tModule


func (_ tModule) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Module.Index", args).Url
}


type tPanel struct {}
var Panel tPanel


func (_ tPanel) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Panel.Index", args).Url
}


type tPlugin struct {}
var Plugin tPlugin


func (_ tPlugin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Plugin.Index", args).Url
}


type tAjax struct {}
var Ajax tAjax


func (_ tAjax) GetCaptcha(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetCaptcha", args).Url
}

func (_ tAjax) Pos(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Ajax.Pos", args).Url
}

func (_ tAjax) GetPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.GetPanel", args).Url
}

func (_ tAjax) DelPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.DelPanel", args).Url
}

func (_ tAjax) AddPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.AddPanel", args).Url
}

func (_ tAjax) GetMessage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetMessage", args).Url
}

func (_ tAjax) ScreenLock(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.ScreenLock", args).Url
}

func (_ tAjax) ScreenUnlock(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Ajax.ScreenUnlock", args).Url
}


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.Index", args).Url
}

func (_ tCaptcha) GetCaptchaId(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.GetCaptchaId", args).Url
}


type tKindeditor struct {}
var Kindeditor tKindeditor


func (_ tKindeditor) Manager(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.Manager", args).Url
}

func (_ tKindeditor) TitleImage(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.TitleImage", args).Url
}

func (_ tKindeditor) Upload(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.Upload", args).Url
}

func (_ tKindeditor) AnnounceImage(
		upload interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "upload", upload)
	return revel.MainRouter.Reverse("Kindeditor.AnnounceImage", args).Url
}


type tPublic struct {}
var Public tPublic


func (_ tPublic) Map(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Public.Map", args).Url
}

func (_ tPublic) CreateHtml(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.CreateHtml", args).Url
}

func (_ tPublic) Search(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Search", args).Url
}

func (_ tPublic) Message(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Message", args).Url
}


type tTest struct {}
var Test tTest


func (_ tTest) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Test.Index", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (_ tAdmin) Add(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Add", args).Url
}

func (_ tAdmin) Edit(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Edit", args).Url
}

func (_ tAdmin) Delete(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Delete", args).Url
}


type tLogs struct {}
var Logs tLogs


func (_ tLogs) Index(
		logs interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "logs", logs)
	return revel.MainRouter.Reverse("Logs.Index", args).Url
}

func (_ tLogs) DelAll(
		logs interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "logs", logs)
	return revel.MainRouter.Reverse("Logs.DelAll", args).Url
}


type tMenu struct {}
var Menu tMenu


func (_ tMenu) Index(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Index", args).Url
}

func (_ tMenu) Add(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Add", args).Url
}

func (_ tMenu) Edit(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Edit", args).Url
}

func (_ tMenu) Delete(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Delete", args).Url
}


type tRole struct {}
var Role tRole


func (_ tRole) Index(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Index", args).Url
}

func (_ tRole) Member(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Member", args).Url
}

func (_ tRole) Add(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Add", args).Url
}

func (_ tRole) Edit(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Edit", args).Url
}

func (_ tRole) SetStatus(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.SetStatus", args).Url
}

func (_ tRole) Delete(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Delete", args).Url
}


type tSetting struct {}
var Setting tSetting


func (_ tSetting) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Setting.Index", args).Url
}


type tTask struct {}
var Task tTask


func (_ tTask) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Task.Index", args).Url
}


type tStyle struct {}
var Style tStyle


func (_ tStyle) Index(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Index", args).Url
}

func (_ tStyle) File(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.File", args).Url
}

func (_ tStyle) Import(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Import", args).Url
}

func (_ tStyle) Setstatus(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Setstatus", args).Url
}

func (_ tStyle) Edit(
		template interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "template", template)
	return revel.MainRouter.Reverse("Style.Edit", args).Url
}


type tGroup struct {}
var Group tGroup


func (_ tGroup) Index(
		user_group interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_group", user_group)
	return revel.MainRouter.Reverse("Group.Index", args).Url
}

func (_ tGroup) Add(
		user_group interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_group", user_group)
	return revel.MainRouter.Reverse("Group.Add", args).Url
}

func (_ tGroup) Edit(
		user_group interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user_group", user_group)
	return revel.MainRouter.Reverse("Group.Edit", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Index(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Index", args).Url
}

func (_ tUser) Add(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Add", args).Url
}

func (_ tUser) Edit(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Edit", args).Url
}

func (_ tUser) Delete(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Delete", args).Url
}

func (_ tUser) Lock(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Lock", args).Url
}

func (_ tUser) Unlock(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Unlock", args).Url
}

func (_ tUser) Move(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Move", args).Url
}

func (_ tUser) UserInfo(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.UserInfo", args).Url
}

func (_ tUser) Login(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Login", args).Url
}

func (_ tUser) Logout(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Logout", args).Url
}

func (_ tUser) EditInfo(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditInfo", args).Url
}

func (_ tUser) AdminPanel(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.AdminPanel", args).Url
}

func (_ tUser) EditPwd(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditPwd", args).Url
}

func (_ tUser) Left(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("User.Left", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tPprof struct {}
var Pprof tPprof


func (_ tPprof) Profile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Profile", args).Url
}

func (_ tPprof) Symbol(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Symbol", args).Url
}

func (_ tPprof) Cmdline(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Cmdline", args).Url
}

func (_ tPprof) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Index", args).Url
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).Url
}


