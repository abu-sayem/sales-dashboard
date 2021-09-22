package controllers

import (
	"salesdashboard/database"
	"salesdashboard/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
		var roles []models.Role
		database.DB.Find(&roles)
		return c.JSON(roles)
	}


func CreatRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil{
		return err
	}

	list := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name: roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GeteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role {
		Id: uint(id),
	}
	database.DB.Preload("permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map
	if err := c.BodyParser(&roleDTO); err != nil{
		return err
	}

	list := roleDTO["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id: uint(id),
		Name: roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role {
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}