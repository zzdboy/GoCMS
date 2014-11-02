// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "admin/app"
	controllers "admin/app/controllers"
	controllers0 "admin/app/controllers/Content"
	controllers1 "admin/app/controllers/Extend"
	controllers2 "admin/app/controllers/Module"
	controllers3 "admin/app/controllers/Panel"
	controllers4 "admin/app/controllers/Plugin"
	controllers5 "admin/app/controllers/Public"
	controllers6 "admin/app/controllers/Setting"
	controllers7 "admin/app/controllers/Style"
	controllers8 "admin/app/controllers/User"
	models "admin/app/models"
	tests "admin/tests"
	controllers12 "github.com/revel/revel/modules/jobs/app/controllers"
	_ "github.com/revel/revel/modules/jobs/app/jobs"
	controllers11 "github.com/revel/revel/modules/pprof/app/controllers"
	controllers10 "github.com/revel/revel/modules/static/app/controllers"
	_ "github.com/revel/revel/modules/testrunner/app"
	controllers9 "github.com/revel/revel/modules/testrunner/app/controllers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					63: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Main",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					133: []string{ 
						"title",
						"admin_info",
						"system_info",
						"panel_list",
					},
					135: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Category)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "category", Type: reflect.TypeOf((**models.Category)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
						"title",
						"categorys",
					},
					42: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "category", Type: reflect.TypeOf((**models.Category)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					64: []string{ 
						"title",
						"categorys",
						"Id",
					},
					68: []string{ 
						"title",
						"categorys",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "category", Type: reflect.TypeOf((**models.Category)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					230: []string{ 
						"title",
						"categorys",
						"category_info",
					},
					236: []string{ 
						"title",
						"categorys",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "category", Type: reflect.TypeOf((**models.Category)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Content)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					29: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Left",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
						"title",
						"categorys",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"title",
						"cid",
						"categorys",
						"article_list",
						"where",
						"pages",
					},
					73: []string{ 
						"title",
						"cid",
						"categorys",
						"article_list",
						"where",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Keywords",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CateNameSearch",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "category", Type: reflect.TypeOf((**models.Category)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Push",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					132: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Remove",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					142: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Comment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					185: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Relationlist",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					214: []string{ 
						"title",
						"cid",
						"category_info",
						"article_list",
						"pages",
						"where",
					},
				},
			},
			&revel.MethodType{
				Name: "AddContent",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					239: []string{ 
						"title",
						"cid",
						"category_info",
						"article_info",
					},
					242: []string{ 
						"title",
						"cid",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					359: []string{ 
						"title",
						"cid",
						"category_info",
						"copyfrom_list",
					},
					362: []string{ 
						"title",
						"cid",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "article", Type: reflect.TypeOf((**models.Article)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					603: []string{ 
						"title",
						"cid",
						"category_info",
						"article_info",
						"copyfrom_list",
					},
					606: []string{ 
						"title",
						"cid",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Focus)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focus", Type: reflect.TypeOf((**models.Focus)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					50: []string{ 
						"title",
						"focus_list",
						"Cate_list",
						"where",
						"pages",
					},
					54: []string{ 
						"title",
						"focus_list",
						"Cate_list",
						"where",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focus", Type: reflect.TypeOf((**models.Focus)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					70: []string{ 
						"title",
						"Cate_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focus", Type: reflect.TypeOf((**models.Focus)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					284: []string{ 
						"title",
						"Cate_list",
						"focus_info",
					},
					286: []string{ 
						"title",
						"Cate_list",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.FocusCate)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focusCate", Type: reflect.TypeOf((**models.FocusCate)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					36: []string{ 
						"title",
						"focusCate_list",
						"pages",
					},
					40: []string{ 
						"title",
						"focusCate_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focusCate", Type: reflect.TypeOf((**models.FocusCate)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					53: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focusCate", Type: reflect.TypeOf((**models.FocusCate)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					122: []string{ 
						"title",
						"focusCate_info",
					},
					124: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "focusCate", Type: reflect.TypeOf((**models.FocusCate)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Copyfrom)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "copyfrom", Type: reflect.TypeOf((**models.Copyfrom)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					60: []string{ 
						"title",
						"copyfrom_list",
						"sitedomain",
						"pages",
					},
					64: []string{ 
						"title",
						"copyfrom_list",
						"sitedomain",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "copyfrom", Type: reflect.TypeOf((**models.Copyfrom)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "copyfrom", Type: reflect.TypeOf((**models.Copyfrom)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					223: []string{ 
						"title",
						"copyfrom_info",
						"Id",
					},
					225: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "copyfrom", Type: reflect.TypeOf((**models.Copyfrom)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Extend)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Announce)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "announce", Type: reflect.TypeOf((**models.Announce)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					36: []string{ 
						"title",
						"announce_list",
						"pages",
					},
					40: []string{ 
						"title",
						"announce_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "announce", Type: reflect.TypeOf((**models.Announce)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					52: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "announce", Type: reflect.TypeOf((**models.Announce)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					152: []string{ 
						"title",
						"announce_info",
					},
					154: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "announce", Type: reflect.TypeOf((**models.Announce)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Complaints)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "complaints", Type: reflect.TypeOf((**models.Complaints)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					35: []string{ 
						"title",
						"complaints_list",
						"pages",
					},
					38: []string{ 
						"title",
						"complaints_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "complaints", Type: reflect.TypeOf((**models.Complaints)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Module)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Panel)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers4.Plugin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					23: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Ajax)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetCaptcha",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Pos",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin_panel", Type: reflect.TypeOf((**models.Admin_Panel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DelPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin_panel", Type: reflect.TypeOf((**models.Admin_Panel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin_panel", Type: reflect.TypeOf((**models.Admin_Panel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetMessage",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ScreenLock",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ScreenUnlock",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Captcha)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetCaptchaId",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Kindeditor)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Manager",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "upload", Type: reflect.TypeOf((**models.Upload)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "TitleImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "upload", Type: reflect.TypeOf((**models.Upload)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Upload",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "upload", Type: reflect.TypeOf((**models.Upload)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AnnounceImage",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "upload", Type: reflect.TypeOf((**models.Upload)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Public)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Map",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
						"title",
						"map_html",
					},
					40: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "CreateHtml",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					54: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Message",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					60: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Test)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					39: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					43: []string{ 
						"title",
						"admin_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					58: []string{ 
						"title",
						"role_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					223: []string{ 
						"title",
						"admin_info",
						"role_list",
					},
					225: []string{ 
						"title",
						"role_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Logs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "logs", Type: reflect.TypeOf((**models.Logs)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
						"title",
						"logs_list",
						"where",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "DelAll",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "logs", Type: reflect.TypeOf((**models.Logs)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Menu)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					37: []string{ 
						"title",
						"menus",
					},
					39: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"title",
						"menus",
						"Id",
					},
					71: []string{ 
						"title",
						"Id",
					},
					86: []string{ 
						"title",
						"menus",
					},
					88: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					223: []string{ 
						"title",
						"menus",
						"menu_info",
					},
					225: []string{ 
						"title",
						"menu_info",
					},
					242: []string{ 
						"title",
						"menus",
					},
					244: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Role)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					36: []string{ 
						"title",
						"role_list",
						"pages",
					},
					40: []string{ 
						"title",
						"role_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Member",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					71: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					75: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					79: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					104: []string{ 
						"title",
						"tree",
					},
					106: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					210: []string{ 
						"title",
						"role_info",
						"tree",
						"Id",
					},
					212: []string{ 
						"title",
						"role_info",
						"Id",
					},
					230: []string{ 
						"title",
						"tree",
					},
					232: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "SetStatus",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Setting)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					24: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Task)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					27: []string{ 
						"title",
						"entries",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers7.Style)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "template", Type: reflect.TypeOf((**models.Template)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					27: []string{ 
						"title",
						"template_list",
					},
				},
			},
			&revel.MethodType{
				Name: "File",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "template", Type: reflect.TypeOf((**models.Template)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					44: []string{ 
						"title",
						"template_info",
					},
					46: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Import",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "template", Type: reflect.TypeOf((**models.Template)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					59: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Setstatus",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "template", Type: reflect.TypeOf((**models.Template)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "template", Type: reflect.TypeOf((**models.Template)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					191: []string{ 
						"title",
						"template_info",
					},
					193: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers8.Group)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user_group", Type: reflect.TypeOf((**models.User_Group)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"title",
						"group_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user_group", Type: reflect.TypeOf((**models.User_Group)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					39: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user_group", Type: reflect.TypeOf((**models.User_Group)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					214: []string{ 
						"title",
						"user_group_info",
					},
					216: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers8.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"title",
						"user_list",
						"group_list",
						"pages",
						"where",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					61: []string{ 
						"title",
						"group_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					254: []string{ 
						"title",
						"id",
						"group_list",
						"user_info",
					},
					256: []string{ 
						"title",
						"id",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Lock",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Unlock",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Move",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UserInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					517: []string{ 
						"title",
						"id",
						"user_info",
					},
					519: []string{ 
						"title",
						"id",
					},
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					532: []string{ 
						"title",
						"CaptchaId",
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "EditInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					690: []string{ 
						"title",
						"admin_info",
					},
					692: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "AdminPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					780: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "EditPwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					803: []string{ 
						"title",
						"admin_info",
					},
					805: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Left",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					906: []string{ 
						"title",
						"left_menu",
					},
					908: []string{ 
						"title",
					},
					926: []string{ 
						"title",
						"left_menu",
					},
					928: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers9.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					78: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers10.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers11.Pprof)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Profile",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Symbol",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Cmdline",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers12.Jobs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Status",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					19: []string{ 
						"entries",
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"admin/app/models.(*Admin).Validate": { 
			42: "a.Username",
			43: "a.Username",
			57: "a.Email",
			58: "a.Email",
			72: "a.Password",
			73: "a.Password",
		},
		"admin/app/models.(*Menu).Validate": { 
			32: "menu.Name",
			33: "menu.Name",
			34: "menu.Pid",
			35: "menu.Url",
			36: "menu.Order",
		},
		"admin/app/models.(*Password).ValidatePassword": { 
			78: "P.Password",
			79: "P.PasswordConfirm",
			81: "P.Password",
			82: "P.Password",
		},
	}
	revel.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
