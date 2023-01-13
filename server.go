package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	SetUp()
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/artistInfoByPage", ArtistInfoByPageHandler)
	e.GET("/albumInfoByPage", AlbumInfoByPageHandler)
	e.GET("/songInfoByPage", SongInfoByPageHandler)

	e.GET("/albumsForArtist", AlbumsForArtistHandler)
	e.GET("/songsForAlbum", SongsForAlbumHandler)

	e.GET("/randomPics", RandomPicsHandler)

	e.GET("/getalbumfirstletterid", GetAlbumFirstLetterIDHandler)
	e.GET("/getalbumsforartisturl", GetAlbumsForArtistURLHandler)
	e.GET("/getalbumsforfirstletterurl", GetAlbumsForFirstLetterURLHandler)
	e.GET("/getselectedalbumid", GetSelectedAlbumIDHandler)
	e.GET("/getartistid", GetArtistIDHandler)
	e.GET("/getartistfirstletterid", GetArtistFirstLetterIDHandler)
	e.GET("/getartistsforfirstletterurl", GetArtistsForFirstLetterURLHandler)
	e.GET("/getsongid", GetSongIDHandler)
	e.GET("/getthumbhttppath", GetThumbHttpPathHandler)

	e.GET("/updatealbumfirstletterid", UpdateAlbumFirstLetterIDHandler)
	e.GET("/updatealbumsforartisturl", UpdateAlbumsForArtistURLHandler)
	e.GET("/updatealbumsforfirstletterurl", UpdateAlbumsForFirstLetterURLHandler)
	e.GET("/updateselectedalbumid", UpdateSelectedAlbumIDHandler)
	e.GET("/updateartistid", UpdateArtistIDHandler)
	e.GET("/updateartistfirstletterid", UpdateArtistFirstLetterIDHandler)
	e.GET("/updateartistsforfirstletterurl", UpdateArtistsForFirstLetterURLHandler)

	e.GET("/artistFirstLetter", ArtistFirstLetterHandler)
	e.GET("/artistsForFirstLetter", ArtistsForFirstLetterHandler)
	e.GET("/artist1", Artist1Handler)
	e.GET("/artist2", Artist2Handler)
	e.GET("/artist3", Artist3Handler)
	e.GET("/artist4", Artist4Handler)
	e.GET("/artist5", Artist5Handler)
	e.GET("/artist6", Artist6Handler)
	e.GET("/artist7", Artist7Handler)
	e.GET("/artist8", Artist8Handler)
	e.GET("/artist9", Artist9Handler)
	e.GET("/artistA", ArtistAHandler)
	e.GET("/artistB", ArtistBHandler)
	e.GET("/artistC", ArtistCHandler)
	e.GET("/artistD", ArtistDHandler)
	e.GET("/artistE", ArtistEHandler)
	e.GET("/artistF", ArtistFHandler)
	e.GET("/artistG", ArtistGHandler)
	e.GET("/artistH", ArtistHHandler)
	e.GET("/artistI", ArtistIHandler)
	e.GET("/artistJ", ArtistJHandler)
	e.GET("/artistK", ArtistKHandler)
	e.GET("/artistL", ArtistLHandler)
	e.GET("/artistM", ArtistMHandler)
	e.GET("/artistN", ArtistNHandler)
	e.GET("/artistO", ArtistOHandler)
	e.GET("/artistP", ArtistPHandler)
	e.GET("/artistQ", ArtistQHandler)
	e.GET("/artistR", ArtistRHandler)
	e.GET("/artistS", ArtistSHandler)
	e.GET("/artistT", ArtistTHandler)
	e.GET("/artistU", ArtistUHandler)
	e.GET("/artistV", ArtistVHandler)
	e.GET("/artistW", ArtistWHandler)
	e.GET("/artistX", ArtistXHandler)
	e.GET("/artistY", ArtistYHandler)
	e.GET("/artistZ", ArtistZHandler)

	e.GET("/albumFirstLetter", AlbumFirstLetterHandler)
	e.GET("/albumsForFirstLetter", AlbumsForFirstLetterHandler)
	e.GET("/album1", Album1Handler)
	e.GET("/album2", Album2Handler)
	e.GET("/album3", Album3Handler)
	e.GET("/album4", Album4Handler)
	e.GET("/album5", Album5Handler)
	e.GET("/album6", Album6Handler)
	e.GET("/album7", Album7Handler)
	e.GET("/album8", Album8Handler)
	e.GET("/album9", Album9Handler)
	e.GET("/albumA", AlbumAHandler)
	e.GET("/albumB", AlbumBHandler)
	e.GET("/albumC", AlbumCHandler)
	e.GET("/albumD", AlbumDHandler)
	e.GET("/albumE", AlbumEHandler)
	e.GET("/albumF", AlbumFHandler)
	e.GET("/albumG", AlbumGHandler)
	e.GET("/albumH", AlbumHHandler)
	e.GET("/albumI", AlbumIHandler)
	e.GET("/albumJ", AlbumJHandler)
	e.GET("/albumK", AlbumKHandler)
	e.GET("/albumL", AlbumLHandler)
	e.GET("/albumM", AlbumMHandler)
	e.GET("/albumN", AlbumNHandler)
	e.GET("/albumO", AlbumOHandler)
	e.GET("/albumP", AlbumPHandler)
	e.GET("/albumQ", AlbumQHandler)
	e.GET("/albumR", AlbumRHandler)
	e.GET("/albumS", AlbumSHandler)
	e.GET("/albumT", AlbumTHandler)
	e.GET("/albumU", AlbumUHandler)
	e.GET("/albumV", AlbumVHandler)
	e.GET("/albumW", AlbumWHandler)
	e.GET("/albumX", AlbumXHandler)
	e.GET("/albumY", AlbumYHandler)
	e.GET("/albumZ", AlbumZHandler)

	e.GET("/songFirstLetter", SongFirstLetterHandler)
	e.GET("/songsForFirstLetter", SongsForFirstLetterHandler)
	e.GET("/song1", Song1Handler)
	e.GET("/song2", Song2Handler)
	e.GET("/song3", Song3Handler)
	e.GET("/song4", Song4Handler)
	e.GET("/song5", Song5Handler)
	e.GET("/song6", Song6Handler)
	e.GET("/song7", Song7Handler)
	e.GET("/song8", Song8Handler)
	e.GET("/song9", Song9Handler)
	e.GET("/songA", SongAHandler)
	e.GET("/songB", SongBHandler)
	e.GET("/songC", SongCHandler)
	e.GET("/songD", SongDHandler)
	e.GET("/songE", SongEHandler)
	e.GET("/songF", SongFHandler)
	e.GET("/songG", SongGHandler)
	e.GET("/songH", SongHHandler)
	e.GET("/songI", SongIHandler)
	e.GET("/songJ", SongJHandler)
	e.GET("/songK", SongKHandler)
	e.GET("/songL", SongLHandler)
	e.GET("/songM", SongMHandler)
	e.GET("/songN", SongNHandler)
	e.GET("/songO", SongOHandler)
	e.GET("/songP", SongPHandler)
	e.GET("/songQ", SongQHandler)
	e.GET("/songR", SongRHandler)
	e.GET("/songS", SongSHandler)
	e.GET("/songT", SongTHandler)
	e.GET("/songU", SongUHandler)
	e.GET("/songV", SongVHandler)
	e.GET("/songW", SongWHandler)
	e.GET("/songX", SongXHandler)
	e.GET("/songY", SongYHandler)
	e.GET("/songZ", SongZHandler)

	e.Static("/static", "static")      //for pics
	e.Static("/music", "fsData/music") //for music
	e.Logger.Fatal(e.Start(":9090"))
}
