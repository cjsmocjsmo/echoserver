package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	// "github.com/bogem/id3v2"
	// "github.com/disintegration/imaging"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}







func CreateTextSearchIndexes(db string, coll string) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "AmpgoFindOne: MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)

	mod := mongo.IndexModel{
		Keys: bson.M{
			"song": "text",
		},
		Options: nil,
	}

	// ind, err := collection.Indexes().CreateOne(ctx, mod)
	ind, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		log.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1) // exit in case of error
	} else {
		// API call returns string of the index name
		fmt.Println("CreateOne() index:", ind)
		fmt.Println("CreateOne() type:", reflect.TypeOf(ind))

		log.Println("CreateOne() index:", ind)
		log.Println("CreateOne() type:", reflect.TypeOf(ind))
	}
}






















func AmpgoFindOne(db string, coll string, filtertype string, filterstring string) map[string]string {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "AmpgoFindOne: MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results map[string]string = make(map[string]string)
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		fmt.Println("AmpgoFindOne: find one has fucked up")
		log.Println(err)
	}
	return results
}

func AmpgoFind(dbb string, collb string, filtertype string, filterstring string) []map[string]string {
	filter := bson.M{}
	if filtertype != "None" && filterstring != "None" {
		filter = bson.M{filtertype: filterstring}
	}
	// filter := bson.D{{filtertype, filterstring}}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "AmpgoFind: MongoDB connection has failed")
	coll := client.Database(dbb).Collection(collb)
	cur, err := coll.Find(context.TODO(), filter)
	CheckError(err, "AmpgoFind: ArtPipeline find has failed")
	var results []map[string]string //all albums for artist to include double entries
	if err = cur.All(context.TODO(), &results); err != nil {
		fmt.Println("AmpgoFind: cur.All has fucked up")
		log.Println(err)
	}
	return results
}

func AmpgoInsertOne(db string, coll string, ablob map[string]string) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "AmpgoInsertOne: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "AmpgoInsertOne has failed")
}

func DeleteOne(client *mongo.Client, ctxx context.Context, dataBase, col string, doc interface{}) (*mongo.DeleteResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result2, err1 := collection.DeleteOne(ctxx, doc)
	return result2, err1
}

func InsertMP3Json(db string, coll string, ablob JsonMP3) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsertMP3Json: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "InsertMP3Json has failed")
}

func InsertJPGJson(db string, coll string, ablob JsonJPG) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsertJPGJson: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "InsertJPGJson has failed")
}

func InsertPagesJson(db string, coll string, ablob JsonPage) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsertPagesJson: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "InsertPagesJson has failed")
}

func InsertArtistIDSJson(db string, coll string, ablob map[string]string) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsertArtistIDSJson: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "InsertArtistIDSJson has failed")
}

func InsertAlbumIDSJson(db string, coll string, ablob AlbID) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsertAlbumIDSJson: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "InsertAlbumIDSJson has failed")
}

func GetAllMP3Objects() (Main2SL []JsonMP3) {
	filter := bson.D{}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "GetAllMP3Objects: MongoDB connection has failed")
	collection := client.Database("maindb").Collection("mp3s")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cur.All(context.Background(), &Main2SL); err != nil {
		fmt.Println("GetAllMP3Objects: cur.All has failed")
		log.Println(err)
	}
	return
}

func GetAllJPGObjects() (Main2S []JsonJPG) {
	filter := bson.D{}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "GetAllJPGObjects: MongoDB connection has failed")
	collection := client.Database("maindb").Collection("jpgs")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cur.All(context.Background(), &Main2S); err != nil {
		fmt.Println("GetAllJPGObjects: cur.All has failed")
		log.Println(err)
	}
	return
}

func AmpgoDistinct(db string, coll string, fieldd string) []string {
	filter := bson.D{}
	opts := options.Distinct().SetMaxTime(2 * time.Second)
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "AmpgoDistinct: MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	DD1, err2 := collection.Distinct(context.TODO(), fieldd, filter, opts)
	CheckError(err2, "AmpgoDistinct: MongoDB distinct album has failed")
	var DAlbum1 []string
	for _, DD := range DD1 {
		zoo := fmt.Sprintf("%s", DD)
		DAlbum1 = append(DAlbum1, zoo)
	}
	return DAlbum1
}

func gArtistInfo(Art string) map[string]string {
	filter := bson.M{"Artist": Art}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "gArtistInfo: MongoDB connection has failed")
	collection := client.Database("ids").Collection("artistids")
	var ArtInfo map[string]string = make(map[string]string)
	err = collection.FindOne(context.Background(), filter).Decode(&ArtInfo)
	if err != nil {
		fmt.Println("gArtistInfo: has failed")
		log.Println(err)
	}
	return ArtInfo
}

func gAlbumInfo(Alb string) map[string]string {
	filter := bson.M{"Album": Alb}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "gAlbumInfo: MongoDB connection has failed")
	collection := client.Database("ids").Collection("albumids")
	var AlbInfo map[string]string = make(map[string]string)
	err = collection.FindOne(context.Background(), filter).Decode(&AlbInfo)
	if err != nil {
		fmt.Println("gAlbumInfo: has failed")
		log.Println(err)
	}
	return AlbInfo
}

func InsArtPipeline(AV1 ArtVieW2) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsArtPipeline: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "artistview", "artistview", &AV1)
	CheckError(err2, "InsArtPipeline: artistview insertion has failed")
}

func InsAlbViewID(MyAlbview AlbVieW2) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	CheckError(err, "InsAlbViewID: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, "albumview", "albumview", &MyAlbview)
	CheckError(err2, "InsAlbViewID: AmpgoInsertOne has failed")
}

func GetMainDbMeta() []map[string]string {
	filter := bson.M{}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "GetMainDbMeta: MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "GetMainDbMeta: allIdx has failed")
	var letters []map[string]string
	if err = cur.All(context.TODO(), &letters); err != nil {
		log.Println(err)
	}
	return letters
}

func ArtViewFindOne(db string, coll string, filtertype string, filterstring string) ArtVIEW {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results ArtVIEW
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		log.Println(err)
	}
	return results
}

func AlbViewFindOne(db string, coll string, filtertype string, filterstring string) AlbVieW2 {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var results AlbVieW2
	err = collection.FindOne(context.Background(), filter).Decode(&results)
	if err != nil {
		log.Println(err)
	}
	return results
}

func FrontMatterFindOne(db string, coll string, filtertype string, filterstring string) FrontMatter {
	filter := bson.M{filtertype: filterstring}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "FrontMatterFindOne: MongoDB connection has failed")
	collection := client.Database(db).Collection(coll)
	var fmresults FrontMatter
	err = collection.FindOne(context.Background(), filter).Decode(&fmresults)
	if err != nil {
		fmt.Println("FrontMatterFindOne: find one has fucked up")
		log.Println(err)
	}
	return fmresults
}
