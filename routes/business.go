package routes

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/62teknologi-test-alfatah/models"
	"github.com/62teknologi-test-alfatah/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	go MainRoute.NewRoute("POST", "/business", CreateBusiness)
	go MainRoute.NewRoute("PUT", "/business/:id", UpdateBusiness)
	go MainRoute.NewRoute("DELETE", "/business/:id", DeletedBusiness)
	go MainRoute.NewRoute("GET", "business/search", BusinessSearch)
}

func CreateBusiness(c *gin.Context) {
	var body models.CreateBusinessBody
	c.Bind(&body)

	business := body.Business
	err := services.Create(&business)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}

	location := body.Location
	location.BusinessID = business.ID
	err = services.Create(&location)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	returnData := models.CreateBusinessBody{
		Business: business,
		Location: location,
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success create data business",
		"data":    returnData,
	})
}

func UpdateBusiness(c *gin.Context) {
	var body models.CreateBusinessBody
	c.Bind(&body)

	business := body.Business
	idParam := c.Param("id")
	id := uuid.MustParse(idParam)
	err := services.UpdateModel(&models.Business{ID: &id}, business, "")
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}

	location := body.Location
	err = services.UpdateModel(&models.Location{BusinessID: &id}, &location, "business_id = ? ", id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success Update data business",
	})
}

func DeletedBusiness(c *gin.Context) {
	idParam := c.Param("id")
	id := uuid.MustParse(idParam)
	now := time.Now()
	err := services.UpdateModel(&models.Business{ID: &id}, &models.Business{DeletedAt: &now}, "")
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}

	err = services.UpdateModel(&models.Location{BusinessID: &id}, &models.Location{DeletedAt: &now}, "business_id = ? ", id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success delete data business",
		"data": map[string]any{
			"id": id,
		},
	})
}

func BusinessSearch(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	term := c.Query("term")
	sort_by := c.Query("sortBy")
	sort := c.Query("sort")
	if sort_by != "distance" {
		sort_by = "b.created_at"
	}
	if strings.ToUpper(sort) != "DESC" {
		sort = "ASC"
	}

	sortQuery := "ORDER BY " + sort_by + " " + sort
	// if term != "" {
	term = "%" + term + "%"
	// }
	category := c.Query("category")
	// if category != "" {
	category = "%" + category + "%"
	// }
	latitude, _ := strconv.ParseFloat(c.Query("latitude"), 64)
	longitude, _ := strconv.ParseFloat(c.Query("longitude"), 64)
	distance, _ := strconv.ParseFloat(c.Query("distance"), 64)

	offset := 0
	if page > 0 {
		offset = page - 1
	}
	if limit == 0 {
		limit = 20
	}
	queryString := `SELECT 
    b.id,
    b.alias,
    b.name,
    b.image_url,
    b.is_closed,
    b.url,
    b.coordinates,
    b.price,
    b.phone,
    b.display_phone,
    b.transactions,
	ST_Distance_Sphere(
    point(?, ?),
    point(JSON_EXTRACT(b.coordinates, '$.latitude'), JSON_EXTRACT(b.coordinates, '$.longitude')))/1000 distance,
    json_arrayagg(json_object("alias", c.alias, "name", c.name)) categories,
    json_object("address1", l.address1, "address2", l.address2, "address3", l.address3,"city", l.city, "zip_code", l.zip_code, "country", l.country, "state", l.state, "display_address", l.display_address) location 
FROM
    businesses b
        LEFT JOIN
    locations l ON l.business_id = b.id
		LEFT JOIN
	categories c ON find_in_set(c.id, b.categories)
	WHERE (b.name like ? or b.alias like ?) 
	AND (c.name like ? or c.alias like ?)
	AND CASE WHEN ? > 0 THEN ST_Distance_Sphere(
		point(?, ?),
		point(JSON_EXTRACT(b.coordinates, '$.latitude'), JSON_EXTRACT(b.coordinates, '$.longitude')))/1000 > ? ELSE TRUE END
    group by b.id ` + sortQuery +
		` limit ?
	offset ? 
	;`
	var returnData []models.BusinessSearch
	err := services.GetAll(queryString, &returnData, latitude, longitude, term, term, category, category, distance, latitude, longitude, distance, limit, offset)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	if returnData == nil {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Success get data business",
			"data":    []any{},
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success get data business",
		"data":    returnData,
	})
}
