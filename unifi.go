package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"github.com/BurntSushi/toml"
	"math"
	gc "github.com/rthornton128/goncurses"
	"net/http/cookiejar"
	"strconv"
)

type Configuration struct {
	Username string
	Password string
	Site     string
	Url      string
}

func main() {

	config := openConfig()

	cookieJar, _ := cookiejar.New(nil);
	client := &http.Client{
		Jar: cookieJar,
	}

	stdscr, _ := gc.Init()
	defer gc.End()

	if !gc.HasColors() {
		panic("Example requires a colour capable terminal")
	}

	// Must be called after Init but before using any colour related functions
	if err := gc.StartColor(); err != nil {
		panic(err)
	}

	maxY, maxX := stdscr.MaxYX()

	maxHeight := int(float64(maxY) * 1)
	maxWidth := int(float64(maxX) * 0.2)

	fmt.Println("maxx", maxX, "maxy", maxY)

	//upload.Color(gc.C_BLUE)
	if err := gc.InitPair(1, gc.C_CYAN, gc.C_CYAN); err != nil {
		panic(err)
	}

	if err := gc.InitPair(2, gc.C_BLUE, gc.C_BLUE); err != nil {
		panic(err)
	}

	if err := gc.InitPair(4, gc.C_BLUE, gc.C_BLACK); err != nil {
		panic(err)
	}

	if err := gc.InitPair(3, gc.C_CYAN, gc.C_BLACK); err != nil {
		panic(err)
	}
	if err := gc.InitPair(5, gc.C_BLACK, gc.C_BLACK); err != nil {
		panic(err)
	}

	upload, err := gc.NewWindow(maxHeight, maxWidth, 0, 10)
	if (err != nil ) {
		panic(err)
	}

	download, err := gc.NewWindow(maxHeight, maxWidth, 0, maxX-maxWidth-10)
	if (err != nil ) {
		panic(err)
	}

	stdscr.MovePrint(maxY, maxX/2, "Latency")

	stdscr.Overlay(upload)
	stdscr.Overlay(download)

	stdscr.Refresh()

	go GetData(config, client, stdscr, upload, download)

	stdscr.GetChar()

}

func GetData(config Configuration, client *http.Client, screen *gc.Window, uploadBar *gc.Window, downloadBar *gc.Window) {
	var maxUp float64 = 0
	var maxDown float64 = 0

	login(config.Url, config.Username, config.Password, client)

	for {
		latency, upload, download := getInfo(config.Url, config.Site, client)
		//keeping the max value
		maxUp = math.Max(upload, maxUp)
		maxDown = math.Max(download, maxDown)

		//getting the speed in mbps
		readableUpload := Round(bytesToMebibit(upload), 0.01)
		readableDownload := Round(bytesToMebibit(download), 0.01)

		maxUploadPercent := (upload / maxUp) * 100
		maxDownloadPercent := (download / maxDown) * 100

		//fmt.Println(latency, "ms");
		//fmt.Println("Upload ", readableUpload, "mbps", maxUploadPercent, "% from max upload")
		//fmt.Println("Download", readableDownload, "mbps", maxDownloadPercent, "% from max download")

		maxY, maxX := screen.MaxYX()

		/*
		uploadBar.Erase()
		uploadBar.Refresh()
		newUploadHeight, newUploadY := CalculateNewHeightAndY(maxUploadPercent, maxY)
		_, uploadWidth := uploadBar.MaxYX()
		_, uploadX := uploadBar.YX()
		uploadBar.Resize(newUploadHeight, uploadWidth)
		uploadBar.MoveWindow(newUploadY, uploadX)
		uploadBar.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE,
			gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
		uploadBar.MovePrint(0, 0, "Upload: "+strconv.FormatFloat(readableUpload, 'f', 2, 64)+"mbps")
		uploadBar.Refresh()
		*/
		uploadText := "Upload: " + strconv.FormatFloat(readableUpload, 'f', 2, 64) + "mbps"
		downloadText := "Download: " + strconv.FormatFloat(readableDownload, 'f', 2, 64) + "mbps";

		screen.Erase()
		screen.Refresh()
		UpdateBar(uploadBar, maxUploadPercent, maxY, 2)

		UpdateBar(downloadBar, maxDownloadPercent, maxY, 1)

		screen.ColorOn(4)
		screen.MovePrint(maxY/2-1, maxX/2-4, uploadText)
		screen.ColorOn(3)
		screen.MovePrint(maxY/2, maxX/2-4, downloadText)
		screen.ColorOff(3)
		screen.MovePrint(maxY/2+1, maxX/2-4, "Latency: "+strconv.FormatFloat(latency, 'f', 0, 64)+"ms")

		screen.Refresh()
		time.Sleep(1 * time.Second)
	}
}

func UpdateBar(bar *gc.Window, percent float64, maxY int, color int16) {

	newUploadHeight, newUploadY := CalculateNewHeightAndY(percent, maxY)
	_, uploadWidth := bar.MaxYX()
	_, uploadX := bar.YX()
	bar.Resize(newUploadHeight, uploadWidth)

	bar.ColorOn(color)
	bar.MoveWindow(newUploadY, uploadX)
	bar.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE,
		gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
	bar.Color(color)
	bar.ColorOff(color)
	bar.SetBackground(gc.ColorPair(color))
	bar.Refresh()

}

func CalculateNewHeightAndY(percent float64, maxY int) (int, int) {

	newHeight := int(float64(maxY) * (percent / 100))
	//fmt.Println("new height", newHeight)
	newY := maxY - newHeight
	return newHeight, newY
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func openConfig() Configuration {
	var conf Configuration
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		// handle error
		fmt.Println("SOMETHING WRONG !")
		panic(err)
	}

	return conf
}

func bytesToMebibit(bytes float64) float64 {
	return bytes / 131072
}

func getInfo(url string, site string, client *http.Client) (float64, float64, float64) {

	resp, err := client.Get(url + "/api/s/" + site + "/stat/health")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var i interface{}
	//fmt.Printf("response %q\n", body)

	if err := json.Unmarshal([]byte(body), &i); err != nil {
		panic(err)
	}

	data := i.(map[string]interface{})["data"].([]interface{})[2].(map[string]interface{})

	latency := data["latency"].(float64)
	upload := data["tx_bytes-r"].(float64)
	download := data["rx_bytes-r"].(float64)
	//return strconv.Atoi(www.([]interface{})["latency"]), www["tx_bytes-r"], www["rx_bytes-r"]
	return latency, upload, download
	//return 0, 0, 0
}

func login(url string, username string, password string, client *http.Client) {

	payload := strings.NewReader("{\n\t\"username\": \"" + username + "\",\n\t\"password\":\"" + password + "\"\n}")
	resp, err := client.Post(url+"/api/login", "application/json", payload)
	if err != nil {
		// handle error
		panic(err)
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response %q\n", body)
}
