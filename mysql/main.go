// package main

// import (
//     "database/sql"
//     "fmt"
//     _ "github.com/go-sql-driver/mysql"
// )

// func main() {
//     dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
//     db, err := sql.Open("mysql", dsn)
//     if err != nil {
//         panic(err)
//     }
//     defer db.Close()

//     err = db.Ping()
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println("连接 MySQL 成功！")
// }


// package main

// import (
//     "database/sql"
//     "fmt"
//     "time"

//     _ "github.com/go-sql-driver/mysql"
// )

// type User struct {
//     ID        int
//     Username  string
//     Password  string
//     Email     string
//     CreatedAt time.Time
//     UpdatedAt time.Time
//     DeletedAt sql.NullTime
// }

// func main() {
//     dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
//     db, err := sql.Open("mysql", dsn)
//     if err != nil {
//         panic(err)
//     }
//     defer db.Close()

//     if err := db.Ping(); err != nil {
//         panic(err)
//     }
//     fmt.Println("✅ 数据库连接成功")

//     // ------------------ 创建用户 ------------------
//     username := "alice"
//     password := "123456"
//     email := "alice@example.com"

//     res, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, password, email)
//     if err != nil {
//         panic(err)
//     }
//     id, _ := res.LastInsertId()
//     fmt.Println("创建用户ID:", id)

//     // ------------------ 查询用户 ------------------
//     var u User
//     err = db.QueryRow("SELECT id, username, password, email, created_at, updated_at, deleted_at FROM users WHERE username=? AND deleted_at IS NULL", "alice").
//         Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("查询用户:", u.Username, u.Email)

//     // ------------------ 更新用户 ------------------
//     newEmail := "alice@newmail.com"
//     _, err = db.Exec("UPDATE users SET email=?, updated_at=NOW() WHERE id=? AND deleted_at IS NULL", newEmail, u.ID)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("更新用户邮箱:", newEmail)

//     // ------------------ 逻辑删除 ------------------
//     _, err = db.Exec("UPDATE users SET deleted_at=NOW() WHERE id=?", u.ID)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("已软删除用户:", u.Username)

//     // ------------------ 查询未删除用户 ------------------
//     rows, err := db.Query("SELECT id, username FROM users WHERE deleted_at IS NULL")
//     if err != nil {
//         panic(err)
//     }
//     defer rows.Close()

//     fmt.Println("未删除用户列表:")
//     for rows.Next() {
//         var id int
//         var name string
//         rows.Scan(&id, &name)
//         fmt.Println(id, name)
//     }

//     // ------------------ 查询包含软删除用户 ------------------
//     rowsAll, err := db.Query("SELECT id, username, deleted_at FROM users")
//     if err != nil {
//         panic(err)
//     }
//     defer rowsAll.Close()

//     fmt.Println("所有用户列表(包含已软删除):")
//     for rowsAll.Next() {
//         var id int
//         var name string
//         var deletedAt sql.NullTime
//         rowsAll.Scan(&id, &name, &deletedAt)
//         fmt.Printf("%d %s deleted_at: %v\n", id, name, deletedAt)
//     }
// }

// package main

// import (
//     "fmt"
//     "time"

//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
//     "gorm.io/gorm/schema"
// )

// // User 对应 users 表，使用 DeletedAt 实现软删除
// type User struct {
//     ID        uint           `gorm:"primaryKey;autoIncrement;comment:用户ID，自增主键"`
//     Username  string         `gorm:"size:50;not null;unique;comment:用户名，唯一标识用户"`
//     Password  string         `gorm:"size:255;not null;comment:用户密码（加密存储）"`
//     Email     string         `gorm:"size:100;comment:用户邮箱，用于联系或找回密码"`
//     CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间，默认为当前时间"`
//     UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间，每次修改自动更新"`
//     DeletedAt gorm.DeletedAt `gorm:"index;comment:逻辑删除标志，软删除"`
// }

// func main() {
//     // 数据库连接
//     dsn := "root:123456@tcp(127.0.0.1:3306)/test2db?charset=utf8mb4&parseTime=True&loc=Local"
//     db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//         NamingStrategy: schema.NamingStrategy{
//             SingularTable: true, // 表名使用单数
//         },
//     })
//     if err != nil {
//         panic("数据库连接失败: " + err.Error())
//     }
//     fmt.Println("✅ 数据库连接成功")

//     // 自动迁移
//     err = db.AutoMigrate(&User{})
//     if err != nil {
//         panic("自动迁移失败: " + err.Error())
//     }
//     fmt.Println("✅ users 表自动创建或更新完成")

//     // ------------------ CRUD 演示 ------------------

//     // 创建
//     user := User{Username: "alice", Password: "123456", Email: "alice@example.com"}
//     db.Create(&user)
//     fmt.Println("创建用户ID:", user.ID)

//     // 查询
//     var u User
//     db.First(&u, "username = ?", "alice")
//     fmt.Println("查询用户:", u.Username, u.Email)

//     // 更新
//     db.Model(&u).Update("Email", "alice@newmail.com")
//     fmt.Println("更新用户邮箱:", u.Email)

//     // 逻辑删除
//     db.Delete(&u)
//     fmt.Println("已软删除用户:", u.Username)

//     // 查询未删除用户
//     var users []User
//     db.Find(&users)
//     fmt.Println("未删除用户数量:", len(users)) // 不包含已软删除用户

//     // 查询包括软删除用户
//     var allUsers []User
//     db.Unscoped().Find(&allUsers)
//     fmt.Println("包含软删除用户数量:", len(allUsers))
// }




// package main

// import (
//     "fmt"
//     "time"

//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
//     "gorm.io/gorm/schema"
// )

// // User 对应 users 表，使用 DeletedAt 实现软删除
// type User struct {
//     ID        uint           `gorm:"primaryKey;autoIncrement;comment:用户ID，自增主键"`
//     Username  string         `gorm:"size:50;not null;unique;comment:用户名，唯一标识用户"`
//     Password  string         `gorm:"size:255;not null;comment:用户密码（加密存储）"`
//     Email     string         `gorm:"size:100;comment:用户邮箱，用于联系或找回密码"`
//     CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间"`
//     UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
//     DeletedAt gorm.DeletedAt `gorm:"index;comment:逻辑删除标志，软删除"`

//     Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 一对一关系
// }

// // Profile 对应 user_profiles 表
// type Profile struct {
//     ID        uint      `gorm:"primaryKey;autoIncrement"`
//     UserID    uint      `gorm:"uniqueIndex;comment:关联用户ID"` // 外键
//     FullName  string    `gorm:"size:100;comment:用户真实姓名"`
//     Bio       string    `gorm:"type:text;comment:用户简介"`
//     Avatar    string    `gorm:"size:255;comment:用户头像URL"`
//     Phone     string    `gorm:"size:20;comment:手机号"`
//     WeChat    string    `gorm:"size:50;comment:微信号"`
//     CreatedAt time.Time `gorm:"autoCreateTime"`
//     UpdatedAt time.Time `gorm:"autoUpdateTime"`
// }

// func main() {
//     // 连接 MySQL 数据库
//     dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
//     db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//         NamingStrategy: schema.NamingStrategy{
//             SingularTable: true, // 表名单数
//         },
//     })
//     if err != nil {
//         panic("数据库连接失败: " + err.Error())
//     }
//     fmt.Println("✅ 数据库连接成功")

//     // 自动迁移 User 和 Profile
//     err = db.AutoMigrate(&User{}, &Profile{})
//     if err != nil {
//         panic("自动迁移失败: " + err.Error())
//     }
//     fmt.Println("✅ users 和 user_profiles 表自动创建或更新完成")

//     // ------------------ CRUD 演示 ------------------

//     // 创建用户及资料
//     user := User{
//         Username: "alice",
//         Password: "123456",
//         Email:    "alice@example.com",
//         Profile: Profile{
//             FullName: "Alice Zhang",
//             Bio:      "Go & MySQL Developer",
//             Avatar:   "https://example.com/avatar/alice.jpg",
//             Phone:    "13800138000",
//             WeChat:   "alicewechat",
//         },
//     }
//     db.Create(&user)
//     fmt.Println("创建用户ID:", user.ID)
//     fmt.Println("创建用户资料ID:", user.Profile.ID)

//     // 查询用户及资料（预加载 Profile）
//     var u User
//     db.Preload("Profile").First(&u, "username = ?", "alice")
//     fmt.Println("查询用户:", u.Username, u.Email)
//     fmt.Println("用户资料:", u.Profile.FullName, u.Profile.Bio, u.Profile.Avatar, u.Profile.Phone, u.Profile.WeChat)

//     // 更新用户资料
//     db.Model(&u.Profile).Updates(Profile{
//         Bio:    "Senior Go Developer",
//         Avatar: "https://example.com/avatar/alice_new.jpg",
//     })
//     fmt.Println("更新用户资料 Bio & Avatar:", u.Profile.Bio, u.Profile.Avatar)

//     // 逻辑删除用户
//     db.Delete(&u)
//     fmt.Println("已软删除用户:", u.Username)

//     // 查询未删除用户及资料
//     var users []User
//     db.Preload("Profile").Find(&users)
//     fmt.Println("未删除用户数量:", len(users))

//     // 查询包括软删除用户
//     var allUsers []User
//     db.Unscoped().Preload("Profile").Find(&allUsers)
//     fmt.Println("包含软删除用户数量:", len(allUsers))
// }

package main

import (
    "fmt"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

// User 对应 users 表，使用 DeletedAt 实现软删除
type User struct {
    ID        uint           `gorm:"primaryKey;autoIncrement;comment:用户ID，自增主键"`
    Username  string         `gorm:"size:50;not null;unique;comment:用户名，唯一标识用户"`
    Password  string         `gorm:"size:255;not null;comment:用户密码（加密存储）"`
    Email     string         `gorm:"size:100;comment:用户邮箱，用于联系或找回密码"`
    CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
    DeletedAt gorm.DeletedAt `gorm:"index;comment:逻辑删除标志，软删除"`

    Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 一对一关系
    Posts   []Post  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 一对多关系
}

// Profile 对应 user_profiles 表
type Profile struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"`
    UserID    uint      `gorm:"uniqueIndex;comment:关联用户ID"` // 外键
    FullName  string    `gorm:"size:100;comment:用户真实姓名"`
    Bio       string    `gorm:"type:text;comment:用户简介"`
    Avatar    string    `gorm:"size:255;comment:用户头像URL"`
    Phone     string    `gorm:"size:20;comment:手机号"`
    WeChat    string    `gorm:"size:50;comment:微信号"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Post 对应 posts 表，表示用户的帖子
type Post struct {
    ID        uint      `gorm:"primaryKey;autoIncrement;comment:帖子ID"`
    UserID    uint      `gorm:"index;comment:关联用户ID"`  // 外键
    Title     string    `gorm:"size:255;not null;comment:帖子标题"`
    Content   string    `gorm:"type:text;comment:帖子内容"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
    // 连接 MySQL 数据库
    dsn := "root:123456@tcp(127.0.0.1:3306)/test3db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true, // 表名单数
        },
    })
    if err != nil {
        panic("数据库连接失败: " + err.Error())
    }
    fmt.Println("✅ 数据库连接成功")

    // 自动迁移 User, Profile 和 Post
    err = db.AutoMigrate(&User{}, &Profile{}, &Post{})
    if err != nil {
        panic("自动迁移失败: " + err.Error())
    }
    fmt.Println("✅ users、user_profiles 和 posts 表自动创建或更新完成")

    // 创建用户及多个帖子
    user := User{
        Username: "bob",
        Password: "123456",
        Email:    "bob@example.com",
        Profile: Profile{
            FullName: "Bob Smith",
            Bio:      "Tech Enthusiast",
            Avatar:   "https://example.com/avatar/bob.jpg",
            Phone:    "13800138001",
            WeChat:   "bobwechat",
        },
        Posts: []Post{
            {Title: "Go Programming", Content: "Go is an open-source programming language."},
            {Title: "MySQL Tutorial", Content: "MySQL is an open-source relational database."},
        },
    }
    db.Create(&user)
    fmt.Println("创建用户ID:", user.ID)

    // 查询用户及帖子（预加载 Profile 和 Posts）
    var u User
    db.Preload("Profile").Preload("Posts").First(&u, "username = ?", "bob")
    fmt.Println("查询用户:", u.Username, u.Email)
    fmt.Println("用户资料:", u.Profile.FullName, u.Profile.Bio, u.Profile.Avatar, u.Profile.Phone, u.Profile.WeChat)
    fmt.Println("用户帖子数量:", len(u.Posts))

    // 查询某个用户的所有帖子
    var posts []Post
    db.Where("user_id = ?", u.ID).Find(&posts)
    fmt.Println("用户的帖子数量:", len(posts))

    // 更新用户的帖子
    db.Model(&u.Posts[0]).Update("Content", "Go is an efficient, statically typed language.")
    fmt.Println("更新后的帖子内容:", u.Posts[0].Content)

    // 逻辑删除帖子
    db.Delete(&u.Posts[0])
    fmt.Println("已软删除帖子:", u.Posts[0].Title)

    // 查询包括软删除的帖子
    var allPosts []Post
    db.Unscoped().Where("user_id = ?", u.ID).Find(&allPosts)
    fmt.Println("包括软删除的帖子数量:", len(allPosts))
}


