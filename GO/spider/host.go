package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ListResp struct {
	Amount      int       `json:"amount"`
	ID          string    `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Type        string    `gorm:"column:type" json:"type"`
	VMType      string    `gorm:"column:vm_type" json:"vmType"`
	ClusterId   string    `gorm:"column:cluster_id" json:"clusterId"`
	AggregateId string    `gorm:"column:aggregate_id" json:"aggregateId" `
	IsDefault   int       `gorm:"column:is_default" json:"isDefault"`
	CreateTime  time.Time `gorm:"column:created_at" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:updated_at" json:"updateTime"`
	OriginName  string    `gorm:"column:origin_name" json:"originName"`
}

type ListResp2 struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	AvailableZoneId string `json:"availableZoneId"`
	ClusterName     string `json:"clusterName"`
	HostGroupName   string `json:"hostGroupName"`
}

func main() {
	hosts := []ListResp{}
	data := `[{
    "amount": 0,
    "id": "30ca5a9196b043b4870952fa936cdfa4",
    "name": "aa",
    "type": "系统组",
    "vmType": "",
    "clusterId": "",
    "aggregateId": "",
    "isDefault": 0,
    "createTime": "2022-05-30T14:19:10+08:00",
    "updateTime": "2022-05-30T14:19:10+08:00",
    "originName": ""
  },
  {
    "amount": 0,
    "id": "fb747a50076e4cd69fb33092de988070",
    "name": "11",
    "type": "系统组",
    "vmType": "",
    "clusterId": "",
    "aggregateId": "",
    "isDefault": 0,
    "createTime": "2022-05-26T14:26:41+08:00",
    "updateTime": "2022-05-26T14:26:41+08:00",
    "originName": ""
  },
  {
    "amount": 9,
    "id": "073129a8-96b7-4295-9dc8-ffdca490",
    "name": "默认主机组",
    "type": "系统组",
    "vmType": "",
    "clusterId": "",
    "aggregateId": "",
    "isDefault": 1,
    "createTime": "2022-05-16T14:30:30+08:00",
    "updateTime": "2022-05-16T14:30:30+08:00",
    "originName": ""
  }
]`
	res := []ListResp{}
	json.Unmarshal([]byte(data), &hosts)

	for _, host := range hosts {
		if host.Name != "默认主机组" {
			res = append(res, host)
		}
	}

	data1 := res
	fmt.Println(res)
	fmt.Println(len(res))
	fmt.Println(data1)

	fmt.Println("--------------------------------")
	res2 := []ListResp2{}
	data2 := `[
		{
            "id": "1763443641304106a4c1b334d8c6a86b",
            "uuid": "697144e8-9ffc-4eb6-9b10-af5b8ee856b3@controller01",
            "name": "controller01",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.83.31",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "fa3013f0-ddbd-4611-851e-c7c0ee90772a",
            "s8": "8d9954cb-3c37-49ab-a6e3-82a36c30c264",
            "isDelete": false,
            "clusterId": "697144e8-9ffc-4eb6-9b10-af5b8ee856b3",
            "clusterName": "ArNet可用区-ArcherOS0525",
            "availableZoneId": "21942fcb-a7f3-4fe4-977f-01f0e14cfb5f",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000124",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "127.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "network,storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-27 09:27:01",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "2ee357c374b94ab1a6fcff7770fc5f4b",
            "uuid": "697144e8-9ffc-4eb6-9b10-af5b8ee856b3@controller03",
            "name": "controller03",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.83.33",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "9137b8fb-1501-4f94-abcb-182cb36f2f11",
            "s8": "8d9954cb-3c37-49ab-a6e3-82a36c30c264",
            "isDelete": false,
            "clusterId": "697144e8-9ffc-4eb6-9b10-af5b8ee856b3",
            "clusterName": "ArNet可用区-ArcherOS0525",
            "availableZoneId": "21942fcb-a7f3-4fe4-977f-01f0e14cfb5f",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000112",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "127.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-27 09:27:01",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "7c963e014cb249cda2b250e77265296a",
            "uuid": "697144e8-9ffc-4eb6-9b10-af5b8ee856b3@controller02",
            "name": "controller02",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.83.32",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "a4e1f118-5b53-4805-8b3b-cbbf76b9957c",
            "s8": "8d9954cb-3c37-49ab-a6e3-82a36c30c264",
            "isDelete": false,
            "clusterId": "697144e8-9ffc-4eb6-9b10-af5b8ee856b3",
            "clusterName": "ArNet可用区-ArcherOS0525",
            "availableZoneId": "21942fcb-a7f3-4fe4-977f-01f0e14cfb5f",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000131",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "127.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "network,storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-27 09:27:01",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "4a2a7f891d2f4cafa72a94eda9d8bbb3",
            "uuid": "7fad0335-9f62-411d-8403-88adf01e5b00@controller02",
            "name": "controller02",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.17.202",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "de2b052a-7f90-42ee-83f2-7c36154db289",
            "s8": "1e314c6a-c46f-49bb-92ea-2d398fac5820",
            "isDelete": false,
            "clusterId": "7fad0335-9f62-411d-8403-88adf01e5b00",
            "clusterName": "SDN可用区-ArcherOS",
            "availableZoneId": "0260c681-30da-41be-a38c-7eed622bf2ae",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000160",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "127.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "network,storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-24 11:31:02",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "61bc3caa7ef44760b76c39437975e457",
            "uuid": "7fad0335-9f62-411d-8403-88adf01e5b00@controller01",
            "name": "controller01",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.17.201",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "83d3fca8-71ca-4492-a7d5-27b6e71bb09b",
            "s8": "1e314c6a-c46f-49bb-92ea-2d398fac5820",
            "isDelete": false,
            "clusterId": "7fad0335-9f62-411d-8403-88adf01e5b00",
            "clusterName": "SDN可用区-ArcherOS",
            "availableZoneId": "0260c681-30da-41be-a38c-7eed622bf2ae",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000077",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "127.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "network,storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-24 11:31:02",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "f19ba5b090fc4971a3cc4b1f0fb32bf8",
            "uuid": "7fad0335-9f62-411d-8403-88adf01e5b00@controller03",
            "name": "controller03",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.17.159",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "290099d3-f760-4d25-b53e-c156288371b3",
            "s8": "1e314c6a-c46f-49bb-92ea-2d398fac5820",
            "isDelete": false,
            "clusterId": "7fad0335-9f62-411d-8403-88adf01e5b00",
            "clusterName": "SDN可用区-ArcherOS",
            "availableZoneId": "0260c681-30da-41be-a38c-7eed622bf2ae",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000042",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "127.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-24 11:31:02",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "64e56db649984321a071acd48dbaca2f",
            "uuid": "b2cd8e29-40cd-42f0-bad9-ab751d71713a@controller01",
            "name": "controller01",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "78",
            "s7": "",
            "ip": "172.118.69.41",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "4ef00624-dccf-47b7-b8df-f823f50bce6b",
            "s8": "6ee4c4e1-91fa-4ca4-b427-27cd22814b20",
            "isDelete": false,
            "clusterId": "b2cd8e29-40cd-42f0-bad9-ab751d71713a",
            "clusterName": "可用区-ArcherOS152",
            "availableZoneId": "0aaf9c71-e651-48c9-989b-a6cfcb02125c",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "6Q2QF52",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "8",
            "memTotal": "63.91",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2640 v3 @ 2.60GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "network,storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-16 14:41:00",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "9f39bd32cfab438f825c3498d4c369af",
            "uuid": "b2cd8e29-40cd-42f0-bad9-ab751d71713a@controller03",
            "name": "controller03",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "54",
            "s7": "",
            "ip": "172.118.69.43",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "5db61625-051c-4dd1-a9f0-8584383e50c2",
            "s8": "6ee4c4e1-91fa-4ca4-b427-27cd22814b20",
            "isDelete": false,
            "clusterId": "b2cd8e29-40cd-42f0-bad9-ab751d71713a",
            "clusterName": "可用区-ArcherOS152",
            "availableZoneId": "0aaf9c71-e651-48c9-989b-a6cfcb02125c",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "2102311CES10G1000068",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "6",
            "memTotal": "63.97",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-16 14:41:00",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        },
        {
            "id": "f8cde112374c4121bfd4d974ac97261b",
            "uuid": "b2cd8e29-40cd-42f0-bad9-ab751d71713a@controller02",
            "name": "controller02",
            "typeName": "Host",
            "zoneId": "",
            "rackId": "",
            "sources": "同步",
            "subType": "ArcherOS",
            "s5": "START",
            "s6": "78",
            "s7": "",
            "ip": "172.118.69.42",
            "s1": "使用中",
            "s2": "节点",
            "s3": "节点",
            "s4": "9ea15efd-f31f-4ee0-9113-1457d16dc841",
            "s8": "6ee4c4e1-91fa-4ca4-b427-27cd22814b20",
            "isDelete": false,
            "clusterId": "b2cd8e29-40cd-42f0-bad9-ab751d71713a",
            "clusterName": "可用区-ArcherOS152",
            "availableZoneId": "0aaf9c71-e651-48c9-989b-a6cfcb02125c",
            "sdnId": "",
            "apiKey": "",
            "deviceId": "3DV24D2",
            "ownerId": "",
            "ownerName": "",
            "vdcId": "",
            "vdcName": "",
            "hostGroupName": "默认主机组",
            "rackName": "",
            "zoneName": "",
            "tags": [],
            "ipmiIp": "",
            "ipmiUsername": "",
            "ipmiPassword": "",
            "cpuNumber": "2",
            "cpuCore": "8",
            "memTotal": "199.78",
            "sshAddr": "",
            "physicalCpu": "Intel(R) Xeon(R) CPU E5-2640 v3 @ 2.60GHz",
            "enabled": "启用",
            "hostHaEnabled": "启用",
            "vmHaEnabled": "禁用",
            "bigPageEnable": "禁用",
            "memMonopolyVal": "50",
            "serviceRole": "network,storage,computer,controller",
            "faultDomain": "",
            "snmpVersion": "",
            "snmpCommunity": "",
            "snmpUsername": "",
            "snmpPort": "",
            "snmpRetries": "",
            "snmpTimeOut": "",
            "hostName": "",
            "brand": "",
            "modelNumber": "",
            "netIp": "",
            "netUserName": "",
            "sshPort": "",
            "vcenterUser": "",
            "vcenterIp": "",
            "vcenterPort": "",
            "createTime": "2022-05-16 14:41:00",
            "updateTime": "2022-05-31 16:24:00",
            "hostGroupId": [
                "073129a8-96b7-4295-9dc8-ffdca490"
            ],
            "diskInfo": null,
            "netInfo": null
        }
    ]`
	json.Unmarshal([]byte(data2), &res2)
	fmt.Println(res2, len(res2))

	availableZoneIds := []string{"0aaf9c71-e651-48c9-989b-a6cfcb02125c", "21942fcb-a7f3-4fe4-977f-01f0e14cfb5f"}
	out2 := []ListResp2{}

	// 初始化map
	set := make(map[string]struct{})

	// 将list内容传递进map,只根据key判断，所以不需要关心value的值，用struct{}{}表示
	for _, value := range availableZoneIds {
		set[value] = struct{}{}
	}

	// 检查元素是否在map
	for _, host := range res2 {
		if _, ok := set[host.AvailableZoneId]; ok {
			out2 = append(out2, host)
			fmt.Printf("%s is in the list \n", host.AvailableZoneId)
		} else {
			fmt.Printf("%s is not in the list \n", host.AvailableZoneId)
		}
	}

	fmt.Println("out2: ", out2, "len: ", len(out2))

}

/*
userid 6a4ce159-4394-4d13-afe3-2880ab00ae42


	user := resource.GetUser(ctx)
	h.Logger.Debug("[xxml] user: %v", user)
	userRole := user.Role
	if userRole == "VDC管理员" {
		h.Logger.Debug("[xxml] userRole is vdc admin")
		res := []*host_group.ListResp{}
		for _, host := range data {
			if host.Name == "默认主机组" {
				res = append(res, host)
			}
		}
		h.Logger.Debug("[xxml] res: ", res)

		h.ResponseSucc(ctx, map[string]interface{}{
			"data":       res,
			"totalCount": total,
			"pageNumber": pNum,
			"pageSize":   pSize,
		})
		return
	}




	type resourceCtrl struct {
	*apiBase
	todayResourceStatsRepo repo.TodayResourceStatsRepo
	logic                  resource.ResourceLogic
	identityService        identity.IdentityService
}

func NewResource(b *apiBase) *resourceCtrl {
	return &resourceCtrl{
		apiBase:                b,
		logic:                  resource.NewResourceLogic(b.DB, b.SrvBase, b.Logger),
		todayResourceStatsRepo: repo.NewTodayResourceStat(b.DB),
	}
}

func (r *resourceCtrl) getTenantavailableZoneIds(userID, userRole string) (availableZoneIds []string, err error) {
	r.Logger.Debug("[xxml] getTenantavailableZoneIds coming...")

	if userRole == "VDC管理员" {
		var vids []string
		user, std := r.identityService.GetUserInfo(&identity.GetUserInfoReq{Uid: userID})
		if std != nil && !std.IsSuccess() {
			return nil, std
		}
		if user == nil && std == nil {
			return nil, cmdbErr.ErrUnexpected
		}
		for _, v := range user.Aggregation.UserVdcs {
			vids = append(vids, v.ID)
		}
		if len(vids) == 1 {
			vdc, std := r.identityService.GetVdc(&identity.GetVdcReq{ID: vids[0]})
			if std != nil && !std.IsSuccess() {
				return nil, std
			}
			availableZoneIds = vdc.Aggregation.AvailableZoneIDs
		}
		return
	}
	return
}




	user := resource.GetUser(ctx)
	if user.Role == "VDC管理员" {
		availableZoneIds, err := r.getTenantavailableZoneIds(user.UserId, user.Role)
		r.Logger.Debug("[xxml] getTenantavailableZoneIds: %v", availableZoneIds)

		if cmdbErr.ErrCheck(err) {
			r.Logger.Error(err)
			r.ResponseFail(ctx, http.StatusInternalServerError, cmdbErr.NewError(err))
			return
		}
		out := []*models.Resource{}

		// 初始化map
		set := make(map[string]struct{})

		// 将list内容传递进map,只根据key判断，所以不需要关心value的值，用struct{}{}表示
		for _, value := range availableZoneIds {
			set[value] = struct{}{}
		}

		// 检查元素是否在map
		for _, host := range data {
			if _, ok := set[host.AvailableZoneId]; ok {
				out = append(out, host)
				fmt.Printf("%s is in the list \n", host.AvailableZoneId)
			} else {
				fmt.Printf("%s is not in the list \n", host.AvailableZoneId)
			}
		}
		r.Logger.Debug("[xxml] out: ", out)

		r.ResponseSucc(ctx, map[string]interface{}{
			"totalCount": out,
			"data":       data,
			"pageNumber": pNum,
			"pageSize":   pSize,
		})
		return
	}

*/
