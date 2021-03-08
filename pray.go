package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

func main() {
	var prayer string
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	flag.StringVar(&prayer, "prayer", trisagion, "choice of prayer")
	flag.Parse()
	switch prayer {
	case "morning":
		fmt.Println(morning)
	case "night":
		fmt.Println(night)
	case "creed":
		fmt.Println(creed)
	default:
		fmt.Println(trisagion)
	}
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := os.Open("chant.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

var (
	trisagion = `In the Name of the Father, and of the Son, and of the Holy Spirit. Amen.

O heavenly King, O Comforter, the Spirit of truth, who art in all
places and fillest all things; Treasury of good things and Giver of
life: Come and dwell in us and cleanse us from every stain, and save our
souls, O gracious Lord.

Holy God, Holy Mighty, Holy Immortal: have mercy on us. (Thrice)

Glory to the Father, and to the Son, and to the Holy Spirit: now and ever and unto ages of ages. Amen.

All-holy Trinity, have mercy on us. Lord, cleanse us from our sins.
Master, pardon our iniquities. Holy God, visit and heal our infirmities
for thy Name's sake.

Lord, have mercy. (Thrice)

Glory to the Father, and to the Son, and to the Holy Spirit: now and ever, and unto ages of ages. Amen.

Our Father, who art in heaven, hallowed be thy Name; thy kingdom come;
thy will be done on earth, as it is in heaven. Give us this day our
daily bread; and forgive us our trespasses, as we forgive those who
trespass against us; and lead us not into temptation, but deliver us
from evil.

Through the prayers of our holy Fathers, Lord Jesus Christ our God, have mercy on us and save us. Amen.`

	morning = `Having arisen from sleep, we fall down before thee, O Blessed One, and
sing to thee, O Mighty One, the Angelic Hymn: Holy, holy, holy art
thou, O God. Through the Theotokos have mercy on us.

Glory to the Father, and to the Son, and to the Holy Spirit:

From my bed and sleep Thou hast raised me: O Lord, enlighten my mind
and my heart, and open my lips that I may praise thee, O Holy Trinity:
Holy, holy, holy art thou, O God. Through the Theotokos have mercy on
us.

Both now and ever, and unto ages of ages. Amen.

Suddenly the Judge shall come, and the deeds of each shall be
revealed: but with fear we cry out in the middle of the night: Holy,
holy, holy art thou, O God. Through the Theotokos have mercy on us.

Lord, have mercy. (12 times)`

	creed = `
I believe in one God, the Father Almighty, Maker of heaven and earth, and of all things visible and invisible;
	
And in one Lord Jesus Christ, the Son of God, the Only-begotten,
Begotten of the Father before all worlds, Light of Light, Very God of
Very God, Begotten, not made; of one essence with the Father, by whom
all things were made:
	
Who for us men and for our salvation came down from heaven, and was
incarnate of the Holy Spirit and the Virgin Mary, and was made man;
	
And was crucified also for us under Pontius Pilate, and suffered and was buried;
	
And the third day He rose again, according to the Scriptures;
	
And ascended into heaven, and sitteth at the right hand of the Father;
	
And He shall come again with glory to judge the quick and the dead, Whose kingdom shall have no end.
	
And I believe in the Holy Spirit, the Lord, and Giver of Life, Who
proceedeth from the Father, Who with the Father and the Son together is
worshiped and glorified, Who spake by the Prophets;
	
And I believe in One Holy Catholic and Apostolic Church.
	
I acknowledge one Baptism for the remission of sins.
	
I look for the Resurrection of the dead.
	
And the Life of the world to come. Amen.`

	night = `Enlighten mine eyes, O Christ God, lest at any time I sleep unto death, lest at any time mine enemy say: I have prevailed against him.

Glory to the Father, and to the Son, and to the Holy Spirit.
	
Be my soul's helper, O God, for I pass through the midst of many snares; deliver me out of them and save me, O Good One, for Thou art the Lover of mankind.
	
Into Thy hands, O Lord Jesus Christ, my God, I commend my spirit. Do thou bless me, have mercy on me and grant me eternal life. Amen.`
)
