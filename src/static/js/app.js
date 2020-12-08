function createGame() {
    $("#createButton").prop("disabled", true)
    $.ajax({
        url: "http://localhost:5500/newGame"
    }).done((gameData) => {
        console.log(gameData.ID)
        $("#joinIDField").val(gameData.ID)
    })
}

$(document).ready(() => {
    $("#joinIDField").val(undefined)
})
