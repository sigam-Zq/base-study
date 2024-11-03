package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
)

func main() {
	log.Printf("mac MD5 : %s \n", getMacMD5())

	// 在老家 2024/10/04 07:33:46 mac MD5 : facf51f91ac0ded86a567212bb3d60d6
	/*
										当前网卡情况

								2024/10/04 07:47:02 append inter name 以太网
								2024/10/04 07:47:02 append string 16 64383a62623a63313a61393a38663a3262
								2024/10/04 07:47:02 append string d8:bb:c1:a9:8f:2b
								2024/10/04 07:47:02 append inter name 以太网 2
								2024/10/04 07:47:02 append string 16 30613a30303a32373a30303a30303a3134
								2024/10/04 07:47:02 append string 0a:00:27:00:00:14
								2024/10/04 07:47:02 append inter name 以太网 3
								2024/10/04 07:47:02 append string 16 30303a66663a32393a33353a66653a3639
								2024/10/04 07:47:02 append string 00:ff:29:35:fe:69
								2024/10/04 07:47:02 append inter name WLAN
								2024/10/04 07:47:02 append string 16 63613a62393a63653a66363a62363a6135
								2024/10/04 07:47:02 append string ca:b9:ce:f6:b6:a5
								2024/10/04 07:47:02 macAddress string [d8:bb:c1:a9:8f:2b 0a:00:27:00:00:14 00:ff:29:35:fe:69 ca:b9:ce:f6:b6:a5]
								2024/10/04 07:47:02 sort after macAddress string [00:ff:29:35:fe:69 0a:00:27:00:00:14 ca:b9:ce:f6:b6:a5 d8:bb:c1:a9:8f:2b]
								2024/10/04 07:47:02 mac MD5 : facf51f91ac0ded86a567212bb3d60d6

							换了个地方连接的wifi
						2024/10/04 22:46:05 append inter name 以太网
						2024/10/04 22:46:05 append string 16 64383a62623a63313a61393a38663a3262
						2024/10/04 22:46:05 append string d8:bb:c1:a9:8f:2b
						2024/10/04 22:46:05 append inter name 以太网 2
						2024/10/04 22:46:05 append string 16 30613a30303a32373a30303a30303a3134
						2024/10/04 22:46:05 append string 0a:00:27:00:00:14
						2024/10/04 22:46:05 append inter name 以太网 3
						2024/10/04 22:46:05 append string 16 30303a66663a32393a33353a66653a3639
						2024/10/04 22:46:05 append string 00:ff:29:35:fe:69
						2024/10/04 22:46:05 append inter name WLAN
						2024/10/04 22:46:05 append string 16 32383a31313a61383a37303a64633a6362
						2024/10/04 22:46:05 append string 	28:11:a8:70:dc:cb
						2024/10/04 22:46:05 macAddress string [d8:bb:c1:a9:8f:2b 0a:00:27:00:00:14 00:ff:29:35:fe:69 28:11:a8:70:dc:cb]
						2024/10/04 22:46:05 sort after macAddress string [00:ff:29:35:fe:69 0a:00:27:00:00:14 28:11:a8:70:dc:cb d8:bb:c1:a9:8f:2b]
						2024/10/04 22:46:05 mac MD5 : 55e0238fea20d2279a0a285c3d29e983




						2024/10/08 09:37:55 append inter name 以太网
				2024/10/08 09:37:55 append string 16 64383a62623a63313a61393a38663a3262
				2024/10/08 09:37:55 append string d8:bb:c1:a9:8f:2b
				2024/10/08 09:37:55 append inter name 以太网 2
				2024/10/08 09:37:55 append string 16 30613a30303a32373a30303a30303a3134
				2024/10/08 09:37:55 append string 0a:00:27:00:00:14
				2024/10/08 09:37:55 append inter name 以太网 3
				2024/10/08 09:37:55 append string 16 30303a66663a32393a33353a66653a3639
				2024/10/08 09:37:55 append string 00:ff:29:35:fe:69
				2024/10/08 09:37:55 append inter name WLAN
				2024/10/08 09:37:55 append string 16 32383a31313a61383a37303a64633a6362
				2024/10/08 09:37:55 append string 28:11:a8:70:dc:cb
				2024/10/08 09:37:55 macAddress string [d8:bb:c1:a9:8f:2b 0a:00:27:00:00:14 00:ff:29:35:fe:69 28:11:a8:70:dc:cb]
				2024/10/08 09:37:55 sort after macAddress string [00:ff:29:35:fe:69 0a:00:27:00:00:14 28:11:a8:70:dc:cb d8:bb:c1:a9:8f:2b]
				2024/10/08 09:37:55 mac MD5 : 55e0238fea20d2279a0a285c3d29e983
				(base)


				在家

				$ ./run
		2024/10/13 11:41:06 append inter name 以太网
		2024/10/13 11:41:06 append string 16 64383a62623a63313a61393a38663a3262
		2024/10/13 11:41:06 append string d8:bb:c1:a9:8f:2b
		2024/10/13 11:41:06 append inter name 以太网 2
		2024/10/13 11:41:06 append string 16 30613a30303a32373a30303a30303a3136
		2024/10/13 11:41:06 append string 0a:00:27:00:00:16
		2024/10/13 11:41:06 append inter name 以太网 3
		2024/10/13 11:41:06 append string 16 30303a66663a32393a33353a66653a3639
		2024/10/13 11:41:06 append string 00:ff:29:35:fe:69
		2024/10/13 11:41:06 append inter name WLAN
		2024/10/13 11:41:06 append string 16 32383a31313a61383a37303a64633a6362
		2024/10/13 11:41:06 append string 28:11:a8:70:dc:cb
		2024/10/13 11:41:06 macAddress string [d8:bb:c1:a9:8f:2b 0a:00:27:00:00:16 00:ff:29:35:fe:69 28:11:a8:70:dc:cb]
		2024/10/13 11:41:06 sort after macAddress string [00:ff:29:35:fe:69 0a:00:27:00:00:16 28:11:a8:70:dc:cb d8:bb:c1:a9:8f:2b]
		2024/10/13 11:41:06 mac MD5 : 0ef23d86ebad4a378b7ac0815f3a319e
	*/
}

func getMacMD5() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	var macAddress []string
	for _, inter := range interfaces {
		// 大于en6的排除
		if strings.HasPrefix(inter.Name, "en") {
			numStr := inter.Name[2:]
			num, _ := strconv.Atoi(numStr)
			if num > 6 {
				log.Printf("Printf inter name %s \n", inter.Name)
				continue
			}
		}
		if strings.HasPrefix(inter.Name, "en") || strings.HasPrefix(inter.Name, "Ethernet") || strings.HasPrefix(inter.Name, "以太网") || strings.HasPrefix(inter.Name, "WLAN") {

			log.Printf("append inter name %s \n", inter.Name)
			log.Printf("append string 16 %x \n", inter.HardwareAddr)
			log.Printf("append string %s \n", inter.HardwareAddr.String())
			macAddress = append(macAddress, inter.HardwareAddr.String())
		}
	}
	log.Printf("macAddress string %s \n", macAddress)
	sort.Strings(macAddress)
	log.Printf("sort after macAddress string %s \n", macAddress)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(macAddress, ","))))
}
