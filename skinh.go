//Copyright (C) Skyfore 2015

//本库依赖于 Skyfore 的 Sys 库 , v2015.8.31+
//请将 SkinH.dll 置于 生成EXE 所在目录

package skinh

import (
	. "github.com/codyguo/sys"
)

var SkinHDll *DLLClass

func init() {
	//您可以在此自行添加搜索目录
	if FileExist("SkinH.dll") {
		SkinHDll = Dll("SkinH.dll")
	} else if FileExist("lib/SkinH.dll") {
		SkinHDll = Dll("lib/SkinH.dll")
	} else {
		panic("载入 SkinH.dll 失败")
	}
}

func Attach(argvs ...string) bool {
	var argvLength int = len(argvs)
	var ret uintptr

	if argvLength == 0 {
		//载入 EXE 目录下的 SkinH.she
		ret = SkinHDll.Call("SkinH_Attach")
	} else if argvLength == 1 {
		ret = SkinHDll.Call("SkinH_AttachEx", TEXT(FullPath(argvs[0])), TEXT(""))
	} else if argvLength == 2 {
		ret = SkinHDll.Call("SkinH_AttachEx", TEXT(FullPath(argvs[0])), TEXT(argvs[1]))
	} else {
		panic("Attach 函数的参数(均可省略)为 @1:皮肤路径(支持相对) @2:皮肤密码")
	}
	return (ret == 0)
}

func AttachEx(skinPath, skinPw string, nHue, nSat, nBri int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_AttachExt", TEXT(FullPath(skinPath)), TEXT(skinPw), uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return (ret == 0)
}

func AttachRes(pShe, dwSize int, nHue, nSat, nBri int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_AttachRes", uintptr(pShe), uintptr(dwSize), uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return (ret == 0)
}

func AttachResEx(lpName, lpType, skinPw string, nHue, nSat, nBri int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_AttachResEx", TEXT(lpName), TEXT(lpType), TEXT(skinPw), uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return (ret == 0)
}

func Detach() bool {
	var ret uintptr = SkinHDll.Call("SkinH_Detach")
	return (ret == 0)
}

func DetachEx(hWnd int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_DetachEx", uintptr(hWnd))
	return (ret == 0)
}

func SetWinAlpha(hWnd, nAlpha int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_SetWindowAlpha", uintptr(hWnd), uintptr(nAlpha))
	return (ret == 0)
}

func AdjustHSV(nHue, nSat, nBri int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_AdjustHSV", uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return (ret == 0)
}

func GetColor(hWnd, x, y int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_GetColor", uintptr(hWnd), uintptr(x), uintptr(y))
	return (ret == 0)
}

func Map(hWnd, nType int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_Map", uintptr(hWnd), uintptr(nType))
	return (ret == 0)
}

func SetAero(open bool) bool {
	var ret uintptr
	if open {
		ret = SkinHDll.Call("SkinH_SetAero", uintptr(1))
	} else {
		ret = SkinHDll.Call("SkinH_SetAero", uintptr(0))
	}
	return (ret == 0)
}

func AdjustAero(nAlpha, nShwDark, nShwSharp, nShwSize, nX, nY, nReg, nGreen, nBlue uintptr) bool {
	var ret uintptr = SkinHDll.Call("SkinH_GetColor", nAlpha, nShwDark, nShwSharp, nShwSize, nX, nY, nReg, nGreen, nBlue)
	return (ret == 0)
}

func SetWinMovable(hWnd int, bMovable bool) bool {
	var ret uintptr
	if bMovable {
		ret = SkinHDll.Call("SkinH_Map", uintptr(hWnd), 1)
	} else {
		ret = SkinHDll.Call("SkinH_Map", uintptr(hWnd), 0)
	}
	return (ret == 0)
}

func SetBkColor(hwnd, nRed, nGreen, nBlue uintptr) bool {
	var ret uintptr = SkinHDll.Call("SkinH_SetBackColor", hwnd, nRed, nGreen, nBlue)
	return (ret == 0)
}

func SetForColor(hwnd, nRed, nGreen, nBlue uintptr) bool {
	var ret uintptr = SkinHDll.Call("SkinH_SetForeColor", hwnd, nRed, nGreen, nBlue)
	return (ret == 0)
}

func LockUpdate(hWnd int, bUpdate bool) bool {
	var ret uintptr
	if bUpdate {
		ret = SkinHDll.Call("SkinH_LockUpdate", uintptr(hWnd), 1)
	} else {
		ret = SkinHDll.Call("SkinH_LockUpdate", uintptr(hWnd), 0)
	}
	return (ret == 0)
}

func SetMenuAlpha(nAlpha int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_SetMenuAlpha", uintptr(nAlpha))
	return (ret == 0)
}

func NineBlt(hDtDC, left, top, right, bottom, nId uintptr) bool {
	var ret uintptr = SkinHDll.Call("SkinH_NineBlt", hDtDC, left, top, right, bottom, nId)
	return (ret == 0)
}

func SetTitleMenuBar(hWnd uintptr, bEnable bool, nTMenuY, nTopOffs, nRightOffs uintptr) bool {
	var ret uintptr
	if bEnable {
		ret = SkinHDll.Call("SkinH_SetTitleMenuBar", hWnd, 1, nTMenuY, nTopOffs, nRightOffs)
	} else {
		ret = SkinHDll.Call("SkinH_SetTitleMenuBar", hWnd, 0, nTMenuY, nTopOffs, nRightOffs)
	}
	return (ret == 0)
}

func SetFont(hWnd, hFont int) bool {
	var ret uintptr = SkinHDll.Call("SkinH_SetFont", uintptr(hWnd), uintptr(hFont))
	return (ret == 0)
}

func SetFontEx(hWnd int, szFace string, nHeight, nWidth, nWeight, nItalic, nUnderline, nStrikeOut uintptr) bool {
	var ret uintptr = SkinHDll.Call("SkinH_SetFontEx", uintptr(hWnd), TEXT(szFace), nHeight, nWidth, nWeight, nItalic, nUnderline, nStrikeOut)
	return (ret == 0)
}

func VerifySign() bool {
	var ret uintptr = SkinHDll.Call("SkinH_VerifySign")
	return (ret == 0)
}
