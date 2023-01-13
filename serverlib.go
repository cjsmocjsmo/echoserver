package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ArtistInfoByPageHandler(c echo.Context) error {
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
		log.Println(err)
	}
	return c.JSON(http.StatusOK, allartist)
}

func AlbumInfoByPageHandler(c echo.Context) error {
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
		log.Println(err)
	}
	return c.JSON(http.StatusOK, allalbums)
}

func SongInfoByPageHandler(c echo.Context) error {
	page := c.QueryParam("page")
	filter := bson.M{"Page": page}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "Artist": 1, "Song": 1, "File_id": 1, "Full_Filename": 1})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "SongInfoByPageHandler find has failed")
	var songs []map[string]string
	if err = cur.All(context.TODO(), &songs); err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, songs)
}

func AlbumsForArtistHandler(c echo.Context) error {
	fmt.Println("Starting AlbumsForArtistHandler")
	var artistid string = c.QueryParam("selected")
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
		log.Println(err)
	}
	return c.JSON(http.StatusOK, allalbum)
}

func SongsForAlbumHandler(c echo.Context) error {
	fmt.Println("Starting SongsForAlbumHandler")
	var albumid string = c.QueryParam("selected")
	filter := bson.M{"AlbumID": albumid}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("maindb").Collection("maindb")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "SongsForAlbumHandler find has failed")
	var allsongs []map[string]string
	if err = cur.All(context.TODO(), &allsongs); err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, allsongs)
}

func shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

func RandomPicsHandler(c echo.Context) error {
	filter := bson.D{{}}
	opts := options.Find()
	opts.SetProjection(bson.M{"_id": 0, "Index": 1})
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed")
	coll := client.Database("coverart").Collection("coverartpages")
	cur, err := coll.Find(context.TODO(), filter, opts)
	CheckError(err, "RandomPicsHandler has failed")
	var indexliststring []map[string]string
	if err = cur.All(context.TODO(), &indexliststring); err != nil {
		log.Println(err)
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
			log.Println(err)
		}
		randpics = append(randpics, rpics["Img_base64_str"])
	}
	return c.JSON(http.StatusOK, randpics)
}

///////////////////////////////////////////////////
///////////////////////////////////////////////////

func ArtistsForFirstLetterHandler(c echo.Context) error {
	fl := c.QueryParam("firstletter")
	var firstletter string
	switch {
	case strings.Contains("1", fl):
		firstletter = "fl1"
	case strings.Contains("2", fl):
		firstletter = "fl2"
	case strings.Contains("3", fl):
		firstletter = "fl3"
	case strings.Contains("4", fl):
		firstletter = "fl4"
	case strings.Contains("5", fl):
		firstletter = "fl5"
	case strings.Contains("6", fl):
		firstletter = "fl6"
	case strings.Contains("7", fl):
		firstletter = "fl7"
	case strings.Contains("8", fl):
		firstletter = "fl8"
	case strings.Contains("9", fl):
		firstletter = "fl9"
	default:
		firstletter = fl
	}
	log.Println("ArtistsForFirstLetterHandler")
	log.Println(firstletter)
	allArtist := AmpgoFind("artistalpha", firstletter, "None", "None")
	return c.JSON(http.StatusOK, allArtist)
}

func AlbumsForFirstLetterHandler(c echo.Context) error {
	fl := c.QueryParam("firstletter")
	var firstletter string
	switch {
	case strings.Contains("1", fl):
		firstletter = "fl1"
	case strings.Contains("2", fl):
		firstletter = "fl2"
	case strings.Contains("3", fl):
		firstletter = "fl3"
	case strings.Contains("4", fl):
		firstletter = "fl4"
	case strings.Contains("5", fl):
		firstletter = "fl5"
	case strings.Contains("6", fl):
		firstletter = "fl6"
	case strings.Contains("7", fl):
		firstletter = "fl7"
	case strings.Contains("8", fl):
		firstletter = "fl8"
	case strings.Contains("9", fl):
		firstletter = "fl9"
	default:
		firstletter = fl
	}
	log.Println("AlbumForFirstLetterHandler")
	log.Println(firstletter)
	allAlbum := AmpgoFind("albumalpha", firstletter, "None", "None") //album && albumid
	log.Println("this is allAlbum")
	log.Println(allAlbum)
	var newalblist []AlbVieW
	for _, alb := range allAlbum {
		albinfo := AlbViewFindOne("albumview", "albumview", "albumID", alb["albumid"])
		var nm AlbVieW
		nm.Artist = albinfo.Artist
		nm.ArtistID = albinfo.ArtistID
		nm.Album = albinfo.Album
		nm.AlbumID = albinfo.AlbumID
		nm.AlbumPage = albinfo.AlbumPage
		nm.NumSongs = albinfo.NumSongs
		nm.ThumbHttpPath = albinfo.ThumbHttpPath
		newalblist = append(newalblist, nm)
	}
	return c.JSON(http.StatusOK, newalblist)
}

func SongsForFirstLetterHandler(c echo.Context) error {
	fl := c.QueryParam("firstletter")
	var firstletter string
	switch {
	case strings.Contains("1", fl):
		firstletter = "fl1"
	case strings.Contains("2", fl):
		firstletter = "fl2"
	case strings.Contains("3", fl):
		firstletter = "fl3"
	case strings.Contains("4", fl):
		firstletter = "fl4"
	case strings.Contains("5", fl):
		firstletter = "fl5"
	case strings.Contains("6", fl):
		firstletter = "fl6"
	case strings.Contains("7", fl):
		firstletter = "fl7"
	case strings.Contains("8", fl):
		firstletter = "fl8"
	case strings.Contains("9", fl):
		firstletter = "fl9"
	default:
		firstletter = fl
	}
	log.Println("SongsForFirstLetterHandler")
	log.Println(firstletter)
	allSong := AmpgoFind("songalpha", firstletter, "None", "None")
	return c.JSON(http.StatusOK, allSong)
}

//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////

func GetAlbumFirstLetterIDHandler(c echo.Context) error {
	aflID := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println("this is aflID")
	log.Println(aflID.AlbumFirstLetterID)
	fmt.Println(aflID.AlbumFirstLetterID)
	return c.JSON(http.StatusOK, aflID.AlbumFirstLetterID)
}

func GetAlbumsForArtistURLHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.AlbumsForArtistURL)
	log.Println(url.AlbumsForArtistURL)
	return c.JSON(http.StatusOK, url.AlbumsForArtistURL)
}

func GetAlbumsForFirstLetterURLHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.AlbumsForFirstLetterURL)
	log.Println(url.AlbumsForFirstLetterURL)
	return c.JSON(http.StatusOK, url.AlbumsForFirstLetterURL)
}

func GetSelectedAlbumIDHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.SelectedAlbumID)
	log.Println(url.SelectedAlbumID)
	return c.JSON(http.StatusOK, url.SelectedAlbumID)
}

func GetArtistIDHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.ArtistID)
	log.Println(url.ArtistID)
	return c.JSON(http.StatusOK, url.ArtistID)
}

func GetArtistFirstLetterIDHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.ArtistFirstLetterID)
	log.Println(url.ArtistFirstLetterID)
	return c.JSON(http.StatusOK, url.ArtistFirstLetterID)
}

func GetArtistsForFirstLetterURLHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.ArtistsForFirstLetterURL)
	log.Println(url.ArtistsForFirstLetterURL)
	return c.JSON(http.StatusOK, url.ArtistsForFirstLetterURL)
}




func GetSongIDHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.SongID)
	log.Println(url.SongID)
	return c.JSON(http.StatusOK, url.SongID)
}

func GetThumbHttpPathHandler(c echo.Context) error {
	url := FrontMatterFindOne("frontmatter", "frontmatter", "fmid", "fm01")
	log.Println(url.ThumbHttpPath)
	log.Println(url.ThumbHttpPath)
	return c.JSON(http.StatusOK, url.ThumbHttpPath)
}

//////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////

func UpdateAlbumFirstLetterIDHandler(c echo.Context) error {
	param := c.QueryParam("albid")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed GetArtistsForFirstLetterURL")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"albumfirstletterid":  param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	log.Println(result)
	CheckError(err2, "MongoDB connection has failed UpdateAlbumFirstLetterIDHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateAlbumsForArtistURLHandler(c echo.Context) error {
	param := c.QueryParam("url")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateAlbumsForArtistURLHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"albumforartisturl": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateAlbumsForArtistURLHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateAlbumsForFirstLetterURLHandler(c echo.Context) error {
	param := c.QueryParam("url")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateAlbumsForFirstLetterURLHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"albumsforfirstletterurl": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateAlbumsForFirstLetterURLHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateSelectedAlbumIDHandler(c echo.Context) error {
	param := c.QueryParam("selalbid")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateSelectedAlbumIDHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"selectedalbumid": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateSelectedAlbumIDHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateArtistIDHandler(c echo.Context) error {
	param := c.QueryParam("artid")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateArtistIDHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"artistid": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateArtistIDHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateArtistFirstLetterIDHandler(c echo.Context) error {
	param := c.QueryParam("artflid")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateArtistFirstLetterIDHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"artistfirstletterid": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateArtistFirstLetterIDHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateArtistsForFirstLetterURLHandler(c echo.Context) error {
	param := c.QueryParam("url")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateArtistsForFirstLetterURLHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"artistsforfirstletterurl": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateArtistsForFirstLetterURLHandler")
	return c.JSON(http.StatusOK, result)
}





func UpdateSongIDHandler(c echo.Context) error {
	param := c.QueryParam("sid")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateSongIDHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"songid": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateSongIDHandler")
	return c.JSON(http.StatusOK, result)
}

func UpdateThumbHttpPathHandler(c echo.Context) error {
	param := c.QueryParam("thumb")
	filter := bson.M{"fmid": "fm01"}
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgodb")
	defer Close(client, ctx, cancel)
	CheckError(err, "MongoDB connection has failed UpdateThumbHttpPathHandler")
	coll := client.Database("frontmatter").Collection("frontmatter")
	update := bson.M{"$set": bson.M{"thumbhttppath": param}}
	result, err2 := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(err2, "MongoDB connection has failed UpdateThumbHttpPathHandler")
	return c.JSON(http.StatusOK, result)
}
