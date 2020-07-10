package main

//UpdateData as
type UpdateData struct {
	Likes, Dislikes int
	//Comments        []Comment
}

/*UpdateDataConstructor build a UpdateData type */
func UpdateDataConstructor(l, d int) (updData UpdateData) {
	updData.Likes = l
	updData.Dislikes = d
	//updData.Comments = comments
	return
}

func getUpdatedData(id int) (updated UpdateData) {
	updated.Likes = GetLikesFromID(id)
	updated.Dislikes = GetDislikesFromID(id)
	return
}
