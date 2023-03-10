package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

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
