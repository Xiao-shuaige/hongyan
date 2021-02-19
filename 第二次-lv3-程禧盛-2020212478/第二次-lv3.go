package main

type Vedio struct {
	praise int
	collect int
	givecoin int
	authorname string
	Vedioname string//五种属性：点赞，收藏，给币，作者名与视频名
}

func (a *Vedio) Praise()  {//方法：点赞
	a.praise++
}
func (a *Vedio) Collect()  {//方法：收藏
	a.collect++
}
func (a *Vedio) Givecoin()  {//方法:投币
	a.givecoin++
}
func (a *Vedio) threeitemsadd()  {//方法：一键三连
	a.collect++
	a.praise++
	a.givecoin++
}
func retvedio (authorname string,vedioname string) Vedio{
	var c string
	var d string
	return Vedio{
		authorname:c,
		Vedioname: d,
	}//函数：返回视频名字与作者名字
}
func main()  {
	a *Vedio{}
	var b Vedio
}