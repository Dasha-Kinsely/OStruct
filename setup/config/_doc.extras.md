## test whether the dbs are connected
base.GET("/test", func(c *gin.Context){
			clientR := config.GetRedis()
			pong, err := clientR.Ping(c).Result()
			fmt.Println(pong, err)
			clientM, _ := config.GetMongoClient().ListDatabaseNames(c, bson.D{{}})
			fmt.Printf("mongo db names are %s\n", clientM)
			clientS, _ := config.GetSqlDB().DB()
			fmt.Printf("mysql db names are %s\n", clientS)
			c.JSON(200, gin.H{
				"message": "all good Saul",
			})
		})
