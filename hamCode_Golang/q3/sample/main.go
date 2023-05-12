package main

import "fmt"

type TLDRDBCached struct {
	AllKeys map[string]string
	Cachedkeys []string
}

func (t *TLDRDBCached) Retrieve(key string) string {

	db := GetConnection()

	err := db.AutoMigrate(&TLDREntity{})
	if err != nil {
		fmt.Println(err.Error())
	}



	// check cache
	for _, v := range t.Cachedkeys{
		if v == key{
			entity := TLDREntity{}
			result := db.Where("Key = ?", key).First(&entity)
			
			if result.Error != nil {
				fmt.Println(result.Error.Error())
			}
			return entity.Val

		}
	}

	// check exist
	check := false
	for _, v := range t.AllKeys{
		if v == key{
			check = true
		}
	} 
	if !check{
		return "Error"
	}

	
	return ""
}


func (t *TLDRDBCached) List() []string {
	check := false
	var resultList1, resultList2 []string
	for _, v := range t.AllKeys{
		for _, vcache := range t.Cachedkeys{
			if v == vcache{
				resultList1 = append(resultList1, v)
				check = true
			}
		}
		if !check {
			resultList2 = append(resultList2, v)
		} else {
			check = false
		}
	}

	for _, v := range resultList2{
		resultList1 = append(resultList1, v)

	}
	return resultList1
}

func NewTLDRDBCached(nonCachedProvider TLDRProvider) TLDRProvider {
	return &TLDRDBCached{
		AllKeys : [][],
	}
}

