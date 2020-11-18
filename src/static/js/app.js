function createGame() {
    var newID = ""
    $.ajax({
        url: "http://localhost:5500/newSession"
    }).done((gameData) => {
        console.log(gameData.ID)
    })
}
