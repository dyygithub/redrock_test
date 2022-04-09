package service

import (
	"fmt"
	"os"
)

func WriteResult(vals string,outfile string)error{
	file,err:=os.Create(outfile)
	if err!=nil{
		fmt.Println("writer",err)
		return err
	}
	defer file.Close()
	file.WriteString(vals)
	return nil
}