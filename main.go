package main
import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type User struct {
	Age      int
	Name     string
	Email    string
	Password string
}

type Service struct {
	Ctx   context.Context
	Users *mongo.Collection
}

func (mainEntity *Service) GetUserByName(name string) (User, error) {
	filter := bson.D{{"name", name}}
	var user User
	err := mainEntity.Users.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

type TokenEntity struct {
	Token string
}

type UserName struct {
	Name string
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		fmt.Println("connected to mongo")
	}
	users := client.Database("task-manager-api").Collection("users")
	service := Service{
		Ctx:   ctx,
		Users: users,
	}
	var user, errFind = service.GetUserByName("Artemka")
	if errFind != nil {
		fmt.Println("cannot find")
	}
	fmt.Println(user)
}
