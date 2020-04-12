package gobizfly

// ServerSecurityGroup contains information of security group of a server.
type ServerSecurityGroup struct {
	Name string `json:"name"`
}

// AttachedVolume contains attached volumes of a server.
type AttachedVolume struct {
	ID string `json:"id"`
}

// Server contains server information.
type Server struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	KeyName         string                 `json:"key_name"`
	UserID          string                 `json:"user_id"`
	ProjectID       string                 `json:"tenant_id"`
	CreatedAt       string                 `json:"created"`
	UpdatedAt       string                 `json:"updated"`
	Status          string                 `json:"status"`
	IPv6            bool                   `json:"ipv6"`
	SecurityGroup   []ServerSecurityGroup  `json:"security_group"`
	Addresses       map[string]interface{} `json:"addresses"`
	Metadata        map[string]string      `json:"metadata"`
	Flavor          map[string]interface{} `json:"flavor"`
	Progress        int                    `json:"progress"`
	AttachedVolumes []AttachedVolume       `json:"os-extended-volumes:volumes_attached"`
}
