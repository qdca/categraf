$(document).ready(function() {
    $('.ui.dropdown').dropdown();

    $('.config.item').click(function() {
        var name = $(this).html();
        load_config(name);
    });

});


function fetch_html(url) {

    fetch(url)
    .then(function(response) {
        response.text().then(function(text) {
            $('#content').html(text);
        });
    })
    .catch(function(error) {
        console.error(error);
    });

}

function update_config(name) {
    var _file = $('#file-content').val();
    $('#dimmer').addClass('active');

    $.ajax({
        type: 'POST',
        url: '/api/config/' + name,
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: JSON.stringify({
            file: _file
        }),
        statusCode: {
            200: function() {

                setTimeout(function() {
                    $('#dimmer').removeClass('active');
                }, 450);

            }
        }
    });

}

function load_config(name) {

    fetch('api/config/' + name)
    .then(function(response) {
        response.text().then(function(text) {
            $('#content').html(text);
        });
    })
    .catch(function(error) {
        console.error(error);
    });

}
