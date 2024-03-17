package conf

type ReplicationConf struct {
	Id     string
	Offset int
	Role   string
}

var Replication *ReplicationConf = &ReplicationConf{}
