package dto



type User struct {
    UserId   string ` json:"UserId" `
    FirstName   string ` json:"FirstName" `
    LastName   string ` json:"LastName" `
    Email   string ` json:"Email" `
    RoleId   string ` json:"RoleId" `
    RoleName   string ` json:"RoleName" `
    Deleted bool `json:"deleted"`}

