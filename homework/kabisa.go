package homework

func Kabisa(year int)(res string){
if year%4 == 0 && year%100 != 0 || year%400 == 0{
	res = "kabisa"
}else{
	res="ksbisa emas"
}
return res
}