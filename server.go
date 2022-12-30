package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	SetUp()
	e := echo.New()
	e.GET("/artistInfoByPage", artistInfoByPageHandler)
	e.GET("/albumInfoByPage", albumInfoByPageHandler)
	e.GET("/songInfoByPage", songInfoByPageHandler)

	e.GET("/albumsForArtist", albumsForArtistHandler)
	e.GET("/songsForAlbum", songsForAlbumHandler)

	e.GET("/randomPics", randomPicsHandler)

	e.GET("/artistFirstLetter", ArtistFirstLetterHandler)
	e.GET("/artist3", Artist3Handler)
	e.GET("/artistA", ArtistAHandler)
	e.GET("/artistB", ArtistBHandler)

	e.GET("/albumFirstLetter", AlbumFirstLetterHandler)
	e.GET("/album2", Album2Handler)
	e.GET("/album3", Album3Handler)
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
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)


	e.GET("/songFirstLetter", SongFirstLetterHandler)
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

	// e.GET("/songD", SongDHandler)
	// e.GET("/songE", SongEHandler)
	// e.GET("/songF", SongFHandler)
	// e.GET("/songG", SongGHandler)
	// e.GET("/songH", SongHHandler)
	// e.GET("/songI", SongIHandler)
	// e.GET("/songJ", SongJHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)
	// e.GET("/album", AlbumHandler)

	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":9090"))
}
