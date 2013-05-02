package main

import (
    "testing"
    "os"
    "fmt"
    "bufio"
)

func TestXYZ(t *testing.T) {

}

func main(){
    xFile:="f:/build.xml"
    fin,err := os.Open(xFile)
    defer fin.Close()
    
    
    if err!= nil {
       fmt.Println(xFile,err)
       return
    }
    
    mbuff:= bufio.NewReader(fin)
    
    for{
      line ,err:= mbuff.ReadString('\n')
      if line=="" {
        fmt.Printf(err.Error())
        break
      } else {
        fmt.Printf("%v",line)
      }
      
    }
    
    return
    
    //分配缓存片段
    buff :=make([]byte, 1024)
    
    for {
       n , _ :=  fin.Read(buff)
       if 0 == n {break}  //文件末尾就退出
       
      fmt.Printf("read file data is:\n[% X]   n=%d\n", buff[:n], n)
      // os.Stdout.Write(buff[:n]) 
       
       
       
    }
}