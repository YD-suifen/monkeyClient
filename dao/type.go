package dao

type SHCpuTable struct {
	HostName string `json:"hostName"`
	PrivateIP string `json:"privateIp"`
	UsedCpu float64 `json:"usedCpu"`
	IdleCpu float64 `json:"idleCpu"`
	TimeUnix int64 `json:"timeUnix"`
}

type SHMemTable struct {
	HostName string `json:"hostName"`
	PrivateIP string `json:"privateIp"`
	Total float64 `json:"total"`
	Used float64 `json:"used"`
	Free float64 `json:"free"`
	TimeUnix int64 `json:"timeUnix"`
}

type SHTcpNetTable struct {
	HostName string `json:"hostName"`
	PrivateIP string `json:"privateIp"`
	AllConn int `json:"allConn"`
	Established int `json:"established"`
	TimeUnix int64 `json:"timeUnix"`
}

type SHDiskTable struct {
	HostName string `json:"hostName"`
	PrivateIP string `json:"privateIp"`
	Disk string `json:"disk"`
	TimeUnix int64 `json:"timeUnix"`
}


//CREATE TABLE `monkey_s_tcpnetdata` (
//`id` int(100) NOT NULL AUTO_INCREMENT,
//`hostName` varchar(100) NOT NULL,
//`privateIp` varchar(100) NOT NULL,
//`allConn` int DEFAULT '0',
//`established` int DEFAULT '0',
//`timeUnix` int(100) NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

//CREATE TABLE `monkey_s_cpudata` (
//`id` int(100) NOT NULL AUTO_INCREMENT,
//`hostName` varchar(100) NOT NULL,
//`privateIp` varchar(100) NOT NULL,
//`usedCpu` float DEFAULT '0',
//`idleCpu` float DEFAULT '0',
//`timeUnix` int(100) NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

//CREATE TABLE `monkey_s_memdata` (
//`id` int(100) NOT NULL AUTO_INCREMENT,
//`hostName` varchar(100) NOT NULL,
//`privateIp` varchar(100) NOT NULL,
//`total` float DEFAULT '0',
//`used` float DEFAULT '0',
//`free` float DEFAULT '0',
//`timeUnix` int(100) NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

//CREATE TABLE `monkey_s_diskdata` (
//`id` int(100) NOT NULL AUTO_INCREMENT,
//`hostName` varchar(100) NOT NULL,
//`privateIp` varchar(100) NOT NULL,
//`disk` varchar(500) DEFAULT '',
//`timeUnix` int(100) NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;