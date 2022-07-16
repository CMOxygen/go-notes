const regUrl = '/reg'

//JSON.parse(Cookies.get("userdata").replace(/'/g, '"'))

$(document).ready(function () {

    $("#login-button").click(function () {
        signIn();
    });
});

function signUn() {

    let usr = {
        username: document.getElementById("reg-username").value,
        password: document.getElementById("reg-password").value
    }

    let json = JSON.stringify(usr);
    console.log(json)

    $.ajax(regUrl, {
            data: json,
            contentType: 'application/json',
            type: 'POST',
            success: function (result) {
                console.log('SUCCESS');
                console.log(result);
                console.log(JSON.parse(result));
                window.location.replace('home.html')
            },
            error: function (e) {
                console.log('ERROR');
                console.log(e);
            }
        }
    );
}