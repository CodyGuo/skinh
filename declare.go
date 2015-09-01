package skinh

//SkinH 错误信息
const (
	SRET_OK             = 0 //成功
	SRET_ERROR          = 1 //失败
	SRET_ERROR_READ     = 2 //皮肤文件读取错误
	SRET_ERROR_PARAM    = 3 //参数错误
	SRET_ERROR_CREATE   = 4 //创建皮肤失败
	SRET_ERROR_FORMAT   = 5 //皮肤格式不正确
	SRET_ERROR_VERSION  = 6 //皮肤版本不兼容
	SRET_ERROR_PASSWORD = 7 //皮肤密钥错误
	SRET_ERROR_INVALID  = 8 //换肤引擎无效
)

const (
	TYPE_UNKNOWN     = 0    //未知类型
	TYPE_ANIMATE     = 1001 //动画控件
	TYPE_CHECKBOX    = 1002 //复选框
	TYPE_COMBOBOX    = 1003 //组合框
	TYPE_COMBOLBOX   = 1004 //组合下拉框
	TYPE_CONTROLBAR  = 1005 //控件栏
	TYPE_DATETIME    = 1006 //日期控件
	TYPE_EDITBOX     = 1007 //文本框
	TYPE_GROUPBOX    = 1008 //分组框
	TYPE_HEADERCTRL  = 1009 //列头控件
	TYPE_HOTKEY      = 1010 //热键控件
	TYPE_IPADDRESS   = 1011 //IP地址控件
	TYPE_LABEL       = 1012 //标签控件
	TYPE_LISTBOX     = 1013 //列表框
	TYPE_LISTVIEW    = 1014 //列表视图
	TYPE_MDICLIENT   = 1015 //MDI客户区
	TYPE_MENU        = 1016 //菜单
	TYPE_MONTHCAL    = 1017 //月历控件
	TYPE_PICTURE     = 1018 //图片框
	TYPE_PROGRESS    = 1019 //进度条
	TYPE_PUSHBUTTON  = 1020 //普通按钮
	TYPE_RADIOBUTTON = 1021 //单选框
	TYPE_REBAR       = 1022 //重组栏
	TYPE_RICHEDIT    = 1023 //富文本框
	TYPE_SCROLLBAR   = 1024 //滚动条
	TYPE_SCROLLCTRL  = 1025 //内置滚动条的控件
	TYPE_SPINCTRL    = 1026 //调节器
	TYPE_STATUSBAR   = 1027 //状态栏
	TYPE_TABCTRL     = 1028 //选择夹
	TYPE_TOOLBAR     = 1029 //工具栏
	TYPE_TOOLBARWND  = 1030 //MFC工具栏窗体
	TYPE_TRACKBAR    = 1031 //滑条控件
	TYPE_TREEVIEW    = 1032 //树形视图
	TYPE_WINDOW      = 1034 //标准窗体
	TYPE_COMCTRL     = 1036 //通用换肤
	TYPE_PAINTCTRL   = 1037 //通用换肤
)
