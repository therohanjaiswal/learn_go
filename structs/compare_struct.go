package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {
	song1 := song{title: "Hawaayein", artist: "arijit singh"}
	song2 := song{title: "jeena jeena", artist: "atif aslam"}

	if song1 == song2 {
		fmt.Println("Songs are equal.")
	} else {
		fmt.Println("Songs are not equal.")
	}

	song1 = song2
	if song1 == song2 {
		fmt.Println("Songs are equal.")
	} else {
		fmt.Println("Songs are not equal.")
	}

	rock := playlist{genre: "indie rock", songs: []song{
		song{title: "Tum hi ho Bandhu", artist: "Kavita Seth"},
		song{title: "Malang", artist: "Siddharth Mahadevan"},
	}}
	// --------------- OR -----------------------------
	// songs := []song{
	// 	{title: "Tum hi ho Bandhu", artist: "Kavita Seth"},
	// 	{title: "Malang", artist: "Siddharth Mahadevan"},
	// }
	// rock := playlist{genre: "indie rock", songs: songs}

	// clone := rock
	// cannot compare clone == rock, since struct contains slices which are not comparables

	// struct values are bare values, cannot be changes with clone variables unless original values changed
	song := rock.songs[0]
	song.title = "Cocktail"
	fmt.Printf("Title not changed: %s\n", rock.songs[0].title)
	rock.songs[0].title = "Cocktail"
	fmt.Printf("Title changed: %s\n\n", rock.songs[0].title)

	// same with loops
	fmt.Printf("%-20s %20s\n", "TITLE", "ARTIST")
	for _, v := range rock.songs {
		fmt.Printf("%-20s %20s\n", v.title, v.artist)
	}

}
