const loginUrl = '/login'

$(document).ready(function () {

    $("#login-button").click(function () {
        signIn();
    });
});

function signIn() {

    let usr = {
        username: document.getElementById("login-username").value,
        password: document.getElementById("login-password").value
    }

    let json = JSON.stringify(usr);
    console.log(json)

    $.ajax(loginUrl, {
            data: json,
            contentType: 'application/json',
            type: 'POST',
            success: function (result) {
                console.log('SUCCESS');
                console.log(result);
                console.log(JSON.parse(result));
            },
            error: function (e) {
                console.log('ERROR');
                console.log(e);
            }
        }
    );
}