function createGame() {
    console.log("Create Game")
}

function joinGame(a, b, c) {
    const gameID = $("#joinIDField").val();

    if(gameID === "") {
        return
    }
    window.location.href = "http://localhost:5500/join/" + $("#joinIDField").val()
}

$(document).ready(( )=> {
    $("#joinForm").trigger("reset")
})
