package dto



type PermissionCategory struct {
    ServiceName   string ` json:"ServiceName" `
    ServiceId   string ` json:"ServiceId" `
    Permissions   []Permission ` json:"Permissions" `
    }

