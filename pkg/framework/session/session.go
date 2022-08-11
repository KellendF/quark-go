package session

import (
	fiber "github.com/gofiber/fiber/v2"
	session "github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

// 设置值
func Set(c *fiber.Ctx, key string, value interface{}) error {

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set(key, value)

	return sess.Save()
}

// 获取值
func Get(c *fiber.Ctx, key string) interface{} {

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	result := sess.Get(key)

	return result
}
