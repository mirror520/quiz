# README

## Usage

```bash
git clone https://github.com/mirror520/quiz
sudo docker-compose up -d
```

## Domain Model

[model/comment](./model/comment/comment.go)

```go
type Comment struct {
	UUID     string    `json:"uuid" gorm:"primaryKey"`
	ParentID string    `json:"parentid"`
	Comment  string    `json:"comment"`
	Author   string    `json:"author"`
	Update   time.Time `json:"update"`
	Favorite bool      `json:"favorite"`
}
```

## Application Service

使用 Go Kit 精神實作微服務

### Service

`Service` 專注在實現業務邏輯，並直接操作領域模型。

此外，通常第一個參數會用來傳遞 `context.Context` 上下文，因為此作業較單純，所以並沒有實現。

[quiz/service](./service/quiz/service.go)

```go
type Service interface {
	CreateComment(*comment.Comment) (*comment.Comment, error)
	GetCommentByUUID(string) (*comment.Comment, error)
	ModifyCommentByUUID(*comment.Comment, string) (*comment.Comment, error)
	RemoveCommentByUUID(string) error
}

type service struct {
	repo comment.Repository
}

func NewService(repo comment.Repository) Service {
	svc := new(service)
	svc.repo = repo
	return svc
}
```

### Service Middleware

此範例較單純為使用到。

可針對 Service 介面再實作 `Middlewares`，如 LoggingMiddleware、ProxyingMiddleware 等。

因 Middleware 係實作相同 Service 介面，故可將 Middlewares 視為 Service。多個單一職責至 Middlwares 則可彈性串接。

```go
type ServiceMiddleware func (Service) Service

func LoggingMiddlware() ServiceMiddleware {
    return func(next Service) Service {
        return &loggingMiddleware{next}
    }
}

type loggingMiddleware struct {
    next Service
}

func (mw *loggingMiddlware) CreateComment(c *comment.Comment) (*comment.Comment, error) {
    // 前處理
    c, err := mw.next(c)
    // 後處理
    return c, err
}

(略)

func main() {
    svc := quiz.NewService()
    svc = quiz.LoggingMiddleware()(svc) // 服務套接中間件
}
```


### Endpoint

`Endpoint` 為提供 `Service` 介面方法之抽象端點，可避免存取時直接相依。

建議學習 GoKit 精神，如 Endpoint 內 req / resp 為已知結構，已可自訂 Endpoint，減少需額外的型態斷言。

[quiz/endpoint](./service/quiz/endpoint.go)

```go
func CreateCommentEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*comment.Comment)
		return svc.CreateComment(req)
	}
}

(略)
```

### Transport

`Transport` 可提供多種服務端點外部傳輸介面，如：HTTP API、gRPC、PubSub 介面等。其主要係經封包解 / 編碼作業後，即送往相同之端點 `Endpoint` 處理。

[quiz/transport](./service/quiz/http_transport.go)

```go
func CreateCommentHandler(e endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *comment.Comment
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrCreateCommentFail.Error())
			return
		}

		resp, err := e(ctx, req)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrCreateCommentFail.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
```

### 多層式包裝

即 `Transport ( Endpoint ( Middleware...( Service ) ) )`

[main](./main.go)

```go
func main() {
    svc := quiz.NewService()
    svc = quiz.LoggingMiddleware()(svc)

    endpoint := quiz.CreateCommentEndpoint(svc)    // 可供內部調用
    handler := quiz.CreateCommentHandler(endpoint) // 可供外部調用
}
```

## Persistent

### Domain Repository

宣告領域模型需提供資料持續化的方法

[model/comment](./model/comment/comment.go)

```go
type Repository interface {
	Store(*Comment) error
	FindCommentByUUID(string) (*Comment, error)
	Remove(string) error
}
```

### 提供領域模型資料持續化

可以使用 `DB`、`INMEM`、`Cache`、`File` 等各種方式實現模型持續化。

其將實作 comment.Repository 介面，故可注入至所需服務中，而服務不需理解其持續化是使用何種方式實現。

[persistent/db](./persistent/db/comment.go)

```go
type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository() comment.Repository {
    // 建立 DB 連線實例
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := new(commentRepository)
	repo.db = db
	return repo
}

func (repo *commentRepository) Store(c *comment.Comment) error {
	var tx *gorm.DB
	if c.UUID == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		c.UUID = id.String()

		tx = repo.db.Create(c)
	} else {
		tx = repo.db.Save(c)
	}

	return tx.Error
}

(略)
```

`注入服務`

[main](./main.go)

```go
func main() {
    // 注入 DB 提供的 comment.Repository 持續化至服務
    {
	    comments := db.NewCommentRepository()
	    svc := quiz.NewService(comments)
    }

    // 注入 INMEM 提供的 comment.Repository 持續化至服務
    {
	    comments := inmem.NewCommentRepository()
	    svc := quiz.NewService(comments)
    }
}
```

## 測試案例

- 針對 `Service` 業務邏輯的測試
[quiz/service_test](./service/quiz/service_test.go)

- 針對 DB 模型資料持續化的測試
[persistent/db](./persistent/db/comment_test.go)
