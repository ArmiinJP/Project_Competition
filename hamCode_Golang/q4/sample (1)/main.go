
package main

func NewBoard(row int, col int) AnnouncementBoard {
	newB := Board{
		allColumn: col,
		allRow: row,
		anon: []Announcements{},
	}

	return &newB
}

func NewPaper(width int, height int, ID int) AnnouncementPaper {
	//newP := Announcements 
	return nil
}

type AnnouncementBoard interface {
	getAnnouncementIDsAt(int, int) []int
}


type AnnouncementPaper interface {
	addTo(AnnouncementBoard, int, int) error
	removeAndGetIDsOnTop() []int
}

//////////////////////////////////////////////////////////
type Announcements struct{
	row int
	column int
	id []int
}

func (a *Announcements) addTo(AB AnnouncementBoard, prow int,pcol int) error{
	// AB.getAnnouncementIDsAt()
	// a.
	return nil
}

func (a *Announcements) removeAndGetIDsOnTop() []int{
	return []int{}
}


//////////////////////////////////////////////////////////
type Board struct{
	allColumn int
	allRow int
	anon []Announcements
}

func (b *Board) getAnnouncementIDsAt(row int, column int) []int{
	for _, value := range b.anon{
		if value.column == column && value.row == row{
			return value.id
		}
	}
	return []int{}
}

