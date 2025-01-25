package guoYangYun

import "errors"

// 自定义错误
var (
	//	身份证格式错误
	ErrIdCardFormat = errors.New("身份证格式错误")
	//	阿里云的国阳云appcode错误
	ErrAppCode = errors.New("appcode错误")
	//  解析失败
	ErrParseUrl = errors.New("错误的请求路径")
	//	初始化请求失败
	ErrInitRequest = errors.New("初始化请求失败")
	//	请求失败
	ErrClientDo = errors.New("请求失败")
	// 验证失败，套餐失效或者次数不足
	ErrAuth = errors.New("验证失败，套餐失效或者次数不足")
	//	读取数据失败
	ErrReadBody = errors.New("读取数据失败")
	//	反序列化错误
	ErrUnmarshal = errors.New("反序列化错误")
	//	校验不通过
	ErrNotExistIdCardOrName = errors.New("校验不通过，身份证或者姓名不存在")
	// 校验失败
	ErrFail = errors.New("校验不通过")
	//	姓名格式错误
	ErrName = errors.New("姓名格式错误")
)
