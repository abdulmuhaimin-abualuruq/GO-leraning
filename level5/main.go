package main

import (
	"fmt"

	"example.com/academy/level5/inventory"
)

const (
	inputPath  = "level5/data/products.txt"
	outputPath = "level5/data/inventory.json"
)

func main() {
	// 1. ReadInventory: قراءة products.txt
	content, err := inventory.ReadInventory(inputPath)
	if err != nil {
		fmt.Printf("خطأ في قراءة ملف المخزون: %v\n", err)
		return
	}

	// 2. ParseInventory: تحويل النص إلى []Product
	products, err := inventory.ParseInventory(content)
	if err != nil {
		fmt.Printf("خطأ في تحليل بيانات المنتجات: %v\n", err)
		return
	}
	fmt.Printf("تم استيراد %d منتجات\n", len(products))

	// 3. SaveJSON: حفظ inventory.json
	if err := inventory.SaveJSON(outputPath, products); err != nil {
		fmt.Printf("خطأ في حفظ البيانات بصيغة JSON: %v\n", err)
		return
	}
	fmt.Printf("تم حفظ البيانات في %s\n", outputPath)

	// 4. LoadJSON: استرجاع المنتجات المحفوظة
	loadedProducts, err := inventory.LoadJSON(outputPath)
	if err != nil {
		fmt.Printf("خطأ في قراءة ملف JSON: %v\n", err)
		return
	}

	// 5. عرض المخزون: طباعة القائمة المسترجعة من JSON
	fmt.Println("=== المخزون ===")
	for i := 0; i < len(loadedProducts); i++ {
		p := loadedProducts[i]
		fmt.Printf("%s | %s | %d\n", p.Code, p.Name, p.Quantity)
	}
}
