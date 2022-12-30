package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func artistInfoByPageHandler(c echo.Context) error {
	page := c.QueryParam("page")
	filter := bson.M{"page": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("artistview").Collection("artistview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "initArtistInfo find has failed")
	var allartist []ArtVIEW
	if err = cur.All(context.TODO(), &allartist); err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, allartist)
}

func albumInfoByPageHandler(c echo.Context) error {
	page := c.QueryParam("page")
	filter := bson.M{"albumpage": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "initAlbumInfo find has failed")
	var allalbums []AlbVieW2
	if err = cur.All(context.TODO(), &allalbums); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is allalbums", allalbums)
	return c.JSON(http.StatusOK, allalbums)
}

func songInfoByPageHandler(c echo.Context) error {
	page := c.QueryParam("page")
	filter := bson.M{"Page": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "Artist": 1, "Song": 1, "File_id": 1, "Full_Filename": 1})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "songInfoByPageHandler find has failed")
	var songs []map[string]string
	if err = cur.All(context.TODO(), &songs); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is songs", songs)
	return c.JSON(http.StatusOK, songs)
}

func albumsForArtistHandler(c echo.Context) error {
	fmt.Println("Starting albumsForArtistHandler")
	var artistid string = c.QueryParam("selected")
	fmt.Printf("%s this is artistid", artistid)
	fmt.Printf("%T this is artistid type", artistid)
	filter := bson.M{"artistID": artistid}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "songs": 0})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("albumview").Collection("albumview")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "initArtistInfo find has failed")
	var allalbum []map[string]string
	if err = cur.All(context.TODO(), &allalbum); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is allalbum-", allalbum)
	return c.JSON(http.StatusOK, allalbum)
}

func songsForAlbumHandler(c echo.Context) error {
	fmt.Println("Starting songsForAlbumHandler")
	var albumid string = c.QueryParam("selected")
	fmt.Printf("%s this is albumid", albumid)
	fmt.Printf("%T this is albumid type", albumid)
	filter := bson.M{"albumID": albumid}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "songsForAlbumHandler find has failed")
	var allsongs []map[string]string
	if err = cur.All(context.TODO(), &allsongs); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s this is allalbum-", allsongs)
	return c.JSON(http.StatusOK, allsongs)
}

func shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

func randomPicsHandler(c echo.Context) error {
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "Index": 1})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("coverart").Collection("coverartpages")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "randomPicsHandler has failed")
	var indexliststring []map[string]string
	if err = cur.All(context.TODO(), &indexliststring); err != nil {
		fmt.Println(err)
	}
	var num_list []int
	for _, idx := range indexliststring {
		indexx := idx["Index"]
		index1, _ := strconv.Atoi(indexx)
		num_list = append(num_list, index1)
	}
	shuffle(num_list)
	var randpics []string
	for _, f := range num_list[:12] {
		ff := strconv.Itoa(f)
		filter := bson.M{"Index": ff}
		client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
		defer Close(client, ctx, cancel)
		CheckError(err, "MongoDB connection has failed")
		collection := client.Database("coverart").Collection("coverartpages")
		var rpics map[string]string
		err = collection.FindOne(context.Background(), filter).Decode(&rpics)
		if err != nil {
			fmt.Println(err)
		}
		randpics = append(randpics, rpics["Img_base64_str"])
	}
	return c.JSON(http.StatusOK, randpics)
}

///////////////////////////////////////////////////
///////////////////////////////////////////////////

func ArtistFirstLetterHandler(c echo.Context) error {
	artistAlphaAll := AmpgoFind("artistalpha", "artistalphalist", "None", "None")
	return c.JSON(http.StatusOK, artistAlphaAll)
}

func Artist1Handler(c echo.Context) error {
	Artist1 := AmpgoFind("artistalpha", "fl1", "None", "None")
	return c.JSON(http.StatusOK, Artist1)
}

func Artist2Handler(c echo.Context) error {
	Artist2 := AmpgoFind("artistalpha", "fl2", "None", "None")
	return c.JSON(http.StatusOK, Artist2)
}

func Artist3Handler(c echo.Context) error {
	Artist3 := AmpgoFind("artistalpha", "fl3", "None", "None")
	return c.JSON(http.StatusOK, Artist3)
}

func Artist4Handler(c echo.Context) error {
	Artist4 := AmpgoFind("artistalpha", "fl4", "None", "None")
	return c.JSON(http.StatusOK, Artist4)
}

func Artist5Handler(c echo.Context) error {
	Artist5 := AmpgoFind("artistalpha", "fl5", "None", "None")
	return c.JSON(http.StatusOK, Artist5)
}

func Artist6Handler(c echo.Context) error {
	Artist6 := AmpgoFind("artistalpha", "fl6", "None", "None")
	return c.JSON(http.StatusOK, Artist6)
}

func Artist7Handler(c echo.Context) error {
	Artist7 := AmpgoFind("artistalpha", "fl7", "None", "None")
	return c.JSON(http.StatusOK, Artist7)
}

func Artist8Handler(c echo.Context) error {
	Artist8 := AmpgoFind("artistalpha", "fl8", "None", "None")
	return c.JSON(http.StatusOK, Artist8)
}

func Artist9Handler(c echo.Context) error {
	Artist9 := AmpgoFind("artistalpha", "fl9", "None", "None")
	return c.JSON(http.StatusOK, Artist9)
}

func ArtistAHandler(c echo.Context) error {
	artistA := AmpgoFind("artistalpha", "A", "None", "None")
	return c.JSON(http.StatusOK, artistA)
}

func ArtistBHandler(c echo.Context) error {
	artistA := AmpgoFind("artistalpha", "B", "None", "None")
	return c.JSON(http.StatusOK, artistA)
}

///////////////////////////////////////////////////
///////////////////////////////////////////////////

func AlbumFirstLetterHandler(c echo.Context) error {
	albumAlphaAll := AmpgoFind("albumalpha", "albumalphalist", "None", "None")
	return c.JSON(http.StatusOK, albumAlphaAll)
}

func Album1Handler(c echo.Context) error {
	Album1 := AmpgoFind("albumalpha", "fl1", "None", "None")
	return c.JSON(http.StatusOK, Album1)
}

func Album2Handler(c echo.Context) error {
	Album2 := AmpgoFind("albumalpha", "fl2", "None", "None")
	return c.JSON(http.StatusOK, Album2)
}

func Album3Handler(c echo.Context) error {
	Album3 := AmpgoFind("albumalpha", "fl3", "None", "None")
	return c.JSON(http.StatusOK, Album3)
}

func Album4Handler(c echo.Context) error {
	Album4 := AmpgoFind("albumalpha", "fl4", "None", "None")
	return c.JSON(http.StatusOK, Album4)
}

func Album5Handler(c echo.Context) error {
	Album5 := AmpgoFind("albumalpha", "fl5", "None", "None")
	return c.JSON(http.StatusOK, Album5)
}

func Album6Handler(c echo.Context) error {
	Album6 := AmpgoFind("albumalpha", "fl6", "None", "None")
	return c.JSON(http.StatusOK, Album6)
}

func Album7Handler(c echo.Context) error {
	Album7 := AmpgoFind("albumalpha", "fl7", "None", "None")
	return c.JSON(http.StatusOK, Album7)
}

func Album8Handler(c echo.Context) error {
	Album8 := AmpgoFind("albumalpha", "fl8", "None", "None")
	return c.JSON(http.StatusOK, Album8)
}

func Album9Handler(c echo.Context) error {
	Album9 := AmpgoFind("albumalpha", "fl9", "None", "None")
	return c.JSON(http.StatusOK, Album9)
}

func AlbumAHandler(c echo.Context) error {
	albumA := AmpgoFind("albumalpha", "A", "None", "None")
	return c.JSON(http.StatusOK, albumA)
}

func AlbumBHandler(c echo.Context) error {
	albumB := AmpgoFind("albumalpha", "B", "None", "None")
	return c.JSON(http.StatusOK, albumB)
}

func AlbumCHandler(c echo.Context) error {
	albumC := AmpgoFind("albumalpha", "C", "None", "None")
	return c.JSON(http.StatusOK, albumC)
}

func AlbumDHandler(c echo.Context) error {
	albumD := AmpgoFind("albumalpha", "D", "None", "None")
	return c.JSON(http.StatusOK, albumD)
}

func AlbumEHandler(c echo.Context) error {
	albumE := AmpgoFind("albumalpha", "E", "None", "None")
	return c.JSON(http.StatusOK, albumE)
}

func AlbumFHandler(c echo.Context) error {
	albumF := AmpgoFind("albumalpha", "F", "None", "None")
	return c.JSON(http.StatusOK, albumF)
}

func AlbumGHandler(c echo.Context) error {
	albumG := AmpgoFind("albumalpha", "G", "None", "None")
	return c.JSON(http.StatusOK, albumG)
}

func AlbumHHandler(c echo.Context) error {
	albumH := AmpgoFind("albumalpha", "H", "None", "None")
	return c.JSON(http.StatusOK, albumH)
}

func AlbumIHandler(c echo.Context) error {
	albumI := AmpgoFind("albumalpha", "I", "None", "None")
	return c.JSON(http.StatusOK, albumI)
}

func AlbumJHandler(c echo.Context) error {
	albumJ := AmpgoFind("albumalpha", "J", "None", "None")
	return c.JSON(http.StatusOK, albumJ)
}

///////////////////////////////////////////////////
///////////////////////////////////////////////////

func SongFirstLetterHandler(c echo.Context) error {
	songAlphaAll := AmpgoFind("songalpha", "songalphalist", "None", "None")
	return c.JSON(http.StatusOK, songAlphaAll)
}

func Song1Handler(c echo.Context) error {
	song1 := AmpgoFind("songalpha", "fl1", "None", "None")
	return c.JSON(http.StatusOK, song1)
}

func Song2Handler(c echo.Context) error {
	song2 := AmpgoFind("songalpha", "fl2", "None", "None")
	return c.JSON(http.StatusOK, song2)
}

func Song3Handler(c echo.Context) error {
	song3 := AmpgoFind("songalpha", "fl3", "None", "None")
	return c.JSON(http.StatusOK, song3)
}

func Song4Handler(c echo.Context) error {
	song4 := AmpgoFind("songalpha", "fl4", "None", "None")
	return c.JSON(http.StatusOK, song4)
}

func Song5Handler(c echo.Context) error {
	song5 := AmpgoFind("songalpha", "fl5", "None", "None")
	return c.JSON(http.StatusOK, song5)
}

func Song6Handler(c echo.Context) error {
	song6 := AmpgoFind("songalpha", "fl6", "None", "None")
	return c.JSON(http.StatusOK, song6)
}

func Song7Handler(c echo.Context) error {
	song7 := AmpgoFind("songalpha", "fl7", "None", "None")
	return c.JSON(http.StatusOK, song7)
}

func Song8Handler(c echo.Context) error {
	song8 := AmpgoFind("songalpha", "fl8", "None", "None")
	return c.JSON(http.StatusOK, song8)
}

func Song9Handler(c echo.Context) error {
	song9 := AmpgoFind("songalpha", "fl9", "None", "None")
	return c.JSON(http.StatusOK, song9)
}

func SongAHandler(c echo.Context) error {
	songA := AmpgoFind("songalpha", "A", "None", "None")
	return c.JSON(http.StatusOK, songA)
}

func SongBHandler(c echo.Context) error {
	songB := AmpgoFind("songalpha", "B", "None", "None")
	return c.JSON(http.StatusOK, songB)
}

func SongCHandler(c echo.Context) error {
	songC := AmpgoFind("songalpha", "C", "None", "None")
	return c.JSON(http.StatusOK, songC)
}