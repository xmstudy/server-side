package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// 创建模型
/*
1. id => 主键
2. 模型名字 => 数据库中 为 小写+s
3. `gorm:"primarykey"` 来标识主键
4. 如果使用Time Uint64 不配置默认时间 可以使用 `gorm:"autoUpdateTime:milli"` | `gorm:"autoUpdateTime:nano"` | `gorm:"autoCreateTime"`
5. 嵌入 embedded =>  Model{ Student Student gorm:embedded；embeddedPrefix:前缀} => 就可以在Model中放置Student 的内容 和 加上指定前缀 author_  （Author）Name
6. <-:create 只创建 <-:update 只更新  , <- 更新和创建
*/
type Model struct {
}

var (
	dsn string
)

func readContent(config *viper.Viper) {
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件...")
		} else {
			fmt.Printf("解析配置文件出错: %v\n", err)
		}
	}
	// 读取配置文件
	dsn = config.GetString("mysql.dsn")
	fmt.Println(dsn)
}
func readYaml() {
	config := viper.New()
	config.AddConfigPath("config")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	readContent(config)
}

func main() {
	readYaml()
	// 连接数据库
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db.AutoMigrate(&Product{})

	// expr 可以使用原生 sql
	// 使用原生 raw
	// db.Raw("select id from users where name = ?",1).Scan(&p)  scan 扫描到p 中

	// before | after Create  钩子 函数
	//  beforeUpdate    Hook
	// beforeDelete

	/*
		func create(db *gorm.DB){}
	*/
	// 创建
	/*
		p := Product{
			Code:  "001",
			Price: 200,
		}
		db.Create(&p)
	*/
	/*----------------------*/
	// Find
	/*
			1. primary Key
				var p Product
				db.First(&p, 1)
			2. 条件
				db.First(&p,"code = ?","001")

				 查
		 result :=db.First()
		 db.Last()
		 db.Take()  limit 1 一样 没有指定主键 就是  1
		 找不到 == errors.Is(result.Error,form.ErrRecordNotFound)

		 db.Where(id == 1  | LIKE %gsx%  | AND | BETWEEN).Find(&p)
		 db.Find(&p,"name = ?","nihao" | User(Age:20 )
		 NOT |  OR
		 db.Not().Find()
		 db.Where().Or()
	*/
	/*------------------*/
	//  update
	/*
			db.First()
			db.Model(&p).Update( "price" , 100 (单个) | Product{"Code":"002","Price":500} | map[string]interface{}{"price":1000,"Code":"1002"} (多个) )


		 update
		 1. db.first(&p);
		 2. p.name = '123'
		 3. db.save(&p)
		 update one
		  db.Model(&p).Where("name = 123").Update("name" : '66')
		 updateMore
		  db.Model(&p.Update(map[string]interface{}{name:'',age:'',"xx":''})

	*/
	/*--------------*/
	/*
	 del 逻辑删除
	 var p Product
	 db.First(&p,1)
	 db.Delete(&p,1)

	 db.Where().Delete()
	*/

	// 关联关系
	// 多对一
	// 多对多
	// 一对一

	// belong to
	// 在 struct  中 直接写 例如  Componry  Componey 就是 外键 和上面 的 en 不同
	// type A struct {
	// 	id   uint
	// 	name string
	// 	B    B
	// 	BID  int 外键 需要 但是 尽量不要使用 外键 会影响效率
	// 	B    B gorm:foreignKey:CompanyRefer 重写 key
	// CompanyRefer  int 为 重写的外键id

	// }
	// type B struct {
	// 	id uint
	// }
	// db.AutoMigrate(&A{},&B{})

	// has one  一对一
	// type A struct {
	// gorm.Model
	// Edit   Edit            gorm:foreignKey;UserName  |
	// }
	// type Edit struct {                               |
	// gorm.Model																			  |
	// 	AId uint //如果是子结构 + ID  会自动创建外键  变为 UserName 就是自动创建——|
	// }

	// 一对多
	// 	struct Person {
	// 		CreditCards []CreditCards
	// 	}
	// type CreditCards struct{
	// 	PersonId int
	// 	Number int
	// }

	// 多对多
	// `gorm:"many2many:user_language(中间表名)"`

	// 自动更新引用关系
	// 默认 createInsert
	// 更新 upadte 需要添加当前更新 FullSaveAssociations:true db.Session(&gorm.Session(FullSaveAssociations:true)).Updates(&user)
	// 关联  db.Model(&user).Association("关联字段").Find(&language)
	// 追加关联  db.Model(&user).Association("关联字段").Append([]languages | &languages)
	// 多态 gorm:polymorphic:Owner

	// 事物 例子: 银行存取 问题 加锁 解锁
	// 1. 事物级别 SkipDefaultTransaction: true,
	// 2. 会话 db.Session(&gorm.Session(SkipDefaultTransaction: true))
	// 手动开启事物
	// tx := db.Begin() 开始
	// tx.Create(&p)
	// err = tx.Create(nil).Error
	// if err != nil {
	// 	tx.Rollback() 回滚
	// }else{
	// tx.Commit()
	// }
	// 自动控制事物
	// db.Transaction(func(tx *gorm.DB) error {
	// err :=tx.Create({}).Error;err != nil {
	// return nil
	// }
	// err :=tx.Create(nil).Error;err != nil {
	// return nil
	// }
	// return nil
	// })
}
