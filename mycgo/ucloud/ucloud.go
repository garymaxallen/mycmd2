package ucloud

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"time"
)

// func main() {
// 	flagProcess()
// }

func flagProcess() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	var listTypeVar string
	listCmd.StringVar(&listTypeVar, "type", "vm", "type, vm, disk, eip, image")
	listCmd.StringVar(&listTypeVar, "t", "vm", "short alias of -type")
	var limitVar int
	listCmd.IntVar(&limitVar, "limit", 5, "limit")
	listCmd.IntVar(&limitVar, "l", 5, "short alias of -limit")
	var offsetVar int
	listCmd.IntVar(&offsetVar, "offset", 0, "offset")
	listCmd.IntVar(&offsetVar, "o", 0, "short alias of -offset")

	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	var showTypeVar string
	showCmd.StringVar(&showTypeVar, "type", "vm", "type, vm, disk, eip, image")
	showCmd.StringVar(&showTypeVar, "t", "vm", "short alias of -type")
	var showidVar string
	showCmd.StringVar(&showidVar, "id", "", "id")
	showCmd.StringVar(&showidVar, "i", "", "short alias of -id")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	var deleteTypeVar string
	deleteCmd.StringVar(&deleteTypeVar, "type", "", "type, vm, disk, eip, image")
	deleteCmd.StringVar(&deleteTypeVar, "t", "", "short alias of -type")
	var deleteidVar string
	deleteCmd.StringVar(&deleteidVar, "id", "", "id")
	deleteCmd.StringVar(&deleteidVar, "i", "", "short alias of -id")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	var createTypeVar string
	createCmd.StringVar(&createTypeVar, "type", "", "type, vm, disk, eip, image")
	createCmd.StringVar(&createTypeVar, "t", "", "short alias of -type")
	var nameVar string
	createCmd.StringVar(&nameVar, "name", "", "name")
	createCmd.StringVar(&nameVar, "n", "", "short alias of -name")
	var imageVar string
	createCmd.StringVar(&imageVar, "image", "", "image id")
	createCmd.StringVar(&imageVar, "i", "", "short alias of -image")
	var cpuVar int
	createCmd.IntVar(&cpuVar, "cpu", 2, "cpu")
	createCmd.IntVar(&cpuVar, "c", 2, "short alias of -cpu")
	var memoryVar int
	createCmd.IntVar(&memoryVar, "memory", 2048, "cpu")
	createCmd.IntVar(&memoryVar, "m", 2048, "short alias of -memory")
	var bootdiskVar int
	createCmd.IntVar(&bootdiskVar, "bootdisk", 40, "bootdisk size")
	createCmd.IntVar(&bootdiskVar, "b", 40, "short alias of -bootdisk")
	var datadiskVar int
	createCmd.IntVar(&datadiskVar, "datadisk", 0, "datadisk size")
	createCmd.IntVar(&datadiskVar, "d", 0, "short alias of -datadisk")

	startvmCmd := flag.NewFlagSet("startvm", flag.ExitOnError)
	var startidVar string
	startvmCmd.StringVar(&startidVar, "id", "", "vm id")
	startvmCmd.StringVar(&startidVar, "i", "", "short alias of -id")

	poweroffvmCmd := flag.NewFlagSet("poweroffvm", flag.ExitOnError)
	var poweroffidVar string
	poweroffvmCmd.StringVar(&poweroffidVar, "id", "", "vm id")
	poweroffvmCmd.StringVar(&poweroffidVar, "i", "", "short alias of -id")

	getvmidbyipCmd := flag.NewFlagSet("getvmidbyip", flag.ExitOnError)
	var ipVar string
	getvmidbyipCmd.StringVar(&ipVar, "ip", "", "ip")
	getvmidbyipCmd.StringVar(&ipVar, "i", "", "short alias of -ip")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommands, list, show, create, delete, poweroffvm, startvm, getvmid, etc.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		if listTypeVar == "vm" {
			listVM(limitVar, offsetVar)
		} else if listTypeVar == "disk" {
			listDisk(limitVar, offsetVar)
		} else if listTypeVar == "eip" {
			listEIP(limitVar, offsetVar)
		} else if listTypeVar == "image" {
			listImage(limitVar, offsetVar)
		}
	case "show":
		showCmd.Parse(os.Args[2:])
		if len(showidVar) == 0 {
			fmt.Println("id is empty")
			os.Exit(1)
		}
		if showTypeVar == "vm" {
			showVM(showidVar)
		} else if showTypeVar == "disk" {
			showDisk(showidVar)
		} else if showTypeVar == "eip" {
			showEIP(showidVar)
		} else if showTypeVar == "image" {
			showImage(showidVar)
		}
	case "create":
		createCmd.Parse(os.Args[2:])
		if len(nameVar) == 0 || len(imageVar) == 0 || cpuVar == 0 || memoryVar == 0 || bootdiskVar == 0 {
			fmt.Println("parameter error")
			os.Exit(1)
		}
		if createTypeVar == "vm" {
			createvm(bootdiskVar, cpuVar, datadiskVar, imageVar, memoryVar, nameVar)
		} else if createTypeVar == "disk" {
			fmt.Println("not yet implemented")
			// createDisk(createTypeVar)
		} else if createTypeVar == "eip" {
			// createEIP(createTypeVar)
			fmt.Println("not yet implemented")
		} else if createTypeVar == "image" {
			// create(createTypeVar)
			fmt.Println("not yet implemented")
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if len(deleteidVar) == 0 || len(deleteTypeVar) == 0 {
			fmt.Println("type or id is empty")
			os.Exit(1)
		}
		if deleteTypeVar == "vm" {
			deleteVM(deleteidVar)
		} else if deleteTypeVar == "disk" {
			deleteDisk(deleteidVar)
		} else if deleteTypeVar == "eip" {
			deleteEIP(deleteidVar)
		} else if deleteTypeVar == "image" {
			fmt.Println("not yet implemented")
			// deleteImage(limitVar, offsetVar)
		}
	case "startvm":
		startvmCmd.Parse(os.Args[2:])
		if len(startidVar) == 0 {
			fmt.Println("id is empty")
			os.Exit(1)
		}
		startVM(startidVar)
	case "poweroffvm":
		poweroffvmCmd.Parse(os.Args[2:])
		if len(poweroffidVar) == 0 {
			fmt.Println("id is empty")
			os.Exit(1)
		}
		poweroffVM(poweroffidVar)
	case "getvmid":
		getvmidbyipCmd.Parse(os.Args[2:])
		if len(ipVar) == 0 {
			fmt.Println("ip is empty")
			os.Exit(1)
		}
		getvmidByEIP(ipVar)
	case "version":
		fmt.Println("version 2023-03-10")
	default:
		fmt.Println("expected subcommands, list, show, create, delete, poweroffvm, startvm, getvmid, etc.")
		os.Exit(1)
	}
}

func startVM(vmid string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "StartVMInstance",
		"Region":    "cn",
		"Zone":      "zone-01",
		"VMID":      vmid,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
	fmt.Println()
}

func poweroffVM(vmid string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "PoweroffVMInstance",
		"Region":    "cn",
		"Zone":      "zone-01",
		"VMID":      vmid,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
	fmt.Println()
}

func getvmidByEIP(ip string) string {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     100,
	}
	ipnum := getipnum()
	count := int(math.Ceil(float64(ipnum) / float64(100)))
	for i := 0; i < count; i++ {
		params["Offset"] = 0
		target := paramsRequest(verify_ac(params))

		if ipnum > params["Limit"].(int) {
			for i := 0; i < params["Limit"].(int); i++ {
				if target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IP"] == ip {
					fmt.Println("EIPID:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["EIPID"])
					fmt.Println("IP:                    ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IP"])
					fmt.Println("BindResourceID:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceID"])
					fmt.Println("BindResourceType:      ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceType"])
					fmt.Println("Status:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Status"])
					fmt.Println()
					return target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceID"].(string)
				}
			}
			params["Offset"] = params["Offset"].(int) + params["Limit"].(int)
			ipnum = ipnum - params["Limit"].(int)
		} else {
			for i := 0; i < ipnum; i++ {
				if target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IP"] == ip {
					// fmt.Println("EIPID:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["EIPID"])
					// fmt.Println("IP:                    ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IP"])
					// fmt.Println("BindResourceID:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceID"])
					// fmt.Println("BindResourceType:      ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceType"])
					// fmt.Println("Status:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Status"])
					// fmt.Println()
					return target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceID"].(string)
				}
			}
		}
	}
	return "error"
}

func getipnum() int {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     1,
		"Offset":    0,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("ipnum: ", int(target.(map[string]interface{})["TotalCount"].(float64)))
	return int(target.(map[string]interface{})["TotalCount"].(float64))
}

func Getipnum() int {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     1,
		"Offset":    0,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("ipnum: ", int(target.(map[string]interface{})["TotalCount"].(float64)))
	return int(target.(map[string]interface{})["TotalCount"].(float64))
}

func deleteVM(vmid string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DeleteVMInstance",
		"Region":    "cn",
		"Zone":      "zone-01",
		"VMID":      vmid,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
	fmt.Println()
}

func createvm(bootdiskVar int, cpuVar int, datadiskVar int, imageVar string, memoryVar int, nameVar string) {
	params := map[string]interface{}{
		"PublicKey":       PublicKey,
		"Action":          "CreateVMInstance",
		"Region":          "cn",
		"Zone":            "zone-01",
		"BootDiskSetType": "Normal",
		// "BootDiskSpace":   40,
		"ChargeType": "Year",
		// "CPU":             1,
		// "DataDiskSetType": "Normal",
		// "DataDiskSpace":   10,
		// "ImageID":         "cn-image-centos-65",
		// "Memory":          2048,
		// "Name":            "gg2023",
		"Password":     "ucloud.cn",
		"SubnetID":     "subnet-MzXZh5_xW",
		"VMType":       "Normal",
		"VPCID":        "vpc-MzXZh5_xW",
		"WANSGID":      "sg-MzXZh5_xW",
		"OperatorName": "wan-bgp-R6KtjeNmR",
		"Bandwidth":    "10000",
		"IPVersion":    "IPv4",
	}
	params["BootDiskSpace"] = bootdiskVar
	params["CPU"] = cpuVar
	if datadiskVar > 0 {
		params["DataDiskSpace"] = datadiskVar
		params["DataDiskSetType"] = "Normal"
	}
	params["ImageID"] = imageVar
	params["Memory"] = memoryVar
	params["Name"] = nameVar
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
	fmt.Println()
}

func showVM(vmid string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeVMInstance",
		"Region":    "cn",
		"Zone":      "zone-01",
		"VMIDs.0":   vmid,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	if len(target.(map[string]interface{})["Infos"].([]interface{})) == 0 {
		log.Fatalln("data length is 0")
	}
	fmt.Println("VMID:            ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["VMID"])
	fmt.Println("Name:            ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["Name"])
	fmt.Println("State:           ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["State"])
	fmt.Println("CPU:             ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["CPU"])
	fmt.Println("Memory:          ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["Memory"])
	fmt.Println("OSName:          ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["OSName"])
	fmt.Println("VMType:          ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["VMType"])
	fmt.Println("IP[0]:           ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["IPInfos"].([]interface{})[0].(map[string]interface{})["IP"])
	fmt.Println("IP[1]:           ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["IPInfos"].([]interface{})[1].(map[string]interface{})["IP"])
	fmt.Println("Disk[0]-DiskID:  ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["DiskID"])
	fmt.Println("Disk[0]-Size:    ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["Size"])
	fmt.Println("Disk[0]-Type:    ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["Type"])
	if len(target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})) > 1 {
		fmt.Println("Disk[1]-DiskID:  ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["DiskID"])
		fmt.Println("Disk[1]-Size:    ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["Size"])
		fmt.Println("Disk[1]-Type:    ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["Type"])
	}
	fmt.Println()
}

func deleteDisk(id string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DeleteDisk",
		"Region":    "cn",
		"Zone":      "zone-01",
		"DiskID":    id,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
}

func listDisk(limit int, offset int) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeDisk",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	for i := 0; i < limit; i++ {
		fmt.Println("DiskID:               ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskID"])
		fmt.Println("DiskStatus:           ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskStatus"])
		fmt.Println("DiskType:             ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskType"])
		fmt.Println("Size:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Size"])
		fmt.Println("AttachResourceID:     ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["AttachResourceID"])
		fmt.Println("AttachResourceType:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["AttachResourceType"])
		fmt.Println("CreateTime:           ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["CreateTime"].(float64)), 0))
		fmt.Println("ExpireTime:           ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ExpireTime"].(float64)), 0))
		fmt.Println()
	}
}

func showDisk(id string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeDisk",
		"Region":    "cn",
		"Zone":      "zone-01",
		"DiskIDs.0": id,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
	fmt.Println()
	if len(target.(map[string]interface{})["Infos"].([]interface{})) == 0 {
		log.Fatalln("data length is 0")
	}
	fmt.Println("DiskID:               ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskID"])
	fmt.Println("DiskStatus:           ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskStatus"])
	fmt.Println("DiskType:             ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["DiskType"])
	fmt.Println("Size:                 ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["Size"])
	fmt.Println("AttachResourceID:     ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["AttachResourceID"])
	fmt.Println("AttachResourceType:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["AttachResourceType"])
	fmt.Println("CreateTime:           ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["CreateTime"].(float64)), 0))
	fmt.Println("ExpireTime:           ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ExpireTime"].(float64)), 0))
	fmt.Println()
}

func verify_ac(params map[string]interface{}) map[string]interface{} {
	params_data := ""
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		params_data = params_data + key + fmt.Sprint(params[key])
	}
	private_key := "e2a5a1cdf459e89d9ff7f52a8bd2e9c035a34ad7"
	params_data = params_data + private_key
	// fmt.Println("params_data: ", params_data)

	hash := sha1.New()
	hash.Write([]byte(params_data))
	signature := hex.EncodeToString(hash.Sum(nil))
	params["Signature"] = signature
	return params
}

var PublicKey = "OmgolGAwCsGsMSo66+L0oDFKFUM6gVVKR0qsKTKwJr/zyCoKHsehIK8Ftq2DIotP"

func showEIP(id string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"EIPIDs.0":  id,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	fmt.Println()
	if len(target.(map[string]interface{})["Infos"].([]interface{})) == 0 {
		log.Fatalln("data length is 0")
	}
	fmt.Println("EIPID:               ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["EIPID"])
	fmt.Println("Bandwidth:           ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["Bandwidth"])
	fmt.Println("BindResourceID:      ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["BindResourceID"])
	fmt.Println("IP:                  ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["IP"])
	fmt.Println("OperatorName:        ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["OperatorName"])
	fmt.Println("Status:              ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["Status"])
	fmt.Println("CreateTime:          ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["CreateTime"].(float64)), 0))
	fmt.Println("ExpireTime:          ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ExpireTime"].(float64)), 0))
	fmt.Println()
}

func deleteEIP(eipid string) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "ReleaseEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"EIPID":     eipid,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("target: ", target)
}

func DeleteEIP(eipid string) string {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "ReleaseEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"EIPID":     eipid,
	}
	target := paramsRequest(verify_ac(params))
	return fmt.Sprintf("%v", target)
}

func listEIP(limit int, offset int) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	count := limit
	for i := 0; i < count; i++ {
		fmt.Println("EIPID:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["EIPID"])
		fmt.Println("IP:                    ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IP"])
		fmt.Println("BindResourceID:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceID"])
		fmt.Println("BindResourceType:      ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceType"])
		fmt.Println("OperatorName:          ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["OperatorName"])
		fmt.Println("Status:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Status"])
		fmt.Println()
	}
}

func ListEIP(limit int, offset int) string {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeEIP",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	result := ""
	target := paramsRequest(verify_ac(params))
	for i := 0; i < len(target.(map[string]interface{})["Infos"].([]interface{})); i++ {
		result = result +
			fmt.Sprintf("%v%v", "EIPID:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["EIPID"]) + "\n" +
			fmt.Sprintf("%v%v", "IP:                    ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IP"]) + "\n" +
			fmt.Sprintf("%v%v", "BindResourceID:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceID"]) + "\n" +
			fmt.Sprintf("%v%v", "BindResourceType:      ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["BindResourceType"]) + "\n" +
			fmt.Sprintf("%v%v", "OperatorName:          ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["OperatorName"]) + "\n" +
			fmt.Sprintf("%v%v", "Status:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Status"]) + "\n" + "\n"
	}
	return result
}

func showImage(id string) {
	params := map[string]interface{}{
		"PublicKey":  PublicKey,
		"Action":     "DescribeImage",
		"Region":     "cn",
		"Zone":       "zone-01",
		"ImageIDs.0": id,
	}
	target := paramsRequest(verify_ac(params))
	fmt.Println("ImageID:     ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageID"])
	fmt.Println("ImageName:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageName"])
	fmt.Println("ImageSize:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageSize"])
	fmt.Println("OSName:      ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["OSName"])
	fmt.Println("ImageType:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageType"])
	fmt.Println("ImageStatus: ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageStatus"])
	fmt.Println()

}

func ShowImage(id string) string {
	params := map[string]interface{}{
		"PublicKey":  PublicKey,
		"Action":     "DescribeImage",
		"Region":     "cn",
		"Zone":       "zone-01",
		"ImageIDs.0": id,
	}
	result := ""
	target := paramsRequest(verify_ac(params))
	result = result +
		fmt.Sprintf("%v%v", "ImageID:     ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageID"]) + "\n" +
		fmt.Sprintf("%v%v", "ImageName:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageName"]) + "\n" +
		fmt.Sprintf("%v%v", "ImageSize:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageSize"]) + "\n" +
		fmt.Sprintf("%v%v", "OSName:      ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["OSName"]) + "\n" +
		fmt.Sprintf("%v%v", "ImageType:   ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageType"]) + "\n" +
		fmt.Sprintf("%v%v", "ImageStatus: ", target.(map[string]interface{})["Infos"].([]interface{})[0].(map[string]interface{})["ImageStatus"]) + "\n" + "\n"
	return result
}

func listImage(limit int, offset int) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeImage",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	// count := stoi(limit)
	for i := 0; i < len(target.(map[string]interface{})["Infos"].([]interface{})); i++ {
		fmt.Println("ImageID:     ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageID"])
		fmt.Println("ImageName:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageName"])
		fmt.Println("ImageSize:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageSize"])
		fmt.Println("OSName:      ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["OSName"])
		fmt.Println("ImageType:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageType"])
		fmt.Println("ImageStatus: ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageStatus"])
		fmt.Println()
	}
}

func ListImage(limit int, offset int) string {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeImage",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	target := paramsRequest(verify_ac(params))
	result := ""
	for i := 0; i < len(target.(map[string]interface{})["Infos"].([]interface{})); i++ {
		result = result +
			fmt.Sprintf("%v%v", "ImageID:     ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageID"]) + "\n" +
			fmt.Sprintf("%v%v", "ImageName:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageName"]) + "\n" +
			fmt.Sprintf("%v%v", "ImageSize:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageSize"]) + "\n" +
			fmt.Sprintf("%v%v", "OSName:      ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["OSName"]) + "\n" +
			fmt.Sprintf("%v%v", "ImageType:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageType"]) + "\n" +
			fmt.Sprintf("%v%v", "ImageStatus: ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ImageStatus"]) + "\n" + "\n"
	}
	return result
}

func listVM(limit int, offset int) {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeVMInstance",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	count := limit
	for i := 0; i < count; i++ {
		fmt.Println("VMID:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["VMID"])
		fmt.Println("Name:                 ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Name"])
		fmt.Println("State:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["State"])
		fmt.Println("CPU:                  ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["CPU"])
		fmt.Println("Memory:               ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Memory"])
		fmt.Println("OSName:               ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["OSName"])
		fmt.Println("IP[0]:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IPInfos"].([]interface{})[0].(map[string]interface{})["IP"])
		fmt.Println("IP[1]:                ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IPInfos"].([]interface{})[1].(map[string]interface{})["IP"])
		fmt.Println("Disk[0]-DiskID:       ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["DiskID"])
		fmt.Println("Disk[0]-Size:         ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["Size"])
		fmt.Println("Disk[0]-Type:         ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["Type"])
		if len(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})) > 1 {
			fmt.Println("Disk[1]-DiskID:       ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["DiskID"])
			fmt.Println("Disk[1]-Size:         ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["Size"])
			fmt.Println("Disk[1]-Type:         ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["Type"])
		}
		fmt.Println("CreateTime:           ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["CreateTime"].(float64)), 0))
		fmt.Println("ExpireTime:           ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ExpireTime"].(float64)), 0))
		fmt.Println()
	}
}

func ListVM(limit int, offset int) string {
	params := map[string]interface{}{
		"PublicKey": PublicKey,
		"Action":    "DescribeVMInstance",
		"Region":    "cn",
		"Zone":      "zone-01",
		"Limit":     limit,
		"Offset":    offset,
	}
	target := paramsRequest(verify_ac(params))
	// fmt.Println("target: ", target)
	result := ""
	for i := 0; i < len(target.(map[string]interface{})["Infos"].([]interface{})); i++ {
		result = result +
			fmt.Sprintf("%v%v", "VMID:         ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["VMID"]) + "\n" +
			fmt.Sprintf("%v%v", "Name:         ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Name"]) + "\n" +
			fmt.Sprintf("%v%v", "State:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["State"]) + "\n" +
			fmt.Sprintf("%v%v", "CPU:          ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["CPU"]) + "\n" +
			fmt.Sprintf("%v%v", "Memory:       ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["Memory"]) + "\n" +
			fmt.Sprintf("%v%v", "OSName:       ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["OSName"]) + "\n" +
			fmt.Sprintf("%v%v", "IP[0]:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IPInfos"].([]interface{})[0].(map[string]interface{})["IP"]) + "\n" +
			fmt.Sprintf("%v%v", "IP[1]:        ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["IPInfos"].([]interface{})[1].(map[string]interface{})["IP"]) + "\n" +
			fmt.Sprintf("%v%v", "Disk[0]-ID:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["DiskID"]) + "\n" +
			fmt.Sprintf("%v%v", "Disk[0]-Size: ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["Size"]) + "\n" +
			fmt.Sprintf("%v%v", "Disk[0]-Type: ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[0].(map[string]interface{})["Type"]) + "\n"
		if len(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})) > 1 {
			result = result +
				fmt.Sprintf("%v%v", "Disk[1]-ID:   ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["DiskID"]) + "\n" +
				fmt.Sprintf("%v%v", "Disk[1]-Size: ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["Size"]) + "\n" +
				fmt.Sprintf("%v%v", "Disk[1]-Type: ", target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["DiskInfos"].([]interface{})[1].(map[string]interface{})["Type"]) + "\n"
		}
		result = result +
			fmt.Sprintf("%v%v", "CreateTime:    ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["CreateTime"].(float64)), 0)) + "\n" +
			fmt.Sprintf("%v%v", "ExpireTime:    ", time.Unix(int64(target.(map[string]interface{})["Infos"].([]interface{})[i].(map[string]interface{})["ExpireTime"].(float64)), 0)) + "\n" + "\n"
	}
	// return fmt.Sprintf("%v", target)
	return result
}

func paramsRequest(params map[string]interface{}) interface{} {
	jsonByte, err := json.Marshal(params)
	if err != nil {
		log.Fatalln("get json from params error: ", err)
	}
	// fmt.Println("params: ", string(jsonByte))

	request, err := http.NewRequest("POST", "http://10.11.104.1/api", bytes.NewReader(jsonByte))
	//request, err := http.NewRequest("POST", "http://192.168.237.2/api", bytes.NewReader(jsonByte))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "text/plain")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	var target interface{}
	json.NewDecoder(response.Body).Decode(&target)
	// fmt.Println("target: ", target)
	return target
}
