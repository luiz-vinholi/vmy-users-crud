package rest

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/benweissmann/memongo"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/luiz-vinholi/vmy-users-crud/src/app"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/databases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	prepareToTest(m)
}

func prepareToTest(m *testing.M) {
	mongoServer, err := memongo.Start("4.0.5")
	defer mongoServer.Stop()
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("MONGODB_URI", mongoServer.URI())
	os.Setenv("MONGODB_DATABASE_NAME", memongo.RandomDatabase())
	databases.ConnectMongoDB()
	app.Load()
	router = setupRouter()
	code := m.Run()
	os.Exit(code)
}

func createDBCollection() *mongo.Collection {
	var result bson.M
	databases.Mongo.RunCommand(context.TODO(), bson.M{"create": "users"}).Decode(&result)
	return databases.Mongo.Collection("users")
}

func insertUserInDB(client *mongo.Collection) string {
	birthDate, _ := time.Parse("2006-01-02", faker.Date())
	data := bson.M{
		"name":      faker.Name(),
		"email":     faker.Email(),
		"password":  faker.Jwt(),
		"birthDate": birthDate.Format("2006-01-02"),
		"address": bson.M{
			"street":  faker.Sentence(),
			"city":    faker.Word(),
			"state":   faker.Word(),
			"country": faker.Word(),
		},
		"createdDate": time.Now(),
	}
	user, _ := client.InsertOne(context.TODO(), data)
	return user.InsertedID.(primitive.ObjectID).Hex()
}

func updateUserInDB(client *mongo.Collection, id string, data map[string]interface{}) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	client.UpdateOne(context.TODO(), filter, bson.M{"$set": data})
}

func findUserInDB(client *mongo.Collection, id string) map[string]interface{} {
	var user map[string]interface{}
	objId, _ := primitive.ObjectIDFromHex(id)
	client.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&user)
	return user
}

func deleteUserInDB(client *mongo.Collection, id string) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	client.DeleteOne(context.TODO(), filter)
}

func deleteUsersInDB(client *mongo.Collection, ids []string) {
	for _, id := range ids {
		go deleteUserInDB(client, id)
	}
}

func parseBodyString(bodyStr string) map[string]interface{} {
	var body map[string]interface{}
	err := json.Unmarshal([]byte(bodyStr), &body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
