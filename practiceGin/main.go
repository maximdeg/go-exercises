package main

import "github.com/gin-gonic/gin"

var ninjaWeapons = map[string]string{
	"ninjaStar": "Beginner Ninja Star - Damage 1",
}

func DeleteWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName, ok := ninjaWeapons[weaponType]
	if !ok {
		c.JSON(404, gin.H{
			weaponType: "",
		})
	}
	return

	delete(ninjaWeapons, weaponType)
	c.JSON(200, gin.H{
		weaponType: weaponName,
	})
}

func GetWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName, ok := ninjaWeapons[weaponType]
	if !ok {
		c.JSON(404, gin.H{
			weaponType: "",
		})
		return
	}
	c.JSON(200, gin.H{
		weaponType: weaponName,
	})
}

func PostWeapon(c *gin.Context) {
	weaponType := c.Query("type")
	weaponName := c.Query("name")

	if len(weaponType) == 0 || len(weaponName) == 0 {
		c.JSON(400, gin.H{
			//"message": "Weapon already exists",
			weaponType: weaponName,
		})
		return
	}

	if _, ok := ninjaWeapons[weaponType]; ok {
		c.JSON(409, gin.H{
			"message": "Weapon already exists",
			//weaponType: weaponName,
		})
		return
	}

	ninjaWeapons[weaponType] = weaponName
	c.JSON(201, gin.H{
		weaponType: weaponName,
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	rGroup := r.Group("/weapon")
	rGroup.GET("", GetWeapon)
	rGroup.POST("", PostWeapon)
	rGroup.DELETE("", DeleteWeapon)

	r.RUN()
}
