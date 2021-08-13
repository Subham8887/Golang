package main

import ("fmt")

func main(){
	jb:=[]string{"shubham","kumar","prateek","singh"}
	fmt.Println(jb)

	mp:=[]string{"vivek","maurya","rinku","yadav"}
	fmt.Println(mp)

	mat:=[][]string{jb,mp}

	fmt.Print(mat)
}