package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	// "strings"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
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
	fmt.Println(fl)
	r, _ := regexp.Compile("artist([1-9]+)")
	var firstletter string
	switch {
	case r.MatchString("artist1"):
		firstletter = "fl1"
	case r.MatchString("artist2"):
		firstletter = "fl2"
	case r.MatchString("artist3"):
		firstletter = "fl3"
	case r.MatchString("artist4"):
		firstletter = "fl4"
	case r.MatchString("artist5"):
		firstletter = "fl5"
	case r.MatchString("artist6"):
		firstletter = "fl6"
	case r.MatchString("artist7"):
		firstletter = "fl7"
	case r.MatchString("artist8"):
		firstletter = "fl8"
	case r.MatchString("artist9"):
		firstletter = "fl9"
	default:
		firstletter = fl
	}

	allArtist := AmpgoFind("artistalpha", firstletter, "None", "None")
	return c.JSON(http.StatusOK, allArtist)
}

func AlbumsForFirstLetterHandler(c echo.Context) error {
	fl := c.QueryParam("firstletter")
	fmt.Println(fl)
	r, _ := regexp.Compile("album([1-9]+)")
	var firstletter string
	switch {
	case r.MatchString("album1"):
		firstletter = "fl1"
	case r.MatchString("album2"):
		firstletter = "fl2"
	case r.MatchString("album3"):
		firstletter = "fl3"
	case r.MatchString("album4"):
		firstletter = "fl4"
	case r.MatchString("album5"):
		firstletter = "fl5"
	case r.MatchString("album6"):
		firstletter = "fl6"
	case r.MatchString("album7"):
		firstletter = "fl7"
	case r.MatchString("album8"):
		firstletter = "fl8"
	case r.MatchString("album9"):
		firstletter = "fl9"
	default:
		firstletter = fl
	}
	allAlbum := AmpgoFind("albumalpha", firstletter, "None", "None")
	return c.JSON(http.StatusOK, allAlbum)
}

func SongsForFirstLetterHandler(c echo.Context) error {
	fl := c.QueryParam("firstletter")
	fmt.Println(fl)
	r, _ := regexp.Compile("song([1-9]+)")
	var firstletter string
	switch {
	case r.MatchString("song1"):
		firstletter = "fl1"
	case r.MatchString("song2"):
		firstletter = "fl2"
	case r.MatchString("song3"):
		firstletter = "fl3"
	case r.MatchString("song4"):
		firstletter = "fl4"
	case r.MatchString("song5"):
		firstletter = "fl5"
	case r.MatchString("song6"):
		firstletter = "fl6"
	case r.MatchString("song7"):
		firstletter = "fl7"
	case r.MatchString("song8"):
		firstletter = "fl8"
	case r.MatchString("song9"):
		firstletter = "fl9"
	default:
		firstletter = fl
	}
	allSong := AmpgoFind("songalpha", firstletter, "None", "None")
	return c.JSON(http.StatusOK, allSong)
}

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
	ArtistA := AmpgoFind("artistalpha", "A", "None", "None")
	return c.JSON(http.StatusOK, ArtistA)
}

func ArtistBHandler(c echo.Context) error {
	ArtistB := AmpgoFind("artistalpha", "B", "None", "None")
	return c.JSON(http.StatusOK, ArtistB)
}

func ArtistCHandler(c echo.Context) error {
	ArtistC := AmpgoFind("artistalpha", "C", "None", "None")
	return c.JSON(http.StatusOK, ArtistC)
}

func ArtistDHandler(c echo.Context) error {
	ArtistD := AmpgoFind("artistalpha", "D", "None", "None")
	return c.JSON(http.StatusOK, ArtistD)
}

func ArtistEHandler(c echo.Context) error {
	ArtistE := AmpgoFind("artistalpha", "E", "None", "None")
	return c.JSON(http.StatusOK, ArtistE)
}

func ArtistFHandler(c echo.Context) error {
	ArtistF := AmpgoFind("artistalpha", "F", "None", "None")
	return c.JSON(http.StatusOK, ArtistF)
}

func ArtistGHandler(c echo.Context) error {
	ArtistG := AmpgoFind("artistalpha", "G", "None", "None")
	return c.JSON(http.StatusOK, ArtistG)
}

func ArtistHHandler(c echo.Context) error {
	ArtistH := AmpgoFind("artistalpha", "H", "None", "None")
	return c.JSON(http.StatusOK, ArtistH)
}

func ArtistIHandler(c echo.Context) error {
	ArtistI := AmpgoFind("artistalpha", "I", "None", "None")
	return c.JSON(http.StatusOK, ArtistI)
}

func ArtistJHandler(c echo.Context) error {
	ArtistJ := AmpgoFind("artistalpha", "J", "None", "None")
	return c.JSON(http.StatusOK, ArtistJ)
}

func ArtistKHandler(c echo.Context) error {
	ArtistK := AmpgoFind("artistalpha", "K", "None", "None")
	return c.JSON(http.StatusOK, ArtistK)
}

func ArtistLHandler(c echo.Context) error {
	ArtistL := AmpgoFind("artistalpha", "L", "None", "None")
	return c.JSON(http.StatusOK, ArtistL)
}

func ArtistMHandler(c echo.Context) error {
	ArtistM := AmpgoFind("artistalpha", "M", "None", "None")
	return c.JSON(http.StatusOK, ArtistM)
}

func ArtistNHandler(c echo.Context) error {
	ArtistN := AmpgoFind("artistalpha", "N", "None", "None")
	return c.JSON(http.StatusOK, ArtistN)
}

func ArtistOHandler(c echo.Context) error {
	ArtistO := AmpgoFind("artistalpha", "O", "None", "None")
	return c.JSON(http.StatusOK, ArtistO)
}

func ArtistPHandler(c echo.Context) error {
	ArtistP := AmpgoFind("artistalpha", "P", "None", "None")
	return c.JSON(http.StatusOK, ArtistP)
}

func ArtistQHandler(c echo.Context) error {
	ArtistQ := AmpgoFind("artistalpha", "Q", "None", "None")
	return c.JSON(http.StatusOK, ArtistQ)
}

func ArtistRHandler(c echo.Context) error {
	ArtistR := AmpgoFind("artistalpha", "R", "None", "None")
	return c.JSON(http.StatusOK, ArtistR)
}

func ArtistSHandler(c echo.Context) error {
	ArtistS := AmpgoFind("artistalpha", "S", "None", "None")
	return c.JSON(http.StatusOK, ArtistS)
}

func ArtistTHandler(c echo.Context) error {
	ArtistT := AmpgoFind("artistalpha", "T", "None", "None")
	return c.JSON(http.StatusOK, ArtistT)
}

func ArtistUHandler(c echo.Context) error {
	ArtistU := AmpgoFind("artistalpha", "U", "None", "None")
	return c.JSON(http.StatusOK, ArtistU)
}

func ArtistVHandler(c echo.Context) error {
	ArtistV := AmpgoFind("artistalpha", "V", "None", "None")
	return c.JSON(http.StatusOK, ArtistV)
}

func ArtistWHandler(c echo.Context) error {
	ArtistW := AmpgoFind("artistalpha", "W", "None", "None")
	return c.JSON(http.StatusOK, ArtistW)
}

func ArtistXHandler(c echo.Context) error {
	ArtistX := AmpgoFind("artistalpha", "X", "None", "None")
	return c.JSON(http.StatusOK, ArtistX)
}

func ArtistYHandler(c echo.Context) error {
	ArtistY := AmpgoFind("artistalpha", "Y", "None", "None")
	return c.JSON(http.StatusOK, ArtistY)
}

func ArtistZHandler(c echo.Context) error {
	ArtistZ := AmpgoFind("artistalpha", "Z", "None", "None")
	return c.JSON(http.StatusOK, ArtistZ)
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

func AlbumKHandler(c echo.Context) error {
	AlbumK := AmpgoFind("albumalpha", "K", "None", "None")
	return c.JSON(http.StatusOK, AlbumK)
}

func AlbumLHandler(c echo.Context) error {
	AlbumL := AmpgoFind("albumalpha", "L", "None", "None")
	return c.JSON(http.StatusOK, AlbumL)
}

func AlbumMHandler(c echo.Context) error {
	AlbumM := AmpgoFind("albumalpha", "M", "None", "None")
	return c.JSON(http.StatusOK, AlbumM)
}

func AlbumNHandler(c echo.Context) error {
	AlbumN := AmpgoFind("albumalpha", "N", "None", "None")
	return c.JSON(http.StatusOK, AlbumN)
}

func AlbumOHandler(c echo.Context) error {
	AlbumO := AmpgoFind("albumalpha", "O", "None", "None")
	return c.JSON(http.StatusOK, AlbumO)
}

func AlbumPHandler(c echo.Context) error {
	AlbumP := AmpgoFind("albumalpha", "P", "None", "None")
	return c.JSON(http.StatusOK, AlbumP)
}

func AlbumQHandler(c echo.Context) error {
	AlbumQ := AmpgoFind("albumalpha", "Q", "None", "None")
	return c.JSON(http.StatusOK, AlbumQ)
}

func AlbumRHandler(c echo.Context) error {
	AlbumR := AmpgoFind("albumalpha", "R", "None", "None")
	return c.JSON(http.StatusOK, AlbumR)
}

func AlbumSHandler(c echo.Context) error {
	AlbumS := AmpgoFind("albumalpha", "S", "None", "None")
	return c.JSON(http.StatusOK, AlbumS)
}

func AlbumTHandler(c echo.Context) error {
	AlbumT := AmpgoFind("albumalpha", "T", "None", "None")
	return c.JSON(http.StatusOK, AlbumT)
}

func AlbumUHandler(c echo.Context) error {
	AlbumU := AmpgoFind("albumalpha", "U", "None", "None")
	return c.JSON(http.StatusOK, AlbumU)
}

func AlbumVHandler(c echo.Context) error {
	AlbumV := AmpgoFind("albumalpha", "V", "None", "None")
	return c.JSON(http.StatusOK, AlbumV)
}

func AlbumWHandler(c echo.Context) error {
	AlbumW := AmpgoFind("albumalpha", "W", "None", "None")
	return c.JSON(http.StatusOK, AlbumW)
}

func AlbumXHandler(c echo.Context) error {
	AlbumX := AmpgoFind("albumalpha", "X", "None", "None")
	return c.JSON(http.StatusOK, AlbumX)
}

func AlbumYHandler(c echo.Context) error {
	AlbumY := AmpgoFind("albumalpha", "Y", "None", "None")
	return c.JSON(http.StatusOK, AlbumY)
}

func AlbumZHandler(c echo.Context) error {
	AlbumZ := AmpgoFind("albumalpha", "Z", "None", "None")
	return c.JSON(http.StatusOK, AlbumZ)
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
	SongA := AmpgoFind("songalpha", "A", "None", "None")
	return c.JSON(http.StatusOK, SongA)
}

func SongBHandler(c echo.Context) error {
	SongB := AmpgoFind("songalpha", "B", "None", "None")
	return c.JSON(http.StatusOK, SongB)
}

func SongCHandler(c echo.Context) error {
	SongC := AmpgoFind("songalpha", "C", "None", "None")
	return c.JSON(http.StatusOK, SongC)
}

func SongDHandler(c echo.Context) error {
	SongD := AmpgoFind("songalpha", "D", "None", "None")
	return c.JSON(http.StatusOK, SongD)
}

func SongEHandler(c echo.Context) error {
	SongE := AmpgoFind("songalpha", "E", "None", "None")
	return c.JSON(http.StatusOK, SongE)
}

func SongFHandler(c echo.Context) error {
	SongF := AmpgoFind("songalpha", "F", "None", "None")
	return c.JSON(http.StatusOK, SongF)
}

func SongGHandler(c echo.Context) error {
	SongG := AmpgoFind("songalpha", "G", "None", "None")
	return c.JSON(http.StatusOK, SongG)
}

func SongHHandler(c echo.Context) error {
	SongH := AmpgoFind("songalpha", "H", "None", "None")
	return c.JSON(http.StatusOK, SongH)
}

func SongIHandler(c echo.Context) error {
	SongI := AmpgoFind("songalpha", "I", "None", "None")
	return c.JSON(http.StatusOK, SongI)
}

func SongJHandler(c echo.Context) error {
	SongJ := AmpgoFind("songalpha", "J", "None", "None")
	return c.JSON(http.StatusOK, SongJ)
}

func SongKHandler(c echo.Context) error {
	SongK := AmpgoFind("songalpha", "K", "None", "None")
	return c.JSON(http.StatusOK, SongK)
}

func SongLHandler(c echo.Context) error {
	SongL := AmpgoFind("songalpha", "L", "None", "None")
	return c.JSON(http.StatusOK, SongL)
}

func SongMHandler(c echo.Context) error {
	SongM := AmpgoFind("songalpha", "M", "None", "None")
	return c.JSON(http.StatusOK, SongM)
}

func SongNHandler(c echo.Context) error {
	SongN := AmpgoFind("songalpha", "N", "None", "None")
	return c.JSON(http.StatusOK, SongN)
}

func SongOHandler(c echo.Context) error {
	SongO := AmpgoFind("songalpha", "O", "None", "None")
	return c.JSON(http.StatusOK, SongO)
}

func SongPHandler(c echo.Context) error {
	SongP := AmpgoFind("songalpha", "P", "None", "None")
	return c.JSON(http.StatusOK, SongP)
}

func SongQHandler(c echo.Context) error {
	SongQ := AmpgoFind("songalpha", "Q", "None", "None")
	return c.JSON(http.StatusOK, SongQ)
}

func SongRHandler(c echo.Context) error {
	SongR := AmpgoFind("songalpha", "R", "None", "None")
	return c.JSON(http.StatusOK, SongR)
}

func SongSHandler(c echo.Context) error {
	SongS := AmpgoFind("songalpha", "S", "None", "None")
	return c.JSON(http.StatusOK, SongS)
}

func SongTHandler(c echo.Context) error {
	SongT := AmpgoFind("songalpha", "T", "None", "None")
	return c.JSON(http.StatusOK, SongT)
}

func SongUHandler(c echo.Context) error {
	SongU := AmpgoFind("songalpha", "U", "None", "None")
	return c.JSON(http.StatusOK, SongU)
}

func SongVHandler(c echo.Context) error {
	SongV := AmpgoFind("songalpha", "V", "None", "None")
	return c.JSON(http.StatusOK, SongV)
}

func SongWHandler(c echo.Context) error {
	SongW := AmpgoFind("songalpha", "W", "None", "None")
	return c.JSON(http.StatusOK, SongW)
}

func SongXHandler(c echo.Context) error {
	SongX := AmpgoFind("songalpha", "X", "None", "None")
	return c.JSON(http.StatusOK, SongX)
}

func SongYHandler(c echo.Context) error {
	SongY := AmpgoFind("songalpha", "Y", "None", "None")
	return c.JSON(http.StatusOK, SongY)
}

func SongZHandler(c echo.Context) error {
	SongZ := AmpgoFind("songalpha", "Z", "None", "None")
	return c.JSON(http.StatusOK, SongZ)
}
