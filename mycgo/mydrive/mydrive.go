package mydrive

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var color_reset = "\033[0m"
var color_green = "\033[32m"
var color_red = "\033[31m"
var color_yellow = "\033[33m"
var color_blue = "\033[34m"
var color_purple = "\033[35m"
var color_cyan = "\033[36m"
var color_gray = "\033[37m"
var color_white = "\033[97m"

// func main() {
// 	mydrive()
// 	// vimtest()
// }

func vimtest() {
	cmd := exec.Command("vim", "test.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", err)
	}
}

func colorPrint(color string, message string) {
	fmt.Println(color, message, color_reset)
}

func mydrive() {
	_, ok := os.LookupEnv("https_proxy")
	if !ok {
		setenv()
	}
	switch os.Args[1] {
	case "list":
		// listByPath(os.Args[2])
		if len(os.Args) == 2 {
			ListById("root")
		} else if len(os.Args) == 3 {
			ListById(os.Args[2])
		}
	case "upload":
		uploadByPath(os.Args[2], os.Args[3])
	case "update":
		updateByPath(os.Args[2], os.Args[3])
	case "updatebyid":
		updateById(os.Args[2], os.Args[3])
	case "print":
		printByPath(os.Args[2])
	case "printbyid":
		printById(os.Args[2])
	case "download":
		if os.Args[2] == "--replace" {
			downloadByPath(os.Args[3], true)
		} else {
			downloadByPath(os.Args[2], false)
		}
	case "downloadbyid":
		if os.Args[2] == "--replace" {
			downloadById(os.Args[3], true)
		} else {
			downloadById(os.Args[2], false)
		}
	case "search":
		search_name(os.Args[2])
	case "searchtext":
		search_text(os.Args[2])
	case "delete":
		deleteByPath(os.Args[2])
	case "deletebyid":
		deleteById(os.Args[2])
	case "show":
		showByName(os.Args[2])
	case "showbyid":
		showById(os.Args[2])
	case "revision":
		revision(os.Args[2])
	case "create":
		createByPath(os.Args[2])
	case "getpathbyId":
		getPathById(os.Args[2])
	case "getidbypath":
		getIdByPath(os.Args[2])
	case "help":
		help()
	case "version":
		fmt.Println("2023-03-07")
	default:
		fmt.Println("no action provided")
	}
}

func help() {
	help := `### list files or folders
    list path
### upload file to folder path
    upload folderPath local_file_path
### download by file path, do not replace existing file
    download fileName
### download by file path to replace existing file
    download --replace path
### download by file id, do not replace existing file
    downloadbyid fileId
### update file by path, file name will be changed to local file name
    update filepath local_file_path
### print file by path
    print filepath
### search file name
    search fileName
### search file content
    searchtext text
### delete file by path
    delete filepath
### delete file by id
    delete fildId
### show file info
    show fileId
### list revisions
    revision fileId
### create file by path from stdin
    create filepath`
	fmt.Println(help)
}

func setenv() {
	os.Setenv("https_proxy", "socks5://127.0.0.1:1080")
	// log.Println("https_proxy:", os.Getenv("https_proxy"))
}

// Retrieve a token, saves the token, then returns the generated client.
// func getClient(config *oauth2.Config) *http.Client {
// 	// The file token.json stores the user's access and refresh tokens, and is
// 	// created automatically when the authorization flow completes for the first
// 	// time.
// 	tokFile := "token.json"
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = getTokenFromWeb(config)
// 		saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }

func getClient(config *oauth2.Config) *http.Client {
	tokenString := `{"access_token":"ya29.a0AVvZVsqNpcnLeFlY3vpPfHsQBRrbV9d_2CbovYMS6SSBPg4ARj-L6hsJYqePD6iwnb2JfqhnXOznE9CJp_YcybjkfcRuff6UGOJQ5eBFIp6zByXMD9UhoeGIMv1H6v8yejJnrQoEh8wk1jObw7EItiOejprDaCgYKAY4SARESFQGbdwaIm5SITL75RN-ZlOlh_kyX7Q0163","token_type":"Bearer","refresh_token":"1//06KC78QedwJ2WCgYIARAAGAYSNwF-L9IrnOtMVSBPUko9yYf13UH0MprK9nYb0KPp_iXVeotnc1CUlj-9OayBH7TdxO0VeusAn1g","expiry":"2023-02-27T21:33:57.8121635+08:00"}`
	tok := &oauth2.Token{}
	err := json.NewDecoder(strings.NewReader(tokenString)).Decode(tok)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken("token.json", tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
//
//	func tokenFromFile(file string) (*oauth2.Token, error) {
//		f, err := os.Open(file)
//		if err != nil {
//			return nil, err
//		}
//		defer f.Close()
//		tok := &oauth2.Token{}
//		err = json.NewDecoder(f).Decode(tok)
//		return tok, err
//	}
func tokenFromFile() (*oauth2.Token, error) {
	tokenFile := `{"access_token":"ya29.a0AVvZVsqNpcnLeFlY3vpPfHsQBRrbV9d_2CbovYMS6SSBPg4ARj-L6hsJYqePD6iwnb2JfqhnXOznE9CJp_YcybjkfcRuff6UGOJQ5eBFIp6zByXMD9UhoeGIMv1H6v8yejJnrQoEh8wk1jObw7EItiOejprDaCgYKAY4SARESFQGbdwaIm5SITL75RN-ZlOlh_kyX7Q0163","token_type":"Bearer","refresh_token":"1//06KC78QedwJ2WCgYIARAAGAYSNwF-L9IrnOtMVSBPUko9yYf13UH0MprK9nYb0KPp_iXVeotnc1CUlj-9OayBH7TdxO0VeusAn1g","expiry":"2023-02-27T21:33:57.8121635+08:00"}`
	tok := &oauth2.Token{}
	err := json.NewDecoder(strings.NewReader(tokenFile)).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getService() *drive.Service {
	ctx := context.Background()
	credentials := `{"installed":{"client_id":"300366883097-63bf1p4a9fgu0m3qaoen3lktndpvar36.apps.googleusercontent.com","project_id":"neon-infinity-335803","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"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","redirect_uris":["http://localhost"]}}`
	b := []byte(credentials)
	// b, err := os.ReadFile("credentials.json")
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return service
}

func revision(fileId string) {
	revisions, err := getService().Revisions.List(fileId).Do()
	if err != nil {
		log.Println("list revision error: ", err)
		return
	}
	for _, revision := range revisions.Revisions {
		fmt.Println(revision)
	}
}

func getFileNameById(fileId string) string {
	file, err := getService().Files.Get(fileId).Fields("*").Do()
	if err != nil {
		log.Fatalln("get file error: ", err)
	}
	return file.Name
}

func getFilesInFolder(folderId string, fileName string) *drive.File {
	fileList, err := getService().Files.List().PageSize(1000).Q("'" + folderId + "' in parents and name = '" + fileName + "'").Fields("*").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	if len(fileList.Files) != 1 {
		log.Fatalln("file count is not 1")
	}
	return fileList.Files[0]
}

func getIdByPath(path string) string {
	if path == "/" {
		return "root"
	}
	var file *drive.File
	paths := strings.FieldsFunc(path,
		func(c rune) bool {
			return c == '/'
		})
	count := len(paths)
	for i := 0; i < count; i++ {
		if i == 0 {
			file = getFilesInFolder("root", paths[i])
		} else {
			file = getFilesInFolder(file.Id, paths[i])
		}
	}
	// fmt.Println("file.Id: ", file.Id)
	// fmt.Println("file.Name: ", file.Name)
	return file.Id
}

func getPathById(fileId string) string {
	file, err := getService().Files.Get(fileId).Fields("*").Do()
	if err != nil {
		log.Fatalln("get file error: ", err)
	}
	path := ""
	for {
		if len(file.Parents) == 0 {
			// path = file.Name + path
			// log.Println("path: ", path)
			return path
		} else if len(file.Parents) > 0 {
			path = "/" + file.Name + path
			file, _ = getService().Files.Get(file.Parents[0]).Fields("*").Do()
		}
	}
}

func downloadByPath(path string, replace bool) {
	downloadById(getIdByPath(path), replace)
}

func downloadById(fileId string, replace bool) {
	file, err := getService().Files.Get(fileId).Do()
	response, err := getService().Files.Get(fileId).Download()
	if err != nil {
		log.Println("download file error: ", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("read response error: ", err)
	}
	if _, err := os.Stat(file.Name); err == nil {
		log.Println("file exists: ", file.Name)
		if replace {
			log.Println("replace file: ", file.Name)
			ioutil.WriteFile(file.Name, body, 0644)
			log.Println("file write success: ", file.Name)
		} else {
			log.Println("do not replace file", file.Name)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Println("file not exist: ", file.Name)
		ioutil.WriteFile(file.Name, body, 0644)
		log.Println("file write success: ", file.Name)
	} else {
		log.Println("file may or may not exist")
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}
}

func localTime(timeString string) string {
	loc, _ := time.LoadLocation("Local")
	mytime1, _ := time.Parse(time.RFC3339, timeString)
	return mytime1.In(loc).Format("2006-01-02 15:04:05")
}

func showByName(fileName string) {
	showById(getFileIdByName(fileName))
}

func showById(fileId string) {
	file, err := getService().Files.Get(fileId).Fields("*").Do()
	if err != nil {
		log.Println("get file error: ", err)
		return
	}
	fmt.Println(color_yellow+"Id                    "+color_reset, color_purple+file.Id+color_reset)
	fmt.Println(color_yellow+"Name                  "+color_reset, color_purple+file.Name+color_reset)
	if len(file.Parents) > 0 {
		fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+file.Parents[0]+color_reset)
	} else {
		fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+"no parent"+color_reset)
	}
	fmt.Println(color_yellow+"Size                  "+color_reset, color_purple+strconv.Itoa(int(file.Size/1024))+"KB"+color_reset)
	fmt.Println(color_yellow+"MimeType              "+color_reset, color_purple+file.MimeType+color_reset)
	fmt.Println(color_yellow+"ModifiedTime          "+color_reset, color_purple+localTime(file.ModifiedTime)+color_reset)
	fmt.Println(color_yellow+"Version               "+color_reset, color_purple+strconv.Itoa(int(file.Version))+color_reset)
}

func deleteByPath(path string) {
	deleteById(getIdByPath(path))
}

func deleteById(fileId string) {
	file, err := getService().Files.Get(fileId).Do()
	if err != nil {
		log.Println("get file error: ", err)
		return
	}
	err = getService().Files.Delete(fileId).Do()
	if err != nil {
		log.Println("delete file error: ", err)
		return
	}
	log.Println("delete file success: ", file.Name)
}

func getFileIdByName(fileName string) string {
	fileList, err := getService().Files.List().PageSize(1000).Q("name = '" + fileName + "'").Fields("*").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	if len(fileList.Files) == 0 {
		log.Fatalln("no file found.")
	} else if len(fileList.Files) > 1 {
		log.Fatalln("multiple files found.")
	} else if len(fileList.Files) == 1 {
		return fileList.Files[0].Id
	} else {
		log.Fatalln("unknown error")
	}
	return ""
}

func printByPath(path string) {
	printById(getIdByPath(path))
}

func printById(fileId string) {
	response, err := getService().Files.Get(fileId).Download()
	if err != nil {
		log.Println("download file error: ", err)
		return
	}
	// Close body on function exit
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("read response error: ", err)
	}
	fmt.Println(string(body))
}

func updateByPath(path string, local_file_path string) {
	updateById(getIdByPath(path), local_file_path)
}

func updateById(fileId string, local_file_path string) {
	service := getService()
	file, err := os.Open(local_file_path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// _, err = service.Files.Update(fileId, &drive.File{MimeType: "text/plain"}).Media(file).Do()
	_, err = service.Files.Update(fileId, &drive.File{Name: filepath.Base(local_file_path)}).Media(file).Do()
	if err != nil {
		log.Println("update file error", err)
	}
	log.Println("file " + filepath.Base(local_file_path) + " updated")
}

func createByPath(path string) {
	createByFolderId(getIdByPath(filepath.Dir(path)), filepath.Base(path))
}

func createByFolderId(folderId string, fileName string) {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Lines:")
	var lines []string
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			// Group Separator (GS ^]): ctrl-]
			if line[0] == '\x1D' {
				break
			}
		}
		lines = append(lines, line)
	}
	// if len(lines) > 0 {
	// 	fmt.Println()
	// 	fmt.Println("Result:")
	// 	for _, line := range lines {
	// 		fmt.Println(line)
	// 	}
	// 	fmt.Println()
	// }
	file, _ := getService().Files.Create(&drive.File{Name: fileName, Parents: []string{folderId}, MimeType: "text/plain"}).Media(strings.NewReader(strings.Join(lines, "\n"))).Do()
	log.Println("created file id: ", file.Id)
	log.Println("created file name: ", file.Name)
}

func uploadByPath(path string, local_file_path string) {
	uploadByFolderId(getIdByPath(path), local_file_path)
}

func uploadByFolderId(folderId string, local_file_path string) {
	service := getService()
	file, err := os.Open(local_file_path)
	if err != nil {
		log.Fatalln(err)
	}
	fileInf, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	_, err = service.Files.Create(&drive.File{Name: filepath.Base(local_file_path), Parents: []string{folderId}}).
		ResumableMedia(context.Background(), file, fileInf.Size(), "text/plain").
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d KB, %d KB\r", now/1024, size/1024) }).
		Do()
	if err != nil {
		log.Println("upload file error", err)
	}
}

func search_text(text string) {
	fileList, err := getService().Files.List().PageSize(1000).Q("fullText contains '" + text + "'").Fields("*").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	if len(fileList.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, file := range fileList.Files {
			fmt.Println(color_yellow+"Id                    "+color_reset, color_purple+file.Id+color_reset)
			fmt.Println(color_yellow+"Name                  "+color_reset, color_purple+file.Name+color_reset)
			fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+file.Parents[0]+color_reset)
			fmt.Println(color_yellow+"Size                  "+color_reset, color_purple+strconv.Itoa(int(file.Size/1024))+"KB"+color_reset)
			fmt.Println(color_yellow+"MimeType              "+color_reset, color_purple+file.MimeType+color_reset)
			fmt.Println(color_yellow+"ModifiedTime          "+color_reset, color_purple+localTime(file.ModifiedTime)+color_reset)
			fmt.Println(color_yellow+"Version               "+color_reset, color_purple+strconv.Itoa(int(file.Version))+color_reset)
			fmt.Println()
		}
	}
}

func search_name(fileName string) {
	// https://developers.google.com/drive/api/guides/search-files
	fileList, err := getService().Files.List().PageSize(1000).Q("name contains '" + fileName + "'").Fields("*").OrderBy("folder,name").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	if len(fileList.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, file := range fileList.Files {
			fmt.Println(color_yellow+"Id                    "+color_reset, color_purple+file.Id+color_reset)
			fmt.Println(color_yellow+"Name                  "+color_reset, color_purple+file.Name+color_reset)
			fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+file.Parents[0]+color_reset)
			fmt.Println(color_yellow+"Path                  "+color_reset, color_purple+getPathById(file.Id)+color_reset)
			fmt.Println(color_yellow+"Size                  "+color_reset, color_purple+strconv.Itoa(int(file.Size/1024))+"KB"+color_reset)
			fmt.Println(color_yellow+"MimeType              "+color_reset, color_purple+file.MimeType+color_reset)
			fmt.Println(color_yellow+"ModifiedTime          "+color_reset, color_purple+localTime(file.ModifiedTime)+color_reset)
			fmt.Println(color_yellow+"Version               "+color_reset, color_purple+strconv.Itoa(int(file.Version))+color_reset)
			fmt.Println()
		}
	}
}

func ListById(fileId string) string {
	file, _ := getService().Files.Get(fileId).Fields("*").Do()
	result := ""
	if file.MimeType == "application/vnd.google-apps.folder" {
		fileList, err := getService().Files.List().PageSize(1000).Q("'" + fileId + "' in parents").Fields("*").OrderBy("folder,name").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}
		if len(fileList.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, item := range fileList.Files {
				result += fmt.Sprintln("Id           ", item.Id) +
					fmt.Sprintln("Name         ", item.Name) +
					fmt.Sprintln("Parent       ", item.Parents[0]) +
					fmt.Sprintln("Size         ", strconv.Itoa(int(item.Size/1024))+"KB") +
					fmt.Sprintln("MimeType     ", item.MimeType) +
					fmt.Sprintln("ModifiedTime ", localTime(item.ModifiedTime)) +
					fmt.Sprintln("Version      ", strconv.Itoa(int(item.Version))) + "\n"
			}
		}
	} else {
		result +=
			fmt.Sprintln("Id           ", file.Id) +
				fmt.Sprintln("Name           ", file.Name)
		if len(file.Parents) > 0 {
			result += fmt.Sprintln("Parent           ", file.Parents[0])
		} else {
			result += fmt.Sprintln("Parent           ", "no parent")
		}
		result +=
			fmt.Sprintln("Size           ", strconv.Itoa(int(file.Size/1024))+"KB") +
				fmt.Sprintln("MimeType           ", file.MimeType) +
				fmt.Sprintln("ModifiedTime           ", localTime(file.ModifiedTime)) +
				fmt.Sprintln("Version           ", strconv.Itoa(int(file.Version))) + "\n"
	}
	return result
}

func listByPath(path string) {
	fileId := getIdByPath(path)
	file, _ := getService().Files.Get(fileId).Fields("*").Do()
	if file.MimeType == "application/vnd.google-apps.folder" {
		fileList, err := getService().Files.List().PageSize(1000).Q("'" + fileId + "' in parents").Fields("*").OrderBy("folder,name").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve files: %v", err)
		}
		if len(fileList.Files) == 0 {
			fmt.Println("No files found.")
		} else {
			for _, item := range fileList.Files {
				fmt.Println(color_yellow+"Id                    "+color_reset, color_purple+item.Id+color_reset)
				fmt.Println(color_yellow+"Name                  "+color_reset, color_purple+item.Name+color_reset)
				fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+item.Parents[0]+color_reset)
				// fmt.Println(color_yellow+"Path                  "+color_reset, color_purple+getPathById(item.Id)+color_reset)
				fmt.Println(color_yellow+"Size                  "+color_reset, color_purple+strconv.Itoa(int(item.Size/1024))+"KB"+color_reset)
				fmt.Println(color_yellow+"MimeType              "+color_reset, color_purple+item.MimeType+color_reset)
				fmt.Println(color_yellow+"ModifiedTime          "+color_reset, color_purple+localTime(item.ModifiedTime)+color_reset)
				fmt.Println(color_yellow+"Version               "+color_reset, color_purple+strconv.Itoa(int(item.Version))+color_reset)
				fmt.Println()
			}
		}
	} else {
		fmt.Println(color_yellow+"Id                    "+color_reset, color_purple+file.Id+color_reset)
		fmt.Println(color_yellow+"Name                  "+color_reset, color_purple+file.Name+color_reset)
		if len(file.Parents) > 0 {
			fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+file.Parents[0]+color_reset)
		} else {
			fmt.Println(color_yellow+"Parent                "+color_reset, color_purple+"no parent"+color_reset)
		}
		// fmt.Println(color_yellow+"Path                  "+color_reset, color_purple+getPathById(file.Id)+color_reset)
		fmt.Println(color_yellow+"Size                  "+color_reset, color_purple+strconv.Itoa(int(file.Size/1024))+"KB"+color_reset)
		fmt.Println(color_yellow+"MimeType              "+color_reset, color_purple+file.MimeType+color_reset)
		fmt.Println(color_yellow+"ModifiedTime          "+color_reset, color_purple+localTime(file.ModifiedTime)+color_reset)
		fmt.Println(color_yellow+"Version               "+color_reset, color_purple+strconv.Itoa(int(file.Version))+color_reset)
	}

}
