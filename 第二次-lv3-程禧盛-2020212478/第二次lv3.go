package main

type Vedio struct {
	praise int
	collect int
	givecoin int
	authorname string
	Vedioname string
}

func (a *Vedio) Praise()  {
	a.praise++
}
func (a *Vedio) Collect()  {
	a.collect++
}
func (a *Vedio) Givecoin()  {
	a.givecoin++
}
func (a *Vedio) threeitemsadd()  {
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
	}
}
func main()  {
	a *Vedio{}
	var b Vedio
}