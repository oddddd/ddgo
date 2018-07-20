package code

/*
返回code
 */
type codeMsg struct {
	code int
	msg  string
	data interface{}
}

func Success1() (int,string){
	//success1 := codeMsg{
	//	code : 1,
	//	msg  : "操作成功",
	//}
	//return success1
	return 1,"123"
}