package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

var SHELL = "/bin/bash"

// start aria2c and py_http.server with shell
func start() {
  aria2_start := "aria2c --enable-rpc --rpc-listen-all=true --rpc-allow-origin-all -c -D"
  py_httpd_start := "nohup python -m http.server --directory AriaNg-source >/dev/null 2>&1 &"

  aria2c_cmd := exec.Command( SHELL, "-c", aria2_start )
  py_httpd_cmd := exec.Command( SHELL, "-c", py_httpd_start )

  if err := aria2c_cmd.Run(); err != nil {
    log.Fatalln("start err ==>",err)
  }
  if err := py_httpd_cmd.Run(); err != nil{
    log.Fatalln("start err ==>",err)
  }
  fmt.Println("start ok, go to http://127.0.0.1:8000")
}

// get py_http.server's pid with lsof and kill it ,then pkill aria2c
func stop()  {
  py_httpd_pid_get_cmd := exec.Command( SHELL, "-c", "lsof -i:8000 | tail -1 | awk -F\" \" '{print $2}'" )
  out, err := py_httpd_pid_get_cmd.CombinedOutput()
  if err != nil {
    log.Fatalln("get python http.server err ==>",err)
  }
  pid_str := string(out)
  pid_str = strings.Replace(pid_str, "\n", "", -1)
  if err := exec.Command( SHELL, "-c", "kill"+" "+pid_str ).Run(); err != nil{
    log.Fatalln("kill python http.server err==>",err)
  }
  if err := exec.Command( SHELL, "-c", "pkill aria2c").Run(); err != nil{
    log.Fatalln("kill aria2c err ==>",err)
  }

  fmt.Println("stop ok")
}

func main()  {
  args()
}

// Handle command-line arguments
func args()  {
  arg1 := os.Args[1]
  switch arg1{
  case "", "-h", "-help", "--help", "help":
    help()
  case "start":
    start()
  case "stop":
    stop()
  case "restart":
    stop()
    time.Sleep(time.Second*3)
    start()
  case "install":
    install()
  case "reinstall":
    reinstall()
  default:
    fmt.Println("error option:", arg1)
    help()
  }
}

// print help 
func help()  {
  fmt.Println("+--------------------------------------------------------+")
  fmt.Println("+ command  ->  AriaNg-service <option>                   +")
  fmt.Println("+--------------------------------------------------------+")
  fmt.Println("+ help       : show help                                 +")
  fmt.Println("+ start      : start AriaNg                              +")
  fmt.Println("+ stop       : stop AriaNg                               +")
  fmt.Println("+ restart    : restart AriaNg                            +")
  fmt.Println("+ install    : install AriaNg least release              +")
  fmt.Println("+ reinstall  : delate old AriaNg release and install new +")
  fmt.Println("+--------------------------------------------------------+")
}

// get mayswind/AriaNg latest release url and release name
func get_release_url() (release_url, release_name string) {
  repository_name := "mayswind/AriaNg"
  tag_name_get := "wget -qO- -t1 -T2 \"https://api.github.com/repos/"+repository_name+"/releases/latest\" | jq -r '.tag_name'"
  release_name_get := "wget -qO- -t1 -T2 \"https://api.github.com/repos/"+repository_name+"/releases/latest\" | jq -r '.assets[].name' | grep \"AllInOne.zip\""

  tag_name_cmd := exec.Command(SHELL, "-c", tag_name_get)
  release_name_cmd := exec.Command(SHELL, "-c", release_name_get)

  //run cmd and get stdout
  tag_name_cmd_out, err := tag_name_cmd.CombinedOutput()
  if err != nil {
    log.Fatalln("get tag_name err ==>",err)
  }
  if tag_name_cmd_out == nil{
    log.Fatalln("get tag_name err ==>","Didn't get tag_name")
  }

  //run cmd and get stdout
  release_name_cmd_out, err := release_name_cmd.CombinedOutput()
  if err != nil {
    log.Fatalln("get release_name err ==>",err)
  }
  if release_name_cmd_out== nil{
    log.Fatalln("get release_name err ==>","Didn't get release_name")
  }

  urlstr := "https://github.com/"+repository_name+"/releases/download/"+string(tag_name_cmd_out)+"/"+string(release_name_cmd_out)
  release_name_str := string(release_name_cmd_out)

  // remove the '\n' in the strings
  urlstr = strings.Replace(urlstr, "\n", "", -1)
  release_name_str = strings.Replace(release_name_str, "\n", "", -1)
  return urlstr, release_name_str
}

// install the latest mayswind/AriaNg
func install()  {
  release_url, release_name := get_release_url()
  fmt.Println("install start ...")

  // http request
  request, err := http.NewRequest("GET", release_url, nil)
  if err != nil {
    log.Fatalln("request err ==>", err)
  }

  // http client response
  response, err := http.DefaultClient.Do(request)
  if err != nil {
    log.Fatalln("response err ==>", err)
  }
  if err := os.Mkdir("AriaNg-source", 0755); err != nil{
    log.Fatalln("mkdir err ==>", err)
  }
  file, err := os.OpenFile("AriaNg-source/"+release_name, os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatalln("download err ==>",err)
  }
  defer file.Close()

  //create a progressbar for downloading
  bar := progressbar.NewOptions(
    int(response.ContentLength),
    progressbar.OptionSetDescription("Downloading"),
    progressbar.OptionShowBytes(true),
  )
  io.Copy(io.MultiWriter(file, bar), response.Body)

  //unzip the zip file to ./AriaNg-source
  if err := exec.Command( SHELL, "-c", "unzip -d ./AriaNg-source "+"AriaNg-source/"+release_name).Run(); err != nil {
    log.Fatalln("\nunzip err ==>", err)
  }
  fmt.Println("\ninstall done !!!")
}

//reinstall the mayswind/AriaNg: remove the old release and install the new
func reinstall()  {
  source_dir := "AriaNg-source"
  if !fileIsExisted(source_dir) {
    log.Fatalln("AriaNg-source dir is not exist")
  }
  os.RemoveAll(source_dir)
  install()
}

//File exists returns true, otherwise false
func fileIsExisted(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}