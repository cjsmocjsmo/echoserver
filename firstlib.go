package main

import (
	"sort"
	"strconv"
	"strings"
)

func ArtistFirst(astring string) string {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
	CheckError(err, "ArtistFirst: Connections has failed")
	defer Close(client, ctx, cancel)

	char := strings.ToUpper(astring[:1])

	switch {
	// case char == "8":
	// 	artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
	// 	item := make(map[string]string)
	// 	item["artist"] = astring
	// 	item["artistid"] = artid["ArtistID"]
	// 	_, erra := InsertOne(client, ctx, "artistalpha", "8", item)
	// 	CheckError(erra, "ArtistFirst: A insertion has failed")
	// 	return "8 Created"

	case char == "3":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, erra := InsertOne(client, ctx, "artistalpha", "fl3", item)
		CheckError(erra, "ArtistFirst: A insertion has failed")
		return "3 Created"

	case char == "A":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, erra := InsertOne(client, ctx, "artistalpha", "A", item)
		CheckError(erra, "ArtistFirst: A insertion has failed")
		return "A Created"

	case char == "B":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errb := InsertOne(client, ctx, "artistalpha", "B", item)
		CheckError(errb, "ArtistFirst: B insertion has failed")
		return "B Created"

	case char == "C":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errc := InsertOne(client, ctx, "artistalpha", "C", item)
		CheckError(errc, "ArtistFirst: C insertion has failed")
		return "C Created"

	case char == "D":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errd := InsertOne(client, ctx, "artistalpha", "D", item)
		CheckError(errd, "ArtistFirst: D insertion has failed")
		return "D Created"

	case char == "E":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, erre := InsertOne(client, ctx, "artistalpha", "E", item)
		CheckError(erre, "ArtistFirst: E insertion has failed")
		return "E Created"

	case char == "F":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errf := InsertOne(client, ctx, "artistalpha", "F", item)
		CheckError(errf, "ArtistFirst: F insertion has failed")
		return "F Created"

	case char == "G":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errg := InsertOne(client, ctx, "artistalpha", "G", item)
		CheckError(errg, "ArtistFirst: G insertion has failed")
		return "G Created"

	case char == "H":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errh := InsertOne(client, ctx, "artistalpha", "H", item)
		CheckError(errh, "ArtistFirst: H insertion has failed")
		return "H Created"

	case char == "I":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, erri := InsertOne(client, ctx, "artistalpha", "I", item)
		CheckError(erri, "ArtistFirst: I insertion has failed")
		return "I Created"

	case char == "J":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errj := InsertOne(client, ctx, "artistalpha", "J", item)
		CheckError(errj, "ArtistFirst: J insertion has failed")
		return "J Created"

	case char == "K":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errk := InsertOne(client, ctx, "artistalpha", "K", item)
		CheckError(errk, "ArtistFirst: K insertion has failed")
		return "K Created"

	case char == "L":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errl := InsertOne(client, ctx, "artistalpha", "L", item)
		CheckError(errl, "ArtistFirst: L insertion has failed")
		return "L Created"

	case char == "M":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errm := InsertOne(client, ctx, "artistalpha", "M", item)
		CheckError(errm, "ArtistFirst: M insertion has failed")
		return "M Created"

	case char == "N":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errn := InsertOne(client, ctx, "artistalpha", "N", item)
		CheckError(errn, "ArtistFirst: N insertion has failed")
		return "N Created"

	case char == "O":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, erro := InsertOne(client, ctx, "artistalpha", "O", item)
		CheckError(erro, "ArtistFirst: O insertion has failed")
		return "O Created"

	case char == "P":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errp := InsertOne(client, ctx, "artistalpha", "P", item)
		CheckError(errp, "ArtistFirst: P insertion has failed")
		return "P Created"

	case char == "Q":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errq := InsertOne(client, ctx, "artistalpha", "Q", item)
		CheckError(errq, "ArtistFirst: Q insertion has failed")
		return "Q Created"

	case char == "R":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errr := InsertOne(client, ctx, "artistalpha", "R", item)
		CheckError(errr, "ArtistFirst: R insertion has failed")
		return "R Created"

	case char == "S":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errs := InsertOne(client, ctx, "artistalpha", "S", item)
		CheckError(errs, "ArtistFirst: S insertion has failed")
		return "S Created"

	case char == "T":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errt := InsertOne(client, ctx, "artistalpha", "T", item)
		CheckError(errt, "ArtistFirst: T insertion has failed")
		return "T Created"

	case char == "U":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, erru := InsertOne(client, ctx, "artistalpha", "U", item)
		CheckError(erru, "ArtistFirst: U insertion has failed")
		return "U Created"

	case char == "V":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errv := InsertOne(client, ctx, "artistalpha", "V", item)
		CheckError(errv, "ArtistFirst: V insertion has failed")
		return "V Created"

	case char == "W":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errw := InsertOne(client, ctx, "artistalpha", "W", item)
		CheckError(errw, "ArtistFirst: W insertion has failed")
		return "W Created"

	case char == "X":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errx := InsertOne(client, ctx, "artistalpha", "X", item)
		CheckError(errx, "ArtistFirst: X insertion has failed")
		return "X Created"

	case char == "Z":
		artid := AmpgoFindOne("ids", "artistids", "Artist", astring)
		item := make(map[string]string)
		item["artist"] = astring
		item["artistid"] = artid["ArtistID"]
		_, errz := InsertOne(client, ctx, "artistalpha", "Z", item)
		CheckError(errz, "ArtistFirst: Z insertion has failed")
		return "Z Created"
	}
	return "None"
}

func AlbumFirst(astring string) string {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
	CheckError(err, "AlbumFirst:  Connections has failed")
	defer Close(client, ctx, cancel)

	char := strings.ToUpper(astring[:1])

	switch {
	case char == "2":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erra := InsertOne(client, ctx, "albumalpha", "fl2", item)
		CheckError(erra, "AlbumFirst: A insertion has failed")
		return "2 Created"

	case char == "3":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erra := InsertOne(client, ctx, "albumalpha", "fl3", item)
		CheckError(erra, "AlbumFirst: A insertion has failed")
		return "3 Created"

	case char == "A":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erra := InsertOne(client, ctx, "albumalpha", "A", item)
		CheckError(erra, "AlbumFirst: A insertion has failed")
		return "A Created"

	case char == "B":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errb := InsertOne(client, ctx, "albumalpha", "B", item)
		CheckError(errb, "AlbumFirst: B insertion has failed")
		return "B Created"

	case char == "C":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errc := InsertOne(client, ctx, "albumalpha", "C", item)
		CheckError(errc, "AlbumFirst: C insertion has failed")
		return "C Created"

	case char == "D":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errd := InsertOne(client, ctx, "albumalpha", "D", item)
		CheckError(errd, "AlbumFirst: D insertion has failed")
		return "D Created"

	case char == "E":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erre := InsertOne(client, ctx, "albumalpha", "E", item)
		CheckError(erre, "AlbumFirst: E insertion has failed")
		return "E Created"

	case char == "F":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errf := InsertOne(client, ctx, "albumalpha", "F", item)
		CheckError(errf, "AlbumFirst: F insertion has failed")
		return "F Created"

	case char == "G":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errg := InsertOne(client, ctx, "albumalpha", "G", item)
		CheckError(errg, "AlbumFirst: G insertion has failed")
		return "G Created"

	case char == "H":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errh := InsertOne(client, ctx, "albumalpha", "H", item)
		CheckError(errh, "AlbumFirst: H insertion has failed")
		return "H Created"

	case char == "I":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erri := InsertOne(client, ctx, "albumalpha", "I", item)
		CheckError(erri, "AlbumFirst: I insertion has failed")
		return "I Created"

	case char == "J":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errj := InsertOne(client, ctx, "albumalpha", "J", item)
		CheckError(errj, "AlbumFirst: J insertion has failed")
		return "J Created"

	case char == "K":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errk := InsertOne(client, ctx, "albumalpha", "K", item)
		CheckError(errk, "AlbumFirst: K insertion has failed")
		return "K Created"

	case char == "L":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errl := InsertOne(client, ctx, "albumalpha", "L", item)
		CheckError(errl, "AlbumFirst: L insertion has failed")
		return "L Created"

	case char == "M":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errm := InsertOne(client, ctx, "albumalpha", "M", item)
		CheckError(errm, "AlbumFirst: M insertion has failed")
		return "M Created"

	case char == "N":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errn := InsertOne(client, ctx, "albumalpha", "N", item)
		CheckError(errn, "AlbumFirst: N insertion has failed")
		return "N Created"

	case char == "O":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erro := InsertOne(client, ctx, "albumalpha", "O", item)
		CheckError(erro, "AlbumFirst: O insertion has failed")
		return "O Created"

	case char == "P":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errp := InsertOne(client, ctx, "albumalpha", "P", item)
		CheckError(errp, "AlbumFirst: P insertion has failed")
		return "P Created"

	case char == "Q":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errq := InsertOne(client, ctx, "albumalpha", "Q", item)
		CheckError(errq, "AlbumFirst: Q insertion has failed")
		return "Q Created"

	case char == "R":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errr := InsertOne(client, ctx, "albumalpha", "R", item)
		CheckError(errr, "AlbumFirst: R insertion has failed")
		return "R Created"

	case char == "S":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errs := InsertOne(client, ctx, "albumalpha", "S", item)
		CheckError(errs, "AlbumFirst: S insertion has failed")
		return "S Created"

	case char == "T":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errt := InsertOne(client, ctx, "albumalpha", "T", item)
		CheckError(errt, "AlbumFirst: T insertion has failed")
		return "T Created"

	case char == "U":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, erru := InsertOne(client, ctx, "albumalpha", "U", item)
		CheckError(erru, "AlbumFirst: U insertion has failed")
		return "U Created"

	case char == "V":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errv := InsertOne(client, ctx, "albumalpha", "V", item)
		CheckError(errv, "AlbumFirst: V insertion has failed")
		return "V Created"

	case char == "W":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errw := InsertOne(client, ctx, "albumalpha", "W", item)
		CheckError(errw, "AlbumFirst: W insertion has failed")
		return "W Created"

	case char == "X":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errx := InsertOne(client, ctx, "albumalpha", "X", item)
		CheckError(errx, "AlbumFirst: X insertion has failed")
		return "X Created"

	case char == "Z":
		albid := AmpgoFindOne("ids", "albumids", "Album", astring)
		item := make(map[string]string)
		item["album"] = astring
		item["albumid"] = albid["AlbumID"]
		_, errz := InsertOne(client, ctx, "albumalpha", "Z", item)
		CheckError(errz, "AlbumFirst: Z insertion has failed")
		return "Z Created"
	}
	return "None"
}

func SongFirst() string {
	client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
	CheckError(err, "SongFirst: Connections has failed")
	defer Close(client, ctx, cancel)

	All1 := AmpgoFind("maindb", "maindb", "Song_first", "1")
	for _, a1 := range All1 {
		_, err = InsertOne(client, ctx, "songalpha", "fl1", a1)
		CheckError(err, "SongFirst: a insertion has failed")
	}
	a1 := len(All1)

	All3 := AmpgoFind("maindb", "maindb", "Song_first", "3")
	for _, a3 := range All3 {
		_, err = InsertOne(client, ctx, "songalpha", "fl3", a3)
		CheckError(err, "SongFirst: a insertion has failed")
	}
	a3 := len(All3)

	All7 := AmpgoFind("maindb", "maindb", "Song_first", "7")
	for _, a7 := range All7 {
		_, err = InsertOne(client, ctx, "songalpha", "fl7", a7)
		CheckError(err, "SongFirst: a insertion has failed")
	}
	a7 := len(All7)

	All9 := AmpgoFind("maindb", "maindb", "Song_first", "9")
	for _, a9 := range All9 {
		_, err = InsertOne(client, ctx, "songalpha", "fl9", a9)
		CheckError(err, "SongFirst: a insertion has failed")
	}
	a9 := len(All9)

	aAll := AmpgoFind("maindb", "maindb", "Song_first", "A")
	for _, a := range aAll {
		_, err = InsertOne(client, ctx, "songalpha", "A", a)
		CheckError(err, "SongFirst: a insertion has failed")
	}
	aa := len(aAll)

	bAll := AmpgoFind("maindb", "maindb", "Song_first", "B")
	for _, b := range bAll {
		_, err = InsertOne(client, ctx, "songalpha", "B", b)
		CheckError(err, "SongFirst: b insertion has failed")
	}
	bb := len(bAll)

	cAll := AmpgoFind("maindb", "maindb", "Song_first", "C")
	for _, c := range cAll {
		_, err = InsertOne(client, ctx, "songalpha", "C", c)
		CheckError(err, "SongFirst: c insertion has failed")
	}
	cc := len(cAll)

	dAll := AmpgoFind("maindb", "maindb", "Song_first", "D")
	for _, d := range dAll {
		_, err = InsertOne(client, ctx, "songalpha", "D", d)
		CheckError(err, "SongFirst: d insertion has failed")
	}
	dd := len(dAll)

	eAll := AmpgoFind("maindb", "maindb", "Song_first", "E")
	for _, e := range eAll {
		_, err = InsertOne(client, ctx, "songalpha", "E", e)
		CheckError(err, "SongFirst: e insertion has failed")
	}
	ee := len(eAll)

	fAll := AmpgoFind("maindb", "maindb", "Song_first", "F")
	for _, f := range fAll {
		_, err = InsertOne(client, ctx, "songalpha", "F", f)
		CheckError(err, "SongFirst: f insertion has failed")
	}
	ff := len(fAll)

	gAll := AmpgoFind("maindb", "maindb", "Song_first", "G")
	gg := len(gAll)
	for _, g := range gAll {
		_, err = InsertOne(client, ctx, "songalpha", "G", g)
		CheckError(err, "SongFirst: g insertion has failed")
	}

	hAll := AmpgoFind("maindb", "maindb", "Song_first", "H")
	hh := len(hAll)
	for _, h := range hAll {
		_, err = InsertOne(client, ctx, "songalpha", "H", h)
		CheckError(err, "SongFirst: h insertion has failed")
	}

	iAll := AmpgoFind("maindb", "maindb", "Song_first", "I")
	ii := len(iAll)
	for _, i := range iAll {
		_, err = InsertOne(client, ctx, "songalpha", "I", i)
		CheckError(err, "SongFirst: i insertion has failed")
	}

	jAll := AmpgoFind("maindb", "maindb", "Song_first", "J")
	jj := len(jAll)
	for _, j := range jAll {
		_, err = InsertOne(client, ctx, "songalpha", "J", j)
		CheckError(err, "SongFirst: j insertion has failed")
	}

	kAll := AmpgoFind("maindb", "maindb", "Song_first", "K")
	kk := len(kAll)
	for _, k := range kAll {
		_, err = InsertOne(client, ctx, "songalpha", "K", k)
		CheckError(err, "SongFirst: k insertion has failed")
	}

	lAll := AmpgoFind("maindb", "maindb", "Song_first", "L")
	ll := len(lAll)
	for _, l := range lAll {
		_, err = InsertOne(client, ctx, "songalpha", "L", l)
		CheckError(err, "SongFirst: l insertion has failed")
	}

	mAll := AmpgoFind("maindb", "maindb", "Song_first", "M")
	mm := len(mAll)
	for _, m := range mAll {
		_, err = InsertOne(client, ctx, "songalpha", "M", m)
		CheckError(err, "SongFirst: m insertion has failed")
	}

	nAll := AmpgoFind("maindb", "maindb", "Song_first", "N")
	nn := len(nAll)
	for _, n := range nAll {
		_, err = InsertOne(client, ctx, "songalpha", "N", n)
		CheckError(err, "SongFirst: n insertion has failed")
	}

	oAll := AmpgoFind("maindb", "maindb", "Song_first", "O")
	oo := len(oAll)
	for _, o := range oAll {
		_, err = InsertOne(client, ctx, "songalpha", "O", o)
		CheckError(err, "SongFirst: o insertion has failed")
	}

	pAll := AmpgoFind("maindb", "maindb", "Song_first", "P")
	pp := len(pAll)
	for _, p := range pAll {
		_, err = InsertOne(client, ctx, "songalpha", "P", p)
		CheckError(err, "SongFirst: p insertion has failed")
	}

	qAll := AmpgoFind("maindb", "maindb", "Song_first", "Q")
	qq := len(qAll)
	for _, q := range qAll {
		_, err = InsertOne(client, ctx, "songalpha", "Q", q)
		CheckError(err, "SongFirst: q insertion has failed")
	}

	rAll := AmpgoFind("maindb", "maindb", "Song_first", "R")
	rr := len(rAll)
	for _, r := range rAll {
		_, err = InsertOne(client, ctx, "songalpha", "R", r)
		CheckError(err, "SongFirst: r insertion has failed")
	}

	sAll := AmpgoFind("maindb", "maindb", "Song_first", "S")
	ss := len(sAll)
	for _, s := range sAll {
		_, err = InsertOne(client, ctx, "songalpha", "S", s)
		CheckError(err, "SongFirst: s insertion has failed")
	}

	tAll := AmpgoFind("maindb", "maindb", "Song_first", "T")
	tt := len(tAll)
	for _, t := range tAll {
		_, err = InsertOne(client, ctx, "songalpha", "T", t)
		CheckError(err, "SongFirst: t insertion has failed")
	}

	uAll := AmpgoFind("maindb", "maindb", "Song_first", "U")
	uu := len(uAll)
	for _, u := range uAll {
		_, err = InsertOne(client, ctx, "songalpha", "U", u)
		CheckError(err, "SongFirst: u insertion has failed")
	}

	vAll := AmpgoFind("maindb", "maindb", "Song_first", "V")
	vv := len(vAll)
	for _, v := range vAll {
		_, err = InsertOne(client, ctx, "songalpha", "V", v)
		CheckError(err, "SongFirst: v insertion has failed")
	}

	wAll := AmpgoFind("maindb", "maindb", "Song_first", "W")
	ww := len(wAll)
	for _, w := range wAll {
		_, err = InsertOne(client, ctx, "songalpha", "W", w)
		CheckError(err, "SongFirst: w insertion has failed")
	}

	xAll := AmpgoFind("maindb", "maindb", "Song_first", "X")
	xx := len(xAll)
	for _, x := range xAll {
		_, err = InsertOne(client, ctx, "songalpha", "X", x)
		CheckError(err, "SongFirst: x insertion has failed")
	}

	yAll := AmpgoFind("maindb", "maindb", "Song_first", "Y")
	yy := len(yAll)
	for _, y := range yAll {
		_, err = InsertOne(client, ctx, "songalpha", "Y", y)
		CheckError(err, "SongFirst: y insertion has failed")
	}

	zAll := AmpgoFind("maindb", "maindb", "Song_first", "Z")
	zz := len(zAll)
	for _, z := range zAll {
		_, err = InsertOne(client, ctx, "songalpha", "Z", z)
		CheckError(err, "SongFirst: z insertion has failed")
	}
	tot := 0
	tototal := []int{a1, a3, a7, a9, aa, bb, cc, dd, ee, ff, gg, hh, ii,
		jj, kk, ll, mm, nn, oo, pp, qq, rr, ss, tt, uu, vv, ww, xx, yy, zz,
	}
	for _, v := range tototal {
		tot += v
	}
	total := strconv.Itoa(tot)
	total2 := make(map[string]string)
	total2["total"] = total

	client, ctx, cancel, err = Connect("mongodb://db:27017/ampgo")
	CheckError(err, "SongFirst: Connections has failed")
	defer Close(client, ctx, cancel)

	_, err = InsertOne(client, ctx, "songtotal", "total", &total2)
	CheckError(err, "SongFirst: z insertion has failed")

	return "Complete"
}

func ArtistAlphaList() {
	theletters := GetMainDbMeta()
	var newlist []string
	for _, fl := range theletters {
		boo := fl["Artist_first"]
		newlist = append(newlist, boo)
	}
	uniquelist := Unique(newlist)
	sort.Strings(uniquelist)
	for _, firstletter := range uniquelist {
		id := "artist" + firstletter
		ufl := make(map[string]string)
		ufl["firstletter"] = firstletter
		ufl["firstletterid"] = id
		ufl["firstletterpath"] = "artist/" + id
		client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
		CheckError(err, "artistalpha:  Connections has failed")
		defer Close(client, ctx, cancel)
		_, erra := InsertOne(client, ctx, "artistalpha", "artistalphalist", ufl)
		CheckError(erra, "artistalpha: A insertion has failed")
	}
	return
}

func AlbumAlphaList() {
	theletters := GetMainDbMeta()
	var newlist []string
	for _, fl := range theletters {
		boo := fl["Album_first"]
		newlist = append(newlist, boo)
	}
	uniquelist := Unique(newlist)
	sort.Strings(uniquelist)
	for _, firstletteralb := range uniquelist {
		albid := "album" + firstletteralb
		albufl := make(map[string]string)
		albufl["firstletter"] = firstletteralb
		albufl["firstletterid"] = albid
		albufl["firstletterpath"] = "album/" + albid
		client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
		CheckError(err, "albumalpha:  Connections has failed")
		defer Close(client, ctx, cancel)
		_, erra := InsertOne(client, ctx, "albumalpha", "albumalphalist", albufl)
		CheckError(erra, "albumalpha: A insertion has failed")
	}
	return
}

func SongAlphaList() {
	theletters := GetMainDbMeta()
	var newlist []string
	for _, fl := range theletters {
		boo := fl["Song_first"]
		newlist = append(newlist, boo)
	}
	uniquelist := Unique(newlist)
	sort.Strings(uniquelist)
	for _, firstletterS := range uniquelist {
		sid := "song" + firstletterS
		sufl := make(map[string]string)
		sufl["firstletter"] = firstletterS
		sufl["firstletterid"] = sid
		sufl["firstletterpath"] = "song/" + sid
		client, ctx, cancel, err := Connect("mongodb://db:27017/ampgo")
		CheckError(err, "songalpha:  Connections has failed")
		defer Close(client, ctx, cancel)
		_, erra := InsertOne(client, ctx, "songalpha", "songalphalist", sufl)
		CheckError(erra, "songalpha: A insertion has failed")
	}
	return
}
