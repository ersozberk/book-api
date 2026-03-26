# 📚 Go REST API: Temiz Mimari

[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/ersozberk/book-api/edit/main/README.md)
[![pt-br](https://img.shields.io/badge/lang-tr-green.svg)](https://github.com/ersozberk/book-api/edit/main/README-tr.md)

![Go [Sürüm](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Mimari](https://img.shields.io/badge/Architecture-Modular%20%2F%20Internal-orange)
![Bağımlılıklar](https://img.shields.io/badge/Dependencies-Zero-brightgreen)

Go'nun standart kütüphanesini (`net/http`) kullanarak tamamen sıfırdan oluşturulmuş modüler bir RESTful API. Bu proje, harici web çerçevelerine güvenmeden, endişeleri farklı katmanlara ayırarak ölçeklenebilir bir Go uygulamasının nasıl yapılandırılacağını göstermektedir.

## ✨ Özellikler

* **Sıfır Bağımlılık:** Yönlendirme ve HTTP işleme için tamamen Go standart kütüphane uygulaması.

* **Temiz Mimari:** Sektör standardı `cmd/` ve `internal/` dizin düzeni kullanılarak düzenlenmiştir.

* **Eşzamanlılık Güvenliği:** Bellek içi veri deposunda eşzamanlı okuma/yazma işlemleri sırasında yarış koşullarını önlemek için `sync.Mutex` kullanır.

* **RESTful Prensipleri:** HTTP yöntemlerine (GET, POST) ve doğru durum kodlarına (200 OK, 201 Created, 400 Bad Request) sıkı bağlılık.

## 📂 Proje Yapısı

```metin
.

├── cmd/
│ └── server/
│ └── main.go # Uygulama giriş noktası ve rota kaydı
├── internal/
│ ├── handlers/
│ │ └── book.go # HTTP işleyicileri ve iş mantığı
│ └── models/
│ └── book.go # Veri yapıları ve etki alanı modelleri
├── go.mod # Go modül tanımları
├── .gitignore # Git yok sayma kuralları
└── README.md # Proje dokümantasyonu
```

## 🛠 Kurulum ve Kullanım
Depoyu klonlayın:

```bash
git clone [https://github.com/ersozberk/book-api.git](https://github.com/ersozberk/book-api.git)
cd book-api
```

API sunucusunu çalıştırın:

```bash
go run cmd/server/main.go

``` Uç noktaları test edin (ayrı bir terminalde):

```bash
# Tüm kitapları getir (GET)
curl -s http://localhost:8080/api/books

# Yeni bir kitap ekle (POST)
curl -s -X POST http://localhost:8080/api/books \
-H "Content-Type: application/json" \
-d '{"id": "2", "title": "Clean Architecture", "author": "Robert C. Martin", "price": 45.00}'
```

Derleme için Üretim:

```bash
go build -o book-server cmd/server/main.go
./book-server
```
