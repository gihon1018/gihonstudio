package repository

import (
	"context"
	"os"
	"testing"

	"gihonstudio/internal/config"
	"gihonstudio/internal/database"
	. "gihonstudio/internal/model"

	"gorm.io/gorm"
)

// 全局预初始化
var (
	globalDB     *gorm.DB
	globalCtx    context.Context
	userAuthRepo *UserAuthRepository
)

func TestMain(m *testing.M) {
	// 加载测试专属配置
	cfg, err := config.LoadConfigFromPath("../../config/database_test.yaml")
	if err != nil {
		panic("加载测试配置失败：" + err.Error())
	}

	// 初始化测试PG连接
	globalDB, err = database.InitPostgreSQL(&cfg.Database)
	if err != nil {
		panic("测试数据库连接失败：" + err.Error())
	}

	// 初始化全局上下文
	globalCtx = context.Background()
	// 初始化仓库
	userAuthRepo = NewUserAuthRepository(globalDB)

	// 执行所有测试
	exitCode := m.Run()

	// 测试全部跑完关闭连接
	sqlDB, _ := globalDB.DB()
	_ = sqlDB.Close()

	os.Exit(exitCode)
}

// 测试创建
func TestUserAuthRepository_Create(t *testing.T) {
	userAuth := &UserAuth{
		Username:     "test_user_002",
		PasswordHash: "123abc_hash",
	}

	err := userAuthRepo.Create(globalCtx, userAuth)
	if err != nil {
		t.Fatalf("创建失败：%v", err)
	}
	t.Logf("创建成功，UID：%d", userAuth.Uid)
}
