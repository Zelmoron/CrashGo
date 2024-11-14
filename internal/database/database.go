package database

import (
	"CaseGo/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

// type MarketPriceResponse struct {
// 	Success     bool   `json:"success"`
// 	LowestPrice string `json:"lowest_price"`
// 	Volume      string `json:"volume"`
// 	MedianPrice string `json:"median_price"`
// }

func New() *Database {
	return &Database{}
}

// func getSteamMarketPrices(appID int, marketHashName string, attempts int) ([]string, error) {
// 	marketHashName = url.QueryEscape(marketHashName)
// 	requestURL := fmt.Sprintf("https://steamcommunity.com/market/priceoverview/?appid=%d&currency=1&market_hash_name=%s", appID, marketHashName)

// 	var prices []string

// 	for i := 0; i < attempts; i++ {
// 		resp, err := http.Get(requestURL)
// 		if err != nil {
// 			fmt.Printf("Request %d failed: %v\n", i+1, err)
// 			continue
// 		}

// 		body, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()

// 		if err != nil {
// 			fmt.Printf("Reading response %d failed: %v\n", i+1, err)
// 			continue
// 		}

// 		var priceResponse MarketPriceResponse
// 		if err := json.Unmarshal(body, &priceResponse); err != nil {
// 			fmt.Printf("Parsing response %d failed: %v\n", i+1, err)
// 			continue
// 		}

// 		if !priceResponse.Success {
// 			fmt.Printf("Request %d was not successful\n", i+1)
// 			continue
// 		}

// 		if priceResponse.LowestPrice != "" {
// 			prices = append(prices, priceResponse.LowestPrice)
// 		}

// 		// Добавляем задержку между запросами
// 		time.Sleep(300 * time.Millisecond)
// 	}

// 	return prices, nil
// }

// func (d *Database) API(item string) {

// 	appID := 730 // CS:GO
// 	itemName := item
// 	attempts := 1

// 	prices, err := getSteamMarketPrices(appID, itemName, attempts)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	if len(prices) == 0 {
// 		fmt.Println("No prices found!")
// 		return
// 	}

// 	for _, price := range prices {
// 		fmt.Printf(price[1:])
// 	}

// }

func (d *Database) CreateTables() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Irkutsk",
		host, user, password, dbname, port)

	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error: ", err)
		panic("Не удалось подключиться к базе данных")

	}

	if err := d.db.AutoMigrate(&models.UserModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}

	if err := d.db.AutoMigrate(&models.InventoryModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	if err := d.db.AutoMigrate(&models.ItemModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	if err := d.db.AutoMigrate(&models.CasesModel{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		panic("Fatal error - dont create databases")
	}
	// Сначала создаем все кейсы
	var case1, case2, case3 models.CasesModel

	// Проверяем, существует ли уже кейс "Решающий момент"
	result := d.db.Where("name = ?", "Решающий момент").First(&case1)
	if result.Error != nil {
		// Если нет, создаем новый
		case1 = models.CasesModel{Name: "Решающий момент", Image: "https://qliquiz.github.io/CaSeGO-front/images/cases/decisive_moment.png"}
		d.db.Create(&case1)

		// Теперь создаем элементы, связанные с кейсами
		items := []models.ItemModel{
			{WeaponName: "PP-Bizon", SkinName: "Ночной бунт", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpotLO_JAlf0Ob3czRY49KJhomEg8j4OrzZgiUD7pUp3rHDp9v00QXj-UtrY2_xJ4aTJAI3aV_QqQe3lL3vg8Tu7s-c1zI97Wr-owub", CaseID: case1.ID},
			{WeaponName: "SG 553", SkinName: "Алоха", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopb3wflFfwOP3YjoXv4-JlYyEn_bLP7LWnn8fuJZwi7GXptqt2FW2-UFuYGDxINfAe1VsNFCC_Ve4w7_ngcDuvZvLmmwj5Hc11_x0mg", CaseID: case1.ID},
			{WeaponName: "XM1014", SkinName: "Оксидное пламя", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgporrf0e1Y07PDdTiVPvYznwL-Ej_7wNoTTmmpL7fp9g-7J4cKj0QW2rktsMW7zItOUJ1c6NwmG8wO7kue90MW4vM_Kz3Ni6CF24XvdgVXp1gKDw8ad", CaseID: case1.ID},
			{WeaponName: "Five-SeveN", SkinName: "Испытание огнём", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgposLOzLhRlxfbGTjpR09q_goWYkuHxPYTZmX9u-sp1tf_I-oDwnGu4ohQ0J3f1ItXHcVI4YlvWrFXrkO7o1JHquMibmyZguykgtnrUyRXm10sdbbM8m7XAHrBtyPY3", CaseID: case1.ID},
			{WeaponName: "P2000", SkinName: "Городская опасность", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovrG1eVcwg8zLZAJSvozmxL-ehfX1PYTZl3FQ-sFOhuDG_Zi72QPi_kQ_Zzz6d4WWdQ9oZ1vUqVa2lOq7hZTv7ZScwCQy6XJ37CqJzQv3309hpG0-UA", CaseID: case1.ID},
			{WeaponName: "MP9", SkinName: "Чёрный песок", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpou6r8FAR17P7YKAJA4N21n7-YlvnwDLfYkWNFppYmjurEpdTz3ATnrhBrYDrycYeTIAVqMlzW-le2k-zth5-_6ZTMyHZ9-n51wvY0MQA", CaseID: case1.ID},
			{WeaponName: "Револьвер R8", SkinName: "Хватка", Type: "milspec", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpopL-zJAt21uH3di59_tmgm4ydkuXLJ6nUl29u5Mx2gv2Poo-milDl-ENuNW_xLIOWJwM4aFyBrwK8lenv1sC975rIzXIxuXZx5WGdwUIffS2-og", CaseID: case1.ID},
			{WeaponName: "Negev", SkinName: "Рыба-лев", Type: "restricted", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpouL-iLhFfwOP3fzhF6cqJmImEmfH9ILPummJW4NE_jLGSp9r03gPi-kQ_ZmjwLNfHelQ4N16BrAK2wO3ogMDu6J3AyHJguT5iuyjn_404lQ", CaseID: case1.ID},
			{WeaponName: "UMP-45", SkinName: "Арктический волк", Type: "restricted", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpoo7e1f1Jf0Ob3ZDBSuImJg4iCg_LLNbrfkVRd4cJ5nqfHo9_02QSw_hY5YmGmJ4aRd1dqNwyDrFi4wrzmhZC77p6bySNl6CQq-z-DyN1nM0Y5", CaseID: case1.ID},
			{WeaponName: "Nova", SkinName: "Nova | Дикая шестёрка", Type: "restricted", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpouLWzKjhnwMzGfitD0924l4iEhf7gNoTdn2xZ_Isl27DDrdqsigHn-kU-ZW6iItPBdAE9NAyBrAW-yea-jMK9us7Jz3QwpGB8ssgBFjCg", CaseID: case1.ID},
			{WeaponName: "Glock-18", SkinName: "Лунная ночь", Type: "restricted", Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgposbaqKAxf0vL3djFN79eJxdi0guX2MrXum2Re5vp3j__E57P5gVO8v109Y2vxI9Cdc1M6ZQyCq1e2kLy90JO1ucnNy3U3vCJ07CnUn0HmiBEYcKUx0m01ug-1", CaseID: case1.ID},
			// {WeaponName: "", SkinName: "", Type: "", Image: "", CaseID: case1.ID},
			// {WeaponName: "", SkinName: "", Type: "", Image: "", CaseID: case1.ID},
			// {WeaponName: "", SkinName: "", Type: "", Image: "", CaseID: case1.ID},
			// {WeaponName: "", SkinName: "", Type: "", Image: "", CaseID: case1.ID},
			// {WeaponName: "", SkinName: "", Type: "", Image: "", CaseID: case1.ID},
		}

		for _, item := range items {

			// d.API(item.WeaponName + " | " + item.SkinName + " (Minimal Wear)")
			d.db.Create(&item)
		}

	}

	// Проверяем, существует ли уже кейс "Гидра"
	result = d.db.Where("name = ?", "Гидра").First(&case2)
	if result.Error != nil {
		// Если нет, создаем новый
		case2 = models.CasesModel{Name: "Гидра", Image: "https://qliquiz.github.io/CaSeGO-front/images/cases/hydra.png"}
		d.db.Create(&case2)
	}

	// Проверяем, существует ли уже кейс "Фальшион"
	result = d.db.Where("name = ?", "Фальшион").First(&case3)
	if result.Error != nil {
		// Если нет, создаем новый
		case3 = models.CasesModel{Name: "Фальшион", Image: "https://qliquiz.github.io/CaSeGO-front/images/cases/falchion.png"}
		d.db.Create(&case3)
	}

	log.Println("Таблицы успешно созданы")

}

func (d *Database) InsertUser(id int, name string) models.UserModel {
	var user models.UserModel
	if err := d.db.Where("telegram_id=?", id).First(&user).Error; err == nil {
		log.Println("Уже есть")
		return user
	}

	user = models.UserModel{
		Name:       name,
		TelegramID: id,
		Coins:      100,
	}
	if err := d.db.Create(&user).Error; err != nil {
		return user
	}
	log.Println("Новый")

	return user

}

func (d *Database) SelectUser(id int) models.UserModel {
	var user models.UserModel

	if err := d.db.Where("telegram_id=?", id).First(&user).Error; err != nil {
		return user
	}

	return user
}

func (d *Database) SelectInventory(id int) []models.InventoryModel {
	var inventory []models.InventoryModel

	if err := d.db.Where("telegram_id=?", id).Find(&inventory).Error; err != nil {
		return []models.InventoryModel{}
	}

	return inventory
}

func (d *Database) SelectCases() []models.CasesModel {
	var cases []models.CasesModel

	if err := d.db.Find(&cases).Error; err != nil {
		return []models.CasesModel{}
	}

	return cases

}

func (d *Database) SelectWeapons(id int) []models.ItemModel {
	var cases models.CasesModel

	if err := d.db.Preload("Items").First(&cases, id).Error; err != nil {
		fmt.Println(1)
		return []models.ItemModel{}
	}
	weaponsData := cases.Items

	return weaponsData

}

func (d *Database) InsertInventory(userId int, itemId int) models.ItemModel {
	var item models.ItemModel

	if err := d.db.First(&item, itemId).Error; err != nil {
		return models.ItemModel{}
	}

	inventory := models.InventoryModel{
		SkinId:     item.ID,
		WeaponName: item.WeaponName,
		SkinName:   item.SkinName,
		Image:      item.Image,
		TelegramID: uint(userId),
		Type:       item.Type,
	}
	if err := d.db.Create(&inventory).Error; err != nil {
		return models.ItemModel{}
	}

	return item
}
