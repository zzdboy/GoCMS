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
	controllers10 "github.com/revel/revel/modules/jobs/app/controllers"
	_ "github.com/revel/revel/modules/jobs/app/jobs"
	controllers9 "github.com/revel/revel/modules/static/app/controllers"
	_ "github.com/revel/revel/modules/testrunner/app"
	controllers11 "github.com/revel/revel/modules/testrunner/app/controllers"
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
					55: []string{ 
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
					106: []string{ 
						"title",
						"admin_info",
						"system_info",
						"panel_list",
					},
					108: []string{ 
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
					27: []string{ 
						"title",
						"categorys",
					},
					29: []string{ 
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
					51: []string{ 
						"title",
						"categorys",
						"Id",
					},
					55: []string{ 
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
					217: []string{ 
						"title",
						"categorys",
						"category_info",
					},
					223: []string{ 
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
					19: []string{ 
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
					30: []string{ 
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
					59: []string{ 
						"title",
						"cid",
						"categorys",
						"article_list",
						"where",
						"pages",
					},
					63: []string{ 
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
					122: []string{ 
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
					132: []string{ 
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
					175: []string{ 
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
					204: []string{ 
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
					229: []string{ 
						"title",
						"cid",
						"category_info",
						"article_info",
					},
					232: []string{ 
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
					347: []string{ 
						"title",
						"cid",
						"category_info",
						"copyfrom_list",
					},
					350: []string{ 
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
					591: []string{ 
						"title",
						"cid",
						"category_info",
						"article_info",
						"copyfrom_list",
					},
					594: []string{ 
						"title",
						"cid",
					},
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
					50: []string{ 
						"title",
						"copyfrom_list",
						"sitedomain",
						"pages",
					},
					54: []string{ 
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
					66: []string{ 
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
					213: []string{ 
						"title",
						"copyfrom_info",
						"Id",
					},
					215: []string{ 
						"title",
					},
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
					13: []string{ 
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
					25: []string{ 
						"title",
						"announce_list",
						"pages",
					},
					29: []string{ 
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
					41: []string{ 
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
					139: []string{ 
						"title",
						"announce_info",
					},
					141: []string{ 
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
	
	revel.RegisterController((*controllers2.Module)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
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
					13: []string{ 
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
					13: []string{ 
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
					28: []string{ 
						"title",
						"map_html",
					},
					30: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "CreateHtml",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Search",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					44: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Message",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					50: []string{ 
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
					28: []string{ 
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
					28: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					32: []string{ 
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
					47: []string{ 
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
					197: []string{ 
						"title",
						"admin_info",
						"role_list",
					},
					199: []string{ 
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
					28: []string{ 
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
					27: []string{ 
						"title",
						"menus",
					},
					29: []string{ 
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
					59: []string{ 
						"title",
						"menus",
						"Id",
					},
					61: []string{ 
						"title",
						"Id",
					},
					76: []string{ 
						"title",
						"menus",
					},
					78: []string{ 
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
					213: []string{ 
						"title",
						"menus",
						"menu_info",
					},
					215: []string{ 
						"title",
						"menu_info",
					},
					232: []string{ 
						"title",
						"menus",
					},
					234: []string{ 
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
					26: []string{ 
						"title",
						"role_list",
						"pages",
					},
					30: []string{ 
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
					61: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					65: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					69: []string{ 
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
					94: []string{ 
						"title",
						"tree",
					},
					96: []string{ 
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
					200: []string{ 
						"title",
						"role_info",
						"tree",
						"Id",
					},
					202: []string{ 
						"title",
						"role_info",
						"Id",
					},
					220: []string{ 
						"title",
						"tree",
					},
					222: []string{ 
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
					14: []string{ 
						"title",
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
					17: []string{ 
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
					34: []string{ 
						"title",
						"template_info",
					},
					36: []string{ 
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
					49: []string{ 
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
					181: []string{ 
						"title",
						"template_info",
					},
					183: []string{ 
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
					18: []string{ 
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
					29: []string{ 
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
					204: []string{ 
						"title",
						"user_group_info",
					},
					206: []string{ 
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
					35: []string{ 
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
					50: []string{ 
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
					228: []string{ 
						"title",
						"id",
						"group_list",
						"user_info",
					},
					230: []string{ 
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
					491: []string{ 
						"title",
						"id",
						"user_info",
					},
					493: []string{ 
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
					506: []string{ 
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
					639: []string{ 
						"title",
						"admin_info",
					},
					641: []string{ 
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
					729: []string{ 
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
					752: []string{ 
						"title",
						"admin_info",
					},
					754: []string{ 
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
					855: []string{ 
						"title",
						"left_menu",
					},
					857: []string{ 
						"title",
					},
					875: []string{ 
						"title",
						"left_menu",
					},
					877: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers9.Static)(nil),
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
	
	revel.RegisterController((*controllers10.Jobs)(nil),
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
	
	revel.RegisterController((*controllers11.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
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
					69: []string{ 
						"error",
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
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"admin/app/models.(*Admin).Validate": { 
			31: "a.Username",
			32: "a.Username",
			46: "a.Email",
			47: "a.Email",
			61: "a.Password",
			62: "a.Password",
		},
		"admin/app/models.(*Menu).Validate": { 
			22: "menu.Name",
			23: "menu.Name",
			24: "menu.Pid",
			25: "menu.Url",
			26: "menu.Order",
		},
		"admin/app/models.(*Password).ValidatePassword": { 
			67: "P.Password",
			68: "P.PasswordConfirm",
			70: "P.Password",
			71: "P.Password",
		},
	}
	revel.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
