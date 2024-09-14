package dto



type Role struct {
    RoleId   string ` json:"RoleId" `
    Name   string ` json:"Name" `
    Description   string ` json:"Description" `
    PermissionCategories   []PermissionCategory ` json:"PermissionCategories" `
    Deleted bool `json:"deleted"`}

