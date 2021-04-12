package main
import "github.com/kataras/iris/v12"

func main(){
	app:= iris.New()
	booksAPI:=app.Party("/books")
	{
		booksAPI.Use(iris.Compression)
		booksAPI.Get("/",list)
		booksAPI.Post("/",create)
	}
	app.Listen(":8080")
}

type Book struct {
	Title string `json:"Title"`
}

func list(ctx iris.Context){
	books:=[]Book{
		{"Hello world"},
		{"Hello go lang"},
		{"Coba saja"},
	}
	ctx.JSON(books)
}

func create(ctx iris.Context){
	var b Book;
	err := ctx.ReadJSON(&b)
	if err!= nil {
		ctx.StopWithProblem(iris.StatusBadRequest,iris.NewProblem().Title("Gagal membuat data baru").DetailErr(err))
		return
	}
	println("Received book:"+b.Title)
	ctx.StatusCode(iris.StatusCreated)
}