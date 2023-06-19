package homework



func OneOrTwo(a,b int,c string)(int,string){
	if c == "one"{
		return a,c
	}else if c == "two"{
		return b,c
	}else{
		return  0,c
	}
	} 