///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2023, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var OFFSET string = os.Getenv("AMPGO_OFFSET")
var OffSet int = ConvertSTR(OFFSET)

func SetUp() {
	StartSetupLogging()
	ti := time.Now()
	log.Println(ti)
	runtime.GOMAXPROCS(runtime.NumCPU())
	var addr string = os.Getenv("AMPGO_MEDIA_METADATA_PATH")
	var address string = addr + "/*.json"
	log.Println(address)
	files, err := filepath.Glob(address)
	if err != nil {
		log.Println(err)
	}

	log.Println("starting walk")
	for _, foo := range files {
		switch {
		case strings.Contains(foo, "mp3"):
			Read_File_mp3(foo)

		case strings.Contains(foo, "jpg"):
			Read_File_jpg(foo)

		case strings.Contains(foo, "page"):
			Read_File_pages(foo)

		case strings.Contains(foo, "Artist_ID"):
			Read_Artist_ID_List(foo)

		case strings.Contains(foo, "Album_ID"):
			Read_Album_ID_List(foo)
		}
	}
	log.Println("walk is complete")

	log.Println("starting GetAllMP3Objects")
	AllObjs := GetAllMP3Objects()
	log.Println("GetAllMP3Objects is complete ")

	log.Println("starting UpdateMainDB")
	var wg3 sync.WaitGroup
	var maindbpage int = 0
	for maindbIDX, blob := range AllObjs {
		if maindbIDX < OffSet {
			maindbpage = 1
		} else if maindbIDX%OffSet == 0 {
			maindbpage++
		} else {
			maindbpage = maindbpage + 0
		}
		wg3.Add(1)
		go func(blob JsonMP3, maindbpage int) {
			UpdateMainDB(blob, maindbpage)
			wg3.Done()
		}(blob, maindbpage)
		wg3.Wait()
	}
	log.Println("UpdateMainDB is complete ")

	log.Println("starting GDistArtist")
	dart := AmpgoDistinct("maindb", "mp3s", "Tags_artist")
	log.Println("GDistArtist is complete ")

	log.Println("starting ArtistFirst ")
	var wg99a sync.WaitGroup
	for _, art := range dart {
		wg99a.Add(1)
		go func(art string) {
			ArtistFirst(art)
			wg99a.Done()
		}(art)
		wg99a.Wait()
	}
	log.Println("ArtistFirst is complete ")

	log.Println("starting GetDistAlbumMeta1")
	dalb := AmpgoDistinct("maindb", "mp3s", "Tags_album")
	log.Println("GetDistAlbumMeta1 is complete ")

	log.Println("starting AlbumFirst ")
	var wg99 sync.WaitGroup
	for _, alb := range dalb {
		wg99.Add(1)
		go func(alb string) {
			AlbumFirst(alb)
			wg99.Done()
		}(alb)
		wg99.Wait()
	}
	log.Println("AlbumFirst is complete ")

	SongFirst()

	// // //AggArtist
	log.Println("starting GDistArtistForAgg")
	DistArtist := GDistArtistForAgg()
	log.Println("GDistArtistForAgg is complete ")

	fmt.Println("starting AggArtist")
	var wg5 sync.WaitGroup
	var artpage int = 0
	for artIdsx, DArtt := range DistArtist {
		if artIdsx < OffSet {
			artpage = 1
		} else if artIdsx%OffSet == 0 {
			artpage++
		} else {
			artpage = artpage + 0
		}

		APL := ArtPipline(DArtt, artpage, artIdsx)

		wg5.Add(1)
		go func(APL ArtVieW2) {
			InsArtPipeline(APL)
			wg5.Done()
		}(APL)
		wg5.Wait()
	}
	fmt.Println("AggArtists is complete")
	log.Println("AggArtists is complete")

	//AggAlbum
	log.Println("AggAlbum has started")

	log.Println("Starting GDistAlbum3")
	DistAlbum := GDistAlbum()

	var wg6 sync.WaitGroup
	var albpage int = 0
	for albIdx, DAlb := range DistAlbum {
		wg6.Add(1)
		if albIdx < OffSet {
			albpage = 1
		} else if albIdx%OffSet == 0 {
			albpage++
		} else {
			albpage = albpage + 0
		}
		APLX := AlbPipeline(DAlb, albpage, albIdx)
		go func(APLX AlbVieW2) {
			InsAlbViewID(APLX)
			wg6.Done()
		}(APLX)
		wg6.Wait()
	}

	CreateRandomPicsDB()
	ArtistAlphaList()
	AlbumAlphaList()
	SongAlphaList()
	CreateFrontMatterDB()

	// CreateRandomPlaylistDB()

	// CreateCurrentPlayListNameDB()

	t2 := time.Now().Sub(ti)
	fmt.Println(t2)
	fmt.Println("THE END")
}
