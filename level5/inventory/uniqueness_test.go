package inventory

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDuplicateCodes(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "منتجان بكودين مختلفين",
			content: "KB-001 | لوحة مفاتيح | 10\nMS-002 | فأرة | 8\n",
			wantErr: false,
		},
		{
			name:    "منتجان بالكود KB-001 نفسه",
			content: "KB-001 | لوحة مفاتيح | 10\nKB-001 | فأرة | 8\n",
			wantErr: true,
		},
		{
			name:    "KB-001 وKB-001",
			content: "KB-001 | لوحة مفاتيح | 10\n KB-001  | فأرة | 8\n",
			wantErr: true,
		},
		{
			name:    "نص فارغ",
			content: "",
			wantErr: false,
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			products, err := ParseInventory(tc.content)

			if tc.wantErr && err == nil {
				t.Errorf("كنا نتوقع خطأ ولكن تم قبوله بنجاح!")
			}

			if !tc.wantErr && err != nil {
				t.Errorf("لم نكن نتوقع خطأ ولكن حدث الخطأ: %v", err)
			}

			if tc.name == "نص فارغ" && len(products) != 0 {
				t.Errorf("توقعنا قائمة فارغة ولكن حصلنا على %d عناصر", len(products))
			}
		})
	}
}

func TestJSONDuplicateCodes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "inventory.json")

	// 1. حفظ منتجات سليمة
	validProducts := []Product{
		{Code: "KB-001", Name: "لوحة مفاتيح", Quantity: 10},
		{Code: "KB-002", Name: "لوحة ثانية", Quantity: 2},
	}
	if err := SaveJSON(path, validProducts); err != nil {
		t.Fatalf("فشل حفظ المنتجات السليمة: %v", err)
	}

	// 2. محاولة حفظ منتجات مكررة
	duplicateProducts := []Product{
		{Code: "KB-001", Name: "لوحة مفاتيح", Quantity: 10},
		{Code: "KB-001", Name: "لوحة ثانية مكررة", Quantity: 5},
	}
	if err := SaveJSON(path, duplicateProducts); err == nil {
		t.Errorf("كنا نتوقع خطأ عند محاولة حفظ كود مكرر ولكن تم قبوله!")
	}

	// 3. التحقق من بقاء البيانات السابقة دون تغيير
	reloaded, err := LoadJSON(path)
	if err != nil || len(reloaded) != 2 || reloaded[0].Code != "KB-001" || reloaded[1].Code != "KB-002" {
		t.Fatalf("البيانات السابقة تم تخريبها أو تغييرها بعد محاولة الحفظ الخاطئة!")
	}

	// 4. رفض JSON خارجي غير صالح به كود مكرر
	corruptPath := filepath.Join(t.TempDir(), "external_duplicate.json")
	rawJSON := `[
		{"code":"KB-001","name":"لوحة مفاتيح","quantity":12},
		{"code":"KB-001","name":"فأرة","quantity":8}
	]`

	if err := os.WriteFile(corruptPath, []byte(rawJSON), 0644); err != nil {
		t.Fatalf("فشل في تحضير ملف الاختبار المؤقت: %v", err)
	}

	if _, err := LoadJSON(corruptPath); err == nil {
		t.Errorf("كنا نتوقع خطأ من LoadJSON لوجود أكواد مكررة في الملف ولكن تم قبوله!")
	}
}
